package main

import (
	"fmt"
	"regexp"
)

func main() {
	// basicTest()
	// testIsIP()
	// web_crawler()
	// do_some_stuff()
	// try_expnad()
	parse_phone_number()
}

func basicTest() {
	soucerString := "John: 2578.34 William: 4567.23 Steve: 5632.18"
	pattern := "[0-9]+.[0-9]+"

	if ok, _ := regexp.Match(pattern, []byte(soucerString)); ok { // second return is error
		fmt.Println("Found")
	}

	re, _ := regexp.Compile(pattern)
	afterReplace := re.ReplaceAllString(soucerString, "##.#")
	fmt.Println(afterReplace)
}

func testIsIP() {
	testCases := []string{
		"163.23.24.111",
		"aa.bb.cc.dd",
		"999.444.555.222",
		"999.444.555.222.666",
		"1.2.3.4"}
	pattern_ip := "^[0-9]{1,3}.[0-9]{1,3}.[0-9]{1,3}.[0-9]{1,3}$"

	for _, testCase := range testCases {
		if match, err := regexp.MatchString(pattern_ip, testCase); err != nil {
			fmt.Println(err)
		} else {
			var yesOrNot string
			if match == false {
				yesOrNot = "not "
			}

			fmt.Printf("Test Case: %s is %sip\n", testCase, yesOrNot)
		}
	}
}

func parse_phone_number() {
	phoneNumber := "(02)2367-8927"

	countryCodePattern := `\(\d{1,2}\)`
	re, _ := regexp.Compile(countryCodePattern)
	find := re.FindString(phoneNumber)
	fmt.Println("Country Code is:", find)

	removeCountryCode := func() string {
		index := re.FindStringIndex(phoneNumber)[0]
		strLen := len("(00)")
		return string([]byte(phoneNumber[index+strLen:]))
	}()

	fmt.Println(removeCountryCode)

	// Expend
}

func try_expand() {
	aabbcc := "aabbcc"
	pattern := `(?P<a1>[a]{2}).*(?P<a2>[b]{2}).*(?P<a3>[c]{2})`

	re, _ := regexp.Compile(pattern)
	for _, sub := range re.FindStringSubmatch(aabbcc) {
		fmt.Println(sub)
	}

	byteaabbcc := []byte(aabbcc)
	var result []byte
	for _, subIndex := range re.FindAllSubmatchIndex(byteaabbcc, -1) {
		result = re.Expand(result, []byte(`$a3--$a2--$a1`), []byte(byteaabbcc), subIndex)
	}

	fmt.Println(string(result))
}
