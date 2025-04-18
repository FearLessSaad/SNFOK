import os
import numpy as np
import tensorflow as tf
from tensorflow.keras.models import Model
from tensorflow.keras.layers import Input, Dense, LSTM, Dropout, BatchNormalization, Concatenate
import json

class ThreatDetectionEngine:
    """
    Main engine that combines anomaly detection and threat classification
    for comprehensive Kubernetes security monitoring.
    """
    
    def __init__(self, config=None):
        """
        Initialize the threat detection engine with configuration.
        
        Args:
            config (dict): Configuration parameters for the engine
        """
        self.config = config or {
            'anomaly_threshold': 0.8,  # Confidence threshold for anomaly detection
            'classification_threshold': 0.7,  # Confidence threshold for classification
            'severity_threshold': 0.5,  # Threshold for high severity
            'use_continual_learning': True,
            'feedback_buffer_size': 1000
        }
        
        # Component models
        self.preprocessor = None
        self.anomaly_detector = None
        self.sequence_detector = None
        self.threat_classifier = None
        self.severity_estimator = None
        self.continual_learning = None
        self.feedback_processor = None
        
        # Detection state
        self.detection_count = 0
        self.false_positive_count = 0
        self.detection_history = []
        self.classifier_feedback_buffer = []  # Added to store classifier feedback
    
    def load_models(self, model_dir):
        """
        Load all component models from directory.
        
        Args:
            model_dir (str): Base directory containing model subdirectories
            
        Returns:
            bool: Success status
        """
        try:
            # Import required modules
            from preprocessing.data_preprocessor import DataPreprocessor
            from models.anomaly_detection import AnomalyDetector, LSTMAnomalyDetector
            from models.threat_classification import ThreatClassifier, ThreatSeverityEstimator
            from models.continual_learning import ContinualLearningFramework, FeedbackProcessor
            
            # Load preprocessor
            if os.path.exists(os.path.join(model_dir, 'preprocessor')):
                self.preprocessor = DataPreprocessor.load(os.path.join(model_dir, 'preprocessor'))
            
            # Load anomaly detector
            if os.path.exists(os.path.join(model_dir, 'anomaly_detector')):
                self.anomaly_detector = AnomalyDetector.load(os.path.join(model_dir, 'anomaly_detector'))
            
            # Load sequence detector
            if os.path.exists(os.path.join(model_dir, 'sequence_detector')):
                self.sequence_detector = LSTMAnomalyDetector.load(os.path.join(model_dir, 'sequence_detector'))
            
            # Load threat classifier
            if os.path.exists(os.path.join(model_dir, 'threat_classifier')):
                self.threat_classifier = ThreatClassifier.load(os.path.join(model_dir, 'threat_classifier'))
            
            # Load severity estimator
            if os.path.exists(os.path.join(model_dir, 'severity_estimator')):
                self.severity_estimator = ThreatSeverityEstimator.load(os.path.join(model_dir, 'severity_estimator'))
            
            # Load continual learning framework if enabled
            if self.config['use_continual_learning']:
                if os.path.exists(os.path.join(model_dir, 'continual_learning')):
                    self.continual_learning = ContinualLearningFramework.load(
                        os.path.join(model_dir, 'continual_learning'),
                        self.anomaly_detector.__class__
                    )
                else:
                    self.continual_learning = ContinualLearningFramework()
                    self.continual_learning.set_base_model(self.anomaly_detector)
                
                # Load feedback processor
                if os.path.exists(os.path.join(model_dir, 'feedback_processor')):
                    self.feedback_processor = FeedbackProcessor.load(
                        os.path.join(model_dir, 'feedback_processor')
                    )
                else:
                    self.feedback_processor = FeedbackProcessor({
                        'max_buffer_size': self.config['feedback_buffer_size']
                    })
                
                self.feedback_processor.set_continual_learning_framework(self.continual_learning)
            
            return True
        
        except Exception as e:
            print(f"Error loading models: {e}")
            return False
    
    def save_models(self, model_dir):
        """
        Save all component models to directory.
        
        Args:
            model_dir (str): Base directory to save models
            
        Returns:
            bool: Success status
        """
        try:
            os.makedirs(model_dir, exist_ok=True)
            
            # Save preprocessor
            if self.preprocessor:
                os.makedirs(os.path.join(model_dir, 'preprocessor'), exist_ok=True)
                self.preprocessor.save(os.path.join(model_dir, 'preprocessor'))
            
            # Save anomaly detector
            if self.anomaly_detector:
                os.makedirs(os.path.join(model_dir, 'anomaly_detector'), exist_ok=True)
                self.anomaly_detector.save(os.path.join(model_dir, 'anomaly_detector'))
            
            # Save sequence detector
            if self.sequence_detector:
                os.makedirs(os.path.join(model_dir, 'sequence_detector'), exist_ok=True)
                self.sequence_detector.save(os.path.join(model_dir, 'sequence_detector'))
            
            # Save threat classifier
            if self.threat_classifier:
                os.makedirs(os.path.join(model_dir, 'threat_classifier'), exist_ok=True)
                self.threat_classifier.save(os.path.join(model_dir, 'threat_classifier'))
            
            # Save severity estimator
            if self.severity_estimator:
                os.makedirs(os.path.join(model_dir, 'severity_estimator'), exist_ok=True)
                self.severity_estimator.save(os.path.join(model_dir, 'severity_estimator'))
            
            # Save continual learning framework
            if self.continual_learning:
                os.makedirs(os.path.join(model_dir, 'continual_learning'), exist_ok=True)
                self.continual_learning.save(os.path.join(model_dir, 'continual_learning'))
            
            # Save feedback processor
            if self.feedback_processor:
                os.makedirs(os.path.join(model_dir, 'feedback_processor'), exist_ok=True)
                self.feedback_processor.save(os.path.join(model_dir, 'feedback_processor'))
            
            # Save config
            with open(os.path.join(model_dir, 'config.json'), 'w') as f:
                json.dump(self.config, f, indent=2)
            
            return True
        
        except Exception as e:
            print(f"Error saving models: {e}")
            return False
    
    def process_event(self, event, event_type):
        """
        Process a single event for threat detection.
        
        Args:
            event (dict): Event data from eBPF collector
            event_type (str): Type of event ('syscall', 'network', 'process', 'file', 'container')
            
        Returns:
            dict: Detection result
        """
        if self.preprocessor is None or self.anomaly_detector is None:
            return {
                'status': 'error',
                'message': 'Models not loaded'
            }
        
        # Preprocess event
        features = self.preprocessor.preprocess_events([event], event_type)
        
        if features is None or len(features) == 0:
            return {
                'status': 'error',
                'message': 'Failed to extract features from event'
            }
        
        # Detect anomaly
        anomaly_scores, is_anomaly = self.anomaly_detector.predict(features.values)
        
        # Initialize result
        result = {
            'event_id': event.get('id', str(self.detection_count)),
            'timestamp': event.get('timestamp', 0),
            'event_type': event_type,
            'anomaly_score': float(anomaly_scores[0]),
            'is_anomaly': bool(is_anomaly[0]),
            'confidence': float(anomaly_scores[0] / self.anomaly_detector.threshold if self.anomaly_detector.threshold > 0 else 0)
        }
        
        # If anomaly detected, classify threat
        if is_anomaly[0]:
            self.detection_count += 1
            
            # Classify threat if classifier is available
            if self.threat_classifier:
                class_probs, pred_class, pred_class_name = self.threat_classifier.predict(features.values)
                
                # Add classification results
                result['threat_classification'] = {
                    'class_id': int(pred_class[0]),
                    'class_name': pred_class_name[0],
                    'confidence': float(class_probs[0][pred_class[0]]),
                    'all_probabilities': {
                        self.threat_classifier.class_names[i]: float(class_probs[0][i])
                        for i in range(len(self.threat_classifier.class_names))
                    }
                }
            
            # Estimate severity if estimator is available
            if self.severity_estimator:
                severity_score = self.severity_estimator.predict(features.values)
                
                # Add severity results
                result['severity'] = {
                    'score': float(severity_score[0]),
                    'level': 'high' if severity_score[0] > self.config['severity_threshold'] else 'medium' if severity_score[0] > self.config['severity_threshold'] / 2 else 'low'
                }
            
            # Add to detection history
            self.detection_history.append({
                'event_id': result['event_id'],
                'timestamp': result['timestamp'],
                'event_type': event_type,
                'anomaly_score': result['anomaly_score'],
                'features': features.values[0].tolist(),
                'classification': result.get('threat_classification', {}).get('class_name', 'unknown'),
                'severity': result.get('severity', {}).get('level', 'unknown')
            })
            
            # Limit history size
            if len(self.detection_history) > 1000:
                self.detection_history.pop(0)
        
        return result
    
    def process_sequence(self, events, event_type, sequence_length=10):
        """
        Process a sequence of events for threat detection.
        
        Args:
            events (list): List of event data from eBPF collector
            event_type (str): Type of events ('syscall', 'network', 'process', 'file', 'container')
            sequence_length (int): Length of sequence to analyze
            
        Returns:
            dict: Detection result
        """
        if self.preprocessor is None or self.sequence_detector is None:
            return {
                'status': 'error',
                'message': 'Models not loaded'
            }
        
        if len(events) < sequence_length:
            return {
                'status': 'error',
                'message': f'Sequence too short, need at least {sequence_length} events'
            }
        
        # Preprocess events
        all_features = self.preprocessor.preprocess_events(events, event_type)
        
        if all_features is None or len(all_features) == 0:
            return {
                'status': 'error',
                'message': 'Failed to extract features from events'
            }
        
        # Create sequences
        sequences = []
        for i in range(len(all_features) - sequence_length + 1):
            seq = all_features.values[i:i+sequence_length]
            sequences.append(seq)
        
        sequences = np.array(sequences)
        
        # Detect anomalies in sequences
        anomaly_scores, is_anomaly = self.sequence_detector.predict(sequences)
        
        # Find the most anomalous sequence
        max_score_idx = np.argmax(anomaly_scores)
        max_score = anomaly_scores[max_score_idx]
        is_max_anomaly = is_anomaly[max_score_idx]
        
        # Get the corresponding events
        anomalous_events_idx = list(range(max_score_idx, max_score_idx + sequence_length))
        anomalous_events = [events[i] for i in anomalous_events_idx]
        
        # Initialize result
        result = {
            'sequence_id': str(self.detection_count),
            'timestamp_start': anomalous_events[0].get('timestamp', 0),
            'timestamp_end': anomalous_events[-1].get('timestamp', 0),
            'event_type': event_type,
            'sequence_length': sequence_length,
            'anomaly_score': float(max_score),
            'is_anomaly': bool(is_max_anomaly),
            'confidence': float(max_score / self.sequence_detector.threshold if self.sequence_detector.threshold > 0 else 0),
            'anomalous_events_indices': anomalous_events_idx
        }
        
        # If anomaly detected, add to detection count
        if is_max_anomaly:
            self.detection_count += 1
            
            # Add to detection history
            self.detection_history.append({
                'sequence_id': result['sequence_id'],
                'timestamp_start': result['timestamp_start'],
                'timestamp_end': result['timestamp_end'],
                'event_type': event_type,
                'anomaly_score': result['anomaly_score'],
                'is_sequence': True,
                'sequence_length': sequence_length
            })
            
            # Limit history size
            if len(self.detection_history) > 1000:
                self.detection_history.pop(0)
        
        return result
    
    def record_feedback(self, detection_id, is_false_positive, actual_class=None):
        """
        Record feedback for a detection to improve model accuracy.
        
        Args:
            detection_id (str): ID of the detection to provide feedback for
            is_false_positive (bool): Whether the detection was a false positive
            actual_class (str, optional): Actual threat class if known
            
        Returns:
            dict: Feedback processing result
        """
        if not self.config['use_continual_learning'] or self.feedback_processor is None:
            return {
                'status': 'error',
                'message': 'Continual learning not enabled'
            }
        
        # Find detection in history
        detection = next((d for d in self.detection_history if d.get('event_id') == detection_id or d.get('sequence_id') == detection_id), None)
        
        if detection is None:
            return {
                'status': 'error',
                'message': 'Detection not found in history'
            }
        
        if is_false_positive:
            # Record false positive for anomaly detector
            self.false_positive_count += 1
            features = detection.get('features')
            if features:
                self.feedback_processor.add_false_positive(
                    np.array(features),
                    detection.get('classification'),
                    detection.get('timestamp', 0)
                )
                return {
                    'status': 'success',
                    'message': 'False positive feedback recorded'
                }
            else:
                return {
                    'status': 'error',
                    'message': 'No features found for feedback'
                }
        elif actual_class is not None:
            # Record misclassification feedback for classifier
            if 'classification' not in detection:
                return {
                    'status': 'error',
                    'message': 'Cannot provide class feedback for sequence detections'
                }
            predicted_class = detection['classification']
            features = detection.get('features')
            if features:
                self.classifier_feedback_buffer.append({
                    'features': features,
                    'predicted_class': predicted_class,
                    'actual_class': actual_class,
                    'timestamp': detection['timestamp']
                })
                if len(self.classifier_feedback_buffer) > self.config['feedback_buffer_size']:
                    self.classifier_feedback_buffer.pop(0)
                return {
                    'status': 'success',
                    'message': 'Misclassification feedback recorded'
                }
            else:
                return {
                    'status': 'error',
                    'message': 'No features found for feedback'
                }
        else:
            return {
                'status': 'error',
                'message': 'Invalid feedback parameters'
            }