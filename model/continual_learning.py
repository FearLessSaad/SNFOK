import os
import numpy as np
import tensorflow as tf
from tensorflow.keras.models import Model, clone_model
from tensorflow.keras.optimizers import Adam
import json
import time
from datetime import datetime

class ContinualLearningFramework:
    """
    Framework for continuous learning and model adaptation in Kubernetes security monitoring.
    Enables models to evolve over time as new data becomes available and threats evolve.
    """
    
    def __init__(self, config=None):
        """
        Initialize the continual learning framework with configuration.
        
        Args:
            config (dict): Configuration parameters for the framework
        """
        self.config = config or {
            'update_frequency': 24,  # Hours between model updates
            'min_samples_for_update': 100,  # Minimum new samples required for update
            'max_history_models': 5,  # Maximum number of historical models to keep
            'performance_threshold': 0.05,  # Maximum allowed performance degradation
            'learning_rate': 0.0005,  # Lower learning rate for incremental updates
            'batch_size': 32,
            'epochs': 10
        }
        
        self.current_model = None
        self.model_history = []
        self.last_update_time = None
        self.new_samples_buffer = {
            'features': [],
            'labels': [],
            'timestamps': []
        }
        self.performance_metrics = []
    
    def set_base_model(self, model):
        """
        Set the initial base model for continual learning.
        
        Args:
            model: Trained model to start with
        """
        self.current_model = model
        self.last_update_time = datetime.now()
        
        # Save initial model as first history entry
        self.model_history.append({
            'model': clone_model(model.model) if hasattr(model, 'model') else clone_model(model),
            'timestamp': self.last_update_time,
            'version': 1,
            'metrics': None
        })
    
    def add_samples(self, features, labels=None, timestamps=None):
        """
        Add new samples to the buffer for future model updates.
        
        Args:
            features (np.ndarray): Feature vectors of new samples
            labels (np.ndarray, optional): Labels for supervised learning
            timestamps (list, optional): Timestamps for each sample
        """
        # Convert features to list if it's a single sample
        if isinstance(features, np.ndarray) and len(features.shape) == 1:
            features = [features]
            if labels is not None and not isinstance(labels, list):
                labels = [labels]
            if timestamps is not None and not isinstance(timestamps, list):
                timestamps = [timestamps]
        
        # Add to buffer
        self.new_samples_buffer['features'].extend(features)
        
        if labels is not None:
            self.new_samples_buffer['labels'].extend(labels)
        
        if timestamps is not None:
            self.new_samples_buffer['timestamps'].extend(timestamps)
        else:
            # Use current time if timestamps not provided
            current_time = datetime.now().timestamp()
            self.new_samples_buffer['timestamps'].extend([current_time] * len(features))
    
    def should_update_model(self):
        """
        Determine if the model should be updated based on time and data criteria.
        
        Returns:
            bool: True if model should be updated
        """
        # Check if enough time has passed
        hours_since_update = (datetime.now() - self.last_update_time).total_seconds() / 3600
        time_condition = hours_since_update >= self.config['update_frequency']
        
        # Check if enough new samples are available
        data_condition = len(self.new_samples_buffer['features']) >= self.config['min_samples_for_update']
        
        return time_condition and data_condition
    
    def update_model(self, validation_data=None):
        """
        Update the model with new data using incremental learning.
        
        Args:
            validation_data (tuple, optional): Data for validating the updated model
            
        Returns:
            dict: Update results including performance metrics
        """
        if not self.should_update_model():
            return {"status": "skipped", "reason": "Update criteria not met"}
        
        if self.current_model is None:
            return {"status": "error", "reason": "No base model set"}
        
        # Prepare data for update
        X_new = np.array(self.new_samples_buffer['features'])
        
        # Handle supervised vs unsupervised updates
        has_labels = len(self.new_samples_buffer['labels']) == len(X_new)
        
        # Clone current model for update
        if hasattr(self.current_model, 'model'):
            # Handle custom model classes that have a .model attribute
            new_model = clone_model(self.current_model.model)
            new_model.compile(
                optimizer=Adam(learning_rate=self.config['learning_rate']),
                loss=self.current_model.model.loss,
                metrics=self.current_model.model.metrics
            )
        else:
            # Handle direct Keras models
            new_model = clone_model(self.current_model)
            new_model.compile(
                optimizer=Adam(learning_rate=self.config['learning_rate']),
                loss=self.current_model.loss,
                metrics=self.current_model.metrics
            )
        
        # Perform update based on model type
        if has_labels:
            # Supervised learning update
            y_new = np.array(self.new_samples_buffer['labels'])
            
            # Train on new data
            history = new_model.fit(
                X_new, y_new,
                epochs=self.config['epochs'],
                batch_size=self.config['batch_size'],
                validation_data=validation_data,
                verbose=0
            )
            
            update_metrics = {
                'final_loss': history.history['loss'][-1],
                'final_val_loss': history.history['val_loss'][-1] if validation_data else None
            }
        else:
            # Unsupervised learning update (e.g., for autoencoders)
            # For autoencoders, X is both input and output
            history = new_model.fit(
                X_new, X_new,
                epochs=self.config['epochs'],
                batch_size=self.config['batch_size'],
                validation_data=validation_data,
                verbose=0
            )
            
            update_metrics = {
                'final_loss': history.history['loss'][-1],
                'final_val_loss': history.history['val_loss'][-1] if validation_data else None
            }
        
        # Evaluate model performance change
        if validation_data:
            old_metrics = self.current_model.evaluate(*validation_data, verbose=0)
            new_metrics = new_model.evaluate(*validation_data, verbose=0)
            
            # Check if performance degraded significantly
            if isinstance(old_metrics, list):
                performance_change = (old_metrics[0] - new_metrics[0]) / old_metrics[0]
            else:
                performance_change = (old_metrics - new_metrics) / old_metrics
                
            update_metrics['performance_change'] = performance_change
            
            # Reject update if performance degraded beyond threshold
            if performance_change < -self.config['performance_threshold']:
                return {
                    "status": "rejected",
                    "reason": f"Performance degraded by {-performance_change:.2%}",
                    "metrics": update_metrics
                }
        
        # Update successful, apply changes
        if hasattr(self.current_model, 'model'):
            # Update model weights for custom model classes
            self.current_model.model = new_model
        else:
            # Update direct Keras model
            self.current_model = new_model
        
        # Record update in history
        self.model_history.append({
            'model': clone_model(new_model),
            'timestamp': datetime.now(),
            'version': len(self.model_history) + 1,
            'metrics': update_metrics
        })
        
        # Limit history size
        if len(self.model_history) > self.config['max_history_models']:
            # Remove oldest model but keep the first (original) model
            self.model_history.pop(1)
        
        # Update timestamp and clear buffer
        self.last_update_time = datetime.now()
        self.performance_metrics.append(update_metrics)
        
        # Clear buffer
        self.new_samples_buffer = {
            'features': [],
            'labels': [],
            'timestamps': []
        }
        
        return {
            "status": "success",
            "version": len(self.model_history),
            "timestamp": self.last_update_time.isoformat(),
            "metrics": update_metrics
        }
    
    def rollback_to_version(self, version):
        """
        Rollback to a previous model version.
        
        Args:
            version (int): Version number to rollback to
            
        Returns:
            bool: Success status
        """
        if version < 1 or version > len(self.model_history):
            return False
        
        # Get historical model
        historical_model = self.model_history[version - 1]
        
        # Apply rollback
        if hasattr(self.current_model, 'model'):
            self.current_model.model = clone_model(historical_model['model'])
            self.current_model.model.set_weights(historical_model['model'].get_weights())
        else:
            self.current_model = clone_model(historical_model['model'])
            self.current_model.set_weights(historical_model['model'].get_weights())
        
        # Add rollback event to history
        self.model_history.append({
            'model': clone_model(historical_model['model']),
            'timestamp': datetime.now(),
            'version': len(self.model_history) + 1,
            'metrics': historical_model['metrics'],
            'rollback_from': version
        })
        
        return True
    
    def get_performance_trend(self):
        """
        Get the performance trend over time.
        
        Returns:
            dict: Performance metrics over time
        """
        if not self.performance_metrics:
            return {"status": "no_data"}
        
        # Extract metrics over time
        timestamps = [m['timestamp'].isoformat() for m in self.model_history[1:] if 'metrics' in m and m['metrics']]
        losses = [m['metrics']['final_loss'] for m in self.model_history[1:] if 'metrics' in m and m['metrics']]
        val_losses = [m['metrics'].get('final_val_loss') for m in self.model_history[1:] 
                     if 'metrics' in m and m['metrics'] and 'final_val_loss' in m['metrics']]
        
        return {
            "timestamps": timestamps,
            "losses": losses,
            "validation_losses": val_losses if any(v is not None for v in val_losses) else None
        }
    
    def save(self, directory):
        """Save framework state to directory"""
        os.makedirs(directory, exist_ok=True)
        
        # Save current model
        if hasattr(self.current_model, 'save'):
            self.current_model.save(os.path.join(directory, 'current_model'))
        elif hasattr(self.current_model, 'model') and hasattr(self.current_model.model, 'save'):
            self.current_model.model.save(os.path.join(directory, 'current_model'))
        
        # Save config and metadata
        metadata = {
            'config': self.config,
            'last_update_time': self.last_update_time.isoformat() if self.last_update_time else None,
            'performance_metrics': self.performance_metrics,
            'model_history': [
                {
                    'version': h['version'],
                    'timestamp': h['timestamp'].isoformat(),
                    'metrics': h['metrics'],
                    'rollback_from': h.get('rollback_from')
                }
                for h in self.model_history
            ]
        }
        
        with open(os.path.join(directory, 'metadata.json'), 'w') as f:
            json.dump(metadata, f, indent=2)
        
        # Save buffer statistics
        buffer_stats = {
            'sample_count': len(self.new_samples_buffer['features']),
            'has_labels': len(self.new_samples_buffer['labels']) > 0,
            'timestamp_range': [
                min(self.new_samples_buffer['timestamps']) if self.new_samples_buffer['timestamps'] else None,
                max(self.new_samples_buffer['timestamps']) if self.new_samples_buffer['timestamps'] else None
            ]
        }
        
        with open(os.path.join(directory, 'buffer_stats.json'), 'w') as f:
            json.dump(buffer_stats, f, indent=2)
    
    @classmethod
    def load(cls, directory, model_class):
        """
        Load framework state from directory.
        
        Args:
            directory (str): Directory containing saved state
            model_class: Class of the model to load
            
        Returns:
            ContinualLearningFramework: Loaded framework
        """
        # Load metadata
        with open(os.path.join(directory, 'metadata.json'), 'r') as f:
            metadata = json.load(f)
        
        # Create instance with saved config
        instance = cls(config=metadata['config'])
        
        # Load current model
        if os.path.exists(os.path.join(directory, 'current_model')):
            if hasattr(model_class, 'load'):
                instance.current_model = model_class.load(os.path.join(directory, 'current_model'))
            else:
                instance.current_model = tf.keras.models.load_model(os.path.join(directory, 'current_model'))
        
        # Restore metadata
        instance.last_update_time = datetime.fromisoformat(metadata['last_update_time']) if metadata['last_update_time'] else None
        instance.performance_metrics = metadata['performance_metrics']
        
        # Note: We don't restore the full model history with actual models,
        # just the metadata about each version
        instance.model_history = [
            {
                'model': None,  # Actual model objects aren't restored
                'version': h['version'],
                'timestamp': datetime.fromisoformat(h['timestamp']),
                'metrics': h['metrics'],
                'rollback_from': h.get('rollback_from')
            }
            for h in metadata['model_history']
        ]
        
        return instance


class FeedbackProcessor:
    """
    Processes user feedback on model predictions to improve model accuracy.
    Handles false positives and false negatives for continuous refinement.
    """
    
    def __init__(self, config=None):
        """
        Initialize the feedback processor with configuration.
        
        Args:
            config (dict): Configuration parameters
        """
        self.config = config or {
            'feedback_weight': 1.5,  # Weight multiplier for feedback samples (not currently used)
            'min_feedback_for_update': 10,  # Minimum feedback items before triggering update
            'max_buffer_size': 1000  # Maximum size of feedback buffer
        }
        
        self.feedback_buffer = {
            'false_positives': {'features': [], 'timestamps': []},
            'false_negatives': {'features': [], 'timestamps': []}
        }
        
        self.continual_learning = None
    
    def set_continual_learning_framework(self, framework):
        """
        Set the continual learning framework for model updates.
        
        Args:
            framework (ContinualLearningFramework): The continual learning instance
        """
        self.continual_learning = framework
    
    def add_false_positive(self, features, timestamp=None):
        """
        Add a false positive sample to the feedback buffer.
        False positives are instances predicted as threats (1) but are actually normal (0).
        
        Args:
            features (np.ndarray): Feature vector of the sample
            timestamp (float, optional): Timestamp of the feedback; defaults to current time
        """
        if timestamp is None:
            timestamp = datetime.now().timestamp()
        self.feedback_buffer['false_positives']['features'].append(features)
        self.feedback_buffer['false_positives']['timestamps'].append(timestamp)
        
        # Limit buffer size by removing oldest entry if necessary
        if len(self.feedback_buffer['false_positives']['features']) > self.config['max_buffer_size']:
            self.feedback_buffer['false_positives']['features'].pop(0)
            self.feedback_buffer['false_positives']['timestamps'].pop(0)
    
    def add_false_negative(self, features, timestamp=None):
        """
        Add a false negative sample to the feedback buffer.
        False negatives are instances predicted as normal (0) but are actually threats (1).
        
        Args:
            features (np.ndarray): Feature vector of the sample
            timestamp (float, optional): Timestamp of the feedback; defaults to current time
        """
        if timestamp is None:
            timestamp = datetime.now().timestamp()
        self.feedback_buffer['false_negatives']['features'].append(features)
        self.feedback_buffer['false_negatives']['timestamps'].append(timestamp)
        
        # Limit buffer size by removing oldest entry if necessary
        if len(self.feedback_buffer['false_negatives']['features']) > self.config['max_buffer_size']:
            self.feedback_buffer['false_negatives']['features'].pop(0)
            self.feedback_buffer['false_negatives']['timestamps'].pop(0)
    
    def process_feedback(self):
        """
        Process accumulated feedback and potentially trigger a model update.
        Adds false positives with label 0 (normal) and false negatives with label 1 (threat)
        to the ContinualLearningFramework's buffer.
        
        Returns:
            dict: Status and message indicating the result of processing
        """
        fp_features = self.feedback_buffer['false_positives']['features']
        fp_timestamps = self.feedback_buffer['false_positives']['timestamps']
        fn_features = self.feedback_buffer['false_negatives']['features']
        fn_timestamps = self.feedback_buffer['false_negatives']['timestamps']
        
        fp_count = len(fp_features)
        fn_count = len(fn_features)
        
        # Check if enough feedback has been accumulated
        if fp_count >= self.config['min_feedback_for_update'] or fn_count >= self.config['min_feedback_for_update']:
            if not self.continual_learning:
                return {"status": "error", "message": "Continual learning framework not set"}
            
            # Add false positives with label 0 (normal)
            if fp_count > 0:
                fp_labels = [0] * fp_count
                self.continual_learning.add_samples(fp_features, labels=fp_labels, timestamps=fp_timestamps)
            
            # Add false negatives with label 1 (threat)
            if fn_count > 0:
                fn_labels = [1] * fn_count
                self.continual_learning.add_samples(fn_features, labels=fn_labels, timestamps=fn_timestamps)
            
            # Clear feedback buffer after processing
            self.feedback_buffer['false_positives']['features'] = []
            self.feedback_buffer['false_positives']['timestamps'] = []
            self.feedback_buffer['false_negatives']['features'] = []
            self.feedback_buffer['false_negatives']['timestamps'] = []
            
            # Check if model update should be triggered
            if self.continual_learning.should_update_model():
                result = self.continual_learning.update_model()
                return result
            else:
                return {"status": "feedback_added", "message": "Feedback added, but update criteria not met"}
        else:
            return {"status": "no_action", "message": "Not enough feedback for update"}
    
    def save(self, directory):
        """
        Save feedback processor state to a directory.
        
        Args:
            directory (str): Directory to save the state
        """
        os.makedirs(directory, exist_ok=True)
        
        # Prepare state dictionary, converting numpy arrays to lists for JSON
        state = {
            'config': self.config,
            'feedback_buffer': {
                'false_positives': {
                    'features': [f.tolist() for f in self.feedback_buffer['false_positives']['features']],
                    'timestamps': self.feedback_buffer['false_positives']['timestamps']
                },
                'false_negatives': {
                    'features': [f.tolist() for f in self.feedback_buffer['false_negatives']['features']],
                    'timestamps': self.feedback_buffer['false_negatives']['timestamps']
                }
            }
        }
        
        # Save to JSON file
        with open(os.path.join(directory, 'feedback_processor.json'), 'w') as f:
            json.dump(state, f, indent=2)
    
    @classmethod
    def load(cls, directory):
        """
        Load feedback processor state from a directory.
        
        Args:
            directory (str): Directory containing the saved state
            
        Returns:
            FeedbackProcessor: Loaded instance
        """
        with open(os.path.join(directory, 'feedback_processor.json'), 'r') as f:
            state = json.load(f)
        
        # Create instance with loaded config
        instance = cls(config=state['config'])
        
        # Restore feedback buffer, converting lists back to numpy arrays
        instance.feedback_buffer['false_positives']['features'] = [
            np.array(f) for f in state['feedback_buffer']['false_positives']['features']
        ]
        instance.feedback_buffer['false_positives']['timestamps'] = state['feedback_buffer']['false_positives']['timestamps']
        instance.feedback_buffer['false_negatives']['features'] = [
            np.array(f) for f in state['feedback_buffer']['false_negatives']['features']
        ]
        instance.feedback_buffer['false_negatives']['timestamps'] = state['feedback_buffer']['false_negatives']['timestamps']
        
        return instance