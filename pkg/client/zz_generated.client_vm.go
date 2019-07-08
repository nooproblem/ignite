
/*
	Note: This file is autogenerated! Do not edit it manually!
	Edit client_vm_template.go instead, and run
	hack/generate-client.sh afterwards.
*/

package client

import (
	api "github.com/weaveworks/ignite/pkg/apis/ignite/v1alpha1"
	meta "github.com/weaveworks/ignite/pkg/apis/meta/v1alpha1"
	"github.com/weaveworks/ignite/pkg/storage"
	"github.com/weaveworks/ignite/pkg/storage/filterer"
)

// VMClient is an interface for accessing VM-specific API objects
type VMClient interface {
	// Get returns the VM matching given UID from the storage
	Get(meta.UID) (*api.VM, error)
	// Set saves the given VM into persistent storage
	Set(*api.VM) error
	// Find returns the VM matching the given filter, filters can
	// match e.g. the Object's Name, UID or a specific property
	Find(filter filterer.BaseFilter) (*api.VM, error)
	// FindAll returns multiple VMs matching the given filter, filters can
	// match e.g. the Object's Name, UID or a specific property
	FindAll(filter filterer.BaseFilter) ([]*api.VM, error)
	// Delete deletes the VM with the given UID from the storage
	Delete(uid meta.UID) error
	// List returns a list of all VMs available
	List() ([]*api.VM, error)
}

// VMs returns the VMClient for the Client instance
func (c *Client) VMs() VMClient {
	if c.vmClient == nil {
		c.vmClient = newVMClient(c.storage)
	}

	return c.vmClient
}

// VMs is a shorthand for accessing VMs using the default client
func VMs() VMClient {
	return DefaultClient.VMs()
}

// vmClient is a struct implementing the VMClient interface
// It uses a shared storage instance passed from the Client together with its own Filterer
type vmClient struct {
	storage  storage.Storage
	filterer *filterer.Filterer
}

// newVMClient builds the vmClient struct using the storage implementation and a new Filterer
func newVMClient(s storage.Storage) VMClient {
	return &vmClient{
		storage:  s,
		filterer: filterer.NewFilterer(s),
	}
}

// Find returns a single VM based on the given Filter
func (c *vmClient) Find(filter filterer.BaseFilter) (*api.VM, error) {
	object, err := c.filterer.Find(api.VMKind, filter)
	if err != nil {
		return nil, err
	}

	return object.(*api.VM), nil
}

// FindAll returns multiple VMs based on the given Filter
func (c *vmClient) FindAll(filter filterer.BaseFilter) ([]*api.VM, error) {
	matches, err := c.filterer.FindAll(api.VMKind, filter)
	if err != nil {
		return nil, err
	}

	results := make([]*api.VM, 0, len(matches))
	for _, item := range matches {
		results = append(results, item.(*api.VM))
	}

	return results, nil
}

// Get returns the VM matching given UID from the storage
func (c *vmClient) Get(uid meta.UID) (*api.VM, error) {
	object, err := c.storage.GetByID(meta.KindVM, uid)
	if err != nil {
		return nil, err
	}

	return object.(*api.VM), nil
}

// Set saves the given VM into the persistent storage
func (c *vmClient) Set(vm *api.VM) error {
	return c.storage.Set(vm)
}

// Delete deletes the VM from the storage
func (c *vmClient) Delete(uid meta.UID) error {
	return c.storage.Delete(meta.KindVM, uid)
}

// List returns a list of all VMs available
func (c *vmClient) List() ([]*api.VM, error) {
	list, err := c.storage.List(meta.KindVM)
	if err != nil {
		return nil, err
	}

	results := make([]*api.VM, 0, len(list))
	for _, item := range list {
		results = append(results, item.(*api.VM))
	}

	return results, nil
}
