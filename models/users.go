package models

import (
	"fmt"

	"github.com/beego/beego/v2/client/orm"
	_ "github.com/lib/pq"
)

type Users struct {
	Id       int    `orm:"column(id);auto;pk"`
	Name     string `orm:"column(name)"`
	Email    string `orm:"column(email)"  valid:"Email"`
	Phone    string `orm:"column(phone)"  valid:"Numeric;MinSize(6);MaxSize(11)"`
	Password string `orm:"column(password)"`
}

func (t *Users) TableName() string {
	return "users"
}

func init() {
	orm.RegisterModel(new(Users))
	orm.RegisterDriver("postgres", orm.DRPostgres)
	orm.RegisterDataBase("default", "postgres", "postgres://postgres:postgres@localhost:5432/testz?sslmode=disable")
}

func AddUsers(m *Users) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

func GetUsersById(id int) (v *Users, err error) {
	o := orm.NewOrm()
	v = &Users{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

func DeleteUsers(id int) (err error) {
	o := orm.NewOrm()
	v := Users{Id: id}
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Users{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
