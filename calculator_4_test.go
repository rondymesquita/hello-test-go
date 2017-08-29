package math

import(
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)
func TestCalculator4(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Calculator Suite")
}

var _ = Describe("Calculator Behaviors", func(){
	Context("Should sum two values", func(){

		var calculator *Calculator

		BeforeEach(func(){
			calculator = &Calculator{}
		})

		It("When both are positive", func(){
			result := calculator.Sum(1.0,2.0)
			Expect(result).Should(Equal(3.0))
		})

		It("When both are negative", func(){
			result := calculator.Sum(-3.0,-2.0)
			Expect(result).Should(Equal(-5.0))
		})

		It("When they have different signals", func(){
			result := calculator.Sum(1.0,-2.0)
			Expect(result).Should(Equal(-1.0))
		})
	})
})
