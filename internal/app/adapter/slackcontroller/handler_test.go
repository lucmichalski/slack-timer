package slackcontroller

import (
	"context"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"proteinreminder/internal/app/driver/di"
	"proteinreminder/internal/app/usecase/updateproteinevent"
	"strings"
	"testing"
)

func TestNewRequestHandler(t *testing.T) {
	cases := []struct {
		name    string
		text    string
		subType string
	}{
		{"set", "set 1", CmdSet},
		{"got", "got", CmdGot},
		{"nil", "invalid", ""},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			// https://golang.org/src/net/http/request_test.go
			body := strings.NewReader(`text=` + c.text)
			httpReq := httptest.NewRequest(http.MethodPost, "/", body)
			httpReq.Header.Set("Content-Type", "application/x-www-form-urlencoded")

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			m := di.NewMockDI(ctrl)
			if c.name == "nil" {
				m.EXPECT().Get(gomock.Eq("UpdateProteinEvent")).MinTimes(0)
			} else {
				m.EXPECT().Get(gomock.Eq("UpdateProteinEvent")).Return(&updateproteinevent.Interactor{})
			}
			di.SetDi(m)

			req, err := NewRequestHandler(httpReq)
			if c.subType != "" {
				assert.NoError(t, err)
				assert.NotNil(t, req)
			}

			if c.subType == CmdGot {
				h, match := req.(*GotRequestHandler)
				assert.True(t, match)
				assert.Equal(t, h.params.Text, c.text)
			} else if c.subType == CmdSet {
				h, match := req.(*SetRequestHandler)
				assert.True(t, match)
				assert.Equal(t, h.params.Text, c.text)
			} else {
				assert.Nil(t, req)
			}
		})
	}
}

func TestMakeErrorCallbackResponseBody(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		msg := "test"
		err := errors.New("test")
		want := `{"message":"test","error":"test"}`

		gotB, gotErr := makeErrorCallbackResponseBody(msg, err)
		assert.NoError(t, gotErr)
		assert.Equal(t, []byte(want), gotB)
	})
}

func TestSlackCallbackSetRequest_validate(t *testing.T) {
	t.Log("TestSlackCallbackGotRequest_validate covers this test.")
}

func TestHandler(t *testing.T) {
	t.Run("wrong method", func(t *testing.T) {
		body := strings.NewReader("")
		httpReq := httptest.NewRequest(http.MethodGet, "/", body)
		httpReq.Header.Set("Content-Type", "application/x-www-form-urlencoded")

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		m := di.NewMockDI(ctrl)
		m.EXPECT().Get(gomock.Eq("UpdateProteinEvent")).Times(0)
		di.SetDi(m)

		w := httptest.NewRecorder()

		Handler(context.TODO(), w, httpReq)

		assert.Equal(t, 404, w.Result().StatusCode)
	})

	t.Run("invalid parameter", func(t *testing.T) {
		want := fmt.Sprintf(`{"message":"parameter error","error":"%s"}`, ErrInvalidRequest.Error())

		body := strings.NewReader("")
		httpReq := httptest.NewRequest(http.MethodPost, "/", body)
		httpReq.Header.Set("Content-Type", "application/x-www-form-urlencoded")

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		m := di.NewMockDI(ctrl)
		m.EXPECT().Get(gomock.Eq("UpdateProteinEvent")).Times(0)
		di.SetDi(m)

		w := httptest.NewRecorder()

		Handler(context.TODO(), w, httpReq)

		assert.Equal(t, http.StatusBadRequest, w.Result().StatusCode)
		assert.Equal(t, want, w.Body.String())

	})

	t.Run("ok", func(t *testing.T) {
		ctx := context.TODO()
		userId := "abc"

		body := strings.NewReader(`text=got&user_id=` + userId)
		httpReq := httptest.NewRequest(http.MethodPost, "/", body)
		httpReq.Header.Set("Content-Type", "application/x-www-form-urlencoded")

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		mu := updateproteinevent.NewMockUsecase(ctrl)
		mu.EXPECT().UpdateTimeToDrink(gomock.Eq(ctx), gomock.Eq(userId), gomock.Any())
		m := di.NewMockDI(ctrl)
		m.EXPECT().Get(gomock.Eq("UpdateProteinEvent")).Return(mu)
		di.SetDi(m)

		w := httptest.NewRecorder()

		Handler(ctx, w, httpReq)

		t.Log(w.Body.String())
		assert.Equal(t, http.StatusOK, w.Result().StatusCode)
	})
}
