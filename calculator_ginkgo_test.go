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

/*
	This describes the scope of your tests
 */
var _ = Describe("Calculator Behaviors", func(){

	/*
		This describe the context.
	 */
	Context("Should sum two values", func(){

		/*
			This is the test indeed
		 */
		It("When both are positive", func(){

		})
	})
})
