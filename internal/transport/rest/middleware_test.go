package rest

import (
	"fmt"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/olenka--91/reminder-app/internal/service"
	mock_service "github.com/olenka--91/reminder-app/internal/service/mocks"
	"github.com/stretchr/testify/assert"
)

func TestHandler_userIdentity(t *testing.T) {
	type mockBehaviour func(s *mock_service.MockAuthorization, token string)

	testTable := []struct {
		name                 string
		headerName           string
		headerValue          string
		token                string
		mockBehaviour        mockBehaviour
		expectedStatusCode   int
		expectedResponceBody string
	}{
		{name: "OK",
			headerName:  "Authorization",
			headerValue: "Bearer token",
			token:       "token",
			mockBehaviour: func(s *mock_service.MockAuthorization, token string) {
				s.EXPECT().ParseToken(token).Return(1, nil)
			},
			expectedStatusCode:   200,
			expectedResponceBody: "1",
		},
		{name: "No header",
			headerName:           "",
			mockBehaviour:        func(s *mock_service.MockAuthorization, token string) {},
			expectedStatusCode:   401,
			expectedResponceBody: `{"message":"empty auth header"}`,
		},
		{name: "Invalid header",
			headerName:           "Authorization",
			headerValue:          "Beaer token",
			mockBehaviour:        func(s *mock_service.MockAuthorization, token string) {},
			expectedStatusCode:   401,
			expectedResponceBody: `{"message":"invalid auth header"}`,
		},
		{name: "Empty header",
			headerName:           "Authorization",
			headerValue:          "Bearer ",
			mockBehaviour:        func(s *mock_service.MockAuthorization, token string) {},
			expectedStatusCode:   401,
			expectedResponceBody: `{"message":"token is empty"}`,
		},
	}
	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			//
			c := gomock.NewController(t)
			defer c.Finish()

			auth := mock_service.NewMockAuthorization(c)
			testCase.mockBehaviour(auth, testCase.token)

			services := &service.Service{Authorization: auth}
			handler := NewHandler(services)

			//test Server
			r := gin.New()
			r.GET("/protected", handler.userIdentity, func(c *gin.Context) {
				id, _ := c.Get(userCtx)
				c.String(200, fmt.Sprintf("%d", id.(int)))
			})

			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "/protected", nil)
			req.Header.Set(testCase.headerName, testCase.headerValue)

			r.ServeHTTP(w, req)

			assert.Equal(t, w.Code, testCase.expectedStatusCode)
			assert.Equal(t, w.Body.String(), testCase.expectedResponceBody)
		},
		)
	}
}

func TestGetUserId(t *testing.T) {
	getCtx := func(id int) *gin.Context {
		ctx := &gin.Context{}
		ctx.Set(userCtx, id)
		return ctx
	}

	testTable := []struct {
		name       string
		ctx        *gin.Context
		id         int
		shouldFail bool
	}{
		{
			name:       "OK",
			id:         1,
			ctx:        getCtx(1),
			shouldFail: false,
		}, {
			name:       "Empty",
			ctx:        &gin.Context{},
			shouldFail: true,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			id, err := getUserId(testCase.ctx)
			if testCase.shouldFail {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, testCase.id, id)
		})
	}

}
