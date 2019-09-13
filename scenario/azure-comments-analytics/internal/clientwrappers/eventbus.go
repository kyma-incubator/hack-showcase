package wrappers

import (
	"github.com/kyma-incubator/hack-showcase/scenario/azure-comments-analytics/internal/k8scomponents"
	eventingv1alpha1 "github.com/kyma-project/kyma/components/event-bus/generated/push/clientset/versioned/typed/eventing.kyma-project.io/v1alpha1"
)

type EventbusWrapper interface {
	Subscription(namespace string) k8scomponents.Subscription
}

type EventbusClient interface {
	Subscriptions(string) eventingv1alpha1.SubscriptionInterface
}
type eventbusWrapper struct {
	client EventbusClient
}

func NewSubscriptionManager(bus EventbusClient) EventbusWrapper {
	return &eventbusWrapper{client: bus}
}

func (s *eventbusWrapper) Subscription(namespace string) k8scomponents.Subscription {
	return k8scomponents.NewSubscription(s.client.Subscriptions(namespace), namespace)
}
