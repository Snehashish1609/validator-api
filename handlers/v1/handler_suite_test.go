package v1_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestValidatorApi(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "ValidatorApi Suite")
}