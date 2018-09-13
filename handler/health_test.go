package splitbillapi_test

import (
	"net/http/httptest"
	"testing"

	splitbillapi "github.com/desdulianto/splitbillapi/handler"
)

func TestHealth(t *testing.T) {
	req := httptest.NewRequest("GET", URL+"/healthz", nil)
	res := httptest.NewRecorder()
	splitbillapi.HealthHandler(res, req)

	result := res.Result()
	body := string(res.Body.Bytes())

	if result.StatusCode != 200 {
		t.Errorf("Health suppose to return 200, got %v", result.StatusCode)
	}

	if body != "ok" {
		t.Errorf("Health suppose to return ok, got %v", body)
	}
}
