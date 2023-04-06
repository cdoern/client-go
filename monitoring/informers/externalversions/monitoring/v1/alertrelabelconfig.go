// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	"context"
	time "time"

	monitoringv1 "github.com/openshift/api/monitoring/v1"
	versioned "github.com/openshift/client-go/monitoring/clientset/versioned"
	internalinterfaces "github.com/openshift/client-go/monitoring/informers/externalversions/internalinterfaces"
	v1 "github.com/openshift/client-go/monitoring/listers/monitoring/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// AlertRelabelConfigInformer provides access to a shared informer and lister for
// AlertRelabelConfigs.
type AlertRelabelConfigInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.AlertRelabelConfigLister
}

type alertRelabelConfigInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewAlertRelabelConfigInformer constructs a new informer for AlertRelabelConfig type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewAlertRelabelConfigInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredAlertRelabelConfigInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredAlertRelabelConfigInformer constructs a new informer for AlertRelabelConfig type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredAlertRelabelConfigInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.MonitoringV1().AlertRelabelConfigs(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.MonitoringV1().AlertRelabelConfigs(namespace).Watch(context.TODO(), options)
			},
		},
		&monitoringv1.AlertRelabelConfig{},
		resyncPeriod,
		indexers,
	)
}

func (f *alertRelabelConfigInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredAlertRelabelConfigInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *alertRelabelConfigInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&monitoringv1.AlertRelabelConfig{}, f.defaultInformer)
}

func (f *alertRelabelConfigInformer) Lister() v1.AlertRelabelConfigLister {
	return v1.NewAlertRelabelConfigLister(f.Informer().GetIndexer())
}
