package validation_test

import (
	"testing"
	. "github.com/onsi/gomega"
	"github.com/peraponglim1125/sut-final-lab/entity"
)

func TestInValidEmployeeCode (t *testing.T){
	g := NewGomegaWithT(t)
	employees := entity.Employees{
		Name:  "Phirapong",
		Salary:   160000.00,
		EmployeeCode: "HR-10245365",  
	}
	ok, err := entity.ValidationEmployees(&employees)
	g.Expect(ok).To(BeFalse())
	g.Expect(err.Error()).To(Equal("EmployeeCode must be 2 uppercase English letters (A-Z) followed by ‘-’ and 4 digits (0-9)"))

}