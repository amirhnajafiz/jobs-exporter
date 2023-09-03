package handler

import "k8s.io/client-go/kubernetes"

type Handler struct {
	Namespace string
	K8SClient *kubernetes.Clientset
}

func (h Handler) Monitor() {
	jobsInterface := h.K8SClient.BatchV1().Jobs(h.Namespace)
}
