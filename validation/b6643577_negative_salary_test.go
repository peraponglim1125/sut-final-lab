package validation_test

import (
	"testing"
	. "github.com/onsi/gomega"
	"github.com/peraponglim1125/sut-final-lab/entity"
)

func TestInValidSalary(t *testing.T){
	g := NewGomegaWithT(t)
	employees := entity.Employees{
		Name:  "Phirapong",
		Salary:   1500000000.0,
		EmployeeCode: "HR-1024",  
	}
	ok, err := entity.ValidationEmployees(&employees)
	g.Expect(ok).To(BeFalse())
	g.Expect(err.Error()).To(Equal("Salary must be between 15000 and 200000"))

}