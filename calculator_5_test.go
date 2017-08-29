package math

import(
. "github.com/onsi/ginkgo"
. "github.com/onsi/gomega"
"testing"
)
func TestCalculator5(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Calculator Suite")
}

var _ = Describe("Calculator Behaviors", func(){

	Context("Should divide one number by another", func(){

		calculator := &Calculator{}

		Context("When both numbers are positive", func(){
			result, err := calculator.Division(4.0,2.0)

			It("And the result is correct", func(){
				Expect(result).Should(Equal(2.0))
			})

			It("And not error have occurred", func(){
				Expect(err).ShouldNot(HaveOccurred())
			})
		})

		Context("When a division by zero occur", func() {
			_, err := calculator.Division(4.0, 0)

			It("And an error should have occurred", func(){
				Expect(err).Should(HaveOccurred())
			})
		})

	})
})
