package k8scomponents

import (
	"fmt"

	v1alpha1 "github.com/kyma-project/kyma/components/event-bus/api/push/eventing.kyma-project.io/v1alpha1"
	"github.com/pkg/errors"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

//Subscription define subscription struct
type Subscription interface {
	Create(body *v1alpha1.Subscription) (*v1alpha1.Subscription, error)
	Delete(name string, options *v1.DeleteOptions) error
	Prepare(id string, lambdaName string) *v1alpha1.Subscription
}

//SubscriptionInterface describe constructors argument and containe Subscriptions method
type SubscriptionInterface interface {
	Create(*v1alpha1.Subscription) (*v1alpha1.Subscription, error)
	Delete(name string, options *v1.DeleteOptions) error
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

func (s *subscription) Create(body *v1alpha1.Subscription) (*v1alpha1.Subscription, error) {
	data, err := s.subscriptionInterface.Create(body)
	if err != nil {
		return nil, errors.Wrap(err, "Can not create Subscription")
	}
	return data, nil
}

func (s *subscription) Delete(name string, options *v1.DeleteOptions) error {
	return s.subscriptionInterface.Delete(name, options)
}

func (s *subscription) Prepare(id string, lambdaName string) *v1alpha1.Subscription {
	return &v1alpha1.Subscription{
		ObjectMeta: v1.ObjectMeta{
			Name:      lambdaName + "-lambda-issuesevent-v1",
			Namespace: s.namespace,
			Labels:    map[string]string{"Function": lambdaName},
		},
		SubscriptionSpec: v1alpha1.SubscriptionSpec{
			Endpoint:                      fmt.Sprintf("%s%s%s%s%s", "http://", lambdaName, ".", s.namespace, ":8080/"),
			EventType:                     "IssuesEvent",
			EventTypeVersion:              "v1",
			IncludeSubscriptionNameHeader: true,
			SourceID:                      fmt.Sprintf("%s%s", id, "-app"),
		},
	}
}
