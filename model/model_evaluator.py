import os
import numpy as np
import pandas as pd
import tensorflow as tf
import matplotlib.pyplot as plt
import seaborn as sns
from sklearn.metrics import confusion_matrix, classification_report, roc_curve, auc, precision_recall_curve
import json

class ModelEvaluator:
    """
    Evaluates AI models for Kubernetes security monitoring.
    Provides comprehensive evaluation metrics and visualizations.
    """
    
    def __init__(self, config=None):
        """
        Initialize the model evaluator with configuration.
        
        Args:
            config (dict): Configuration parameters for evaluation
        """
        self.config = config or {
            'threshold_range': np.linspace(0.01, 0.99, 50),
            'output_dir': 'evaluation_results',
            'save_plots': True
        }
        
        os.makedirs(self.config['output_dir'], exist_ok=True)
    
    def evaluate_anomaly_detector(self, model, X_test, y_test=None, dataset_name='test'):
        """
        Evaluate an anomaly detection model.
        
        Args:
            model: Anomaly detection model to evaluate
            X_test: Test data features
            y_test: Optional ground truth labels (1 for anomaly, 0 for normal)
            dataset_name: Name of the dataset for reporting
            
        Returns:
            dict: Evaluation metrics
        """
        # Get anomaly scores
        anomaly_scores, is_anomaly = model.predict(X_test)
        
        results = {
            'dataset': dataset_name,
            'num_samples': len(X_test),
            'anomaly_count': np.sum(is_anomaly),
            'anomaly_rate': float(np.mean(is_anomaly)),
            'score_stats': {
                'min': float(np.min(anomaly_scores)),
                'max': float(np.max(anomaly_scores)),
                'mean': float(np.mean(anomaly_scores)),
                'median': float(np.median(anomaly_scores)),
                'std': float(np.std(anomaly_scores)),
                'threshold': float(model.threshold)
            }
        }
        
        # If ground truth is available, calculate classification metrics
        if y_test is not None:
            # Ensure y_test is binary
            y_test_binary = (y_test > 0).astype(int)
            
            # Calculate metrics at current threshold
            results['classification_metrics'] = self._calculate_binary_metrics(
                y_test_binary, is_anomaly, anomaly_scores
            )
            
            # Find optimal threshold
            optimal_threshold, optimal_metrics = self._find_optimal_threshold(
                y_test_binary, anomaly_scores
            )
            
            results['optimal_threshold'] = float(optimal_threshold)
            results['optimal_metrics'] = optimal_metrics
            
            # Generate plots if configured
            if self.config['save_plots']:
                self._plot_anomaly_detection_results(
                    y_test_binary, anomaly_scores, dataset_name
                )
        
        return results
    
    def evaluate_classifier(self, model, X_test, y_test, class_names=None, dataset_name='test'):
        """
        Evaluate a classification model.
        
        Args:
            model: Classification model to evaluate
            X_test: Test data features
            y_test: Test data labels (one-hot encoded)
            class_names: Names of the classes
            dataset_name: Name of the dataset for reporting
            
        Returns:
            dict: Evaluation metrics
        """
        # Get predictions
        y_pred_probs, y_pred_classes, y_pred_names = model.predict(X_test)
        
        # Convert one-hot encoded y_test to class indices
        y_test_classes = np.argmax(y_test, axis=1)
        
        # Use provided class names or get them from the model
        if class_names is None:
            if hasattr(model, 'class_names'):
                class_names = model.class_names
            else:
                class_names = [f"Class {i}" for i in range(y_test.shape[1])]
        
        # Calculate metrics
        conf_matrix = confusion_matrix(y_test_classes, y_pred_classes)
        class_report = classification_report(y_test_classes, y_pred_classes, 
                                            target_names=class_names, output_dict=True)
        
        # Calculate per-class ROC and PR curves
        roc_curves = {}
        pr_curves = {}
        
        for i, class_name in enumerate(class_names):
            # One-vs-rest ROC curve
            fpr, tpr, _ = roc_curve((y_test[:, i] > 0.5).astype(int), y_pred_probs[:, i])
            roc_auc = auc(fpr, tpr)
            
            roc_curves[class_name] = {
                'fpr': fpr.tolist(),
                'tpr': tpr.tolist(),
                'auc': float(roc_auc)
            }
            
            # Precision-Recall curve
            precision, recall, _ = precision_recall_curve((y_test[:, i] > 0.5).astype(int), y_pred_probs[:, i])
            pr_auc = auc(recall, precision)
            
            pr_curves[class_name] = {
                'precision': precision.tolist(),
                'recall': recall.tolist(),
                'auc': float(pr_auc)
            }
        
        results = {
            'dataset': dataset_name,
            'num_samples': len(X_test),
            'accuracy': float(class_report['accuracy']),
            'macro_avg_precision': float(class_report['macro avg']['precision']),
            'macro_avg_recall': float(class_report['macro avg']['recall']),
            'macro_avg_f1': float(class_report['macro avg']['f1-score']),
            'weighted_avg_precision': float(class_report['weighted avg']['precision']),
            'weighted_avg_recall': float(class_report['weighted avg']['recall']),
            'weighted_avg_f1': float(class_report['weighted avg']['f1-score']),
            'per_class_metrics': {
                class_name: {
                    'precision': float(class_report[class_name]['precision']),
                    'recall': float(class_report[class_name]['recall']),
                    'f1_score': float(class_report[class_name]['f1-score']),
                    'support': int(class_report[class_name]['support'])
                }
                for class_name in class_names
            },
            'confusion_matrix': conf_matrix.tolist(),
            'roc_curves': roc_curves,
            'pr_curves': pr_curves
        }
        
        # Generate plots if configured
        if self.config['save_plots']:
            self._plot_classification_results(
                y_test_classes, y_pred_classes, y_pred_probs, class_names, dataset_name
            )
        
        return results
    
    def evaluate_severity_estimator(self, model, X_test, y_test, dataset_name='test'):
        """
        Evaluate a severity estimation model.
        
        Args:
            model: Severity estimation model to evaluate
            X_test: Test data features
            y_test: Test data severity scores (0-1)
            dataset_name: Name of the dataset for reporting
            
        Returns:
            dict: Evaluation metrics
        """
        # Get predictions
        y_pred = model.predict(X_test)
        
        # Calculate metrics
        mse = np.mean(np.square(y_test - y_pred))
        rmse = np.sqrt(mse)
        mae = np.mean(np.abs(y_test - y_pred))
        
        # Calculate R-squared
        ss_total = np.sum(np.square(y_test - np.mean(y_test)))
        ss_residual = np.sum(np.square(y_test - y_pred))
        r_squared = 1 - (ss_residual / ss_total)
        
        # Calculate error distribution
        errors = y_test - y_pred
        
        results = {
            'dataset': dataset_name,
            'num_samples': len(X_test),
            'mse': float(mse),
            'rmse': float(rmse),
            'mae': float(mae),
            'r_squared': float(r_squared),
            'error_stats': {
                'min': float(np.min(errors)),
                'max': float(np.max(errors)),
                'mean': float(np.mean(errors)),
                'median': float(np.median(errors)),
                'std': float(np.std(errors))
            }
        }
        
        # Generate plots if configured
        if self.config['save_plots']:
            self._plot_regression_results(y_test, y_pred, dataset_name)
        
        return results
    
    def _calculate_binary_metrics(self, y_true, y_pred, scores=None):
        """Calculate metrics for binary classification"""
        # True positives, false positives, true negatives, false negatives
        tp = np.sum((y_true == 1) & (y_pred == 1))
        fp = np.sum((y_true == 0) & (y_pred == 1))
        tn = np.sum((y_true == 0) & (y_pred == 0))
        fn = np.sum((y_true == 1) & (y_pred == 0))
        
        # Calculate metrics
        accuracy = (tp + tn) / (tp + fp + tn + fn)
        precision = tp / (tp + fp) if (tp + fp) > 0 else 0
        recall = tp / (tp + fn) if (tp + fn) > 0 else 0
        f1 = 2 * precision * recall / (precision + recall) if (precision + recall) > 0 else 0
        
        metrics = {
            'accuracy': float(accuracy),
            'precision': float(precision),
            'recall': float(recall),
            'f1_score': float(f1),
            'true_positives': int(tp),
            'false_positives': int(fp),
            'true_negatives': int(tn),
            'false_negatives': int(fn)
        }
        
        # Calculate ROC and PR curves if scores are provided
        if scores is not None:
            fpr, tpr, _ = roc_curve(y_true, scores)
            roc_auc = auc(fpr, tpr)
            
            precision_curve, recall_curve, _ = precision_recall_curve(y_true, scores)
            pr_auc = auc(recall_curve, precision_curve)
            
            metrics['roc_curve'] = {
                'fpr': fpr.tolist(),
                'tpr': tpr.tolist(),
                'auc': float(roc_auc)
            }
            
            metrics['pr_curve'] = {
                'precision': precision_curve.tolist(),
                'recall': recall_curve.tolist(),
                'auc': float(pr_auc)
            }
        
        return metrics
    
    def _find_optimal_threshold(self, y_true, scores):
        """Find optimal threshold for anomaly detection"""
        best_f1 = 0
        best_threshold = 0
        best_metrics = None
        
        for threshold in self.config['threshold_range']:
            y_pred = (scores > threshold).astype(int)
            metrics = self._calculate_binary_metrics(y_true, y_pred)
            
            if metrics['f1_score'] > best_f1:
                best_f1 = metrics['f1_score']
                best_threshold = threshold
                best_metrics = metrics
        
        return best_threshold, best_metrics
    
    def _plot_anomaly_detection_results(self, y_true, scores, dataset_name):
        """Generate plots for anomaly detection results"""
        # Create output directory
        plot_dir = os.path.join(self.config['output_dir'], 'anomaly_detection', dataset_name)
        os.makedirs(plot_dir, exist_ok=True)
        
        # Plot score distribution
        plt.figure(figsize=(10, 6))
        sns.histplot(scores[y_true == 0], label='Normal', alpha=0.5, bins=50)
        sns.histplot(scores[y_true == 1], label='Anomaly', alpha=0.5, bins=50)
        plt.xlabel('Anomaly Score')
        plt.ylabel('Count')
        plt.title('Distribution of Anomaly Scores')
        plt.legend()
        plt.savefig(os.path.join(plot_dir, 'score_distribution.png'))
        plt.close()
        
        # Plot ROC curve
        fpr, tpr, _ = roc_curve(y_true, scores)
        roc_auc = auc(fpr, tpr)
        
        plt.figure(figsize=(10, 6))
        plt.plot(fpr, tpr, label=f'ROC Curve (AUC = {roc_auc:.3f})')
        plt.plot([0, 1], [0, 1], 'k--')
        plt.xlabel('False Positive Rate')
        plt.ylabel('True Positive Rate')
        plt.title('Receiver Operating Characteristic (ROC) Curve')
        plt.legend(loc='lower right')
        plt.savefig(os.path.join(plot_dir, 'roc_curve.png'))
        plt.close()
        
        # Plot Precision-Recall curve
        precision, recall, _ = precision_recall_curve(y_true, scores)
        pr_auc = auc(recall, precision)
        
        plt.figure(figsize=(10, 6))
        plt.plot(recall, precision, label=f'PR Curve (AUC = {pr_auc:.3f})')
        plt.xlabel('Recall')
        plt.ylabel('Precision')
        plt.title('Precision-Recall Curve')
        plt.legend(loc='lower left')
        plt.savefig(os.path.join(plot_dir, 'pr_curve.png'))
        plt.close()
        
        # Plot F1 score vs threshold
        f1_scores = []
        thresholds = self.config['threshold_range']
        
        for threshold in thresholds:
            y_pred = (scores > threshold).astype(int)
            metrics = self._calculate_binary_metrics(y_true, y_pred)
            f1_scores.append(metrics['f1_score'])
        
        plt.figure(figsize=(10, 6))
        plt.plot(thresholds, f1_scores)
        plt.xlabel('Threshold')
        plt.ylabel('F1 Score')
        plt.title('F1 Score vs Threshold')
        plt.savefig(os.path.join(plot_dir, 'f1_vs_threshold.png'))
        plt.close()
    
    def _plot_classification_results(self, y_true, y_pred, y_pred_probs, class_names, dataset_name):
        """Generate plots for classification results"""
        # Create output directory
        plot_dir = os.path.join(self.config['output_dir'], 'classification', dataset_name)
        os.makedirs(plot_dir, exist_ok=True)
        
        # Plot confusion matrix
        plt.figure(figsize=(10, 8))
        cm = confusion_matrix(y_true, y_pred)
        sns.heatmap(cm, annot=True, fmt='d', cmap='Blues', xticklabels=class_names, yticklabels=class_names)
        plt.xlabel('Predicted')
        plt.ylabel('True')
        plt.title('Confusion Matrix')
        plt.savefig(os.path.join(plot_dir, 'confusion_matrix.png'))
        plt.close()
        
        # Plot ROC curves for each class
        plt.figure(figsize=(10, 8))
        
        for i, class_name in enumerate(class_names):
            # One-vs-rest ROC curve
            fpr, tpr, _ = roc_curve((y_true == i).astype(int), y_pred_probs[:, i])
            roc_auc = auc(fpr, tpr)
            
            plt.plot(fpr, tpr, label=f'{class_name} (AUC = {roc_auc:.3f})')
        
        plt.plot([0, 1], [0, 1], 'k--')
        plt.xlabel('False Positive Rate')
        plt.ylabel('True Positive Rate')
        plt.title('ROC Curves (One-vs-Rest)')
        plt.legend(loc='lower right')
        plt.savefig(os.path.join(plot_dir, 'roc_curves.png'))
        plt.close()
        
        # Plot class distribution
        plt.figure(figsize=(10, 6))
        sns.countplot(x=y_true)
        plt.xticks(range(len(class_names)), class_names)
        plt.xlabel('Class')
        plt.ylabel('Count')
        plt.title('Class Distribution')
        plt.savefig(os.path.join(plot_dir, 'class_distribution.png'))
        plt.close()
    
    def _plot_regression_results(self, y_true, y_pred, dataset_name):
        """Generate plots for regression results"""
        # Create output directory
        plot_dir = os.path.join(self.config['output_dir'], 'regression', dataset_name)
        os.makedirs(plot_dir, exist_ok=True)
        
        # Plot predicted vs actual
        plt.figure(figsize=(10, 8))
        plt.scatter(y_true, y_pred, alpha=0.5)
        plt.plot([0, 1], [0, 1], 'r--')
        plt.xlabel('Actual Severity')
        plt.ylabel('Predicted Severity')
        plt.title('Predicted vs Actual Severity')
        plt.savefig(os.path.join(plot_dir, 'predicted_vs_actual.png'))
        plt.close()
        
        # Plot error distribution
        errors = y_true - y_pred
        
        plt.figure(figsize=(10, 6))
        sns.histplot(errors, bins=50, kde=True)
        plt.xlabel('Error (Actual - Predicted)')
        plt.ylabel('Frequency')
        plt.title('Distribution of Prediction Errors')
        plt.savefig(os.path.join(plot_dir, 'error_distribution.png'))
        plt.close()