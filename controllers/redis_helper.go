package controllers

import (
	"context"
	coreV1 "k8s.io/api/core/v1"
	v1 "kubebuilder-demo/api/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func CreateRedis(client client.Client, redisConfig *v1.Redis) error {
	newPod := &coreV1.Pod{}
	newPod.Name = redisConfig.Name
	newPod.Namespace = redisConfig.Namespace
	newPod.Spec.Containers = []coreV1.Container{
		{
			Name:            redisConfig.Name,
			Image:           "redis:5-alpine",
			ImagePullPolicy: coreV1.PullIfNotPresent,
			Ports: []coreV1.ContainerPort{
				{
					ContainerPort: int32(redisConfig.Spec.Port),
				},
			},
		},
	}
	return client.Create(context.Background(), newPod)
}
