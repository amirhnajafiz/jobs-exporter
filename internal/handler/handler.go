package handler

import (
	"context"
	"encoding/json"
	"log"

	"github.com/segmentio/kafka-go"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type Handler struct {
	Namespace string
	KafkaConn *kafka.Conn
	K8SClient *kubernetes.Clientset
}

func (h Handler) Monitor() {
	ctx := context.Background()

	for {
		// get list of namespace jobs
		jobs, err := h.K8SClient.BatchV1().Jobs(h.Namespace).List(ctx, v1.ListOptions{})
		if err != nil {
			log.Println(err)
		}

		// create messages array
		var messages []kafka.Message

		// iterate over existing jobs
		for _, job := range jobs.Items {
			p := &Pack{
				Name:      job.GetName(),
				Namespace: job.GetNamespace(),
				Created:   job.GetCreationTimestamp().Time,
				Status:    job.Status.String(),
			}

			bytes, er := json.Marshal(&p)
			if er != nil {
				log.Println(er)
			}

			messages = append(messages, kafka.Message{
				Value: bytes,
			})
		}

		// publish over kafka
		if _, err = h.KafkaConn.WriteMessages(messages...); err != nil {
			log.Println(err)
		}
	}
}
