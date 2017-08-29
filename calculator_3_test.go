package math

import(
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)
/*
	This makes the link between Gomega, Ginkgo and Golang Testing
 */
func TestCalculator3(t *testing.T) {
	//This enables all fail raised by Gomega could be passed to Golang
	RegisterFailHandler(Fail)

	//This enables Ginkgo to run the tests
	RunSpecs(t, "Calculator Suite")
}

var _ = Describe("Calculator Behaviors", func(){
	Context("Should sum two values", func(){

		It("When both are positive", func(){
			calculator := &Calculator{}
			result := calculator.Sum(1.0,2.0)
			Expect(result).Should(Equal(3.0))
		})

	})
})
