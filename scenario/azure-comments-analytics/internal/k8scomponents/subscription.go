package k8scomponents

import (
	"fmt"

	"github.com/kyma-incubator/hack-showcase/scenario/azure-comments-analytics/internal/apperrors"
	v1alpha1 "github.com/kyma-project/kyma/components/event-bus/api/push/eventing.kyma-project.io/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type Subscription interface {
	Create(body *v1alpha1.Subscription) (*v1alpha1.Subscription, apperrors.AppError)
	GetSubscription(id string) *v1alpha1.Subscription
}

type SubscriptionInterface interface {
	Create(*v1alpha1.Subscription) (*v1alpha1.Subscription, error)
}

type subscription struct {
	subscriptionInterface SubscriptionInterface
	namespace             string
}

//NewSubscription create new instance of subscription structure
func NewSubscription(sub SubscriptionInterface, nspace string) Subscription {
	return subscription{
		subscriptionInterface: sub,
		namespace:             nspace,
	}
}

func (s subscription) Create(body *v1alpha1.Subscription) (*v1alpha1.Subscription, apperrors.AppError) {
	data, err := s.subscriptionInterface.Create(body)
	if err != nil {
		return nil, apperrors.WrongInput("Can not create subscription: %s", err)
	}
	return data, nil
}

func (s subscription) GetSubscription(id string) *v1alpha1.Subscription {
	return &v1alpha1.Subscription{
		ObjectMeta: v1.ObjectMeta{
			Name:      "lambda-julia-lambda-issuesevent-v1",
			Namespace: s.namespace,
			Labels:    map[string]string{"Function": "julia-lambda"},
		},
		SubscriptionSpec: v1alpha1.SubscriptionSpec{
			Endpoint:                      fmt.Sprintf("%s%s%s", "http://julia-lambda.", s.namespace, ":8080/"),
			EventType:                     "IssuesEvent",
			EventTypeVersion:              "v1",
			IncludeSubscriptionNameHeader: true,
			SourceID:                      fmt.Sprintf("%s%s", id, "-app"),
		},
	}
}
