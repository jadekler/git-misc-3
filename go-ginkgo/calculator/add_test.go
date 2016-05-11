package calculator_test

import (
	. "git-misc-3/go-ginkgo/calculator"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Calculator", func() {
	Describe("Add", func() {
		It("works", func() {
			Expect(Add(3, 5)).To(Equal(8))
		})
	})
})
