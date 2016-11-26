package algorithmutil

import (
	"fmt"

	"github.com/onsi/ginkgo"
)

// GinkgoTestReporter simple warpper to reference testing.T in ginkgo since gomock requires testing.T
// https://github.com/onsi/ginkgo/issues/9
type GinkgoTestReporter struct {
}

// Errorf ginkgo error fail
func (g GinkgoTestReporter) Errorf(format string, args ...interface{}) {
	ginkgo.Fail(fmt.Sprintf(format, args))
}

// Fatalf ginkgo fatal fail
func (g GinkgoTestReporter) Fatalf(format string, args ...interface{}) {
	ginkgo.Fail(fmt.Sprintf(format, args))
}
