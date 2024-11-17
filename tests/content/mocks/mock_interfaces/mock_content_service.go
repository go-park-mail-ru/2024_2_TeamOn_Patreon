// Source: C:\Users\HYPERPC\GolandProjects\2024_2_TeamOn_Patreon\internal\content\controller\interfaces\content_behavior.go

package mock_interfaces

import (
	context "context"
	reflect "reflect"

	models "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/content/pkg/models"
	models0 "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/service/models"
	gomock "github.com/golang/mock/gomock"
)

// MockContentBehavior is a mock of ContentBehavior interface.
type MockContentBehavior struct {
	ctrl     *gomock.Controller
	recorder *MockContentBehaviorMockRecorder
}

// MockContentBehaviorMockRecorder is the mock recorder for MockContentBehavior.
type MockContentBehaviorMockRecorder struct {
	mock *MockContentBehavior
}

// NewMockContentBehavior creates a new mock instance.
func NewMockContentBehavior(ctrl *gomock.Controller) *MockContentBehavior {
	mock := &MockContentBehavior{ctrl: ctrl}
	mock.recorder = &MockContentBehaviorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockContentBehavior) EXPECT() *MockContentBehaviorMockRecorder {
	return m.recorder
}

// CreatePost mocks base method.
func (m *MockContentBehavior) CreatePost(ctx context.Context, userID, title, content string, layer int) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePost", ctx, userID, title, content, layer)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreatePost indicates an expected call of CreatePost.
func (mr *MockContentBehaviorMockRecorder) CreatePost(ctx, userID, title, content, layer interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePost", reflect.TypeOf((*MockContentBehavior)(nil).CreatePost), ctx, userID, title, content, layer)
}

// DeletePost mocks base method.
func (m *MockContentBehavior) DeletePost(ctx context.Context, userID, postID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeletePost", ctx, userID, postID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeletePost indicates an expected call of DeletePost.
func (mr *MockContentBehaviorMockRecorder) DeletePost(ctx, userID, postID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePost", reflect.TypeOf((*MockContentBehavior)(nil).DeletePost), ctx, userID, postID)
}

// GetAuthorPosts mocks base method.
func (m *MockContentBehavior) GetAuthorPosts(ctx context.Context, userID, authorID string, opt *models0.FeedOpt) ([]*models.Post, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAuthorPosts", ctx, userID, authorID, opt)
	ret0, _ := ret[0].([]*models.Post)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAuthorPosts indicates an expected call of GetAuthorPosts.
func (mr *MockContentBehaviorMockRecorder) GetAuthorPosts(ctx, userID, authorID, opt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAuthorPosts", reflect.TypeOf((*MockContentBehavior)(nil).GetAuthorPosts), ctx, userID, authorID, opt)
}

// GetFeedSubscription mocks base method.
func (m *MockContentBehavior) GetFeedSubscription(ctx context.Context, userID string, opt *models0.FeedOpt) ([]*models.Post, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFeedSubscription", ctx, userID, opt)
	ret0, _ := ret[0].([]*models.Post)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFeedSubscription indicates an expected call of GetFeedSubscription.
func (mr *MockContentBehaviorMockRecorder) GetFeedSubscription(ctx, userID, opt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFeedSubscription", reflect.TypeOf((*MockContentBehavior)(nil).GetFeedSubscription), ctx, userID, opt)
}

// GetFile mocks base method.
func (m *MockContentBehavior) GetFile(ctx context.Context, userID, postID string) ([]*models.Media, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFile", ctx, userID, postID)
	ret0, _ := ret[0].([]*models.Media)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFile indicates an expected call of GetFile.
func (mr *MockContentBehaviorMockRecorder) GetFile(ctx, userID, postID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFile", reflect.TypeOf((*MockContentBehavior)(nil).GetFile), ctx, userID, postID)
}

// GetPopularPosts mocks base method.
func (m *MockContentBehavior) GetPopularPosts(ctx context.Context, userID string, opt *models0.FeedOpt) ([]*models.Post, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPopularPosts", ctx, userID, opt)
	ret0, _ := ret[0].([]*models.Post)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPopularPosts indicates an expected call of GetPopularPosts.
func (mr *MockContentBehaviorMockRecorder) GetPopularPosts(ctx, userID, opt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPopularPosts", reflect.TypeOf((*MockContentBehavior)(nil).GetPopularPosts), ctx, userID, opt)
}

// LikePost mocks base method.
func (m *MockContentBehavior) LikePost(ctx context.Context, userID, postID string) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LikePost", ctx, userID, postID)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LikePost indicates an expected call of LikePost.
func (mr *MockContentBehaviorMockRecorder) LikePost(ctx, userID, postID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LikePost", reflect.TypeOf((*MockContentBehavior)(nil).LikePost), ctx, userID, postID)
}

// UpdatePost mocks base method.
func (m *MockContentBehavior) UpdatePost(ctx context.Context, userID, postID, title, about string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePost", ctx, userID, postID, title, about)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdatePost indicates an expected call of UpdatePost.
func (mr *MockContentBehaviorMockRecorder) UpdatePost(ctx, userID, postID, title, about interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePost", reflect.TypeOf((*MockContentBehavior)(nil).UpdatePost), ctx, userID, postID, title, about)
}
