// Package goalgorithms provides basic algorithm and data structure implementations
package goalgorithms_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestSorts(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Sorts Suite")
}
