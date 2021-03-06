// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/openshift/api/servicecertsigner/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ServiceCertSignerOperatorConfigLister helps list ServiceCertSignerOperatorConfigs.
// All objects returned here must be treated as read-only.
type ServiceCertSignerOperatorConfigLister interface {
	// List lists all ServiceCertSignerOperatorConfigs in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ServiceCertSignerOperatorConfig, err error)
	// Get retrieves the ServiceCertSignerOperatorConfig from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.ServiceCertSignerOperatorConfig, error)
	ServiceCertSignerOperatorConfigListerExpansion
}

// serviceCertSignerOperatorConfigLister implements the ServiceCertSignerOperatorConfigLister interface.
type serviceCertSignerOperatorConfigLister struct {
	indexer cache.Indexer
}

// NewServiceCertSignerOperatorConfigLister returns a new ServiceCertSignerOperatorConfigLister.
func NewServiceCertSignerOperatorConfigLister(indexer cache.Indexer) ServiceCertSignerOperatorConfigLister {
	return &serviceCertSignerOperatorConfigLister{indexer: indexer}
}

// List lists all ServiceCertSignerOperatorConfigs in the indexer.
func (s *serviceCertSignerOperatorConfigLister) List(selector labels.Selector) (ret []*v1alpha1.ServiceCertSignerOperatorConfig, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ServiceCertSignerOperatorConfig))
	})
	return ret, err
}

// Get retrieves the ServiceCertSignerOperatorConfig from the index for a given name.
func (s *serviceCertSignerOperatorConfigLister) Get(name string) (*v1alpha1.ServiceCertSignerOperatorConfig, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("servicecertsigneroperatorconfig"), name)
	}
	return obj.(*v1alpha1.ServiceCertSignerOperatorConfig), nil
}
