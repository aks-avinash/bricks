// Copyright © 2018 by PACE Telematics GmbH. All rights reserved.
// Created at 2018/09/04 by Vincent Landgraf

package jsonapi

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/pace/bricks/maintenance/metric"
)

func TestCaptureStatus(t *testing.T) {
	rec := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/test/1234567", nil)

	handler := func(w http.ResponseWriter, r *http.Request) {
		w = NewMetric("simple", "/test/{id}", w, r)
		w.WriteHeader(204)
	}

	handler(rec, req)

	resp := rec.Result()
	defer resp.Body.Close()

	if resp.StatusCode != 204 {
		t.Errorf("Failed to return correct 204 response status, got: %v", resp.StatusCode)
	}

	rec = httptest.NewRecorder()
	req = httptest.NewRequest("GET", "/metrics", nil)
	metric.Handler().ServeHTTP(rec, req)

	body := rec.Body.String()
	for _, metric := range []string{
		"pace_api_http_request_duration",
		"pace_api_http_request_duration_seconds",
		"pace_api_http_size_bytes",
	} {
		if !strings.Contains(body, metric) {
			t.Errorf("Expected pace api metrics got: %v", body)
		}
	}
}
