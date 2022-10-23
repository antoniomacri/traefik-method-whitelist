package main

import (
	"context"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestWithRecorder(t *testing.T) {
	ctx := context.Background()

	Convey("Given a plugin instance with allowed methods GET and POST and specific message", t, func() {
		mockHandler := &MockHandler{}
		mockHandler.On("ServeHTTP", mock.Anything, mock.Anything).Return()

		cfg := CreateConfig()
		cfg.Message = "This Method Is Not Allowed"
		cfg.Methods = []string{"GET", "POST"}

		handler, err := New(ctx, mockHandler, cfg, "pluginName")

		SoMsg("Plugin initialization error", err, ShouldBeNil)

		Convey("Given an HTTP recorder", func() {
			recorder := httptest.NewRecorder()

			Convey("When a GET request arrives", func() {
				request, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://localhost", nil)
				SoMsg("Request creation error", err, ShouldBeNil)
				handler.ServeHTTP(recorder, request)

				Convey("Then the request is allowed", func() {
					So(mockHandler.AssertCalled(t, "ServeHTTP", mock.Anything, request), ShouldBeTrue)
					So(recorder.Result().StatusCode, ShouldEqual, http.StatusOK)
				})
			})
			Convey("When a POST request arrives", func() {
				request, err := http.NewRequestWithContext(ctx, http.MethodPost, "http://localhost", nil)
				SoMsg("Request creation error", err, ShouldBeNil)
				handler.ServeHTTP(recorder, request)

				Convey("Then the request is allowed", func() {
					So(mockHandler.AssertCalled(t, "ServeHTTP", mock.Anything, request), ShouldBeTrue)
					So(recorder.Result().StatusCode, ShouldEqual, http.StatusOK)
				})
			})
			Convey("When a HEAD request arrives", func() {
				request, err := http.NewRequestWithContext(ctx, http.MethodHead, "http://localhost", nil)
				SoMsg("Request creation error", err, ShouldBeNil)
				handler.ServeHTTP(recorder, request)

				Convey("Then the request is NOT allowed", func() {
					So(mockHandler.AssertNotCalled(t, "ServeHTTP", mock.Anything, mock.Anything), ShouldBeTrue)
					So(recorder.Result().StatusCode, ShouldEqual, http.StatusMethodNotAllowed)
					So(recorder.Body.String(), ShouldEqual, "This Method Is Not Allowed")
				})
			})
		})
	})

	Convey("Given a plugin instance with allowed methods GET and POST and no message", t, func() {
		mockHandler := &MockHandler{}
		mockHandler.On("ServeHTTP", mock.Anything, mock.Anything).Return()

		cfg := CreateConfig()
		cfg.Methods = []string{"GET", "POST"}

		handler, err := New(ctx, mockHandler, cfg, "pluginName")

		SoMsg("Plugin initialization error", err, ShouldBeNil)

		Convey("Given an HTTP recorder", func() {
			recorder := httptest.NewRecorder()

			Convey("When a GET request arrives", func() {
				request, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://localhost", nil)
				SoMsg("Request creation error", err, ShouldBeNil)
				handler.ServeHTTP(recorder, request)

				Convey("Then the request is allowed", func() {
					So(mockHandler.AssertCalled(t, "ServeHTTP", mock.Anything, request), ShouldBeTrue)
					So(recorder.Result().StatusCode, ShouldEqual, http.StatusOK)
				})
			})
			Convey("When a POST request arrives", func() {
				request, err := http.NewRequestWithContext(ctx, http.MethodPost, "http://localhost", nil)
				SoMsg("Request creation error", err, ShouldBeNil)
				handler.ServeHTTP(recorder, request)

				Convey("Then the request is allowed", func() {
					So(mockHandler.AssertCalled(t, "ServeHTTP", mock.Anything, request), ShouldBeTrue)
					So(recorder.Result().StatusCode, ShouldEqual, http.StatusOK)
				})
			})
			Convey("When a HEAD request arrives", func() {
				request, err := http.NewRequestWithContext(ctx, http.MethodHead, "http://localhost", nil)
				SoMsg("Request creation error", err, ShouldBeNil)
				handler.ServeHTTP(recorder, request)

				Convey("Then the request is NOT allowed", func() {
					So(mockHandler.AssertNotCalled(t, "ServeHTTP", mock.Anything, mock.Anything), ShouldBeTrue)
					So(recorder.Result().StatusCode, ShouldEqual, http.StatusMethodNotAllowed)
					So(recorder.Body.String(), ShouldEqual, "Method Not Allowed")
				})
			})
		})
	})
}

type MockHandler struct {
	mock.Mock
}

func (m *MockHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	m.Called(w, r)
}
