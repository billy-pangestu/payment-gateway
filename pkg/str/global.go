package str

import (
	"bytes"
	"errors"
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// StringToBool ...
func StringToBool(data string) bool {
	res, err := strconv.ParseBool(data)
	if err != nil {
		res = false
	}

	return res
}

// StringToBoolString ...
func StringToBoolString(data string) string {
	res, err := strconv.ParseBool(data)
	if err != nil {
		return "false"
	}

	return strconv.FormatBool(res)
}

// StringToInt ...
func StringToInt(data string) int {
	res, err := strconv.Atoi(data)
	if err != nil {
		res = 0
	}

	return res
}

// StringToFloat ...
func StringToFloat(data string) float64 {
	res, err := strconv.ParseFloat(data, 64)
	if err != nil {
		res = 0
	}

	return res
}

// BoolToInt ...
func BoolToInt(data bool) int {
	if data {
		return 1
	}

	return 0
}

// IntToBool ...
func IntToBool(data int) bool {
	if data == 1 {
		return true
	}

	return false
}

// Float64ToString ...
func Float64ToString(data float64) string {
	return fmt.Sprintf("%g", data)
}

// Contains ...
func Contains(slices []string, comparizon string) bool {
	for _, a := range slices {
		if a == comparizon {
			return true
		}
	}

	return false
}

// CensorName ...
func CensorName(name string, censorFactor float64) string {
	words := strings.Fields(name)
	for i, s := range words {
		if len(s) > 0 {
			toCensor := int(math.Floor(float64(len(s)) * censorFactor))
			s = s[0:len(s)-toCensor] + strings.Repeat("*", toCensor)
		}
		words[i] = s
	}
	return strings.Join(words, " ")
}

// EmptyString ...
func EmptyString(text string) *string {
	if text == "" {
		return nil
	}
	return &text
}

// EmptyInt ...
func EmptyInt(number int) *int {
	if number == 0 {
		return nil
	}
	return &number
}

// CheckNpwpFormat ...
func CheckNpwpFormat(text string) bool {
	if len(text) < 20 {
		return false
	}

	part1 := text[0:2]
	part2 := text[3:6]
	part3 := text[7:10]
	part4 := text[11:12]
	part5 := text[13:16]
	part6 := text[17:20]

	if _, err := strconv.Atoi(part1); err != nil {
		return false
	}
	if _, err := strconv.Atoi(part2); err != nil {
		return false
	}
	if _, err := strconv.Atoi(part3); err != nil {
		return false
	}
	if _, err := strconv.Atoi(part4); err != nil {
		return false
	}
	if _, err := strconv.Atoi(part5); err != nil {
		return false
	}
	if _, err := strconv.Atoi(part6); err != nil {
		return false
	}

	return true
}

// GetNpwpNumber ...
func GetNpwpNumber(text string) (res string, err error) {
	if len(text) < 20 {
		return res, errors.New("Invalid npwp number format")
	}

	part1 := text[0:2]
	part2 := text[3:6]
	part3 := text[7:10]
	part4 := text[11:12]
	part5 := text[13:16]
	part6 := text[17:20]

	if _, err := strconv.Atoi(part1); err != nil {
		return res, err
	}
	if _, err := strconv.Atoi(part2); err != nil {
		return res, err
	}
	if _, err := strconv.Atoi(part3); err != nil {
		return res, err
	}
	if _, err := strconv.Atoi(part4); err != nil {
		return res, err
	}
	if _, err := strconv.Atoi(part5); err != nil {
		return res, err
	}
	if _, err := strconv.Atoi(part6); err != nil {
		return res, err
	}

	res = part1 + part2 + part3 + part4 + part5 + part6

	return res, err
}

// DecodeNpwpNumber ...
func DecodeNpwpNumber(text string) (res string) {
	if len(text) != 15 {
		return text
	}

	part1 := text[0:2]
	part2 := text[2:5]
	part3 := text[5:8]
	part4 := text[8:9]
	part5 := text[9:12]
	part6 := text[12:15]
	res = part1 + "." + part2 + "." + part3 + "." + part4 + "-" + part5 + "." + part6

	return res
}

// CheckDate ...
func CheckDate(data string, inPast bool) bool {
	d, err := time.Parse("2006-01-02", data)
	if err != nil {
		return false
	}

	if inPast {
		if d.After(time.Now()) {
			return false
		}
	}

	return true
}

// CheckNumeric ...
func CheckNumeric(data string, positive bool) bool {
	res, err := strconv.Atoi(data)
	if err != nil {
		return false
	}
	if positive && res < 0 {
		return false
	}

	return true
}

// PadNumberWithZero ...
func PadNumberWithZero(value int, pad string) string {
	return fmt.Sprintf("%0"+pad+"d", value)
}

// InterfaceStringToString ...
func InterfaceStringToString(data interface{}, key string) string {
	if data == nil || key == "" {
		return ""
	}

	res := fmt.Sprintf("%v", data.(map[string]interface{})[key])
	if res == "<nil>" {
		res = ""
	}

	return res
}

// InsertDash ...
func InsertDash(sourceText string, index int, symbol rune) string {
	var buffer bytes.Buffer
	for i, rune := range sourceText {
		buffer.WriteRune(rune)

		if i == index {
			buffer.WriteRune(symbol)
		}
	}

	return buffer.String()
}

// RemoveSymbol ...
func RemoveSymbol(text string) string {
	reg, _ := regexp.Compile("[^a-zA-Z0-9 ]+")

	return reg.ReplaceAllString(text, "")
}

// CheckEmail ...
func CheckEmail(text string) bool {
	re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

	return re.MatchString(text)
}

// CensorPhoneFormat ...
func CensorPhoneFormat(name string, censorFactor float64, firstDigit int) string {
	if len(name) < firstDigit+1 {
		return name
	}

	first := name[0:firstDigit]
	name = name[firstDigit:len(name)]
	words := strings.Fields(name)
	for i, s := range words {
		if len(s) > 0 {
			toCensor := int(math.Floor(float64(len(s)) * censorFactor))
			s = strings.Repeat("*", toCensor) + "-" + s[toCensor:len(s)]
		}
		words[i] = s
	}

	return "+" + first + "-" + strings.Join(words, " ")
}

// CensorString ...
func CensorString(name string, censorFactor float64, firstDigit int) string {
	if len(name) < firstDigit+1 {
		return name
	}

	first := name[0:firstDigit]
	name = name[firstDigit:len(name)]
	words := strings.Fields(name)
	for i, s := range words {
		if len(s) > 0 {
			toCensor := int(math.Floor(float64(len(s)) * censorFactor))
			s = strings.Repeat("*", toCensor) + s[toCensor:len(s)]
		}
		words[i] = s
	}

	return first + strings.Join(words, " ")
}

// DefaultData ...
func DefaultData(data, defaultData string) string {
	if data == "" {
		return defaultData
	}

	return data
}

// DefaultDataInWhitelist ...
func DefaultDataInWhitelist(data string, whitelist []string, defaultData string) string {
	if data == "" || !Contains(whitelist, data) {
		return defaultData
	}

	return data
}

// EmptyErr ...
func EmptyErr(err error) string {
	if err != nil {
		return err.Error()
	}

	return ""
}

// FirstWords ...
func FirstWords(value string, count int) string {
	// Loop over all indexes in the string.
	for i := range value {
		// If we encounter a space, reduce the count.
		if value[i] == ' ' {
			count--
			// When no more words required, return a substring.
			if count == 0 {
				return value[0:i]
			}
		}
	}

	// Return the entire string.
	return value
}

// Unique ...
func Unique(intSlice []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}

	return list
}

// DefaultString ...
func DefaultString(data, def string) string {
	if data != "" {
		return data
	}

	return def
}

// Normalize ...
func Normalize(data string) string {
	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		return ""
	}

	return reg.ReplaceAllString(data, "")
}

// GetLast ...
func GetLast(data, indexString string) string {
	dataArr := strings.Split(data, indexString)
	if len(dataArr) == 0 {
		return ""
	}

	return dataArr[len(dataArr)-1]
}

// GetExtentionByURL ...
func GetExtentionByURL(name string) string {
	nameArr := strings.Split(name, "?")
	if len(nameArr) == 0 {
		return ""
	}

	extentionArr := strings.Split(nameArr[0], ".")

	return "." + extentionArr[len(extentionArr)-1]
}
