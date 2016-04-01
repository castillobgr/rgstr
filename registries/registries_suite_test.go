package registries_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestRegistries(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Registries Suite")
}
