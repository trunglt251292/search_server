package util

import (
	"encoding/json"
	"regexp"
	"strings"

	"github.com/thoas/go-funk"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Stringify ...
func Stringify(data interface{}) string {
	b, _ := json.Marshal(data)
	return string(b)
}

// ToBytes ...
func ToBytes(data interface{}) []byte {
	b, _ := json.Marshal(data)
	return b
}

// PhoneNumber ...
type PhoneNumber struct {
	Number      string
	CountryCode string
	Full        string
}

const (
	RegexPhoneNumber = `^(0|\+84|84)(\d{9}|\d{10})$`
)

// PhoneNumberIsValid check phone number is valid
func PhoneNumberIsValid(phone string) bool {
	r := regexp.MustCompile(RegexPhoneNumber)
	return r.MatchString(phone)
}

// PhoneNumberFormatFromPhone ...
func PhoneNumberFormatFromPhone(phone string) PhoneNumber {
	var phoneNumber PhoneNumber
	isValid := PhoneNumberIsValid(phone)
	if !isValid {
		return phoneNumber
	}

	phone = strings.Replace(phone, " ", "", -1)

	switch true {
	case string(phone[0]) == "0":
		phone = strings.Replace(phone, "0", "", 1)
	case phone[0:3] == "+84":
		phone = strings.Replace(phone, "+84", "", 1)
	case phone[0:2] == "84":
		phone = strings.Replace(phone, "84", "", 1)
	}
	// remove 0, 84, +84
	//if string(phone[0]) == "0" {
	//	phone = strings.Replace(phone, "0", "", 1)
	//}
	//
	//if phone[0:3] == "+84" {
	//	phone = strings.Replace(phone, "+84", "", 1)
	//}
	//
	//if phone[0:2] == "84" && len(phone) == 11 {
	//	phone = strings.Replace(phone, "84", "", 1)
	//}

	phoneNumber.Number = phone
	phoneNumber.CountryCode = "+84"
	phoneNumber.Full = phoneNumber.CountryCode + phoneNumber.Number
	return phoneNumber
}

// ConvertStringsToObjectIDs ...
func ConvertStringsToObjectIDs(strValues []string) []primitive.ObjectID {
	return funk.Map(strValues, func(item string) primitive.ObjectID {
		return GetAppIDFromHex(item)
	}).([]primitive.ObjectID)
}

// GetAppIDFromHex ...
func GetAppIDFromHex(s string) primitive.ObjectID {
	id, _ := primitive.ObjectIDFromHex(s)
	return id
}
