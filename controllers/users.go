package controllers

import (
	"fmt"
	"testz/models"

	"crypto/sha256"
	"encoding/json"

	"github.com/beego/beego/logs"
	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/validation"
)

type UsersController struct {
	beego.Controller
}

func init() {
	InitLoger()
}

func InitLoger() {
	logs.NewLogger(10000)
	logs.SetLogger(logs.AdapterFile, `{"filename":"testz.log","level":7,"maxlines":0,"maxsize":0,"daily":true,"maxdays":10,"color":true}`)
}

func (c *UsersController) Get() {
	c.TplName = "users/create.tpl"
}

func (c *UsersController) Create() {
	var (
		v   models.Users
		err error
	)
	c.TplName = "users/create.tpl"
	c.Data["model"] = &v

	if c.Ctx.Request.Method == "POST" {
		if err = c.ParseForm(&v); err != nil {
			logs.Error("Ошибка ParseForm:" + err.Error())
			return
		}
		valid := validation.Validation{}
		isValid, _ := valid.Valid(v)
		if !isValid {
			logs.Error("Ошибка validation err")
			fmt.Printf("Проверьте введенные данные: isValid ")
			c.Data["errors"] = valid.ErrorsMap
			return
		}
		h := sha256.New()
		h.Write([]byte(v.Password))
		v.Password = fmt.Sprintf("%x", h.Sum(nil))
		if _, err = models.AddUsers(&v); err != nil {
			logs.Error("Ошибка сохранения в БД:" + err.Error())
			fmt.Printf("Ошибка сохранения в БД %s ", err)
			return
		}
		b, _ := json.Marshal(v)
		c.Redirect("/users/index", 302)
		ConnectRb(b)
	}
}

func (c *UsersController) Index() {
	o := orm.NewOrm()
	var index []*models.Users
	num, err := o.QueryTable("users").All(&index)
	fmt.Printf("Returned Rows Num: %d, %s", num, err)
	if err != nil {
		logs.Error("Ошибка QueryTable:" + err.Error())
		panic(err)
	}
	c.Data["index"] = index
	c.TplName = "users/index.tpl"
}

func (c *UsersController) Delete() {
	id, err := c.GetInt("Id")
	if err != nil {
		fmt.Printf("err  %s", err)
		logs.Error("Ошибка GetInt:" + err.Error())
		return
	}
	if err := models.DeleteUsers(id); err != nil {
		fmt.Printf("err  %s", err)
		logs.Error("Ошибка DeleteUsers:" + err.Error())
		return
	}
	logs.Info("Пользователь удалён")
	c.Redirect("/users/index", 302)
}
