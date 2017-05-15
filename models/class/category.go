package class

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type Category struct {
	Id        uint64 `orm:"auto;pk"`
	Name      string `orm:"size(20)"`
	Status    uint8  `orm:"default(0)"`
	Uid       uint64
	CreatedAt time.Time `orm:"auto_now_add;type(datetime)"`
	UpdatedAt time.Time `orm:"auto_now;type(datetime)"`
}

func init() {
	orm.RegisterModel(new(Category))
}
