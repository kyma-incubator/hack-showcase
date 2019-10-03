package k8scomponents_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/kyma-incubator/github-slack-connectors/scenario/github-issue-sentiment-analysis/internal/k8scomponents"
	"github.com/kyma-incubator/github-slack-connectors/scenario/github-issue-sentiment-analysis/internal/k8scomponents/mocks"
	v1alpha1 "github.com/kyma-project/kyma/components/event-bus/api/push/eventing.kyma-project.io/v1alpha1"
	"github.com/stretchr/testify/assert"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

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
		assert.Nil(t, data)
	})
}

func TestDeleteSubscription(t *testing.T) {
	t.Run("should return ServiceBinding", func(t *testing.T) {
		//given
		name := "name"
		namespace := "namespace"
		options := &v1.DeleteOptions{}
		mockClient := &mocks.SubscriptionInterface{}
		mockClient.On("Delete", name, options).Return(nil)
		//when
		err := k8scomponents.NewSubscription(mockClient, namespace).Delete(name, options)

		//then
		assert.NoError(t, err)

	})
}

func TestPrepare(t *testing.T) {
	t.Run("should return ServiceBindingUsage", func(t *testing.T) {
		//given
		namespace := "namespace"
		id := "github-repo"
		lambdaName := "name"
		body := &v1alpha1.Subscription{
			ObjectMeta: v1.ObjectMeta{
				Name:      lambdaName + "-lambda-issuesevent-v1",
				Namespace: namespace,
				Labels:    map[string]string{"Function": lambdaName},
			},
			SubscriptionSpec: v1alpha1.SubscriptionSpec{
				Endpoint:                      fmt.Sprintf("%s%s%s%s%s", "http://", lambdaName, ".", namespace, ":8080/"),
				EventType:                     "IssuesEvent",
				EventTypeVersion:              "v1",
				IncludeSubscriptionNameHeader: true,
				SourceID:                      fmt.Sprintf("%s%s", id, "-app"),
			},
		}
		mockClient := &mocks.SubscriptionInterface{}

		//when
		sub := k8scomponents.NewSubscription(mockClient, namespace).Prepare(id, lambdaName)

		//then
		assert.Equal(t, body, sub)

	})
}
