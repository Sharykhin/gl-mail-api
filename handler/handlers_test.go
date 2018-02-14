package handler

import (
	"net/http"
	"testing"

	"github.com/gavv/httpexpect"
)

func TestPong(t *testing.T) {
	t.Run("status ok", func(t *testing.T) {
		e := httpexpect.WithConfig(httpexpect.Config{
			Client: &http.Client{
				Transport: httpexpect.NewBinder(Handler()),
				Jar:       httpexpect.NewJar(),
			},
			Reporter: httpexpect.NewAssertReporter(t),
		})
		c := e.Request(http.MethodGet, "/ping")
		c.Expect().Status(http.StatusOK)
	})

	t.Run("response pong", func(t *testing.T) {
		e := httpexpect.WithConfig(httpexpect.Config{
			Client: &http.Client{
				Transport: httpexpect.NewBinder(Handler()),
				Jar:       httpexpect.NewJar(),
			},
			Reporter: httpexpect.NewAssertReporter(t),
		})
		c := e.Request(http.MethodGet, "/ping")
		c.Expect().Body().Equal("pong")
	})

	t.Run("content type text/plain", func(t *testing.T) {
		e := httpexpect.WithConfig(httpexpect.Config{
			Client: &http.Client{
				Transport: httpexpect.NewBinder(Handler()),
				Jar:       httpexpect.NewJar(),
			},
			Reporter: httpexpect.NewAssertReporter(t),
		})
		c := e.Request(http.MethodGet, "/ping")
		c.Expect().Header("Content-Type").Match("text/plain")

	})
}

func TestGetFailedMailsList(t *testing.T) {
	t.Run("bad auth response", func(t *testing.T) {
		e := httpexpect.WithConfig(httpexpect.Config{
			Client: &http.Client{
				Transport: httpexpect.NewBinder(Handler()),
				Jar:       httpexpect.NewJar(),
			},
			Reporter: httpexpect.NewAssertReporter(t),
		})
		c := e.Request(http.MethodGet, "/failed-mails").Expect().Status(http.StatusUnauthorized).JSON().Object()
		c.Value("success").Equal(false)
		c.Value("data").Equal(nil)
		c.Value("error").Equal("Unauthorized")
	})
}

func TestCreateFailedMail(t *testing.T) {
	t.Run("bad request: action is required", func(t *testing.T) {
		e := httpexpect.WithConfig(httpexpect.Config{
			Client: &http.Client{
				Transport: httpexpect.NewBinder(Handler()),
				Jar:       httpexpect.NewJar(),
			},
			Reporter: httpexpect.NewAssertReporter(t),
		})

		payload := map[string]interface{}{
			"action": "",
		}

		c := e.Request(http.MethodPost, "/failed-mails").
			WithJSON(payload).
			Expect().Status(http.StatusBadRequest).JSON().Object()
		c.Value("success").Equal(false)
		c.Value("data").Equal(nil)
		c.Value("error").Equal("action is required")
	})

	t.Run("bad request: payload is required", func(t *testing.T) {
		e := httpexpect.WithConfig(httpexpect.Config{
			Client: &http.Client{
				Transport: httpexpect.NewBinder(Handler()),
				Jar:       httpexpect.NewJar(),
			},
			Reporter: httpexpect.NewAssertReporter(t),
		})

		payload := map[string]interface{}{
			"action":  "register",
			"payload": nil,
		}

		c := e.Request(http.MethodPost, "/failed-mails").
			WithJSON(payload).
			Expect().Status(http.StatusBadRequest).JSON().Object()
		c.Value("success").Equal(false)
		c.Value("data").Equal(nil)
		c.Value("error").Equal("payload is required")
	})

	t.Run("bad request: reason is required", func(t *testing.T) {
		e := httpexpect.WithConfig(httpexpect.Config{
			Client: &http.Client{
				Transport: httpexpect.NewBinder(Handler()),
				Jar:       httpexpect.NewJar(),
			},
			Reporter: httpexpect.NewAssertReporter(t),
		})

		payload := map[string]interface{}{
			"action":  "register",
			"payload": map[string]interface{}{},
			"reason":  nil,
		}

		c := e.Request(http.MethodPost, "/failed-mails").
			WithJSON(payload).
			Expect().Status(http.StatusBadRequest).JSON().Object()
		c.Value("success").Equal(false)
		c.Value("data").Equal(nil)
		c.Value("error").Equal("reason is required")
	})
}
