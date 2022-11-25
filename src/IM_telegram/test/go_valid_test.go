package test

import (
	"fmt"
	"github.com/asaskevich/govalidator"
	"testing"
)

func TestBase(t *testing.T) {
	ip4 := "192.168.1.1"
	fmt.Println(govalidator.IsIPv4(ip4))

	digit := 5
	fmt.Println(govalidator.InRangeInt(digit, 0, 10))
}

type ValidTest struct {
	Name  string `valid:"in(bean|kipple)"`
	Sex   string `valid:"in(male|female)"`
	Age   uint   `valid:"range(0|150)"`
	ID    uint   `valid:"matches(2021140[0-9]{3})"`
	phone uint   `valid:"matches((13[0-9]|14[01456879]|15[0-35-9]|16[2567]|17[0-8]|18[0-9]|19[0-35-9])\\d{8})"`
	email string `valid:"email,omitempty"`
}

func TestStruct(t *testing.T) {
	bean := ValidTest{
		Name:  "bea",
		Sex:   "alien",
		Age:   99,
		ID:    2021140777,
		phone: 13322221111,
		email: "5465@qq.com",
	}
	validateStruct, err := govalidator.ValidateStruct(bean)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(validateStruct)
}
