package http

import (
	"context"
	"github.com/atkrad/wait4x/pkg/log"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func TestHttpInvalidAddress(t *testing.T) {
	logger, _ := log.NewLogrus(logrus.DebugLevel.String(), ioutil.Discard)

	hc := NewHTTP("http://not-exists.tld")
	hc.SetLogger(logger)

	assert.Equal(t, false, hc.Check(context.TODO()))
}

func TestHttpValidAddress(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer ts.Close()

	logger, _ := log.NewLogrus(logrus.DebugLevel.String(), ioutil.Discard)

	hc := NewHTTP(ts.URL)
	hc.SetLogger(logger)

	assert.Equal(t, true, hc.Check(context.TODO()))
}

func TestHttpInvalidStatusCode(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer ts.Close()

	logger, _ := log.NewLogrus(logrus.DebugLevel.String(), ioutil.Discard)

	hc := NewHTTP(ts.URL, WithExpectStatusCode(http.StatusCreated))
	hc.SetLogger(logger)

	assert.Equal(t, false, hc.Check(context.TODO()))
}

func TestHttpValidStatusCode(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer ts.Close()

	logger, _ := log.NewLogrus(logrus.DebugLevel.String(), ioutil.Discard)

	hc := NewHTTP(ts.URL, WithExpectStatusCode(http.StatusOK))
	hc.SetLogger(logger)

	assert.Equal(t, true, hc.Check(context.TODO()))
}

func TestHttpInvalidBody(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Wait4X"))
	}))
	defer ts.Close()

	logger, _ := log.NewLogrus(logrus.DebugLevel.String(), ioutil.Discard)

	hc := NewHTTP(ts.URL, WithExpectBody("FooBar"))
	hc.SetLogger(logger)

	assert.Equal(t, false, hc.Check(context.TODO()))
}

func TestHttpValidBody(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Wait4X is the best CLI tools. Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nulla volutpat arcu malesuada lacus vulputate feugiat. Etiam vitae sem quis ligula consequat euismod. Lorem ipsum dolor sit amet, consectetur adipiscing elit. Phasellus fringilla sapien non lacus volutpat sollicitudin. Donec sollicitudin sit amet purus ac rutrum. Nam nunc orci, luctus a sagittis."))
	}))
	defer ts.Close()

	logger, _ := log.NewLogrus(logrus.DebugLevel.String(), ioutil.Discard)

	hc := NewHTTP(ts.URL, WithExpectBody("Wait4X.+best.+tools"))
	hc.SetLogger(logger)

	assert.Equal(t, true, hc.Check(context.TODO()))
}
