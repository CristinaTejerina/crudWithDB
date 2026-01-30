package httpadapter

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"crudWithDB/internal/application"
	"crudWithDB/internal/domain"
)

func setupRouter(repo *UserRepoMock) *gin.Engine {
	gin.SetMode(gin.TestMode)

	service := application.NewUserService(repo)
	handler := NewUserHTTPHandler(service)

	r := gin.New()
	handler.RegisterRoutes(r)

	return r
}

func TestCreateUserOK(t *testing.T) {
	repo := new(UserRepoMock)
	router := setupRouter(repo)

	user := domain.User{ID: "1", Name: "Ana", Email: "ana@test.com"}
	body, _ := json.Marshal(user)

	repo.On("Create", user).
		Return(nil).
		Once()

	req, _ := http.NewRequest(http.MethodPost, "/users", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	repo.AssertExpectations(t)
}

func TestCreateUserBadRequest(t *testing.T) {
	repo := new(UserRepoMock)
	router := setupRouter(repo)

	req, _ := http.NewRequest(http.MethodPost, "/users", bytes.NewBuffer([]byte("{bad json")))
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestCreateUserInternalError(t *testing.T) {
	repo := new(UserRepoMock)
	router := setupRouter(repo)

	user := domain.User{ID: "1", Name: "Ana", Email: "ana@test.com"}
	body, _ := json.Marshal(user)

	repo.On("Create", mock.Anything).
		Return(errors.New("not found")).
		Once()

	req, _ := http.NewRequest(http.MethodPost, "/users", bytes.NewBuffer(body))
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func TestGetUserOK(t *testing.T) {
	repo := new(UserRepoMock)
	router := setupRouter(repo)

	user := domain.User{ID: "1", Name: "Ana", Email: "ana@test.com"}

	repo.On("GetByID", "1").
		Return(user, nil).
		Once()

	req, _ := http.NewRequest(http.MethodGet, "/users/1", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var resp domain.User
	_ = json.Unmarshal(w.Body.Bytes(), &resp)

	assert.Equal(t, user, resp)
	repo.AssertExpectations(t)
}

func TestGetUserBadRequest(t *testing.T) {
	repo := new(UserRepoMock)
	router := setupRouter(repo)

	repo.On("GetByID", "1").
		Return(domain.User{}, sql.ErrNoRows).
		Once()

	req, _ := http.NewRequest(http.MethodGet, "/users/1", bytes.NewBuffer([]byte("{bad json")))
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestGetUserNotFound(t *testing.T) {
	repo := new(UserRepoMock)
	router := setupRouter(repo)

	repo.On("GetByID", "9").
		Return(domain.User{}, sql.ErrNoRows).
		Once()

	req, _ := http.NewRequest(http.MethodGet, "/users/9", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestGetUserInternalError(t *testing.T) {
	repo := new(UserRepoMock)
	router := setupRouter(repo)

	repo.On("GetByID", mock.Anything).
		Return(domain.User{}, errors.New("not found")).
		Once()

	req, _ := http.NewRequest(http.MethodGet, "/users/1", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func TestUpdateUserOK(t *testing.T) {
	repo := new(UserRepoMock)
	router := setupRouter(repo)

	bodyUser := domain.User{Name: "Nuevo", Email: "nuevo@test.com"}
	expected := domain.User{ID: "1", Name: "Nuevo", Email: "nuevo@test.com"}

	body, _ := json.Marshal(bodyUser)

	repo.On("Update", expected).
		Return(nil).
		Once()

	req, _ := http.NewRequest(http.MethodPut, "/users/1", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	repo.AssertExpectations(t)
}

func TestUpdateUserNotFound(t *testing.T) {
	repo := new(UserRepoMock)
	router := setupRouter(repo)

	user := domain.User{ID: "1", Name: "x", Email: "x@test.com"}
	body, _ := json.Marshal(user)

	repo.On("Update", mock.Anything).
		Return(sql.ErrNoRows).
		Once()

	req, _ := http.NewRequest(http.MethodPut, "/users/1", bytes.NewBuffer(body))
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestUpdateUserInternalError(t *testing.T) {
	repo := new(UserRepoMock)
	router := setupRouter(repo)

	user := domain.User{ID: "1", Name: "x", Email: "x@test.com"}
	body, _ := json.Marshal(user)

	repo.On("Update", mock.Anything).
		Return(errors.New("not found")).
		Once()

	req, _ := http.NewRequest(http.MethodPut, "/users/1", bytes.NewBuffer(body))
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func TestUpdateUserBadRequest(t *testing.T) {
	repo := new(UserRepoMock)
	router := setupRouter(repo)

	repo.On("Update", mock.Anything).
		Return(sql.ErrNoRows).
		Once()

	req, _ := http.NewRequest(http.MethodPut, "/users/1", bytes.NewBuffer([]byte("{bad json")))
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestDeleteUserOK(t *testing.T) {
	repo := new(UserRepoMock)
	router := setupRouter(repo)

	repo.On("Delete", "1").
		Return(nil).
		Once()

	req, _ := http.NewRequest(http.MethodDelete, "/users/1", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	repo.AssertExpectations(t)
}

func TestDeleteUserNotFound(t *testing.T) {
	repo := new(UserRepoMock)
	router := setupRouter(repo)

	repo.On("Delete", "1").
		Return(errors.New("not found")).
		Once()

	req, _ := http.NewRequest(http.MethodDelete, "/users/1", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func TestDeleteUserBadRequest(t *testing.T) {
	repo := new(UserRepoMock)
	router := setupRouter(repo)

	repo.On("Delete", "1").
		Return(sql.ErrNoRows).
		Once()

	req, _ := http.NewRequest(http.MethodDelete, "/users/1", bytes.NewBuffer([]byte("{bad json")))
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
}
