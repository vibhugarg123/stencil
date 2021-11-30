// Code generated by mockery 2.9.0. DO NOT EDIT.

package mocks

import (
	context "context"

	stencilv1beta1 "github.com/odpf/stencil/server/odpf/stencil/v1beta1"
	mock "github.com/stretchr/testify/mock"
)

// StencilServiceServer is an autogenerated mock type for the StencilServiceServer type
type StencilServiceServer struct {
	mock.Mock
}

// CreateNamespace provides a mock function with given fields: _a0, _a1
func (_m *StencilServiceServer) CreateNamespace(_a0 context.Context, _a1 *stencilv1beta1.CreateNamespaceRequest) (*stencilv1beta1.CreateNamespaceResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *stencilv1beta1.CreateNamespaceResponse
	if rf, ok := ret.Get(0).(func(context.Context, *stencilv1beta1.CreateNamespaceRequest) *stencilv1beta1.CreateNamespaceResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*stencilv1beta1.CreateNamespaceResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *stencilv1beta1.CreateNamespaceRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateSchema provides a mock function with given fields: _a0, _a1
func (_m *StencilServiceServer) CreateSchema(_a0 context.Context, _a1 *stencilv1beta1.CreateSchemaRequest) (*stencilv1beta1.CreateSchemaResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *stencilv1beta1.CreateSchemaResponse
	if rf, ok := ret.Get(0).(func(context.Context, *stencilv1beta1.CreateSchemaRequest) *stencilv1beta1.CreateSchemaResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*stencilv1beta1.CreateSchemaResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *stencilv1beta1.CreateSchemaRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteNamespace provides a mock function with given fields: _a0, _a1
func (_m *StencilServiceServer) DeleteNamespace(_a0 context.Context, _a1 *stencilv1beta1.DeleteNamespaceRequest) (*stencilv1beta1.DeleteNamespaceResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *stencilv1beta1.DeleteNamespaceResponse
	if rf, ok := ret.Get(0).(func(context.Context, *stencilv1beta1.DeleteNamespaceRequest) *stencilv1beta1.DeleteNamespaceResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*stencilv1beta1.DeleteNamespaceResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *stencilv1beta1.DeleteNamespaceRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteSchema provides a mock function with given fields: _a0, _a1
func (_m *StencilServiceServer) DeleteSchema(_a0 context.Context, _a1 *stencilv1beta1.DeleteSchemaRequest) (*stencilv1beta1.DeleteSchemaResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *stencilv1beta1.DeleteSchemaResponse
	if rf, ok := ret.Get(0).(func(context.Context, *stencilv1beta1.DeleteSchemaRequest) *stencilv1beta1.DeleteSchemaResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*stencilv1beta1.DeleteSchemaResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *stencilv1beta1.DeleteSchemaRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteVersion provides a mock function with given fields: _a0, _a1
func (_m *StencilServiceServer) DeleteVersion(_a0 context.Context, _a1 *stencilv1beta1.DeleteVersionRequest) (*stencilv1beta1.DeleteVersionResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *stencilv1beta1.DeleteVersionResponse
	if rf, ok := ret.Get(0).(func(context.Context, *stencilv1beta1.DeleteVersionRequest) *stencilv1beta1.DeleteVersionResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*stencilv1beta1.DeleteVersionResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *stencilv1beta1.DeleteVersionRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetLatestSchema provides a mock function with given fields: _a0, _a1
func (_m *StencilServiceServer) GetLatestSchema(_a0 context.Context, _a1 *stencilv1beta1.GetLatestSchemaRequest) (*stencilv1beta1.GetLatestSchemaResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *stencilv1beta1.GetLatestSchemaResponse
	if rf, ok := ret.Get(0).(func(context.Context, *stencilv1beta1.GetLatestSchemaRequest) *stencilv1beta1.GetLatestSchemaResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*stencilv1beta1.GetLatestSchemaResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *stencilv1beta1.GetLatestSchemaRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetNamespace provides a mock function with given fields: _a0, _a1
func (_m *StencilServiceServer) GetNamespace(_a0 context.Context, _a1 *stencilv1beta1.GetNamespaceRequest) (*stencilv1beta1.GetNamespaceResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *stencilv1beta1.GetNamespaceResponse
	if rf, ok := ret.Get(0).(func(context.Context, *stencilv1beta1.GetNamespaceRequest) *stencilv1beta1.GetNamespaceResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*stencilv1beta1.GetNamespaceResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *stencilv1beta1.GetNamespaceRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetSchema provides a mock function with given fields: _a0, _a1
func (_m *StencilServiceServer) GetSchema(_a0 context.Context, _a1 *stencilv1beta1.GetSchemaRequest) (*stencilv1beta1.GetSchemaResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *stencilv1beta1.GetSchemaResponse
	if rf, ok := ret.Get(0).(func(context.Context, *stencilv1beta1.GetSchemaRequest) *stencilv1beta1.GetSchemaResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*stencilv1beta1.GetSchemaResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *stencilv1beta1.GetSchemaRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetSchemaMetadata provides a mock function with given fields: _a0, _a1
func (_m *StencilServiceServer) GetSchemaMetadata(_a0 context.Context, _a1 *stencilv1beta1.GetSchemaMetadataRequest) (*stencilv1beta1.GetSchemaMetadataResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *stencilv1beta1.GetSchemaMetadataResponse
	if rf, ok := ret.Get(0).(func(context.Context, *stencilv1beta1.GetSchemaMetadataRequest) *stencilv1beta1.GetSchemaMetadataResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*stencilv1beta1.GetSchemaMetadataResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *stencilv1beta1.GetSchemaMetadataRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListNamespaces provides a mock function with given fields: _a0, _a1
func (_m *StencilServiceServer) ListNamespaces(_a0 context.Context, _a1 *stencilv1beta1.ListNamespacesRequest) (*stencilv1beta1.ListNamespacesResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *stencilv1beta1.ListNamespacesResponse
	if rf, ok := ret.Get(0).(func(context.Context, *stencilv1beta1.ListNamespacesRequest) *stencilv1beta1.ListNamespacesResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*stencilv1beta1.ListNamespacesResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *stencilv1beta1.ListNamespacesRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListSchemas provides a mock function with given fields: _a0, _a1
func (_m *StencilServiceServer) ListSchemas(_a0 context.Context, _a1 *stencilv1beta1.ListSchemasRequest) (*stencilv1beta1.ListSchemasResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *stencilv1beta1.ListSchemasResponse
	if rf, ok := ret.Get(0).(func(context.Context, *stencilv1beta1.ListSchemasRequest) *stencilv1beta1.ListSchemasResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*stencilv1beta1.ListSchemasResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *stencilv1beta1.ListSchemasRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListVersions provides a mock function with given fields: _a0, _a1
func (_m *StencilServiceServer) ListVersions(_a0 context.Context, _a1 *stencilv1beta1.ListVersionsRequest) (*stencilv1beta1.ListVersionsResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *stencilv1beta1.ListVersionsResponse
	if rf, ok := ret.Get(0).(func(context.Context, *stencilv1beta1.ListVersionsRequest) *stencilv1beta1.ListVersionsResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*stencilv1beta1.ListVersionsResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *stencilv1beta1.ListVersionsRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateNamespace provides a mock function with given fields: _a0, _a1
func (_m *StencilServiceServer) UpdateNamespace(_a0 context.Context, _a1 *stencilv1beta1.UpdateNamespaceRequest) (*stencilv1beta1.UpdateNamespaceResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *stencilv1beta1.UpdateNamespaceResponse
	if rf, ok := ret.Get(0).(func(context.Context, *stencilv1beta1.UpdateNamespaceRequest) *stencilv1beta1.UpdateNamespaceResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*stencilv1beta1.UpdateNamespaceResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *stencilv1beta1.UpdateNamespaceRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateSchemaMetadata provides a mock function with given fields: _a0, _a1
func (_m *StencilServiceServer) UpdateSchemaMetadata(_a0 context.Context, _a1 *stencilv1beta1.UpdateSchemaMetadataRequest) (*stencilv1beta1.UpdateSchemaMetadataResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *stencilv1beta1.UpdateSchemaMetadataResponse
	if rf, ok := ret.Get(0).(func(context.Context, *stencilv1beta1.UpdateSchemaMetadataRequest) *stencilv1beta1.UpdateSchemaMetadataResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*stencilv1beta1.UpdateSchemaMetadataResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *stencilv1beta1.UpdateSchemaMetadataRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// mustEmbedUnimplementedStencilServiceServer provides a mock function with given fields:
func (_m *StencilServiceServer) mustEmbedUnimplementedStencilServiceServer() {
	_m.Called()
}
