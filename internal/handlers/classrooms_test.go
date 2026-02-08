package handlers

import (
	"net/http/httptest"
	"testing"
)

func TestListClassrooms(t *testing.T) {
	rec := httptest.NewRecorder()
	ListClassrooms(rec, httptest.NewRequest("GET", "/api/classrooms", nil))
	if rec.Code != 200 {
		t.Fatalf("status %d", rec.Code)
	}
}
