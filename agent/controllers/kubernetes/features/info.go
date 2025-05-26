package features

import (
	"context"
	"fmt"

	"github.com/FearLessSaad/SNFOK/shared/agent_dto"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// GetWorkerNodes retrieves all worker nodes with their IP addresses and hostnames
func GetWorkerNodes(clientset *kubernetes.Clientset) ([]agent_dto.WorkerNodeInfo, error) {
	// List all nodes
	nodes, err := clientset.CoreV1().Nodes().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, fmt.Errorf("failed to list nodes: %v", err)
	}

	var workerNodes []agent_dto.WorkerNodeInfo
	for _, node := range nodes.Items {
		// Skip control plane nodes
		if isControlPlaneNode(node.Labels) {
			continue
		}

		// Extract node information
		nodeInfo := agent_dto.WorkerNodeInfo{
			Name: node.Name,
		}

		// Get IP addresses and hostname from node addresses
		for _, addr := range node.Status.Addresses {
			switch addr.Type {
			case "InternalIP":
				nodeInfo.InternalIP = addr.Address
			case "ExternalIP":
				nodeInfo.ExternalIP = addr.Address
			case "Hostname":
				nodeInfo.Hostname = addr.Address
			}
		}

		// Fallback to node.Name if hostname is empty
		if nodeInfo.Hostname == "" {
			nodeInfo.Hostname = node.Name
		}

		workerNodes = append(workerNodes, nodeInfo)
	}

	return workerNodes, nil
}

// GetAllNamespaces retrieves the names of all namespaces in the Kubernetes cluster.
func GetAllNamespaces(clientset *kubernetes.Clientset) ([]string, error) {
	// List all namespaces
	namespaces, err := clientset.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, fmt.Errorf("failed to list namespaces: %v", err)
	}

	// Extract namespace names
	var namespaceNames []string
	for _, ns := range namespaces.Items {
		namespaceNames = append(namespaceNames, ns.Name)
	}

	return namespaceNames, nil
}

// GetNamespaceResources retrieves all running resources in the specified namespace
func GetNamespaceResources(clientset *kubernetes.Clientset, namespace string) (agent_dto.NamespaceResources, error) {
	var resources agent_dto.NamespaceResources

	// List Pods
	pods, err := clientset.CoreV1().Pods(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return agent_dto.NamespaceResources{}, fmt.Errorf("failed to list pods in namespace %s: %v", namespace, err)
	}
	for _, pod := range pods.Items {
		// Collect container details
		var containers []agent_dto.ContainerInfo
		for i, container := range pod.Spec.Containers {
			// Default status
			status := "Unknown"
			// Check status from ContainerStatuses if available
			if i < len(pod.Status.ContainerStatuses) {
				containerStatus := pod.Status.ContainerStatuses[i]
				if containerStatus.State.Running != nil {
					status = "Running"
				} else if containerStatus.State.Waiting != nil {
					status = fmt.Sprintf("Waiting (%s)", containerStatus.State.Waiting.Reason)
				} else if containerStatus.State.Terminated != nil {
					status = fmt.Sprintf("Terminated (%s)", containerStatus.State.Terminated.Reason)
				}
			}
			containers = append(containers, agent_dto.ContainerInfo{
				Name:   container.Name,
				Image:  container.Image,
				Status: status,
			})
		}

		resources.Pods = append(resources.Pods, agent_dto.PodInfo{
			Name:       pod.Name,
			Phase:      string(pod.Status.Phase),
			Containers: containers,
		})
	}

	// List Deployments
	deployments, err := clientset.AppsV1().Deployments(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return agent_dto.NamespaceResources{}, fmt.Errorf("failed to list deployments in namespace %s: %v", namespace, err)
	}
	for _, dep := range deployments.Items {
		resources.Deployments = append(resources.Deployments, agent_dto.DeploymentInfo{
			Name:              dep.Name,
			Replicas:          *dep.Spec.Replicas,
			AvailableReplicas: dep.Status.AvailableReplicas,
		})
	}

	// List Services
	services, err := clientset.CoreV1().Services(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return agent_dto.NamespaceResources{}, fmt.Errorf("failed to list services in namespace %s: %v", namespace, err)
	}
	for _, svc := range services.Items {
		resources.Services = append(resources.Services, agent_dto.ServiceInfo{
			Name:      svc.Name,
			Type:      string(svc.Spec.Type),
			ClusterIP: svc.Spec.ClusterIP,
		})
	}

	// List StatefulSets
	statefulSets, err := clientset.AppsV1().StatefulSets(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return agent_dto.NamespaceResources{}, fmt.Errorf("failed to list statefulsets in namespace %s: %v", namespace, err)
	}
	for _, sts := range statefulSets.Items {
		resources.StatefulSets = append(resources.StatefulSets, agent_dto.StatefulSetInfo{
			Name:              sts.Name,
			Replicas:          *sts.Spec.Replicas,
			AvailableReplicas: sts.Status.AvailableReplicas,
		})
	}

	// List DaemonSets
	daemonSets, err := clientset.AppsV1().DaemonSets(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return agent_dto.NamespaceResources{}, fmt.Errorf("failed to list daemonsets in namespace %s: %v", namespace, err)
	}
	for _, ds := range daemonSets.Items {
		resources.DaemonSets = append(resources.DaemonSets, agent_dto.DaemonSetInfo{
			Name:                   ds.Name,
			DesiredNumberScheduled: ds.Status.DesiredNumberScheduled,
			NumberReady:            ds.Status.NumberReady,
		})
	}

	return resources, nil
}

// isControlPlaneNode checks if a node is a control plane node based on labels
func isControlPlaneNode(labels map[string]string) bool {
	// Check common control plane labels
	for key, value := range labels {
		if (key == "node-role.kubernetes.io/control-plane" && value == "") ||
			(key == "node-role.kubernetes.io/master" && value == "") {
			return true
		}
	}
	return false
}

// CountAllRunningPods counts all pods in the "Running" phase across all namespaces
func CountAllRunningPods(clientset *kubernetes.Clientset) (int, error) {
	// Get all namespaces
	namespaces, err := GetAllNamespaces(clientset)
	if err != nil {
		return 0, fmt.Errorf("failed to get namespaces: %v", err)
	}

	// Initialize counter
	totalRunningPods := 0

	// Iterate through each namespace
	for _, namespace := range namespaces {
		// List pods in the namespace
		pods, err := clientset.CoreV1().Pods(namespace).List(context.TODO(), metav1.ListOptions{})
		if err != nil {
			return 0, fmt.Errorf("failed to list pods in namespace %s: %v", namespace, err)
		}

		// Count pods in "Running" phase
		for _, pod := range pods.Items {
			if pod.Status.Phase == "Running" {
				totalRunningPods++
			}
		}
	}

	return totalRunningPods, nil
}
