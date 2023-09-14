package handler

import (
	"context"
	"log"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type Handler struct {
	Namespace string
	K8SClient *kubernetes.Clientset
}

func (h Handler) Monitor() {
	ctx := context.Background()

	for {
		jobs, err := h.K8SClient.BatchV1().Jobs(h.Namespace).List(ctx, v1.ListOptions{})
		if err != nil {
			log.Println(err)
		}

		for _, job := range jobs.Items {
			p := &Pack{
				Name:      job.GetName(),
				Namespace: job.GetNamespace(),
				Created:   job.GetCreationTimestamp().Time,
				Status:    job.Status.String(),
			}
		}
	}
}
