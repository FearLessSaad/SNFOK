import os
import numpy as np
import tensorflow as tf
from tensorflow.keras.models import Model
from tensorflow.keras.layers import Input, Dense, Embedding, Conv1D, MaxPooling1D, GlobalMaxPooling1D, Flatten
from tensorflow.keras.callbacks import EarlyStopping, ModelCheckpoint
from sklearn.metrics import classification_report, confusion_matrix

class ThreatClassifier:
    """
    Multi-class classifier for identifying specific threat types in Kubernetes environments.
    Uses a deep convolutional neural network to classify sequences into threat categories.
    """
    
    def __init__(self):
        """
        Initialize the threat classifier.
        """
        self.model = None
        self.vocabulary_size = None
        self.sequence_length = None
        self.num_classes = None
        self.class_names = None
        
    def build_model(self, vocabulary_size, sequence_length, num_classes, class_names=None):
        """
        Build the classification model with the specified CNN architecture.
        
        Args:
            vocabulary_size (int): Size of the vocabulary for the Embedding layer (e.g., len(word_indexes) + 1)
            sequence_length (int): Length of input sequences (e.g., 805)
            num_classes (int): Number of threat classes to predict (e.g., 5)
            class_names (list): Names of the threat classes
            
        Returns:
            Model: Compiled Keras model
        """
        self.vocabulary_size = vocabulary_size
        self.sequence_length = sequence_length
        self.num_classes = num_classes
        self.class_names = class_names or [f"class_{i}" for i in range(num_classes)]
        
        # Input layer for sequence data
        input_layer = Input(shape=(sequence_length,), dtype='int32')
        
        # Embedding layer to convert token indices to dense vectors
        x = Embedding(input_dim=vocabulary_size, output_dim=72)(input_layer)
        
        # Convolutional and pooling layers
        x = Conv1D(filters=128, kernel_size=7, strides=2, activation='relu')(x)
        x = MaxPooling1D(pool_size=2, strides=2)(x)
        x = Conv1D(filters=192, kernel_size=3, activation='relu')(x)
        x = MaxPooling1D(pool_size=2, strides=2)(x)
        x = Conv1D(filters=128, kernel_size=1, activation='relu')(x)
        x = Conv1D(filters=256, kernel_size=3, activation='relu')(x)
        x = Conv1D(filters=256, kernel_size=1, activation='relu')(x)
        x = MaxPooling1D(pool_size=2, strides=2)(x)
        x = Conv1D(filters=128, kernel_size=1, activation='relu')(x)
        x = Conv1D(filters=256, kernel_size=3, activation='relu')(x)
        x = Conv1D(filters=256, kernel_size=1, activation='relu')(x)
        x = MaxPooling1D(pool_size=2, strides=2)(x)
        x = Conv1D(filters=128, kernel_size=2, activation='relu')(x)
        x = Conv1D(filters=192, kernel_size=3, activation='relu')(x)
        x = MaxPooling1D(pool_size=2, strides=2)(x)
        x = Conv1D(filters=128, kernel_size=2, activation='relu')(x)
        x = Conv1D(filters=256, kernel_size=1, activation='relu')(x)
        x = MaxPooling1D(pool_size=2, strides=2)(x)
        x = Conv1D(filters=128, kernel_size=2, activation='relu')(x)
        x = Conv1D(filters=191, kernel_size=1, activation='relu')(x)
        x = MaxPooling1D(pool_size=2, strides=2)(x)
        
        # Global pooling and flattening
        x = GlobalMaxPooling1D()(x)
        x = Flatten()(x)
        
        # Dense layers
        x = Dense(512, activation='relu')(x)
        x = Dense(1024, activation='relu')(x)
        x = Dense(1024, activation='relu')(x)
        x = Dense(512, activation='relu')(x)
        
        # Output layer for multi-class classification
        output_layer = Dense(num_classes, activation='softmax')(x)
        
        # Create and compile the model
        model = Model(input_layer, output_layer)
        model.compile(
            optimizer=tf.keras.optimizers.legacy.Adam(),
            loss='categorical_crossentropy',
            metrics=['accuracy']
        )
        
        self.model = model
        return model
    
    def train(self, X_train, y_train, epochs=50, batch_size=32, validation_split=0.2, class_weights=None, save_dir=None):
        """
        Train the threat classification model.
        
        Args:
            X_train (np.ndarray): Training features (shape: [samples, sequence_length])
            y_train (np.ndarray): Training labels (one-hot encoded, shape: [samples, num_classes])
            epochs (int): Number of training epochs
            batch_size (int): Batch size for training
            validation_split (float): Fraction of data to use for validation
            class_weights (dict): Weights for each class to handle imbalance
            save_dir (str): Directory to save model checkpoints
            
        Returns:
            History: Training history
        """
        if self.model is None:
            raise ValueError("Model not built. Call build_model first with appropriate parameters.")
        
        callbacks = [
            EarlyStopping(monitor='val_loss', patience=5, restore_best_weights=True)
        ]
        
        if save_dir:
            os.makedirs(save_dir, exist_ok=True)
            callbacks.append(
                ModelCheckpoint(
                    os.path.join(save_dir, 'threat_classifier_checkpoint.h5'),
                    monitor='val_loss',
                    save_best_only=True
                )
            )
        
        history = self.model.fit(
            X_train, y_train,
            epochs=epochs,
            batch_size=batch_size,
            validation_split=validation_split,
            class_weight=class_weights,
            callbacks=callbacks,
            verbose=1
        )
        
        return history
    
    def predict(self, X):
        """
        Predict threat classes for input data.
        
        Args:
            X (np.ndarray): Input features (shape: [samples, sequence_length])
            
        Returns:
            tuple: (class_probabilities, predicted_classes, predicted_class_names)
        """
        if self.model is None:
            raise ValueError("Model has not been trained yet")
        
        class_probs = self.model.predict(X)
        predicted_classes = np.argmax(class_probs, axis=1)
        predicted_class_names = [self.class_names[idx] for idx in predicted_classes]
        
        return class_probs, predicted_classes, predicted_class_names
    
    def evaluate(self, X_test, y_test):
        """
        Evaluate the model on test data.
        
        Args:
            X_test (np.ndarray): Test features (shape: [samples, sequence_length])
            y_test (np.ndarray): Test labels (one-hot encoded, shape: [samples, num_classes])
            
        Returns:
            dict: Evaluation metrics
        """
        if self.model is None:
            raise ValueError("Model has not been trained yet")
        
        y_pred_probs = self.model.predict(X_test)
        y_pred = np.argmax(y_pred_probs, axis=1)
        y_true = np.argmax(y_test, axis=1)
        
        report = classification_report(y_true, y_pred, target_names=self.class_names, output_dict=True)
        conf_matrix = confusion_matrix(y_true, y_pred)
        
        class_metrics = {}
        for i, class_name in enumerate(self.class_names):
            class_metrics[class_name] = {
                'precision': report[class_name]['precision'],
                'recall': report[class_name]['recall'],
                'f1-score': report[class_name]['f1-score'],
                'support': report[class_name]['support']
            }
        
        metrics = {
            'accuracy': report['accuracy'],
            'macro_avg_precision': report['macro avg']['precision'],
            'macro_avg_recall': report['macro avg']['recall'],
            'macro_avg_f1': report['macro avg']['f1-score'],
            'weighted_avg_precision': report['weighted avg']['precision'],
            'weighted_avg_recall': report['weighted avg']['recall'],
            'weighted_avg_f1': report['weighted avg']['f1-score'],
            'class_metrics': class_metrics,
            'confusion_matrix': conf_matrix.tolist()
        }
        
        return metrics
    
    def save(self, directory):
        """
        Save model to directory.
        
        Args:
            directory (str): Directory path to save the model and metadata
        """
        if self.model is None:
            raise ValueError("No model to save")
        
        os.makedirs(directory, exist_ok=True)
        self.model.save(os.path.join(directory, 'threat_classifier.h5'))
        
        with open(os.path.join(directory, 'config.json'), 'w') as f:
            import json
            json.dump({
                'vocabulary_size': self.vocabulary_size,
                'sequence_length': self.sequence_length,
                'num_classes': self.num_classes,
                'class_names': self.class_names
            }, f)
    
    @classmethod
    def load(cls, directory):
        """
        Load model from directory.
        
        Args:
            directory (str): Directory path containing the saved model and metadata
            
        Returns:
            ThreatClassifier: Loaded instance
        """
        import json
        
        with open(os.path.join(directory, 'config.json'), 'r') as f:
            saved_data = json.load(f)
        
        instance = cls()
        instance.vocabulary_size = saved_data['vocabulary_size']
        instance.sequence_length = saved_data['sequence_length']
        instance.num_classes = saved_data['num_classes']
        instance.class_names = saved_data['class_names']
        instance.model = tf.keras.models.load_model(os.path.join(directory, 'threat_classifier.h5'))
        
        return instance

class ThreatSeverityEstimator:
    def load(self, directory):
        pass