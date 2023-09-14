package handler

import "errors"

var (
	ErrJobPulling   = errors.New("failed to pull jobs from kubernetes namespace")
	ErrPackBuild    = errors.New("internal error in building job information pack")
	ErrKafkaPublish = errors.New("failed to push batch over kafka")
)
