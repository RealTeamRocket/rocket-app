package server_test

import (
	"testing"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestDatabaseIntegration(t *testing.T) {
    RegisterFailHandler(Fail)
    RunSpecs(t, "Database Integration Tests Suite")
}
