// Code generated by mockery 2.9.0. DO NOT EDIT.

package mocks

import (
	context "context"

	grpc "google.golang.org/grpc"

	mock "github.com/stretchr/testify/mock"

	stencilv1beta1 "github.com/odpf/stencil/server/odpf/stencil/v1beta1"
)

// StencilServiceClient is an autogenerated mock type for the StencilServiceClient type
type StencilServiceClient struct {
	mock.Mock
}

// CreateNamespace provides a mock function with given fields: ctx, in, opts
func (_m *StencilServiceClient) CreateNamespace(ctx context.Context, in *stencilv1beta1.CreateNamespaceRequest, opts ...grpc.CallOption) (*stencilv1beta1.CreateNamespaceResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *stencilv1beta1.CreateNamespaceResponse
	if rf, ok := ret.Get(0).(func(context.Context, *stencilv1beta1.CreateNamespaceRequest, ...grpc.CallOption) *stencilv1beta1.CreateNamespaceResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*stencilv1beta1.CreateNamespaceResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *stencilv1beta1.CreateNamespaceRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateSchema provides a mock function with given fields: ctx, in, opts
func (_m *StencilServiceClient) CreateSchema(ctx context.Context, in *stencilv1beta1.CreateSchemaRequest, opts ...grpc.CallOption) (*stencilv1beta1.CreateSchemaResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *stencilv1beta1.CreateSchemaResponse
	if rf, ok := ret.Get(0).(func(context.Context, *stencilv1beta1.CreateSchemaRequest, ...grpc.CallOption) *stencilv1beta1.CreateSchemaResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*stencilv1beta1.CreateSchemaResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *stencilv1beta1.CreateSchemaRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteNamespace provides a mock function with given fields: ctx, in, opts
func (_m *StencilServiceClient) DeleteNamespace(ctx context.Context, in *stencilv1beta1.DeleteNamespaceRequest, opts ...grpc.CallOption) (*stencilv1beta1.DeleteNamespaceResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *stencilv1beta1.DeleteNamespaceResponse
	if rf, ok := ret.Get(0).(func(context.Context, *stencilv1beta1.DeleteNamespaceRequest, ...grpc.CallOption) *stencilv1beta1.DeleteNamespaceResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*stencilv1beta1.DeleteNamespaceResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *stencilv1beta1.DeleteNamespaceRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteSchema provides a mock function with given fields: ctx, in, opts
func (_m *StencilServiceClient) DeleteSchema(ctx context.Context, in *stencilv1beta1.DeleteSchemaRequest, opts ...grpc.CallOption) (*stencilv1beta1.DeleteSchemaResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *stencilv1beta1.DeleteSchemaResponse
	if rf, ok := ret.Get(0).(func(context.Context, *stencilv1beta1.DeleteSchemaRequest, ...grpc.CallOption) *stencilv1beta1.DeleteSchemaResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*stencilv1beta1.DeleteSchemaResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *stencilv1beta1.DeleteSchemaRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteVersion provides a mock function with given fields: ctx, in, opts
func (_m *StencilServiceClient) DeleteVersion(ctx context.Context, in *stencilv1beta1.DeleteVersionRequest, opts ...grpc.CallOption) (*stencilv1beta1.DeleteVersionResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *stencilv1beta1.DeleteVersionResponse
	if rf, ok := ret.Get(0).(func(context.Context, *stencilv1beta1.DeleteVersionRequest, ...grpc.CallOption) *stencilv1beta1.DeleteVersionResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*stencilv1beta1.DeleteVersionResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *stencilv1beta1.DeleteVersionRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetLatestSchema provides a mock function with given fields: ctx, in, opts
func (_m *StencilServiceClient) GetLatestSchema(ctx context.Context, in *stencilv1beta1.GetLatestSchemaRequest, opts ...grpc.CallOption) (*stencilv1beta1.GetLatestSchemaResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *stencilv1beta1.GetLatestSchemaResponse
	if rf, ok := ret.Get(0).(func(context.Context, *stencilv1beta1.GetLatestSchemaRequest, ...grpc.CallOption) *stencilv1beta1.GetLatestSchemaResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*stencilv1beta1.GetLatestSchemaResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *stencilv1beta1.GetLatestSchemaRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetNamespace provides a mock function with given fields: ctx, in, opts
func (_m *StencilServiceClient) GetNamespace(ctx context.Context, in *stencilv1beta1.GetNamespaceRequest, opts ...grpc.CallOption) (*stencilv1beta1.GetNamespaceResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *stencilv1beta1.GetNamespaceResponse
	if rf, ok := ret.Get(0).(func(context.Context, *stencilv1beta1.GetNamespaceRequest, ...grpc.CallOption) *stencilv1beta1.GetNamespaceResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*stencilv1beta1.GetNamespaceResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *stencilv1beta1.GetNamespaceRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetSchema provides a mock function with given fields: ctx, in, opts
func (_m *StencilServiceClient) GetSchema(ctx context.Context, in *stencilv1beta1.GetSchemaRequest, opts ...grpc.CallOption) (*stencilv1beta1.GetSchemaResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *stencilv1beta1.GetSchemaResponse
	if rf, ok := ret.Get(0).(func(context.Context, *stencilv1beta1.GetSchemaRequest, ...grpc.CallOption) *stencilv1beta1.GetSchemaResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*stencilv1beta1.GetSchemaResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *stencilv1beta1.GetSchemaRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetSchemaMetadata provides a mock function with given fields: ctx, in, opts
func (_m *StencilServiceClient) GetSchemaMetadata(ctx context.Context, in *stencilv1beta1.GetSchemaMetadataRequest, opts ...grpc.CallOption) (*stencilv1beta1.GetSchemaMetadataResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *stencilv1beta1.GetSchemaMetadataResponse
	if rf, ok := ret.Get(0).(func(context.Context, *stencilv1beta1.GetSchemaMetadataRequest, ...grpc.CallOption) *stencilv1beta1.GetSchemaMetadataResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*stencilv1beta1.GetSchemaMetadataResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *stencilv1beta1.GetSchemaMetadataRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListNamespaces provides a mock function with given fields: ctx, in, opts
func (_m *StencilServiceClient) ListNamespaces(ctx context.Context, in *stencilv1beta1.ListNamespacesRequest, opts ...grpc.CallOption) (*stencilv1beta1.ListNamespacesResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *stencilv1beta1.ListNamespacesResponse
	if rf, ok := ret.Get(0).(func(context.Context, *stencilv1beta1.ListNamespacesRequest, ...grpc.CallOption) *stencilv1beta1.ListNamespacesResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*stencilv1beta1.ListNamespacesResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *stencilv1beta1.ListNamespacesRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListSchemas provides a mock function with given fields: ctx, in, opts
func (_m *StencilServiceClient) ListSchemas(ctx context.Context, in *stencilv1beta1.ListSchemasRequest, opts ...grpc.CallOption) (*stencilv1beta1.ListSchemasResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *stencilv1beta1.ListSchemasResponse
	if rf, ok := ret.Get(0).(func(context.Context, *stencilv1beta1.ListSchemasRequest, ...grpc.CallOption) *stencilv1beta1.ListSchemasResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*stencilv1beta1.ListSchemasResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *stencilv1beta1.ListSchemasRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListVersions provides a mock function with given fields: ctx, in, opts
func (_m *StencilServiceClient) ListVersions(ctx context.Context, in *stencilv1beta1.ListVersionsRequest, opts ...grpc.CallOption) (*stencilv1beta1.ListVersionsResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *stencilv1beta1.ListVersionsResponse
	if rf, ok := ret.Get(0).(func(context.Context, *stencilv1beta1.ListVersionsRequest, ...grpc.CallOption) *stencilv1beta1.ListVersionsResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*stencilv1beta1.ListVersionsResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *stencilv1beta1.ListVersionsRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateNamespace provides a mock function with given fields: ctx, in, opts
func (_m *StencilServiceClient) UpdateNamespace(ctx context.Context, in *stencilv1beta1.UpdateNamespaceRequest, opts ...grpc.CallOption) (*stencilv1beta1.UpdateNamespaceResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *stencilv1beta1.UpdateNamespaceResponse
	if rf, ok := ret.Get(0).(func(context.Context, *stencilv1beta1.UpdateNamespaceRequest, ...grpc.CallOption) *stencilv1beta1.UpdateNamespaceResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*stencilv1beta1.UpdateNamespaceResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *stencilv1beta1.UpdateNamespaceRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateSchemaMetadata provides a mock function with given fields: ctx, in, opts
func (_m *StencilServiceClient) UpdateSchemaMetadata(ctx context.Context, in *stencilv1beta1.UpdateSchemaMetadataRequest, opts ...grpc.CallOption) (*stencilv1beta1.UpdateSchemaMetadataResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *stencilv1beta1.UpdateSchemaMetadataResponse
	if rf, ok := ret.Get(0).(func(context.Context, *stencilv1beta1.UpdateSchemaMetadataRequest, ...grpc.CallOption) *stencilv1beta1.UpdateSchemaMetadataResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*stencilv1beta1.UpdateSchemaMetadataResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *stencilv1beta1.UpdateSchemaMetadataRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
