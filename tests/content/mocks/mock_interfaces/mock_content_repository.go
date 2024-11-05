// Source: C:\Users\HYPERPC\GolandProjects\2024_2_TeamOn_Patreon\internal\content\service\interfaces\content_repository.go

package mock_interfaces

import (
	context "context"
	reflect "reflect"

	models "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/content/pkg/models"
	uuid "github.com/gofrs/uuid"
	gomock "github.com/golang/mock/gomock"
)

// MockContentRepository is a mock of ContentRepository interface.
type MockContentRepository struct {
	ctrl     *gomock.Controller
	recorder *MockContentRepositoryMockRecorder
}

// MockContentRepositoryMockRecorder is the mock recorder for MockContentRepository.
type MockContentRepositoryMockRecorder struct {
	mock *MockContentRepository
}

// NewMockContentRepository creates a new mock instance.
func NewMockContentRepository(ctrl *gomock.Controller) *MockContentRepository {
	mock := &MockContentRepository{ctrl: ctrl}
	mock.recorder = &MockContentRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockContentRepository) EXPECT() *MockContentRepositoryMockRecorder {
	return m.recorder
}

// CheckCustomLayer mocks base method.
func (m *MockContentRepository) CheckCustomLayer(ctx context.Context, authorId uuid.UUID, layer int) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckCustomLayer", ctx, authorId, layer)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckCustomLayer indicates an expected call of CheckCustomLayer.
func (mr *MockContentRepositoryMockRecorder) CheckCustomLayer(ctx, authorId, layer interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckCustomLayer", reflect.TypeOf((*MockContentRepository)(nil).CheckCustomLayer), ctx, authorId, layer)
}

// DeleteLikePost mocks base method.
func (m *MockContentRepository) DeleteLikePost(ctx context.Context, userId, postID uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteLikePost", ctx, userId, postID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteLikePost indicates an expected call of DeleteLikePost.
func (mr *MockContentRepositoryMockRecorder) DeleteLikePost(ctx, userId, postID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteLikePost", reflect.TypeOf((*MockContentRepository)(nil).DeleteLikePost), ctx, userId, postID)
}

// DeletePost mocks base method.
func (m *MockContentRepository) DeletePost(ctx context.Context, postID uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeletePost", ctx, postID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeletePost indicates an expected call of DeletePost.
func (mr *MockContentRepositoryMockRecorder) DeletePost(ctx, postID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePost", reflect.TypeOf((*MockContentRepository)(nil).DeletePost), ctx, postID)
}

// GetAuthorByPost mocks base method.
func (m *MockContentRepository) GetAuthorByPost(postID uuid.UUID) (uuid.UUID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAuthorByPost", postID)
	ret0, _ := ret[0].(uuid.UUID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAuthorByPost indicates an expected call of GetAuthorByPost.
func (mr *MockContentRepositoryMockRecorder) GetAuthorByPost(postID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAuthorByPost", reflect.TypeOf((*MockContentRepository)(nil).GetAuthorByPost), postID)
}

// GetAuthorOfPost mocks base method.
func (m *MockContentRepository) GetAuthorOfPost(ctx context.Context, postID uuid.UUID) (uuid.UUID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAuthorOfPost", ctx, postID)
	ret0, _ := ret[0].(uuid.UUID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAuthorOfPost indicates an expected call of GetAuthorOfPost.
func (mr *MockContentRepositoryMockRecorder) GetAuthorOfPost(ctx, postID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAuthorOfPost", reflect.TypeOf((*MockContentRepository)(nil).GetAuthorOfPost), ctx, postID)
}

// GetAuthorPostsForAnon mocks base method.
func (m *MockContentRepository) GetAuthorPostsForAnon(ctx context.Context, authorId uuid.UUID, offset, limit int) ([]*models.Post, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAuthorPostsForAnon", ctx, authorId, offset, limit)
	ret0, _ := ret[0].([]*models.Post)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAuthorPostsForAnon indicates an expected call of GetAuthorPostsForAnon.
func (mr *MockContentRepositoryMockRecorder) GetAuthorPostsForAnon(ctx, authorId, offset, limit interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAuthorPostsForAnon", reflect.TypeOf((*MockContentRepository)(nil).GetAuthorPostsForAnon), ctx, authorId, offset, limit)
}

// GetAuthorPostsForLayer mocks base method.
func (m *MockContentRepository) GetAuthorPostsForLayer(ctx context.Context, layer int, authorId uuid.UUID, offset, limit int) ([]*models.Post, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAuthorPostsForLayer", ctx, layer, authorId, offset, limit)
	ret0, _ := ret[0].([]*models.Post)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAuthorPostsForLayer indicates an expected call of GetAuthorPostsForLayer.
func (mr *MockContentRepositoryMockRecorder) GetAuthorPostsForLayer(ctx, layer, authorId, offset, limit interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAuthorPostsForLayer", reflect.TypeOf((*MockContentRepository)(nil).GetAuthorPostsForLayer), ctx, layer, authorId, offset, limit)
}

// GetAuthorPostsForMe mocks base method.
func (m *MockContentRepository) GetAuthorPostsForMe(ctx context.Context, authorId uuid.UUID, offset, limit int) ([]*models.Post, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAuthorPostsForMe", ctx, authorId, offset, limit)
	ret0, _ := ret[0].([]*models.Post)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAuthorPostsForMe indicates an expected call of GetAuthorPostsForMe.
func (mr *MockContentRepositoryMockRecorder) GetAuthorPostsForMe(ctx, authorId, offset, limit interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAuthorPostsForMe", reflect.TypeOf((*MockContentRepository)(nil).GetAuthorPostsForMe), ctx, authorId, offset, limit)
}

// GetIsLikedForPosts mocks base method.
func (m *MockContentRepository) GetIsLikedForPosts(ctx context.Context, UserId uuid.UUID, posts []*models.Post) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetIsLikedForPosts", ctx, UserId, posts)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetIsLikedForPosts indicates an expected call of GetIsLikedForPosts.
func (mr *MockContentRepositoryMockRecorder) GetIsLikedForPosts(ctx, UserId, posts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetIsLikedForPosts", reflect.TypeOf((*MockContentRepository)(nil).GetIsLikedForPosts), ctx, UserId, posts)
}

// GetPopularPosts mocks base method.
func (m *MockContentRepository) GetPopularPosts(ctx context.Context, offset, limits int) ([]*models.Post, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPopularPosts", ctx, offset, limits)
	ret0, _ := ret[0].([]*models.Post)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPopularPosts indicates an expected call of GetPopularPosts.
func (mr *MockContentRepositoryMockRecorder) GetPopularPosts(ctx, offset, limits interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPopularPosts", reflect.TypeOf((*MockContentRepository)(nil).GetPopularPosts), ctx, offset, limits)
}

// GetPopularPostsForUser mocks base method.
func (m *MockContentRepository) GetPopularPostsForUser(ctx context.Context, userId uuid.UUID, offset, limits int) ([]*models.Post, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPopularPostsForUser", ctx, userId, offset, limits)
	ret0, _ := ret[0].([]*models.Post)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPopularPostsForUser indicates an expected call of GetPopularPostsForUser.
func (mr *MockContentRepositoryMockRecorder) GetPopularPostsForUser(ctx, userId, offset, limits interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPopularPostsForUser", reflect.TypeOf((*MockContentRepository)(nil).GetPopularPostsForUser), ctx, userId, offset, limits)
}

// GetPostLayerBuPostId mocks base method.
func (m *MockContentRepository) GetPostLayerBuPostId(ctx context.Context, postID uuid.UUID) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPostLayerBuPostId", ctx, postID)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPostLayerBuPostId indicates an expected call of GetPostLayerBuPostId.
func (mr *MockContentRepositoryMockRecorder) GetPostLayerBuPostId(ctx, postID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPostLayerBuPostId", reflect.TypeOf((*MockContentRepository)(nil).GetPostLayerBuPostId), ctx, postID)
}

// GetPostLikeId mocks base method.
func (m *MockContentRepository) GetPostLikeId(ctx context.Context, userId, postID uuid.UUID) (uuid.UUID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPostLikeId", ctx, userId, postID)
	ret0, _ := ret[0].(uuid.UUID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPostLikeId indicates an expected call of GetPostLikeId.
func (mr *MockContentRepositoryMockRecorder) GetPostLikeId(ctx, userId, postID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPostLikeId", reflect.TypeOf((*MockContentRepository)(nil).GetPostLikeId), ctx, userId, postID)
}

// GetPostLikes mocks base method.
func (m *MockContentRepository) GetPostLikes(ctx context.Context, postID uuid.UUID) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPostLikes", ctx, postID)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPostLikes indicates an expected call of GetPostLikes.
func (mr *MockContentRepositoryMockRecorder) GetPostLikes(ctx, postID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPostLikes", reflect.TypeOf((*MockContentRepository)(nil).GetPostLikes), ctx, postID)
}

// GetSubscriptionPostsForUser mocks base method.
func (m *MockContentRepository) GetSubscriptionPostsForUser(ctx context.Context, userId uuid.UUID, offset, limits int) ([]*models.Post, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSubscriptionPostsForUser", ctx, userId, offset, limits)
	ret0, _ := ret[0].([]*models.Post)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSubscriptionPostsForUser indicates an expected call of GetSubscriptionPostsForUser.
func (mr *MockContentRepositoryMockRecorder) GetSubscriptionPostsForUser(ctx, userId, offset, limits interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSubscriptionPostsForUser", reflect.TypeOf((*MockContentRepository)(nil).GetSubscriptionPostsForUser), ctx, userId, offset, limits)
}

// GetUserLayerOfAuthor mocks base method.
func (m *MockContentRepository) GetUserLayerOfAuthor(ctx context.Context, userId, authorId uuid.UUID) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserLayerOfAuthor", ctx, userId, authorId)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserLayerOfAuthor indicates an expected call of GetUserLayerOfAuthor.
func (mr *MockContentRepositoryMockRecorder) GetUserLayerOfAuthor(ctx, userId, authorId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserLayerOfAuthor", reflect.TypeOf((*MockContentRepository)(nil).GetUserLayerOfAuthor), ctx, userId, authorId)
}

// GetUserRole mocks base method.
func (m *MockContentRepository) GetUserRole(ctx context.Context, userId uuid.UUID) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserRole", ctx, userId)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserRole indicates an expected call of GetUserRole.
func (mr *MockContentRepositoryMockRecorder) GetUserRole(ctx, userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserRole", reflect.TypeOf((*MockContentRepository)(nil).GetUserRole), ctx, userId)
}

// InsertLikePost mocks base method.
func (m *MockContentRepository) InsertLikePost(ctx context.Context, userId, postID uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertLikePost", ctx, userId, postID)
	ret0, _ := ret[0].(error)
	return ret0
}

// InsertLikePost indicates an expected call of InsertLikePost.
func (mr *MockContentRepositoryMockRecorder) InsertLikePost(ctx, userId, postID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertLikePost", reflect.TypeOf((*MockContentRepository)(nil).InsertLikePost), ctx, userId, postID)
}

// InsertPost mocks base method.
func (m *MockContentRepository) InsertPost(ctx context.Context, userId, postId uuid.UUID, title, content string, layer int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertPost", ctx, userId, postId, title, content, layer)
	ret0, _ := ret[0].(error)
	return ret0
}

// InsertPost indicates an expected call of InsertPost.
func (mr *MockContentRepositoryMockRecorder) InsertPost(ctx, userId, postId, title, content, layer interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertPost", reflect.TypeOf((*MockContentRepository)(nil).InsertPost), ctx, userId, postId, title, content, layer)
}

// UpdateContentOfPost mocks base method.
func (m *MockContentRepository) UpdateContentOfPost(ctx context.Context, postID uuid.UUID, content string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateContentOfPost", ctx, postID, content)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateContentOfPost indicates an expected call of UpdateContentOfPost.
func (mr *MockContentRepositoryMockRecorder) UpdateContentOfPost(ctx, postID, content interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateContentOfPost", reflect.TypeOf((*MockContentRepository)(nil).UpdateContentOfPost), ctx, postID, content)
}

// UpdateTitleOfPost mocks base method.
func (m *MockContentRepository) UpdateTitleOfPost(ctx context.Context, postID uuid.UUID, title string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateTitleOfPost", ctx, postID, title)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateTitleOfPost indicates an expected call of UpdateTitleOfPost.
func (mr *MockContentRepositoryMockRecorder) UpdateTitleOfPost(ctx, postID, title interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTitleOfPost", reflect.TypeOf((*MockContentRepository)(nil).UpdateTitleOfPost), ctx, postID, title)
}
