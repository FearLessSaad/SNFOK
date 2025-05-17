package k8sclient

import (
	"fmt"
	"log"
	"path/filepath"
	"sync"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

// singleton instance and synchronization primitives
var (
	clientsetInstance *kubernetes.Clientset
	once              sync.Once
	initError         error
)

// GetClientset returns the singleton Kubernetes clientset instance.
// It initializes the clientset on the first call using the default kubeconfig path (~/.kube/config).
// If initialization fails, it returns the error.
func GetClientset() (*kubernetes.Clientset, error) {
	once.Do(func() {
		// Build the Kubernetes client configuration
		kubeconfig := filepath.Join(homedir.HomeDir(), ".kube", "config")
		config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
		if err != nil {
			initError = err
			log.Printf("Failed to load kubeconfig: %v", err)
			return
		}

		// Create a new Kubernetes clientset
		clientsetInstance, err = kubernetes.NewForConfig(config)
		if err != nil {
			initError = err
			log.Printf("Failed to create Kubernetes client: %v", err)
			return
		}
	})

	if initError != nil {
		return nil, initError
	}
	if clientsetInstance == nil {
		return nil, fmt.Errorf("clientset initialization failed")
	}
	return clientsetInstance, nil
}

// GetClientsetWithConfig returns the singleton Kubernetes clientset instance,
// initialized with a custom kubeconfig path.
// It ensures the clientset is only initialized once, even with multiple calls.
func GetClientsetWithConfig(kubeconfigPath string) (*kubernetes.Clientset, error) {
	once.Do(func() {
		// Build the Kubernetes client configuration
		config, err := clientcmd.BuildConfigFromFlags("", kubeconfigPath)
		if err != nil {
			initError = err
			log.Printf("Failed to load kubeconfig from %s: %v", kubeconfigPath, err)
			return
		}

		// Create a new Kubernetes clientset
		clientsetInstance, err = kubernetes.NewForConfig(config)
		if err != nil {
			initError = err
			log.Printf("Failed to create Kubernetes client: %v", err)
			return
		}
	})

	if initError != nil {
		return nil, initError
	}
	if clientsetInstance == nil {
		return nil, fmt.Errorf("clientset initialization failed")
	}
	return clientsetInstance, nil
}

func GetClusterName(kubeconfigPath string) (string, error) {
	if kubeconfigPath == "" {
		kubeconfigPath = filepath.Join(homedir.HomeDir(), ".kube", "config")
	}

	// Load the kubeconfig file
	config, err := clientcmd.LoadFromFile(kubeconfigPath)
	if err != nil {
		return "", fmt.Errorf("failed to load kubeconfig from %s: %w", kubeconfigPath, err)
	}

	// Get the current context
	currentContext := config.CurrentContext
	if currentContext == "" {
		return "", fmt.Errorf("no current context found in kubeconfig")
	}

	// Find the context definition
	ctx, exists := config.Contexts[currentContext]
	if !exists {
		return "", fmt.Errorf("context %s not found in kubeconfig", currentContext)
	}

	// Return the cluster name
	if ctx.Cluster == "" {
		return "", fmt.Errorf("no cluster specified for context %s", currentContext)
	}
	return ctx.Cluster, nil
}
