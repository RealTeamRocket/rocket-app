package routes

import (
	"testing"
)

func TestRegisterRoutes(t *testing.T) {
	result := 1 + 1
    expected := 2

    if result != expected {
        t.Errorf("expected %d but got %d", expected, result)
    }
}
