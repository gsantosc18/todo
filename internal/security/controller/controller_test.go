package controller

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/gsantosc18/todo/test/mock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func getTestContext(w *httptest.ResponseRecorder) *gin.Context {
	gin.SetMode(gin.TestMode)

	ctx, _ := gin.CreateTestContext(w)

	ctx.Request = &http.Request{
		Header: map[string][]string{
			"Content-Type": {"application/json"},
		},
	}

	return ctx
}

var mockUser userLogin = userLogin{
	Email:    "admin",
	Password: "admin",
}

func TestLogin(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	w := httptest.NewRecorder()
	ctx := getTestContext(w)
	token := "encrypted-bearer"
	expectedToken := fmt.Sprintf("{\"token\":\"%s\"}", token)
	userJson, _ := json.Marshal(mockUser)

	ctx.Request.Body = io.NopCloser(strings.NewReader(string(userJson)))

	service := mock.NewMockTokenService(ctrl)

	service.EXPECT().
		NewToken(gomock.Any()).
		Return(token, nil).
		Times(1)

	controller := NewSecurityController(service)

	controller.LoginController(ctx)

	result, _ := io.ReadAll(w.Body)

	assert.Equal(t, expectedToken, string(result))
	assert.Equal(t, http.StatusOK, w.Code)
}
