package user

import "time"

// https://github.com/beego/admin/blob/master/src/models/UserModel.go
type UserDjango struct {
	ID          int64     `orm:"auto;pk" json:"id"`
	Password    string    `orm:"size(128)" json:"-"`
	LastLogin   time.Time `orm:"auto_now_add;column(last_login);type(datetime)" json:"-"`
	IsSuperuser int8      `orm:"column(is_superuser)" json:"-"`
	Username    string    `orm:"size(30)" valid:"Required;MaxSize(30);MinSize(6)" json:"-"`
	FirstName   string    `orm:"size(30);column(first_name)" json:"-"`
	LastName    string    `orm:"size(30);column(last_name)" json:"-"`
	Email       string    `orm:"size(75)" json:"email"`
	IsStaff     int8      `orm:"column(is_staff)" json:"-"`
	IsActive    int8      `orm:"column(is_active)" json:"-"`
	DateJoined  time.Time `orm:"auto_now_add;column(date_joined);type(datetime)" json:"-"`
	// Profile     *Profile   `orm:"rel(one)"` // OneToOne relation
	// DecPassword string    `orm:"-" json:"test"`
}

func (this *UserDjango) TableName() string {
	return "auth_user"
}
