import os
import numpy as np
import tensorflow as tf
from tensorflow.keras.models import Model
from tensorflow.keras.layers import Input, Dense, LSTM, Dropout, BatchNormalization
from tensorflow.keras.callbacks import EarlyStopping, ModelCheckpoint

class AnomalyDetector:
    """
    Anomaly detection model for Kubernetes security monitoring.
    Uses autoencoder architecture to detect anomalies in telemetry data.
    """
    
    def __init__(self, config=None):
        """
        Initialize the anomaly detector with configuration.
        
        Args:
            config (dict): Configuration parameters for the model
        """
        self.config = config or {
            'encoding_dim': 32,
            'hidden_dim': 64,
            'dropout_rate': 0.2,
            'learning_rate': 0.001,
            'threshold_percentile': 95  # Percentile for anomaly threshold
        }
        self.model = None
        self.threshold = None
        self.input_dim = None
        
    def build_model(self, input_dim):
        """
        Build the autoencoder model for anomaly detection.
        
        Args:
            input_dim (int): Dimension of input features
            
        Returns:
            Model: Compiled Keras model
        """
        self.input_dim = input_dim
        
        # Input layer
        input_layer = Input(shape=(input_dim,))
        
        # Encoder
        encoded = Dense(self.config['hidden_dim'], activation='relu')(input_layer)
        encoded = BatchNormalization()(encoded)
        encoded = Dropout(self.config['dropout_rate'])(encoded)
        encoded = Dense(self.config['encoding_dim'], activation='relu')(encoded)
        
        # Decoder
        decoded = Dense(self.config['hidden_dim'], activation='relu')(encoded)
        decoded = BatchNormalization()(decoded)
        decoded = Dropout(self.config['dropout_rate'])(decoded)
        decoded = Dense(input_dim, activation='sigmoid')(decoded)
        
        # Autoencoder model
        autoencoder = Model(input_layer, decoded)
        
        # Compile model
        autoencoder.compile(
            optimizer=tf.keras.optimizers.Adam(learning_rate=self.config['learning_rate']),
            loss='mean_squared_error'
        )
        
        self.model = autoencoder
        return autoencoder
    
    def train(self, X_train, epochs=50, batch_size=32, validation_split=0.2, save_dir=None):
        """
        Train the anomaly detection model.
        
        Args:
            X_train (np.ndarray): Training data
            epochs (int): Number of training epochs
            batch_size (int): Batch size for training
            validation_split (float): Fraction of data to use for validation
            save_dir (str): Directory to save model checkpoints
            
        Returns:
            History: Training history
        """
        if self.model is None:
            self.build_model(X_train.shape[1])
        
        callbacks = [
            EarlyStopping(monitor='val_loss', patience=5, restore_best_weights=True)
        ]
        
        if save_dir:
            os.makedirs(save_dir, exist_ok=True)
            callbacks.append(
                ModelCheckpoint(
                    os.path.join(save_dir, 'autoencoder_checkpoint.h5'),
                    monitor='val_loss',
                    save_best_only=True
                )
            )
        
        history = self.model.fit(
            X_train, X_train,
            epochs=epochs,
            batch_size=batch_size,
            validation_split=validation_split,
            callbacks=callbacks,
            verbose=1
        )
        
        # Calculate reconstruction error on training data
        reconstructions = self.model.predict(X_train)
        mse = np.mean(np.power(X_train - reconstructions, 2), axis=1)
        
        # Set threshold based on percentile of reconstruction errors
        self.threshold = np.percentile(mse, self.config['threshold_percentile'])
        
        return history
    
    def predict(self, X):
        """
        Predict anomalies in data.
        
        Args:
            X (np.ndarray): Input data
            
        Returns:
            tuple: (anomaly_scores, is_anomaly)
        """
        if self.model is None:
            raise ValueError("Model has not been trained yet")
        
        reconstructions = self.model.predict(X)
        mse = np.mean(np.power(X - reconstructions, 2), axis=1)
        
        # Flag as anomaly if reconstruction error > threshold
        is_anomaly = mse > self.threshold
        
        return mse, is_anomaly
    
    def save(self, directory):
        """Save model to directory"""
        if self.model is None:
            raise ValueError("No model to save")
        
        os.makedirs(directory, exist_ok=True)
        
        # Save model architecture and weights
        self.model.save(os.path.join(directory, 'anomaly_detector.h5'))
        
        # Save threshold and config
        np.save(os.path.join(directory, 'threshold.npy'), self.threshold)
        
        # Save config
        with open(os.path.join(directory, 'config.json'), 'w') as f:
            import json
            json.dump({
                'config': self.config,
                'input_dim': self.input_dim
            }, f)
    
    @classmethod
    def load(cls, directory):
        """Load model from directory"""
        import json
        
        # Load config
        with open(os.path.join(directory, 'config.json'), 'r') as f:
            saved_data = json.load(f)
        
        # Create instance with saved config
        instance = cls(config=saved_data['config'])
        instance.input_dim = saved_data['input_dim']
        
        # Load model
        instance.model = tf.keras.models.load_model(os.path.join(directory, 'anomaly_detector.h5'))
        
        # Load threshold
        instance.threshold = np.load(os.path.join(directory, 'threshold.npy'))
        
        return instance


class LSTMAnomalyDetector:
    """
    LSTM-based anomaly detection model for sequence data in Kubernetes security monitoring.
    Detects anomalies in time series of events.
    """
    
    def __init__(self, config=None):
        """
        Initialize the LSTM anomaly detector with configuration.
        
        Args:
            config (dict): Configuration parameters for the model
        """
        self.config = config or {
            'lstm_units': 64,
            'dense_units': 32,
            'dropout_rate': 0.2,
            'learning_rate': 0.001,
            'threshold_percentile': 95  # Percentile for anomaly threshold
        }
        self.model = None
        self.threshold = None
        self.input_dim = None
        self.sequence_length = None
        
    def build_model(self, sequence_length, input_dim):
        """
        Build the LSTM autoencoder model for anomaly detection.
        
        Args:
            sequence_length (int): Length of input sequences
            input_dim (int): Dimension of input features at each time step
            
        Returns:
            Model: Compiled Keras model
        """
        self.input_dim = input_dim
        self.sequence_length = sequence_length
        
        # Input layer
        input_layer = Input(shape=(sequence_length, input_dim))
        
        # Encoder
        encoded = LSTM(self.config['lstm_units'], return_sequences=True)(input_layer)
        encoded = Dropout(self.config['dropout_rate'])(encoded)
        encoded = LSTM(self.config['dense_units'])(encoded)
        
        # Decoder (RepeatVector and LSTM to reconstruct sequence)
        decoded = tf.keras.layers.RepeatVector(sequence_length)(encoded)
        decoded = LSTM(self.config['dense_units'], return_sequences=True)(decoded)
        decoded = Dropout(self.config['dropout_rate'])(decoded)
        decoded = LSTM(self.config['lstm_units'], return_sequences=True)(decoded)
        decoded = Dense(input_dim)(decoded)
        
        # LSTM autoencoder model
        autoencoder = Model(input_layer, decoded)
        
        # Compile model
        autoencoder.compile(
            optimizer=tf.keras.optimizers.Adam(learning_rate=self.config['learning_rate']),
            loss='mean_squared_error'
        )
        
        self.model = autoencoder
        return autoencoder
    
    def train(self, X_train, epochs=50, batch_size=32, validation_split=0.2, save_dir=None):
        """
        Train the LSTM anomaly detection model.
        
        Args:
            X_train (np.ndarray): Training data with shape (samples, sequence_length, features)
            epochs (int): Number of training epochs
            batch_size (int): Batch size for training
            validation_split (float): Fraction of data to use for validation
            save_dir (str): Directory to save model checkpoints
            
        Returns:
            History: Training history
        """
        if self.model is None:
            self.build_model(X_train.shape[1], X_train.shape[2])
        
        callbacks = [
            EarlyStopping(monitor='val_loss', patience=5, restore_best_weights=True)
        ]
        
        if save_dir:
            os.makedirs(save_dir, exist_ok=True)
            callbacks.append(
                ModelCheckpoint(
                    os.path.join(save_dir, 'lstm_autoencoder_checkpoint.h5'),
                    monitor='val_loss',
                    save_best_only=True
                )
            )
        
        history = self.model.fit(
            X_train, X_train,
            epochs=epochs,
            batch_size=batch_size,
            validation_split=validation_split,
            callbacks=callbacks,
            verbose=1
        )
        
        # Calculate reconstruction error on training data
        reconstructions = self.model.predict(X_train)
        mse = np.mean(np.power(X_train - reconstructions, 2), axis=(1, 2))
        
        # Set threshold based on percentile of reconstruction errors
        self.threshold = np.percentile(mse, self.config['threshold_percentile'])
        
        return history
    
    def predict(self, X):
        """
        Predict anomalies in sequence data.
        
        Args:
            X (np.ndarray): Input data with shape (samples, sequence_length, features)
            
        Returns:
            tuple: (anomaly_scores, is_anomaly)
        """
        if self.model is None:
            raise ValueError("Model has not been trained yet")
        
        reconstructions = self.model.predict(X)
        mse = np.mean(np.power(X - reconstructions, 2), axis=(1, 2))
        
        # Flag as anomaly if reconstruction error > threshold
        is_anomaly = mse > self.threshold
        
        return mse, is_anomaly
    
    def save(self, directory):
        """Save model to directory"""
        if self.model is None:
            raise ValueError("No model to save")
        
        os.makedirs(directory, exist_ok=True)
        
        # Save model architecture and weights
        self.model.save(os.path.join(directory, 'lstm_anomaly_detector.h5'))
        
        # Save threshold and config
        np.save(os.path.join(directory, 'threshold.npy'), self.threshold)
        
        # Save config
        with open(os.path.join(directory, 'config.json'), 'w') as f:
            import json
            json.dump({
                'config': self.config,
                'input_dim': self.input_dim,
                'sequence_length': self.sequence_length
            }, f)
    
    @classmethod
    def load(cls, directory):
        """Load model from directory"""
        import json
        
        # Load config
        with open(os.path.join(directory, 'config.json'), 'r') as f:
            saved_data = json.load(f)
        
        # Create instance with saved config
        instance = cls(config=saved_data['config'])
        instance.input_dim = saved_data['input_dim']
        instance.sequence_length = saved_data['sequence_length']
        
        # Load model
        instance.model = tf.keras.models.load_model(os.path.join(directory, 'lstm_anomaly_detector.h5'))
        
        # Load threshold
        instance.threshold = np.load(os.path.join(directory, 'threshold.npy'))
        
        return instance
