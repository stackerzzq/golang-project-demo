package class

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type Post struct {
	Id        uint64 `orm:"auto;pk"`
	Title     string `orm:"size(50)"`
	Desc      string `orm:"type(text)"`
	Content   string `orm:"type(text)"`
	Status    uint8  `orm:"default(0)"`
	Uid       uint64
	CreatedAt time.Time `orm:"auto_now_add;type(datetime)"`
	UpdatedAt time.Time `orm:"auto_now;type(datetime)"`
}

func init() {
	orm.RegisterModel(new(Post))
}
