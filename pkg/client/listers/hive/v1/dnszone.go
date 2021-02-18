// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/openshift/hive/apis/hive/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// DNSZoneLister helps list DNSZones.
// All objects returned here must be treated as read-only.
type DNSZoneLister interface {
	// List lists all DNSZones in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.DNSZone, err error)
	// DNSZones returns an object that can list and get DNSZones.
	DNSZones(namespace string) DNSZoneNamespaceLister
	DNSZoneListerExpansion
}

// dNSZoneLister implements the DNSZoneLister interface.
type dNSZoneLister struct {
	indexer cache.Indexer
}

// NewDNSZoneLister returns a new DNSZoneLister.
func NewDNSZoneLister(indexer cache.Indexer) DNSZoneLister {
	return &dNSZoneLister{indexer: indexer}
}

// List lists all DNSZones in the indexer.
func (s *dNSZoneLister) List(selector labels.Selector) (ret []*v1.DNSZone, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.DNSZone))
	})
	return ret, err
}

// DNSZones returns an object that can list and get DNSZones.
func (s *dNSZoneLister) DNSZones(namespace string) DNSZoneNamespaceLister {
	return dNSZoneNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// DNSZoneNamespaceLister helps list and get DNSZones.
// All objects returned here must be treated as read-only.
type DNSZoneNamespaceLister interface {
	// List lists all DNSZones in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.DNSZone, err error)
	// Get retrieves the DNSZone from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.DNSZone, error)
	DNSZoneNamespaceListerExpansion
}

// dNSZoneNamespaceLister implements the DNSZoneNamespaceLister
// interface.
type dNSZoneNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all DNSZones in the indexer for a given namespace.
func (s dNSZoneNamespaceLister) List(selector labels.Selector) (ret []*v1.DNSZone, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.DNSZone))
	})
	return ret, err
}

// Get retrieves the DNSZone from the indexer for a given namespace and name.
func (s dNSZoneNamespaceLister) Get(name string) (*v1.DNSZone, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("dnszone"), name)
	}
	return obj.(*v1.DNSZone), nil
}
