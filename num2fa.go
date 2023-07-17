package main

import (
	"strings"
)

var units = []string{
	"",
	"یک",
	"دو",
	"سه",
	"چهار",
	"پنج",
	"شش",
	"هفت",
	"هشت",
	"نه",
	"ده",
	"یازده",
	"دوازده",
	"سیزده",
	"چهارده",
	"پانزده",
	"شانزده",
	"هفده",
	"هجده",
	"نوزده",
}

var tens = []string{
	"",
	"ده",
	"بیست",
	"سی",
	"چهل",
	"پنجاه",
	"شصت",
	"هفتاد",
	"هشتاد",
	"نود",
}

var hundreds = []string{
	"",
	"صد",
	"دویست",
	"سیصد",
	"چهارصد",
	"پانصد",
	"ششصد",
	"هفتصد",
	"هشتصد",
	"نهصد",
}

var chunks = []string{
	"",
	"هزار",
	"میلیون",
	"میلیارد",
	"بیلیون",
	"بیلیارد",
	"تریلیون",
	"تریلیارد",
	"کوآدریلیون",
	"کادریلیارد",
	"کوینتیلیون",
	"کوانتینیارد",
	"سکستیلیون",
	"سکستیلیارد",
	"سپتیلیون",
	"سپتیلیارد",
	"اکتیلیون",
	"اکتیلیارد",
	"نانیلیون",
	"نانیلیارد",
	"دسیلیون",
}

func threeDigitToPersianWord(number int) string {
	result := make([]string, 0, 3)
	chunkUnits := number % 10
	chunkTens := (number / 10) % 10
	chunkHundreds := (number / 100) % 10

	if chunkHundreds != 0 {
		result = append(result, hundreds[chunkHundreds])
	}

	if chunkTens == 1 {
		result = append(result, units[chunkUnits+10])
	} else {
		if chunkTens != 0 {
			result = append(result, tens[chunkTens])
		}
		if chunkUnits != 0 {
			result = append(result, units[chunkUnits])
		}
	}

	return strings.Join(result, " و ")
}

func Num2fa(number int) string {
	if number == 0 {
		return "صفر"
	}

	negative := false

	if number < 0 {
		negative = true
		number *= -1
	}

	count := 0
	result := make([]string, 0, 3)

	for number > 0 {
		chunk := number % 1000
		if chunk == 0 {
			number /= 1000
			count++
			continue
		}

		chunkWords := threeDigitToPersianWord(chunk)
		if chunk != 0 {
			result = append([]string{chunkWords + " " + chunks[count]}, result...)
		} else {
			result = append([]string{chunkWords}, result...)
		}
		number /= 1000
		count++
	}

	resultStr := strings.Join(result, " و ")
	resultStr = strings.Trim(resultStr, " ")

	if negative {
		resultStr = "منفی " + resultStr
	}

	return resultStr
}
