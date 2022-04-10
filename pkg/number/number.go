package number

import (
	"github.com/leekchan/accounting"
	"strings"
)

// UniqueInt ...
func UniqueInt(intSlice []int) []int {
	keys := make(map[int]bool)
	list := []int{}
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

// SpellNumberRec ...
func SpellNumberRec(number int, index int) string {
	var num int
	if number < 0 {
		num = number * -1
	} else {
		num = number
	}

	huruf := []string{"", "satu", "dua", "tiga", "empat", "lima", "enam", "tujuh", "delapan", "sembilan", "sepuluh", "sebelas"}
	var temp string
	if num < 12 {
		if index == 0 {
			temp = strings.Title(huruf[num])
		} else {
			temp = " " + huruf[num]
		}
	} else if num < 20 {
		temp = SpellNumberRec(num-10, index) + " belas"
	} else if num < 100 {
		temp = SpellNumberRec(num/10, index) + " puluh" + SpellNumberRec(num%10, 1)
	} else if num < 200 {
		if index == 0 {
			temp = "seratus" + SpellNumberRec(num-100, 1)
		} else {
			temp = " seratus" + SpellNumberRec(num-100, 1)
		}
	} else if num < 1000 {
		temp = SpellNumberRec(num/100, index) + " ratus" + SpellNumberRec(num%100, 1)
	} else if num < 2000 {
		if index == 0 {
			temp = "seribu" + SpellNumberRec(num-1000, 1)
		} else {
			temp = " seribu" + SpellNumberRec(num-1000, 1)
		}
	} else if num < 1000000 {
		temp = SpellNumberRec(num/1000, index) + "ribu" + SpellNumberRec(num%1000, 1)
	} else if num < 1000000000 {
		temp = SpellNumberRec(num/1000000, index) + " juta" + SpellNumberRec(num%1000000, 1)
	}

	return temp
}

// FormatCurrency ...
func FormatCurrency(value float64, currencySymbol, thousandSeparator, decimalSeparator string, precision int) string {
	if currencySymbol == "IDR" {
		currencySymbol = "Rp "
	}
	ac := accounting.Accounting{Symbol: currencySymbol, Precision: precision, Thousand: thousandSeparator, Decimal: decimalSeparator}

	return ac.FormatMoney(value)
}
