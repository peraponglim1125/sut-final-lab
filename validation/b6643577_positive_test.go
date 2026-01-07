package validation_test

import (
	"testing"
	. "github.com/onsi/gomega"
	"github.com/peraponglim1125/sut-final-lab/entity"
)

func TestValid(t *testing.T){
	g := NewGomegaWithT(t)
	employees := entity.Employees{
		Name:  "Phirapong",
		Salary:   15500.00,
		EmployeeCode: "HR-1024",  
	}
	ok, err := entity.ValidationEmployees(&employees)
	g.Expect(ok).To(BeTrue())
	g.Expect(err).To(BeNil())

}