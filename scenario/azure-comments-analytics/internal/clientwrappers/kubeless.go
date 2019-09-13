package wrappers

import (
	"github.com/kubeless/kubeless/pkg/client/clientset/versioned/typed/kubeless/v1beta1"
	"github.com/kyma-incubator/hack-showcase/scenario/azure-comments-analytics/internal/k8scomponents"
)

type KubelessWrapper interface {
	Function(namespace string) k8scomponents.Function
}

type KubelessClient interface {
	Functions(string) v1beta1.FunctionInterface
}

type kubelessWrapper struct {
	client KubelessClient
}

func NewKubelessWrapper(client KubelessClient) KubelessWrapper {
	return &kubelessWrapper{client: client}
}

func (s *kubelessWrapper) Function(namespace string) k8scomponents.Function {
	return k8scomponents.NewFunction(s.client.Functions(namespace), namespace)
}
