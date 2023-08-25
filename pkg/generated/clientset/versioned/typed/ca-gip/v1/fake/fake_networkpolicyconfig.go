// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	cagipv1 "github.com/ca-gip/kubi/pkg/apis/ca-gip/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeNetworkPolicyConfigs implements NetworkPolicyConfigInterface
type FakeNetworkPolicyConfigs struct {
	Fake *FakeCagipV1
}

var networkpolicyconfigsResource = schema.GroupVersionResource{Group: "cagip.github.com", Version: "v1", Resource: "networkpolicyconfigs"}

var networkpolicyconfigsKind = schema.GroupVersionKind{Group: "cagip.github.com", Version: "v1", Kind: "NetworkPolicyConfig"}

// Get takes name of the networkPolicyConfig, and returns the corresponding networkPolicyConfig object, and an error if there is any.
func (c *FakeNetworkPolicyConfigs) Get(ctx context.Context, name string, options v1.GetOptions) (result *cagipv1.NetworkPolicyConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(networkpolicyconfigsResource, name), &cagipv1.NetworkPolicyConfig{})
	if obj == nil {
		return nil, err
	}
	return obj.(*cagipv1.NetworkPolicyConfig), err
}

// List takes label and field selectors, and returns the list of NetworkPolicyConfigs that match those selectors.
func (c *FakeNetworkPolicyConfigs) List(ctx context.Context, opts v1.ListOptions) (result *cagipv1.NetworkPolicyConfigList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(networkpolicyconfigsResource, networkpolicyconfigsKind, opts), &cagipv1.NetworkPolicyConfigList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &cagipv1.NetworkPolicyConfigList{ListMeta: obj.(*cagipv1.NetworkPolicyConfigList).ListMeta}
	for _, item := range obj.(*cagipv1.NetworkPolicyConfigList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested networkPolicyConfigs.
func (c *FakeNetworkPolicyConfigs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(networkpolicyconfigsResource, opts))
}

// Create takes the representation of a networkPolicyConfig and creates it.  Returns the server's representation of the networkPolicyConfig, and an error, if there is any.
func (c *FakeNetworkPolicyConfigs) Create(ctx context.Context, networkPolicyConfig *cagipv1.NetworkPolicyConfig, opts v1.CreateOptions) (result *cagipv1.NetworkPolicyConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(networkpolicyconfigsResource, networkPolicyConfig), &cagipv1.NetworkPolicyConfig{})
	if obj == nil {
		return nil, err
	}
	return obj.(*cagipv1.NetworkPolicyConfig), err
}

// Update takes the representation of a networkPolicyConfig and updates it. Returns the server's representation of the networkPolicyConfig, and an error, if there is any.
func (c *FakeNetworkPolicyConfigs) Update(ctx context.Context, networkPolicyConfig *cagipv1.NetworkPolicyConfig, opts v1.UpdateOptions) (result *cagipv1.NetworkPolicyConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(networkpolicyconfigsResource, networkPolicyConfig), &cagipv1.NetworkPolicyConfig{})
	if obj == nil {
		return nil, err
	}
	return obj.(*cagipv1.NetworkPolicyConfig), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeNetworkPolicyConfigs) UpdateStatus(ctx context.Context, networkPolicyConfig *cagipv1.NetworkPolicyConfig, opts v1.UpdateOptions) (*cagipv1.NetworkPolicyConfig, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(networkpolicyconfigsResource, "status", networkPolicyConfig), &cagipv1.NetworkPolicyConfig{})
	if obj == nil {
		return nil, err
	}
	return obj.(*cagipv1.NetworkPolicyConfig), err
}

// Delete takes name of the networkPolicyConfig and deletes it. Returns an error if one occurs.
func (c *FakeNetworkPolicyConfigs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(networkpolicyconfigsResource, name), &cagipv1.NetworkPolicyConfig{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeNetworkPolicyConfigs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(networkpolicyconfigsResource, listOpts)

	_, err := c.Fake.Invokes(action, &cagipv1.NetworkPolicyConfigList{})
	return err
}

// Patch applies the patch and returns the patched networkPolicyConfig.
func (c *FakeNetworkPolicyConfigs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *cagipv1.NetworkPolicyConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(networkpolicyconfigsResource, name, pt, data, subresources...), &cagipv1.NetworkPolicyConfig{})
	if obj == nil {
		return nil, err
	}
	return obj.(*cagipv1.NetworkPolicyConfig), err
}