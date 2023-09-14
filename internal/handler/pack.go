package handler

import "time"

type Pack struct {
	Name      string    `json:"name"`
	Namespace string    `json:"namespace"`
	Status    string    `json:"status"`
	Logs      string    `json:"logs"`
	Created   time.Time `json:"created"`
}
