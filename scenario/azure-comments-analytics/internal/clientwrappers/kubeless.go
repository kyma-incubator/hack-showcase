package wrappers

import (
	"github.com/kubeless/kubeless/pkg/client/clientset/versioned/typed/kubeless/v1beta1"
	"github.com/kyma-incubator/github-slack-connectors/scenario/azure-comments-analytics/internal/k8scomponents"
)

//KubelessWrapper is a wrapper dedicated for kubeless ClientSet
type KubelessWrapper interface {
	Function(namespace string) k8scomponents.Function
}

//KubelessClient describe constructors argument
type KubelessClient interface {
	Functions(string) v1beta1.FunctionInterface
}

type kubelessWrapper struct {
	client KubelessClient
}

//NewKubelessClient create and return kubelessWrapper
func NewKubelessClient(client KubelessClient) KubelessWrapper {
	return &kubelessWrapper{client: client}
}

func (s *kubelessWrapper) Function(namespace string) k8scomponents.Function {
	return k8scomponents.NewFunction(s.client.Functions(namespace), namespace)
}
