package validator_ext_test

import (
	"fmt"
	ve "github.com/darwinOrg/go-validator-ext"
	"testing"
)

type User struct {
	Name      string   `binding:"required,regex=^[a-zA-Z]+$"`
	Age       int      `binding:"max=35,min=22"`
	Hobby     []string `binding:"maxLength=4,minLength=2"`
	Email     string   `binding:"email"`
	Mobile    string   `binding:"isMobile"`
	Country   string   `binding:"isCountry"`
	Birth     string   `binding:"isDate"`
	CreatedAt string   `binding:"isDatetime"`
	Status    int32    `binding:"oneof=0 1"`
	Works     []*Work  `binding:"dive"`
}

type Work struct {
	Name string `binding:"maxLength=10,unique=Name"`
}

var user = &User{
	Name:      "cb",
	Age:       30,
	Hobby:     []string{"1", "2", "3"},
	Mobile:    "+8615901431753",
	Country:   "BONAIRE_SINT_EUSTATIUS_AND_SABA",
	Email:     "123456789@163.com",
	Birth:     "1990-10-16",
	CreatedAt: "2023-03-23 18:00:00",
	Status:    0,
	Works: []*Work{
		{Name: "abc1"},
		{Name: "abc2"},
		{Name: "swergwsegawegwgwg"},
		{Name: "jtrhsgberhehe"},
	},
}

func TestNotNull(t *testing.T) {
	user.Name = ""
	doValidate(user)
}

func TestRegex(t *testing.T) {
	user.Name = "cb01"
	doValidate(user)
}

func TestMaxValue(t *testing.T) {
	user.Age = 36
	doValidate(user)
}

func TestMinValue(t *testing.T) {
	user.Age = 20
	doValidate(user)
}

func TestMaxLength(t *testing.T) {
	//user.Hobby = []string{"1", "2", "3", "4", "5"}
	doValidate(user)
}

func TestMinLength(t *testing.T) {
	user.Hobby = []string{"1"}
	doValidate(user)
}

func TestIsMobile(t *testing.T) {
	user.Mobile = "+861590"
	doValidate(user)
}

func TestIsCountry(t *testing.T) {
	user.Country = "XXX"
	doValidate(user)
}

func TestIsEmail(t *testing.T) {
	user.Email = "123456789"
	doValidate(user)
}

func TestIsDate(t *testing.T) {
	user.Birth = "1990/10/16"
	doValidate(user)
}

func TestIsDateTime(t *testing.T) {
	user.CreatedAt = "2023-03-23"
	doValidate(user)
}

func TestOneOf(t *testing.T) {
	user.Status = 2
	doValidate(user)
}

func doValidate(user *User) {
	err := ve.NewCustomValidator().Struct(user)
	transErrors := ve.TranslateError(err, "zh")
	for _, v := range transErrors {
		fmt.Println(v)
	}
}
