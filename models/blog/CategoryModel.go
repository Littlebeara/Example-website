
package blog

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type Category struct {
	Id         int64
	Alias      string    `orm:"unique"`
	Title      string    `orm:"size(100)"`
	Content    string    `orm:"type(text);null"` //内容
	Createtime time.Time `orm:"type(datetime);null"`
	Updatetime time.Time `orm:"type(datetime);null"`
	Sort       int64     `orm:"null"`       //排序
	Status     int64     `orm:"default(2)"` //1开启 2关闭
	Siteid     int64     `orm:"default(0)"` //站点ID
	Type       int64     `orm:"default(0)"` //0表示文章 1表示相册
	Image      string    `orm:"null"`       //图片地址，加密，最后为了速度并没有加密
	Pid        int64     `orm:"default(0)"` //父类id
}

func (m *Category) TableName() string {
	return "category"
}

func (m *Category) Read(fields ...string) error {
	if err := orm.NewOrm().Read(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *Category) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *Category) Delete() error {
	if _, err := orm.NewOrm().Delete(m); err != nil {
		return err
	}
	return nil
}

func (m *Category) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(m)
}

func (m *Category) Insert() error {
	if _, err := orm.NewOrm().Insert(m); err != nil {
		return err
	}
	return nil
}
