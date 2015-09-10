package card

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func ValidCard(card string, pref string) (b bool, res string) {
	p := regexp.MustCompile("^(\\d{10})$")
	r := p.FindStringSubmatch(card)
	if len(r) > 0 {
		return true, fmt.Sprintf("%s%s", pref, r[1])
	} else {
		return false, card
	}
}

// Возвращает номер карты без нулей в переди
// @return Card format 0001234567 -> 1234567
func GetCard(card string) (res string) {
	v, _ := strconv.ParseInt(card, 10, 32)
	return fmt.Sprintf("%d", v)
}

// Возвращает номер карты в полном формате
// @return true/false, Card format 0001234567
func GetFullCard(card string) (bool, string) {
	cl := len(card)
	if cl > 10 {
		return false, card
	} else if cl < 10 {
		lenCard := 10 - cl
		var data []string
		for i := 0; i < lenCard; i++ {
			data = append(data, "0")
		}
		data = append(data, card)
		if ok, vcard := ValidCard(strings.Join(data, ""), ""); !ok {
			return false, card
		} else {
			return true, vcard
		}
	}
	if ok, vcard := ValidCard(card, ""); !ok {
		return false, card
	} else {
		return true, vcard
	}
}

func GetPersonalNumber(pn string) string {
	return pn
}
