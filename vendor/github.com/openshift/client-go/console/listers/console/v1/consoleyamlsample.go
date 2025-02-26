// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	consolev1 "github.com/openshift/api/console/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	listers "k8s.io/client-go/listers"
	cache "k8s.io/client-go/tools/cache"
)

// ConsoleYAMLSampleLister helps list ConsoleYAMLSamples.
// All objects returned here must be treated as read-only.
type ConsoleYAMLSampleLister interface {
	// List lists all ConsoleYAMLSamples in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*consolev1.ConsoleYAMLSample, err error)
	// Get retrieves the ConsoleYAMLSample from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*consolev1.ConsoleYAMLSample, error)
	ConsoleYAMLSampleListerExpansion
}

// consoleYAMLSampleLister implements the ConsoleYAMLSampleLister interface.
type consoleYAMLSampleLister struct {
	listers.ResourceIndexer[*consolev1.ConsoleYAMLSample]
}

// NewConsoleYAMLSampleLister returns a new ConsoleYAMLSampleLister.
func NewConsoleYAMLSampleLister(indexer cache.Indexer) ConsoleYAMLSampleLister {
	return &consoleYAMLSampleLister{listers.New[*consolev1.ConsoleYAMLSample](indexer, consolev1.Resource("consoleyamlsample"))}
}
