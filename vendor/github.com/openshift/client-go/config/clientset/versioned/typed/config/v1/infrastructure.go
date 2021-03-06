// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	"time"

	v1 "github.com/openshift/api/config/v1"
	scheme "github.com/openshift/client-go/config/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// InfrastructuresGetter has a method to return a InfrastructureInterface.
// A group's client should implement this interface.
type InfrastructuresGetter interface {
	Infrastructures() InfrastructureInterface
}

// InfrastructureInterface has methods to work with Infrastructure resources.
type InfrastructureInterface interface {
	Create(*v1.Infrastructure) (*v1.Infrastructure, error)
	Update(*v1.Infrastructure) (*v1.Infrastructure, error)
	UpdateStatus(*v1.Infrastructure) (*v1.Infrastructure, error)
	Delete(name string, options *metav1.DeleteOptions) error
	DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error
	Get(name string, options metav1.GetOptions) (*v1.Infrastructure, error)
	List(opts metav1.ListOptions) (*v1.InfrastructureList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.Infrastructure, err error)
	InfrastructureExpansion
}

// infrastructures implements InfrastructureInterface
type infrastructures struct {
	client rest.Interface
}

// newInfrastructures returns a Infrastructures
func newInfrastructures(c *ConfigV1Client) *infrastructures {
	return &infrastructures{
		client: c.RESTClient(),
	}
}

// Get takes name of the infrastructure, and returns the corresponding infrastructure object, and an error if there is any.
func (c *infrastructures) Get(name string, options metav1.GetOptions) (result *v1.Infrastructure, err error) {
	result = &v1.Infrastructure{}
	err = c.client.Get().
		Resource("infrastructures").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Infrastructures that match those selectors.
func (c *infrastructures) List(opts metav1.ListOptions) (result *v1.InfrastructureList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.InfrastructureList{}
	err = c.client.Get().
		Resource("infrastructures").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested infrastructures.
func (c *infrastructures) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("infrastructures").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a infrastructure and creates it.  Returns the server's representation of the infrastructure, and an error, if there is any.
func (c *infrastructures) Create(infrastructure *v1.Infrastructure) (result *v1.Infrastructure, err error) {
	result = &v1.Infrastructure{}
	err = c.client.Post().
		Resource("infrastructures").
		Body(infrastructure).
		Do().
		Into(result)
	return
}

// Update takes the representation of a infrastructure and updates it. Returns the server's representation of the infrastructure, and an error, if there is any.
func (c *infrastructures) Update(infrastructure *v1.Infrastructure) (result *v1.Infrastructure, err error) {
	result = &v1.Infrastructure{}
	err = c.client.Put().
		Resource("infrastructures").
		Name(infrastructure.Name).
		Body(infrastructure).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *infrastructures) UpdateStatus(infrastructure *v1.Infrastructure) (result *v1.Infrastructure, err error) {
	result = &v1.Infrastructure{}
	err = c.client.Put().
		Resource("infrastructures").
		Name(infrastructure.Name).
		SubResource("status").
		Body(infrastructure).
		Do().
		Into(result)
	return
}

// Delete takes name of the infrastructure and deletes it. Returns an error if one occurs.
func (c *infrastructures) Delete(name string, options *metav1.DeleteOptions) error {
	return c.client.Delete().
		Resource("infrastructures").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *infrastructures) DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("infrastructures").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched infrastructure.
func (c *infrastructures) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.Infrastructure, err error) {
	result = &v1.Infrastructure{}
	err = c.client.Patch(pt).
		Resource("infrastructures").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
