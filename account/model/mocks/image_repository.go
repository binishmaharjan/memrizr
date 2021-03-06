package mocks

import (
	"context"
	"github.com/stretchr/testify/mock"
	"mime/multipart"
)

// MockImageRepository is a mock type for model.ImageRepository
type MockImageRepository struct {
	mock.Mock
}

// UpdateProfile is mock of representations of ImageRepository Update Profile
func (m *MockImageRepository) UpdateProfile(ctx context.Context, objName string, imageFile multipart.File) (string, error) {
	ret := m.Called(ctx, objName, imageFile)

	var r0 string
	if ret.Get(0) != nil {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if ret.Get(1) != nil {
		r1 = ret.Get(1).(error)
	}

	return r0, r1
}

// DeleteProfile is mock of representations of ImageRepository DeleteProfile
func (m *MockImageRepository) DeleteProfile(ctx context.Context, objName string) error {
	ret := m.Called(ctx, objName)

	var r0 error
	if ret.Get(0) != nil {
		r0 = ret.Get(0).(error)
	}

	return r0
}
