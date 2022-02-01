// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package cim

import (
	"context"
	"io"
	"sync"
)

// Ensure, that ContextInformationManagerMock does implement ContextInformationManager.
// If this is not the case, regenerate this file with moq.
var _ ContextInformationManager = &ContextInformationManagerMock{}

// ContextInformationManagerMock is a mock implementation of ContextInformationManager.
//
// 	func TestSomethingThatUsesContextInformationManager(t *testing.T) {
//
// 		// make and configure a mocked ContextInformationManager
// 		mockedContextInformationManager := &ContextInformationManagerMock{
// 			CreateEntityFunc: func(ctx context.Context, tenant string, entityType string, entityID string, body io.Reader) (*CreateEntityResult, error) {
// 				panic("mock out the CreateEntity method")
// 			},
// 			QueryEntitiesFunc: func(ctx context.Context, tenant string, entityTypes []string, entityAttributes []string, query string) (*QueryEntitiesResult, error) {
// 				panic("mock out the QueryEntities method")
// 			},
// 			RetrieveEntityFunc: func(ctx context.Context, tenant string, entityID string) (Entity, error) {
// 				panic("mock out the RetrieveEntity method")
// 			},
// 			UpdateEntityAttributesFunc: func(ctx context.Context, tenant string, entityID string, body io.Reader) error {
// 				panic("mock out the UpdateEntityAttributes method")
// 			},
// 		}
//
// 		// use mockedContextInformationManager in code that requires ContextInformationManager
// 		// and then make assertions.
//
// 	}
type ContextInformationManagerMock struct {
	// CreateEntityFunc mocks the CreateEntity method.
	CreateEntityFunc func(ctx context.Context, tenant string, entityType string, entityID string, body io.Reader) (*CreateEntityResult, error)

	// QueryEntitiesFunc mocks the QueryEntities method.
	QueryEntitiesFunc func(ctx context.Context, tenant string, entityTypes []string, entityAttributes []string, query string) (*QueryEntitiesResult, error)

	// RetrieveEntityFunc mocks the RetrieveEntity method.
	RetrieveEntityFunc func(ctx context.Context, tenant string, entityID string) (Entity, error)

	// UpdateEntityAttributesFunc mocks the UpdateEntityAttributes method.
	UpdateEntityAttributesFunc func(ctx context.Context, tenant string, entityID string, body io.Reader) error

	// calls tracks calls to the methods.
	calls struct {
		// CreateEntity holds details about calls to the CreateEntity method.
		CreateEntity []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Tenant is the tenant argument value.
			Tenant string
			// EntityType is the entityType argument value.
			EntityType string
			// EntityID is the entityID argument value.
			EntityID string
			// Body is the body argument value.
			Body io.Reader
		}
		// QueryEntities holds details about calls to the QueryEntities method.
		QueryEntities []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Tenant is the tenant argument value.
			Tenant string
			// EntityTypes is the entityTypes argument value.
			EntityTypes []string
			// EntityAttributes is the entityAttributes argument value.
			EntityAttributes []string
			// Query is the query argument value.
			Query string
		}
		// RetrieveEntity holds details about calls to the RetrieveEntity method.
		RetrieveEntity []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Tenant is the tenant argument value.
			Tenant string
			// EntityID is the entityID argument value.
			EntityID string
		}
		// UpdateEntityAttributes holds details about calls to the UpdateEntityAttributes method.
		UpdateEntityAttributes []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Tenant is the tenant argument value.
			Tenant string
			// EntityID is the entityID argument value.
			EntityID string
			// Body is the body argument value.
			Body io.Reader
		}
	}
	lockCreateEntity           sync.RWMutex
	lockQueryEntities          sync.RWMutex
	lockRetrieveEntity         sync.RWMutex
	lockUpdateEntityAttributes sync.RWMutex
}

// CreateEntity calls CreateEntityFunc.
func (mock *ContextInformationManagerMock) CreateEntity(ctx context.Context, tenant string, entityType string, entityID string, body io.Reader) (*CreateEntityResult, error) {
	if mock.CreateEntityFunc == nil {
		panic("ContextInformationManagerMock.CreateEntityFunc: method is nil but ContextInformationManager.CreateEntity was just called")
	}
	callInfo := struct {
		Ctx        context.Context
		Tenant     string
		EntityType string
		EntityID   string
		Body       io.Reader
	}{
		Ctx:        ctx,
		Tenant:     tenant,
		EntityType: entityType,
		EntityID:   entityID,
		Body:       body,
	}
	mock.lockCreateEntity.Lock()
	mock.calls.CreateEntity = append(mock.calls.CreateEntity, callInfo)
	mock.lockCreateEntity.Unlock()
	return mock.CreateEntityFunc(ctx, tenant, entityType, entityID, body)
}

// CreateEntityCalls gets all the calls that were made to CreateEntity.
// Check the length with:
//     len(mockedContextInformationManager.CreateEntityCalls())
func (mock *ContextInformationManagerMock) CreateEntityCalls() []struct {
	Ctx        context.Context
	Tenant     string
	EntityType string
	EntityID   string
	Body       io.Reader
} {
	var calls []struct {
		Ctx        context.Context
		Tenant     string
		EntityType string
		EntityID   string
		Body       io.Reader
	}
	mock.lockCreateEntity.RLock()
	calls = mock.calls.CreateEntity
	mock.lockCreateEntity.RUnlock()
	return calls
}

// QueryEntities calls QueryEntitiesFunc.
func (mock *ContextInformationManagerMock) QueryEntities(ctx context.Context, tenant string, entityTypes []string, entityAttributes []string, query string) (*QueryEntitiesResult, error) {
	if mock.QueryEntitiesFunc == nil {
		panic("ContextInformationManagerMock.QueryEntitiesFunc: method is nil but ContextInformationManager.QueryEntities was just called")
	}
	callInfo := struct {
		Ctx              context.Context
		Tenant           string
		EntityTypes      []string
		EntityAttributes []string
		Query            string
	}{
		Ctx:              ctx,
		Tenant:           tenant,
		EntityTypes:      entityTypes,
		EntityAttributes: entityAttributes,
		Query:            query,
	}
	mock.lockQueryEntities.Lock()
	mock.calls.QueryEntities = append(mock.calls.QueryEntities, callInfo)
	mock.lockQueryEntities.Unlock()
	return mock.QueryEntitiesFunc(ctx, tenant, entityTypes, entityAttributes, query)
}

// QueryEntitiesCalls gets all the calls that were made to QueryEntities.
// Check the length with:
//     len(mockedContextInformationManager.QueryEntitiesCalls())
func (mock *ContextInformationManagerMock) QueryEntitiesCalls() []struct {
	Ctx              context.Context
	Tenant           string
	EntityTypes      []string
	EntityAttributes []string
	Query            string
} {
	var calls []struct {
		Ctx              context.Context
		Tenant           string
		EntityTypes      []string
		EntityAttributes []string
		Query            string
	}
	mock.lockQueryEntities.RLock()
	calls = mock.calls.QueryEntities
	mock.lockQueryEntities.RUnlock()
	return calls
}

// RetrieveEntity calls RetrieveEntityFunc.
func (mock *ContextInformationManagerMock) RetrieveEntity(ctx context.Context, tenant string, entityID string) (Entity, error) {
	if mock.RetrieveEntityFunc == nil {
		panic("ContextInformationManagerMock.RetrieveEntityFunc: method is nil but ContextInformationManager.RetrieveEntity was just called")
	}
	callInfo := struct {
		Ctx      context.Context
		Tenant   string
		EntityID string
	}{
		Ctx:      ctx,
		Tenant:   tenant,
		EntityID: entityID,
	}
	mock.lockRetrieveEntity.Lock()
	mock.calls.RetrieveEntity = append(mock.calls.RetrieveEntity, callInfo)
	mock.lockRetrieveEntity.Unlock()
	return mock.RetrieveEntityFunc(ctx, tenant, entityID)
}

// RetrieveEntityCalls gets all the calls that were made to RetrieveEntity.
// Check the length with:
//     len(mockedContextInformationManager.RetrieveEntityCalls())
func (mock *ContextInformationManagerMock) RetrieveEntityCalls() []struct {
	Ctx      context.Context
	Tenant   string
	EntityID string
} {
	var calls []struct {
		Ctx      context.Context
		Tenant   string
		EntityID string
	}
	mock.lockRetrieveEntity.RLock()
	calls = mock.calls.RetrieveEntity
	mock.lockRetrieveEntity.RUnlock()
	return calls
}

// UpdateEntityAttributes calls UpdateEntityAttributesFunc.
func (mock *ContextInformationManagerMock) UpdateEntityAttributes(ctx context.Context, tenant string, entityID string, body io.Reader) error {
	if mock.UpdateEntityAttributesFunc == nil {
		panic("ContextInformationManagerMock.UpdateEntityAttributesFunc: method is nil but ContextInformationManager.UpdateEntityAttributes was just called")
	}
	callInfo := struct {
		Ctx      context.Context
		Tenant   string
		EntityID string
		Body     io.Reader
	}{
		Ctx:      ctx,
		Tenant:   tenant,
		EntityID: entityID,
		Body:     body,
	}
	mock.lockUpdateEntityAttributes.Lock()
	mock.calls.UpdateEntityAttributes = append(mock.calls.UpdateEntityAttributes, callInfo)
	mock.lockUpdateEntityAttributes.Unlock()
	return mock.UpdateEntityAttributesFunc(ctx, tenant, entityID, body)
}

// UpdateEntityAttributesCalls gets all the calls that were made to UpdateEntityAttributes.
// Check the length with:
//     len(mockedContextInformationManager.UpdateEntityAttributesCalls())
func (mock *ContextInformationManagerMock) UpdateEntityAttributesCalls() []struct {
	Ctx      context.Context
	Tenant   string
	EntityID string
	Body     io.Reader
} {
	var calls []struct {
		Ctx      context.Context
		Tenant   string
		EntityID string
		Body     io.Reader
	}
	mock.lockUpdateEntityAttributes.RLock()
	calls = mock.calls.UpdateEntityAttributes
	mock.lockUpdateEntityAttributes.RUnlock()
	return calls
}
