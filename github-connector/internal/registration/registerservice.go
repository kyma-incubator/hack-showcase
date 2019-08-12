package registration

import (
	"os"
	"path/filepath"
	"time"

	"github.com/kyma-incubator/hack-showcase/github-connector/internal/apperrors"
	log "github.com/sirupsen/logrus"
)

const (
	retryDelay                = 5 * time.Second
	retriesCount              = 10
	specificationURL          = "https://raw.githubusercontent.com/colunira/github-openapi/master/githubopenAPI.json"
	applicationRegistryPrefix = "http://application-registry-external-api.kyma-integration.svc.cluster.local:8081/"
	applicationRegistrySuffix = "-app/v1/metadata/services"
)

var path, _ = filepath.Abs("./../github-connector/internal/registration/configs/githubasyncAPI.json")

//ServiceRegister is an interface containing all necessary functions required to register a service in Kyma Application Registry
type ServiceRegister interface {
	RegisterService() (string, error)
}

type serviceRegister struct {
	register ServiceRegister
}

//NewServiceRegister creates a serviceRegister instance with the passed in interface
func NewServiceRegister() serviceRegister {
	return serviceRegister{}
}

//RegisterService - register service in Kyma and get a response
func (r serviceRegister) RegisterService() (string, apperrors.AppError) {

	jsonBody, err := BuildServiceDetails(specificationURL, path)
	if err != nil {
		return "", apperrors.Internal("While building service details json: %s", err)
	}

	id, err := jsonBody.requestWithRetries()
	if err != nil {
		return "", apperrors.Internal("While trying to register service: %s", err.Error())
	}
	return id, nil
}

func (jsonBody *ServiceDetails) requestWithRetries() (string, error) {
	var id string
	var err error
	register := NewRegisterRequestSender()
	var applicationRegistryURL = applicationRegistryPrefix + os.Getenv("GITHUB_CONNECTOR_NAME") + applicationRegistrySuffix
	for i := 0; i < retriesCount; i++ {
		time.Sleep(retryDelay)
		id, err = register.Do(*jsonBody, applicationRegistryURL)
		if err == nil {
			break
		}
		log.Warn(err.Error())
	}
	if err != nil {
		return "", apperrors.UpstreamServerCallFailed("While sending service registration request: %s", err)
	}
	return id, nil
}
