package controllers

import (
	"github.com/astaxie/beego"
)

// BaseController operations for Base
type BaseController struct {
	beego.Controller
	Ok      bool        `json:"code"`
	Content interface{} `json:"message"`
}
