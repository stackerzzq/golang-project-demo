package class

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type Media struct {
	Id        uint64 `orm:"auto:pk"`
	Key       string
	Val       string `orm:"type(text)"`
	Status    uint8  `orm:"default(0)"`
	Uid       uint64
	CreatedAt time.Time `orm:"auto_now_add;type(datetime)"`
	UpdatedAt time.Time `orm:"auto_now;type(datetime)"`
}

func init() {
	orm.RegisterModel(new(Media))
}
