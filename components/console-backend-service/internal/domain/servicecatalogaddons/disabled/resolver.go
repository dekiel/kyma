// Code generated by failery v1.0.0. DO NOT EDIT.

package disabled

import (
	context "context"

	gqlschema "github.com/kyma-project/kyma/components/console-backend-service/internal/gqlschema"
)

// Resolver is an autogenerated failing mock type for the Resolver type
type Resolver struct {
	err error
}

// NewResolver creates a new Resolver type instance
func NewResolver(err error) *Resolver {
	return &Resolver{err: err}
}

// AddAddonsConfigurationRepositories provides a failing mock function with given fields: ctx, name, namespace, repositories
func (_m *Resolver) AddAddonsConfigurationRepositories(ctx context.Context, name string, namespace string, repositories []*gqlschema.AddonsConfigurationRepositoryInput) (*gqlschema.AddonsConfiguration, error) {
	var r0 *gqlschema.AddonsConfiguration
	var r1 error
	r1 = _m.err

	return r0, r1
}

// AddAddonsConfigurationURLs provides a failing mock function with given fields: ctx, name, namespace, urls
func (_m *Resolver) AddAddonsConfigurationURLs(ctx context.Context, name string, namespace string, urls []string) (*gqlschema.AddonsConfiguration, error) {
	var r0 *gqlschema.AddonsConfiguration
	var r1 error
	r1 = _m.err

	return r0, r1
}

// AddClusterAddonsConfigurationRepositories provides a failing mock function with given fields: ctx, name, repositories
func (_m *Resolver) AddClusterAddonsConfigurationRepositories(ctx context.Context, name string, repositories []*gqlschema.AddonsConfigurationRepositoryInput) (*gqlschema.AddonsConfiguration, error) {
	var r0 *gqlschema.AddonsConfiguration
	var r1 error
	r1 = _m.err

	return r0, r1
}

// AddClusterAddonsConfigurationURLs provides a failing mock function with given fields: ctx, name, urls
func (_m *Resolver) AddClusterAddonsConfigurationURLs(ctx context.Context, name string, urls []string) (*gqlschema.AddonsConfiguration, error) {
	var r0 *gqlschema.AddonsConfiguration
	var r1 error
	r1 = _m.err

	return r0, r1
}

// AddonsConfigurationEventSubscription provides a failing mock function with given fields: ctx, namespace
func (_m *Resolver) AddonsConfigurationEventSubscription(ctx context.Context, namespace string) (<-chan *gqlschema.AddonsConfigurationEvent, error) {
	var r0 <-chan *gqlschema.AddonsConfigurationEvent
	var r1 error
	r1 = _m.err

	return r0, r1
}

// AddonsConfigurationsQuery provides a failing mock function with given fields: ctx, namespace, first, offset
func (_m *Resolver) AddonsConfigurationsQuery(ctx context.Context, namespace string, first *int, offset *int) ([]*gqlschema.AddonsConfiguration, error) {
	var r0 []*gqlschema.AddonsConfiguration
	var r1 error
	r1 = _m.err

	return r0, r1
}

// ClusterAddonsConfigurationEventSubscription provides a failing mock function with given fields: ctx
func (_m *Resolver) ClusterAddonsConfigurationEventSubscription(ctx context.Context) (<-chan *gqlschema.ClusterAddonsConfigurationEvent, error) {
	var r0 <-chan *gqlschema.ClusterAddonsConfigurationEvent
	var r1 error
	r1 = _m.err

	return r0, r1
}

// ClusterAddonsConfigurationsQuery provides a failing mock function with given fields: ctx, first, offset
func (_m *Resolver) ClusterAddonsConfigurationsQuery(ctx context.Context, first *int, offset *int) ([]*gqlschema.AddonsConfiguration, error) {
	var r0 []*gqlschema.AddonsConfiguration
	var r1 error
	r1 = _m.err

	return r0, r1
}

// CreateAddonsConfiguration provides a failing mock function with given fields: ctx, name, namespace, repositories, urls, labels
func (_m *Resolver) CreateAddonsConfiguration(ctx context.Context, name string, namespace string, repositories []*gqlschema.AddonsConfigurationRepositoryInput, urls []string, labels gqlschema.Labels) (*gqlschema.AddonsConfiguration, error) {
	var r0 *gqlschema.AddonsConfiguration
	var r1 error
	r1 = _m.err

	return r0, r1
}

// CreateClusterAddonsConfiguration provides a failing mock function with given fields: ctx, name, repositories, urls, labels
func (_m *Resolver) CreateClusterAddonsConfiguration(ctx context.Context, name string, repositories []*gqlschema.AddonsConfigurationRepositoryInput, urls []string, labels gqlschema.Labels) (*gqlschema.AddonsConfiguration, error) {
	var r0 *gqlschema.AddonsConfiguration
	var r1 error
	r1 = _m.err

	return r0, r1
}

// CreateServiceBindingUsageMutation provides a failing mock function with given fields: ctx, namespace, input
func (_m *Resolver) CreateServiceBindingUsageMutation(ctx context.Context, namespace string, input *gqlschema.CreateServiceBindingUsageInput) (*gqlschema.ServiceBindingUsage, error) {
	var r0 *gqlschema.ServiceBindingUsage
	var r1 error
	r1 = _m.err

	return r0, r1
}

// DeleteAddonsConfiguration provides a failing mock function with given fields: ctx, name, namespace
func (_m *Resolver) DeleteAddonsConfiguration(ctx context.Context, name string, namespace string) (*gqlschema.AddonsConfiguration, error) {
	var r0 *gqlschema.AddonsConfiguration
	var r1 error
	r1 = _m.err

	return r0, r1
}

// DeleteClusterAddonsConfiguration provides a failing mock function with given fields: ctx, name
func (_m *Resolver) DeleteClusterAddonsConfiguration(ctx context.Context, name string) (*gqlschema.AddonsConfiguration, error) {
	var r0 *gqlschema.AddonsConfiguration
	var r1 error
	r1 = _m.err

	return r0, r1
}

// DeleteServiceBindingUsageMutation provides a failing mock function with given fields: ctx, serviceBindingUsageName, namespace
func (_m *Resolver) DeleteServiceBindingUsageMutation(ctx context.Context, serviceBindingUsageName string, namespace string) (*gqlschema.DeleteServiceBindingUsageOutput, error) {
	var r0 *gqlschema.DeleteServiceBindingUsageOutput
	var r1 error
	r1 = _m.err

	return r0, r1
}

// DeleteServiceBindingUsagesMutation provides a failing mock function with given fields: ctx, serviceBindingUsageNames, namespace
func (_m *Resolver) DeleteServiceBindingUsagesMutation(ctx context.Context, serviceBindingUsageNames []string, namespace string) ([]*gqlschema.DeleteServiceBindingUsageOutput, error) {
	var r0 []*gqlschema.DeleteServiceBindingUsageOutput
	var r1 error
	r1 = _m.err

	return r0, r1
}

// ListBindableResources provides a failing mock function with given fields: ctx, namespace
func (_m *Resolver) ListBindableResources(ctx context.Context, namespace string) ([]*gqlschema.BindableResourcesOutputItem, error) {
	var r0 []*gqlschema.BindableResourcesOutputItem
	var r1 error
	r1 = _m.err

	return r0, r1
}

// ListUsageKinds provides a failing mock function with given fields: ctx, first, offset
func (_m *Resolver) ListUsageKinds(ctx context.Context, first *int, offset *int) ([]*gqlschema.UsageKind, error) {
	var r0 []*gqlschema.UsageKind
	var r1 error
	r1 = _m.err

	return r0, r1
}

// RemoveAddonsConfigurationRepositories provides a failing mock function with given fields: ctx, name, namespace, urls
func (_m *Resolver) RemoveAddonsConfigurationRepositories(ctx context.Context, name string, namespace string, urls []string) (*gqlschema.AddonsConfiguration, error) {
	var r0 *gqlschema.AddonsConfiguration
	var r1 error
	r1 = _m.err

	return r0, r1
}

// RemoveAddonsConfigurationURLs provides a failing mock function with given fields: ctx, name, namespace, urls
func (_m *Resolver) RemoveAddonsConfigurationURLs(ctx context.Context, name string, namespace string, urls []string) (*gqlschema.AddonsConfiguration, error) {
	var r0 *gqlschema.AddonsConfiguration
	var r1 error
	r1 = _m.err

	return r0, r1
}

// RemoveClusterAddonsConfigurationRepositories provides a failing mock function with given fields: ctx, name, urls
func (_m *Resolver) RemoveClusterAddonsConfigurationRepositories(ctx context.Context, name string, urls []string) (*gqlschema.AddonsConfiguration, error) {
	var r0 *gqlschema.AddonsConfiguration
	var r1 error
	r1 = _m.err

	return r0, r1
}

// RemoveClusterAddonsConfigurationURLs provides a failing mock function with given fields: ctx, name, urls
func (_m *Resolver) RemoveClusterAddonsConfigurationURLs(ctx context.Context, name string, urls []string) (*gqlschema.AddonsConfiguration, error) {
	var r0 *gqlschema.AddonsConfiguration
	var r1 error
	r1 = _m.err

	return r0, r1
}

// ResyncAddonsConfiguration provides a failing mock function with given fields: ctx, name, namespace
func (_m *Resolver) ResyncAddonsConfiguration(ctx context.Context, name string, namespace string) (*gqlschema.AddonsConfiguration, error) {
	var r0 *gqlschema.AddonsConfiguration
	var r1 error
	r1 = _m.err

	return r0, r1
}

// ResyncClusterAddonsConfiguration provides a failing mock function with given fields: ctx, name
func (_m *Resolver) ResyncClusterAddonsConfiguration(ctx context.Context, name string) (*gqlschema.AddonsConfiguration, error) {
	var r0 *gqlschema.AddonsConfiguration
	var r1 error
	r1 = _m.err

	return r0, r1
}

// ServiceBindingUsageEventSubscription provides a failing mock function with given fields: ctx, namespace, resourceKind, resourceName
func (_m *Resolver) ServiceBindingUsageEventSubscription(ctx context.Context, namespace string, resourceKind *string, resourceName *string) (<-chan *gqlschema.ServiceBindingUsageEvent, error) {
	var r0 <-chan *gqlschema.ServiceBindingUsageEvent
	var r1 error
	r1 = _m.err

	return r0, r1
}

// ServiceBindingUsageQuery provides a failing mock function with given fields: ctx, name, namespace
func (_m *Resolver) ServiceBindingUsageQuery(ctx context.Context, name string, namespace string) (*gqlschema.ServiceBindingUsage, error) {
	var r0 *gqlschema.ServiceBindingUsage
	var r1 error
	r1 = _m.err

	return r0, r1
}

// ServiceBindingUsagesOfInstanceQuery provides a failing mock function with given fields: ctx, instanceName, env
func (_m *Resolver) ServiceBindingUsagesOfInstanceQuery(ctx context.Context, instanceName string, env string) ([]*gqlschema.ServiceBindingUsage, error) {
	var r0 []*gqlschema.ServiceBindingUsage
	var r1 error
	r1 = _m.err

	return r0, r1
}

// ServiceBindingUsagesQuery provides a failing mock function with given fields: ctx, namespace, resourceKind, resourceName
func (_m *Resolver) ServiceBindingUsagesQuery(ctx context.Context, namespace string, resourceKind *string, resourceName *string) ([]*gqlschema.ServiceBindingUsage, error) {
	var r0 []*gqlschema.ServiceBindingUsage
	var r1 error
	r1 = _m.err

	return r0, r1
}

// UpdateAddonsConfiguration provides a failing mock function with given fields: ctx, name, namespace, repositories, urls, labels
func (_m *Resolver) UpdateAddonsConfiguration(ctx context.Context, name string, namespace string, repositories []*gqlschema.AddonsConfigurationRepositoryInput, urls []string, labels gqlschema.Labels) (*gqlschema.AddonsConfiguration, error) {
	var r0 *gqlschema.AddonsConfiguration
	var r1 error
	r1 = _m.err

	return r0, r1
}

// UpdateClusterAddonsConfiguration provides a failing mock function with given fields: ctx, name, repositories, urls, labels
func (_m *Resolver) UpdateClusterAddonsConfiguration(ctx context.Context, name string, repositories []*gqlschema.AddonsConfigurationRepositoryInput, urls []string, labels gqlschema.Labels) (*gqlschema.AddonsConfiguration, error) {
	var r0 *gqlschema.AddonsConfiguration
	var r1 error
	r1 = _m.err

	return r0, r1
}
