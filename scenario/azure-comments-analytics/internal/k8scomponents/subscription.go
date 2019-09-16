package k8scomponents

import (
	"fmt"

	"github.com/kyma-incubator/hack-showcase/scenario/azure-comments-analytics/internal/apperrors"
	v1alpha1 "github.com/kyma-project/kyma/components/event-bus/api/push/eventing.kyma-project.io/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

//Subscription define subscription struct
type Subscription interface {
	Create(body *v1alpha1.Subscription) (*v1alpha1.Subscription, apperrors.AppError)
	GetEventBody(id string) *v1alpha1.Subscription
}

//SubscriptionInterface describe constructors argument and containe Subscriptions method
type SubscriptionInterface interface {
	Create(*v1alpha1.Subscription) (*v1alpha1.Subscription, error)
}

type subscription struct {
	subscriptionInterface SubscriptionInterface
	namespace             string
}

//NewSubscription create new instance of subscription structure
func NewSubscription(sub SubscriptionInterface, nspace string) Subscription {
	return &subscription{
		subscriptionInterface: sub,
		namespace:             nspace,
	}
}

func (s *subscription) Create(body *v1alpha1.Subscription) (*v1alpha1.Subscription, apperrors.AppError) {
	data, err := s.subscriptionInterface.Create(body)
	if err != nil {
		return nil, apperrors.WrongInput("Can not create Subscription: %s", err)
	}
	return data, nil
}

func (s *subscription) GetEventBody(id string) *v1alpha1.Subscription {
	return &v1alpha1.Subscription{
		ObjectMeta: v1.ObjectMeta{
			Name:      "lambda-" + id[7:] + "-lambda-issuesevent-v1",
			Namespace: s.namespace,
			Labels:    map[string]string{"Function": id[7:] + "-lambda"},
		},
		SubscriptionSpec: v1alpha1.SubscriptionSpec{
			Endpoint:                      fmt.Sprintf("%s%s%s%s%s", "http://", id[7:], "-lambda.", s.namespace, ":8080/"),
			EventType:                     "IssuesEvent",
			EventTypeVersion:              "v1",
			IncludeSubscriptionNameHeader: true,
			SourceID:                      fmt.Sprintf("%s%s", id, "-app"),
		},
	}
}
