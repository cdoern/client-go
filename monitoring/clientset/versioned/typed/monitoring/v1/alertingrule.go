// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	"context"
	json "encoding/json"
	"fmt"
	"time"

	v1 "github.com/openshift/api/monitoring/v1"
	monitoringv1 "github.com/openshift/client-go/monitoring/applyconfigurations/monitoring/v1"
	scheme "github.com/openshift/client-go/monitoring/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// AlertingRulesGetter has a method to return a AlertingRuleInterface.
// A group's client should implement this interface.
type AlertingRulesGetter interface {
	AlertingRules(namespace string) AlertingRuleInterface
}

// AlertingRuleInterface has methods to work with AlertingRule resources.
type AlertingRuleInterface interface {
	Create(ctx context.Context, alertingRule *v1.AlertingRule, opts metav1.CreateOptions) (*v1.AlertingRule, error)
	Update(ctx context.Context, alertingRule *v1.AlertingRule, opts metav1.UpdateOptions) (*v1.AlertingRule, error)
	UpdateStatus(ctx context.Context, alertingRule *v1.AlertingRule, opts metav1.UpdateOptions) (*v1.AlertingRule, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.AlertingRule, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.AlertingRuleList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.AlertingRule, err error)
	Apply(ctx context.Context, alertingRule *monitoringv1.AlertingRuleApplyConfiguration, opts metav1.ApplyOptions) (result *v1.AlertingRule, err error)
	ApplyStatus(ctx context.Context, alertingRule *monitoringv1.AlertingRuleApplyConfiguration, opts metav1.ApplyOptions) (result *v1.AlertingRule, err error)
	AlertingRuleExpansion
}

// alertingRules implements AlertingRuleInterface
type alertingRules struct {
	client rest.Interface
	ns     string
}

// newAlertingRules returns a AlertingRules
func newAlertingRules(c *MonitoringV1Client, namespace string) *alertingRules {
	return &alertingRules{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the alertingRule, and returns the corresponding alertingRule object, and an error if there is any.
func (c *alertingRules) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.AlertingRule, err error) {
	result = &v1.AlertingRule{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("alertingrules").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AlertingRules that match those selectors.
func (c *alertingRules) List(ctx context.Context, opts metav1.ListOptions) (result *v1.AlertingRuleList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.AlertingRuleList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("alertingrules").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested alertingRules.
func (c *alertingRules) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("alertingrules").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a alertingRule and creates it.  Returns the server's representation of the alertingRule, and an error, if there is any.
func (c *alertingRules) Create(ctx context.Context, alertingRule *v1.AlertingRule, opts metav1.CreateOptions) (result *v1.AlertingRule, err error) {
	result = &v1.AlertingRule{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("alertingrules").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(alertingRule).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a alertingRule and updates it. Returns the server's representation of the alertingRule, and an error, if there is any.
func (c *alertingRules) Update(ctx context.Context, alertingRule *v1.AlertingRule, opts metav1.UpdateOptions) (result *v1.AlertingRule, err error) {
	result = &v1.AlertingRule{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("alertingrules").
		Name(alertingRule.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(alertingRule).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *alertingRules) UpdateStatus(ctx context.Context, alertingRule *v1.AlertingRule, opts metav1.UpdateOptions) (result *v1.AlertingRule, err error) {
	result = &v1.AlertingRule{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("alertingrules").
		Name(alertingRule.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(alertingRule).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the alertingRule and deletes it. Returns an error if one occurs.
func (c *alertingRules) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("alertingrules").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *alertingRules) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("alertingrules").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched alertingRule.
func (c *alertingRules) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.AlertingRule, err error) {
	result = &v1.AlertingRule{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("alertingrules").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

// Apply takes the given apply declarative configuration, applies it and returns the applied alertingRule.
func (c *alertingRules) Apply(ctx context.Context, alertingRule *monitoringv1.AlertingRuleApplyConfiguration, opts metav1.ApplyOptions) (result *v1.AlertingRule, err error) {
	if alertingRule == nil {
		return nil, fmt.Errorf("alertingRule provided to Apply must not be nil")
	}
	patchOpts := opts.ToPatchOptions()
	data, err := json.Marshal(alertingRule)
	if err != nil {
		return nil, err
	}
	name := alertingRule.Name
	if name == nil {
		return nil, fmt.Errorf("alertingRule.Name must be provided to Apply")
	}
	result = &v1.AlertingRule{}
	err = c.client.Patch(types.ApplyPatchType).
		Namespace(c.ns).
		Resource("alertingrules").
		Name(*name).
		VersionedParams(&patchOpts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

// ApplyStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
func (c *alertingRules) ApplyStatus(ctx context.Context, alertingRule *monitoringv1.AlertingRuleApplyConfiguration, opts metav1.ApplyOptions) (result *v1.AlertingRule, err error) {
	if alertingRule == nil {
		return nil, fmt.Errorf("alertingRule provided to Apply must not be nil")
	}
	patchOpts := opts.ToPatchOptions()
	data, err := json.Marshal(alertingRule)
	if err != nil {
		return nil, err
	}

	name := alertingRule.Name
	if name == nil {
		return nil, fmt.Errorf("alertingRule.Name must be provided to Apply")
	}

	result = &v1.AlertingRule{}
	err = c.client.Patch(types.ApplyPatchType).
		Namespace(c.ns).
		Resource("alertingrules").
		Name(*name).
		SubResource("status").
		VersionedParams(&patchOpts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}