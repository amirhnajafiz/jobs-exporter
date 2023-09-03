package handler

import (
	"context"
	"log"

	"k8s.io/client-go/kubernetes"
)

type Handler struct {
	Namespace string
	K8SClient *kubernetes.Clientset
}

func (h Handler) Monitor() {
	ctx := context.Background()
	jobsInterface := h.K8SClient.BatchV1().Jobs(h.Namespace)

	for {
		jobs, err := jobsInterface.List(ctx, nil)
		if err != nil {
			log.Println(err)
		}
	}
}
