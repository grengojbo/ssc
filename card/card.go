package card

import (
	"database/sql"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/grengojbo/ssc/user"
)

type Cards struct {
	ID             int64          `orm:"column(userid);auto;pk" json:"id"`
	PersonalNumber string         `orm:"column(badgenumber);size(20);index" json:"personal_number"`
	Card           sql.NullString `orm:"column(Card);size(20);index;null" json:"card"`   // 3 UserInfo (Card) Card Number(Internal)
	FirstName           sql.NullString `orm:"column(name);size(24);null" json:"name"`         // 2 First Name
	LastNname          sql.NullString `orm:"column(lastname);size(20);null" json:"lastname"` // 1 Last Name
	Father         *user.User     `orm:"column(father_id);null;index;rel(fk)" json:"father"`
	ChangeTime     time.Time      `orm:"auto_now;null;type(datetime)" json:"change_time"`
	CreateTime     time.Time      `orm:"auto_now_add;null;type(datetime)" json:"create_time"`
	Street         sql.NullString `orm:"column(street);size(100);null" json:"street"`
	Gender         sql.NullString `orm:"column(gender);size(2);null" json:"gender"`
	HomeAddress    sql.NullString `orm:"column(homeaddress);size(100);null" json:"homeaddress"`
}

func (this *Cards) TableName() string {
	return "userinfo"
}

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
