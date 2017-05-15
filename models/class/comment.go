package class

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type Comment struct {
	Id        uint64 `orm:"auto;pk"`
	Content   string `orm:"type(text)"`
	Status    uint8  `orm:"default(0)"`
	Uid       uint64
	Pid       uint64
	CreatedAt time.Time `orm:"auto_now_add;type(datetime)"`
}

func init() {
	orm.RegisterModel(new(Comment))
}
