
/*
	Note: This file is autogenerated! Do not edit it manually!
	Edit client_kernel_template.go instead, and run
	hack/generate-client.sh afterwards.
*/

package client

import (
	api "github.com/weaveworks/ignite/pkg/apis/ignite/v1alpha1"
	meta "github.com/weaveworks/ignite/pkg/apis/meta/v1alpha1"
	"github.com/weaveworks/ignite/pkg/storage"
	"github.com/weaveworks/ignite/pkg/storage/filterer"
)

// KernelClient is an interface for accessing Kernel-specific API objects
type KernelClient interface {
	// Get returns the Kernel matching given UID from the storage
	Get(meta.UID) (*api.Kernel, error)
	// Set saves the given Kernel into persistent storage
	Set(*api.Kernel) error
	// Find returns the Kernel matching the given filter, filters can
	// match e.g. the Object's Name, UID or a specific property
	Find(filter filterer.BaseFilter) (*api.Kernel, error)
	// FindAll returns multiple Kernels matching the given filter, filters can
	// match e.g. the Object's Name, UID or a specific property
	FindAll(filter filterer.BaseFilter) ([]*api.Kernel, error)
	// Delete deletes the Kernel with the given UID from the storage
	Delete(uid meta.UID) error
	// List returns a list of all Kernels available
	List() ([]*api.Kernel, error)
}

// Kernels returns the KernelClient for the Client instance
func (c *Client) Kernels() KernelClient {
	if c.kernelClient == nil {
		c.kernelClient = newKernelClient(c.storage)
	}

	return c.kernelClient
}

// Kernels is a shorthand for accessing Kernels using the default client
func Kernels() KernelClient {
	return DefaultClient.Kernels()
}

// kernelClient is a struct implementing the KernelClient interface
// It uses a shared storage instance passed from the Client together with its own Filterer
type kernelClient struct {
	storage  storage.Storage
	filterer *filterer.Filterer
}

// newKernelClient builds the kernelClient struct using the storage implementation and a new Filterer
func newKernelClient(s storage.Storage) KernelClient {
	return &kernelClient{
		storage:  s,
		filterer: filterer.NewFilterer(s),
	}
}

// Find returns a single Kernel based on the given Filter
func (c *kernelClient) Find(filter filterer.BaseFilter) (*api.Kernel, error) {
	object, err := c.filterer.Find(api.KernelKind, filter)
	if err != nil {
		return nil, err
	}

	return object.(*api.Kernel), nil
}

// FindAll returns multiple Kernels based on the given Filter
func (c *kernelClient) FindAll(filter filterer.BaseFilter) ([]*api.Kernel, error) {
	matches, err := c.filterer.FindAll(api.KernelKind, filter)
	if err != nil {
		return nil, err
	}

	results := make([]*api.Kernel, 0, len(matches))
	for _, item := range matches {
		results = append(results, item.(*api.Kernel))
	}

	return results, nil
}

// Get returns the Kernel matching given UID from the storage
func (c *kernelClient) Get(uid meta.UID) (*api.Kernel, error) {
	object, err := c.storage.GetByID(meta.KindKernel, uid)
	if err != nil {
		return nil, err
	}

	return object.(*api.Kernel), nil
}

// Set saves the given Kernel into the persistent storage
func (c *kernelClient) Set(kernel *api.Kernel) error {
	return c.storage.Set(kernel)
}

// Delete deletes the Kernel from the storage
func (c *kernelClient) Delete(uid meta.UID) error {
	return c.storage.Delete(meta.KindKernel, uid)
}

// List returns a list of all Kernels available
func (c *kernelClient) List() ([]*api.Kernel, error) {
	list, err := c.storage.List(meta.KindKernel)
	if err != nil {
		return nil, err
	}

	results := make([]*api.Kernel, 0, len(list))
	for _, item := range list {
		results = append(results, item.(*api.Kernel))
	}

	return results, nil
}
