import os
import json
import numpy as np
import pandas as pd
from sklearn.preprocessing import StandardScaler, OneHotEncoder
from sklearn.compose import ColumnTransformer
from sklearn.pipeline import Pipeline

class DataPreprocessor:
    """
    Preprocesses raw telemetry data from eBPF collectors for machine learning models.
    Handles feature extraction, normalization, and encoding for different event types.
    """
    
    def __init__(self, config=None):
        """
        Initialize the preprocessor with configuration.
        
        Args:
            config (dict): Configuration parameters for preprocessing
        """
        self.config = config or {}
        self.scalers = {}
        self.encoders = {}
        self.feature_columns = {}
        
        # Define feature columns for each event type
        self._define_feature_columns()
        
    def _define_feature_columns(self):
        """Define the feature columns for each event type"""
        
        # Syscall features
        self.feature_columns['syscall'] = {
            'numerical': ['syscall_id', 'return_value', 'error_code', 'duration_ns'],
            'categorical': ['syscall_name', 'container_id'],
            'text': ['arguments'],
            'timestamp': 'timestamp'
        }
        
        # Network features
        self.feature_columns['network'] = {
            'numerical': ['source_port', 'destination_port', 'bytes_sent', 'bytes_received'],
            'categorical': ['protocol', 'connection_state', 'container_id', 'interface_name'],
            'ip_addresses': ['source_ip', 'destination_ip'],
            'timestamp': 'timestamp'
        }
        
        # Process features
        self.feature_columns['process'] = {
            'numerical': ['process_id', 'parent_process_id', 'user_id', 'group_id', 'exit_code'],
            'categorical': ['event_type', 'container_id', 'namespace', 'pod_name'],
            'text': ['command', 'arguments'],
            'timestamp': ['start_time', 'end_time']
        }
        
        # File features
        self.feature_columns['file'] = {
            'numerical': ['process_id', 'user_id', 'group_id', 'result', 'permissions', 'bytes_accessed'],
            'categorical': ['operation', 'container_id', 'namespace', 'pod_name'],
            'text': ['path'],
            'timestamp': 'timestamp'
        }
        
        # Container features
        self.feature_columns['container'] = {
            'numerical': ['cpu_limit', 'memory_limit'],
            'categorical': ['event_type', 'network_mode', 'namespace', 'pod_name', 'node_name'],
            'text': ['image_name', 'command'],
            'list': ['capabilities', 'mounts'],
            'timestamp': 'timestamp'
        }
    
    def preprocess_events(self, events, event_type):
        """
        Preprocess events of a specific type.
        
        Args:
            events (list): List of event dictionaries
            event_type (str): Type of events ('syscall', 'network', 'process', 'file', 'container')
            
        Returns:
            pd.DataFrame: Preprocessed features ready for model input
        """
        if not events:
            return None
            
        # Convert to DataFrame
        df = pd.DataFrame(events)
        
        # Extract features based on event type
        if event_type in self.feature_columns:
            return self._extract_features(df, event_type)
        else:
            raise ValueError(f"Unknown event type: {event_type}")
    
    def _extract_features(self, df, event_type):
        """Extract and transform features for a specific event type"""
        
        features = self.feature_columns[event_type]
        result_dfs = []
        
        # Process numerical features
        if features.get('numerical') and all(col in df.columns for col in features['numerical']):
            num_df = df[features['numerical']].copy()
            
            # Create or use existing scaler
            if f"{event_type}_num_scaler" not in self.scalers:
                self.scalers[f"{event_type}_num_scaler"] = StandardScaler()
                num_features = self.scalers[f"{event_type}_num_scaler"].fit_transform(num_df)
            else:
                num_features = self.scalers[f"{event_type}_num_scaler"].transform(num_df)
                
            num_df = pd.DataFrame(
                num_features, 
                columns=[f"{col}_scaled" for col in features['numerical']]
            )
            result_dfs.append(num_df)
        
        # Process categorical features
        if features.get('categorical') and all(col in df.columns for col in features['categorical']):
            cat_df = df[features['categorical']].copy()
            
            # Create or use existing encoder
            if f"{event_type}_cat_encoder" not in self.encoders:
                self.encoders[f"{event_type}_cat_encoder"] = {}
                
                for col in features['categorical']:
                    encoder = OneHotEncoder(sparse_output=False, handle_unknown='ignore')
                    values = cat_df[col].values.reshape(-1, 1)
                    encoded = encoder.fit_transform(values)
                    self.encoders[f"{event_type}_cat_encoder"][col] = encoder
                    
                    # Create DataFrame with encoded values
                    encoded_df = pd.DataFrame(
                        encoded,
                        columns=[f"{col}_{val}" for val in encoder.categories_[0]]
                    )
                    result_dfs.append(encoded_df)
            else:
                for col in features['categorical']:
                    encoder = self.encoders[f"{event_type}_cat_encoder"][col]
                    values = cat_df[col].values.reshape(-1, 1)
                    encoded = encoder.transform(values)
                    
                    # Create DataFrame with encoded values
                    encoded_df = pd.DataFrame(
                        encoded,
                        columns=[f"{col}_{val}" for val in encoder.categories_[0]]
                    )
                    result_dfs.append(encoded_df)
        
        # Process IP addresses (for network events)
        if features.get('ip_addresses') and all(col in df.columns for col in features['ip_addresses']):
            for col in features['ip_addresses']:
                # Convert IP to numerical representation
                ip_features = df[col].apply(self._ip_to_features)
                ip_df = pd.DataFrame(
                    ip_features.tolist(),
                    columns=[f"{col}_octet_{i}" for i in range(4)]
                )
                
                # Scale IP features
                if f"{event_type}_{col}_scaler" not in self.scalers:
                    self.scalers[f"{event_type}_{col}_scaler"] = StandardScaler()
                    ip_scaled = self.scalers[f"{event_type}_{col}_scaler"].fit_transform(ip_df)
                else:
                    ip_scaled = self.scalers[f"{event_type}_{col}_scaler"].transform(ip_df)
                    
                ip_df = pd.DataFrame(
                    ip_scaled,
                    columns=[f"{col}_octet_{i}_scaled" for i in range(4)]
                )
                result_dfs.append(ip_df)
        
        # Process text features (basic processing)
        if features.get('text') and all(col in df.columns for col in features['text']):
            for col in features['text']:
                # For command and arguments, extract length and presence of special characters
                if col in ['command', 'arguments', 'path']:
                    if isinstance(df[col].iloc[0], list):
                        # Join lists into strings
                        text_data = df[col].apply(lambda x: ' '.join(x) if isinstance(x, list) else str(x))
                    else:
                        text_data = df[col].astype(str)
                    
                    # Extract basic text features
                    text_df = pd.DataFrame({
                        f"{col}_length": text_data.str.len(),
                        f"{col}_word_count": text_data.str.split().str.len(),
                        f"{col}_special_chars": text_data.str.count(r'[^a-zA-Z0-9\s]')
                    })
                    
                    # Scale text features
                    if f"{event_type}_{col}_text_scaler" not in self.scalers:
                        self.scalers[f"{event_type}_{col}_text_scaler"] = StandardScaler()
                        text_scaled = self.scalers[f"{event_type}_{col}_text_scaler"].fit_transform(text_df)
                    else:
                        text_scaled = self.scalers[f"{event_type}_{col}_text_scaler"].transform(text_df)
                        
                    text_df = pd.DataFrame(
                        text_scaled,
                        columns=[f"{c}_scaled" for c in text_df.columns]
                    )
                    result_dfs.append(text_df)
        
        # Process timestamp features
        if features.get('timestamp'):
            timestamp_cols = [features['timestamp']] if isinstance(features['timestamp'], str) else features['timestamp']
            
            for col in timestamp_cols:
                if col in df.columns:
                    # Convert timestamp to datetime and extract features
                    df[col] = pd.to_datetime(df[col], unit='ns')
                    time_df = pd.DataFrame({
                        f"{col}_hour": df[col].dt.hour,
                        f"{col}_day": df[col].dt.day,
                        f"{col}_weekday": df[col].dt.weekday,
                        f"{col}_month": df[col].dt.month
                    })
                    
                    # One-hot encode time features
                    for time_col in time_df.columns:
                        if f"{event_type}_{time_col}_encoder" not in self.encoders:
                            encoder = OneHotEncoder(sparse_output=False, handle_unknown='ignore')
                            values = time_df[time_col].values.reshape(-1, 1)
                            encoded = encoder.fit_transform(values)
                            self.encoders[f"{event_type}_{time_col}_encoder"] = encoder
                            
                            # Create DataFrame with encoded values
                            encoded_df = pd.DataFrame(
                                encoded,
                                columns=[f"{time_col}_{val}" for val in encoder.categories_[0]]
                            )
                            result_dfs.append(encoded_df)
                        else:
                            encoder = self.encoders[f"{event_type}_{time_col}_encoder"]
                            values = time_df[time_col].values.reshape(-1, 1)
                            encoded = encoder.transform(values)
                            
                            # Create DataFrame with encoded values
                            encoded_df = pd.DataFrame(
                                encoded,
                                columns=[f"{time_col}_{val}" for val in encoder.categories_[0]]
                            )
                            result_dfs.append(encoded_df)
        
        # Combine all feature DataFrames
        if result_dfs:
            # Add index to all dataframes to ensure proper concatenation
            for i, df_part in enumerate(result_dfs):
                df_part.index = range(len(df_part))
            
            return pd.concat(result_dfs, axis=1)
        else:
            return pd.DataFrame()
    
    def _ip_to_features(self, ip):
        """Convert IP address to numerical features"""
        try:
            return [int(x) for x in ip.split('.')]
        except:
            return [0, 0, 0, 0]
    
    def save(self, directory):
        """Save preprocessor state to directory"""
        os.makedirs(directory, exist_ok=True)
        
        # Save feature columns
        with open(os.path.join(directory, 'feature_columns.json'), 'w') as f:
            json.dump(self.feature_columns, f)
        
        # Save scalers and encoders (using pickle through sklearn)
        for name, scaler in self.scalers.items():
            pd.to_pickle(scaler, os.path.join(directory, f'{name}.pkl'))
        
        for name, encoder_dict in self.encoders.items():
            if isinstance(encoder_dict, dict):
                for col, encoder in encoder_dict.items():
                    pd.to_pickle(encoder, os.path.join(directory, f'{name}_{col}.pkl'))
            else:
                pd.to_pickle(encoder_dict, os.path.join(directory, f'{name}.pkl'))
    
    @classmethod
    def load(cls, directory):
        """Load preprocessor state from directory"""
        preprocessor = cls()
        
        # Load feature columns
        with open(os.path.join(directory, 'feature_columns.json'), 'r') as f:
            preprocessor.feature_columns = json.load(f)
        
        # Load scalers and encoders
        for filename in os.listdir(directory):
            if filename.endswith('.pkl'):
                path = os.path.join(directory, filename)
                if '_cat_encoder_' in filename:
                    # Handle nested encoder dictionaries
                    parts = filename.replace('.pkl', '').split('_cat_encoder_')
                    event_type = parts[0]
                    col = parts[1]
                    
                    if f"{event_type}_cat_encoder" not in preprocessor.encoders:
                        preprocessor.encoders[f"{event_type}_cat_encoder"] = {}
                    
                    preprocessor.encoders[f"{event_type}_cat_encoder"][col] = pd.read_pickle(path)
                elif '_cat_encoder' in filename:
                    name = filename.replace('.pkl', '')
                    preprocessor.encoders[name] = {}
                elif '_encoder' in filename:
                    name = filename.replace('.pkl', '')
                    preprocessor.encoders[name] = pd.read_pickle(path)
                elif '_scaler' in filename:
                    name = filename.replace('.pkl', '')
                    preprocessor.scalers[name] = pd.read_pickle(path)
        
        return preprocessor
