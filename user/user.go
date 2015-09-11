package user

import "database/sql"

// NODeny
// id             : id учетной записи клиента
// name           : (теперь) логин/учетное имя
// passwd         : зашифрованный пароль
// grp            : группа (см. табл. user_grp)
// contract       : договор
// contract_date  : timestamp заключения договора
// state          : состояние доступа вкл/выкл
// balance        : баланс
// limit_balance  : граница отключения
// block_if_limit : отключать ли запись при достижении границы отключения
// modify_time    : timestamp последней модификации записи
// fio            : ФИО
// discount       : % скидки
// cstate         : техническое состояние (настроить/ремонт/вирусы...)
// cstate_time    : timestamp последнего изменения поля cstate
// comment        : комментарий
// lstate         : нужна авторизация/всегда онлайн;
// address        : адрес доставки счетов
// mail           : 1 - доставлять счета по почте
type User struct {
	ID           int64          `orm:"auto;pk" json:"id"`
	Name         string         `orm:"size(64)" json:"name"`
	Passwd       string         `orm:"size(64)" json:"-"`
	Grp          int            `orm:"default(0);index"`
	Contract     string         `orm:"type(text)" json:"-"`
	ContractDate int            `orm:"default(0)" json:"-"`
	State        string         `orm:"size(3);index" json:"-"`
	Balance      float64        `orm:"digits(12);decimals(2);default(0)" json:"-"`
	LimitBalance float64        `orm:"digits(12);decimals(2);default(0)" json:"-"`
	BlockIfLimit int            `orm:"default(1)" json:"-"`
	ModifyTime   int            `orm:"default(0)" json:"-"`
	Fio          string         `orm:"type(text)" json:"fio"`
	Discount     int            `orm:"default(0)" json:"-"`
	Cstate       int            `orm:"default(0)" json:"-"`
	CstateTime   int            `orm:"default(0)" json:"-"`
	Comment      string         `orm:"type(text)" json:"-"`
	Lstate       int            `orm:"default(0)" json:"-"`
	Address      string         `orm:"type(text);null" json:"adress"`
	Mail         int8           `orm:"default(0);index" json:"email"`
	Rands        sql.NullString `orm:"size(10)null" json:"-"`
	User         *UserDjango    `orm:"rel(one);null;index" json:"info"`
	// Profile     *Profile   `orm:"rel(one)"` // OneToOne relation
	// UserId int `orm:"null;index"`
}

func (this *User) TableName() string {
	return "users"
}
