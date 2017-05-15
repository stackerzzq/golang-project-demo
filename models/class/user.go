package class

import (
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type User struct {
	Id        uint64 `orm:"auto;pk"`
	Username  string `orm:"size(20)"`
	Password  string `orm:"size(40)"`
	Email     string `orm:"size(20)"`
	Gender    uint8  `orm:"default(0)"`
	Headimg   string
	CreatedAt time.Time `orm:"auto_now_add;type(datetime)"`
	UpdatedAt time.Time `orm:"auto_now;type(datetime)"`
}

func (u User) Create() (id uint64, err error) {
	o := orm.NewOrm()
	if id, err = o.Insert(&u); err != nil {
		beego.Info(err)
	}
	return
}

func (u User) Delete() (num uint64, err error) {
	o := orm.NewOrm()
	if num, err = o.Delete(&u); err != nil {
		beego.Info(err)
	}
	return
}

func (u User) Update(fields ...string) (err error) {
	o := orm.NewOrm()
	if _, err = o.Update(&u, fields); err != nil {
		beego.Info(err)
	}
	return
}

func (u User) QueryAll() (us interface{}, err error) {
	return
}

func (u User) QueryOne() (u User, err error) {
	o := orm.NewOrm()
	err = o.Read(&u)
	return
}

func init() {
	orm.RegisterModel(new(User))
}
