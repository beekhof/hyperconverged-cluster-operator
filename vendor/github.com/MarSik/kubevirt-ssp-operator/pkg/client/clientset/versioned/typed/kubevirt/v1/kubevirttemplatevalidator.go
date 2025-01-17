/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/MarSik/kubevirt-ssp-operator/pkg/apis/kubevirt/v1"
	scheme "github.com/MarSik/kubevirt-ssp-operator/pkg/client/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// KubevirtTemplateValidatorsGetter has a method to return a KubevirtTemplateValidatorInterface.
// A group's client should implement this interface.
type KubevirtTemplateValidatorsGetter interface {
	KubevirtTemplateValidators(namespace string) KubevirtTemplateValidatorInterface
}

// KubevirtTemplateValidatorInterface has methods to work with KubevirtTemplateValidator resources.
type KubevirtTemplateValidatorInterface interface {
	Create(*v1.KubevirtTemplateValidator) (*v1.KubevirtTemplateValidator, error)
	Update(*v1.KubevirtTemplateValidator) (*v1.KubevirtTemplateValidator, error)
	Delete(name string, options *metav1.DeleteOptions) error
	DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error
	Get(name string, options metav1.GetOptions) (*v1.KubevirtTemplateValidator, error)
	List(opts metav1.ListOptions) (*v1.KubevirtTemplateValidatorList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.KubevirtTemplateValidator, err error)
	KubevirtTemplateValidatorExpansion
}

// kubevirtTemplateValidators implements KubevirtTemplateValidatorInterface
type kubevirtTemplateValidators struct {
	client rest.Interface
	ns     string
}

// newKubevirtTemplateValidators returns a KubevirtTemplateValidators
func newKubevirtTemplateValidators(c *KubevirtV1Client, namespace string) *kubevirtTemplateValidators {
	return &kubevirtTemplateValidators{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the kubevirtTemplateValidator, and returns the corresponding kubevirtTemplateValidator object, and an error if there is any.
func (c *kubevirtTemplateValidators) Get(name string, options metav1.GetOptions) (result *v1.KubevirtTemplateValidator, err error) {
	result = &v1.KubevirtTemplateValidator{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("kubevirttemplatevalidators").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of KubevirtTemplateValidators that match those selectors.
func (c *kubevirtTemplateValidators) List(opts metav1.ListOptions) (result *v1.KubevirtTemplateValidatorList, err error) {
	result = &v1.KubevirtTemplateValidatorList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("kubevirttemplatevalidators").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested kubevirtTemplateValidators.
func (c *kubevirtTemplateValidators) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("kubevirttemplatevalidators").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a kubevirtTemplateValidator and creates it.  Returns the server's representation of the kubevirtTemplateValidator, and an error, if there is any.
func (c *kubevirtTemplateValidators) Create(kubevirtTemplateValidator *v1.KubevirtTemplateValidator) (result *v1.KubevirtTemplateValidator, err error) {
	result = &v1.KubevirtTemplateValidator{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("kubevirttemplatevalidators").
		Body(kubevirtTemplateValidator).
		Do().
		Into(result)
	return
}

// Update takes the representation of a kubevirtTemplateValidator and updates it. Returns the server's representation of the kubevirtTemplateValidator, and an error, if there is any.
func (c *kubevirtTemplateValidators) Update(kubevirtTemplateValidator *v1.KubevirtTemplateValidator) (result *v1.KubevirtTemplateValidator, err error) {
	result = &v1.KubevirtTemplateValidator{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("kubevirttemplatevalidators").
		Name(kubevirtTemplateValidator.Name).
		Body(kubevirtTemplateValidator).
		Do().
		Into(result)
	return
}

// Delete takes name of the kubevirtTemplateValidator and deletes it. Returns an error if one occurs.
func (c *kubevirtTemplateValidators) Delete(name string, options *metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("kubevirttemplatevalidators").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *kubevirtTemplateValidators) DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("kubevirttemplatevalidators").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched kubevirtTemplateValidator.
func (c *kubevirtTemplateValidators) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.KubevirtTemplateValidator, err error) {
	result = &v1.KubevirtTemplateValidator{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("kubevirttemplatevalidators").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
