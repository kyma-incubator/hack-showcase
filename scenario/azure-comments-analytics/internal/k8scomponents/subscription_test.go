package k8scomponents_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/kyma-incubator/hack-showcase/scenario/azure-comments-analytics/internal/k8scomponents"
	"github.com/kyma-incubator/hack-showcase/scenario/azure-comments-analytics/internal/k8scomponents/mocks"
	v1alpha1 "github.com/kyma-project/kyma/components/event-bus/api/push/eventing.kyma-project.io/v1alpha1"
	"github.com/stretchr/testify/assert"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func subscriptionNiller() *v1alpha1.Subscription {
	return nil
}

func TestCreateSubscription(t *testing.T) {
	t.Run("should create Binding, return new bindingUsage and nil", func(t *testing.T) {
		//given
		subscription := &v1alpha1.Subscription{}
		mockClient := &mocks.SubscriptionInterface{}
		mockClient.On("Create", subscription).Return(subscription, nil)

		//when
		data, err := k8scomponents.NewSubscription(mockClient, "default").Create(subscription)

		//then
		assert.NoError(t, err)
		assert.Equal(t, subscription, data)
	})

	t.Run("should return nil and error when cannot create BindingUsage", func(t *testing.T) {
		//given
		subscription := &v1alpha1.Subscription{}
		mockClient := &mocks.SubscriptionInterface{}
		mockClient.On("Create", subscription).Return(nil, errors.New("error text"))

		//when
		data, err := k8scomponents.NewSubscription(mockClient, "default").Create(subscription)

		//then
		assert.Error(t, err)
		assert.Equal(t, subscriptionNiller(), data)
	})
}

func TestGetEventBodySubscription(t *testing.T) {
	t.Run("should return ServiceBindingUsage", func(t *testing.T) {
		//given
		namespace := "namespace"
		id := "github-repo"
		body := &v1alpha1.Subscription{
			ObjectMeta: v1.ObjectMeta{
				Name:      "lambda-repo-lambda-issuesevent-v1",
				Namespace: namespace,
				Labels:    map[string]string{"Function": "julia-lambda"},
			},
			SubscriptionSpec: v1alpha1.SubscriptionSpec{
				Endpoint:                      fmt.Sprintf("%s%s%s", "http://repo-lambda.", namespace, ":8080/"),
				EventType:                     "IssuesEvent",
				EventTypeVersion:              "v1",
				IncludeSubscriptionNameHeader: true,
				SourceID:                      fmt.Sprintf("%s%s", id, "-app"),
			},
		}
		mockClient := &mocks.SubscriptionInterface{}

		//when
		sub := k8scomponents.NewSubscription(mockClient, namespace).GetEventBody(id)

		//then
		assert.Equal(t, body, sub)

	})
}
