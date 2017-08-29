package math

import(
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

func TestCalculator1(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Calculator Suite")
}


var _ = Describe("Calculator Behaviors", func(){

	Context("Should sum two values", func(){

		It("When both are positive", func(){

		})
	})
})
