package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["app/controllers:CategoryController"] = append(beego.GlobalControllerRouter["app/controllers:CategoryController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			Params:           nil})

	beego.GlobalControllerRouter["app/controllers:CategoryController"] = append(beego.GlobalControllerRouter["app/controllers:CategoryController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params:           nil})

	beego.GlobalControllerRouter["app/controllers:CategoryController"] = append(beego.GlobalControllerRouter["app/controllers:CategoryController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			Params:           nil})

	beego.GlobalControllerRouter["app/controllers:CategoryController"] = append(beego.GlobalControllerRouter["app/controllers:CategoryController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params:           nil})

	beego.GlobalControllerRouter["app/controllers:CategoryController"] = append(beego.GlobalControllerRouter["app/controllers:CategoryController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params:           nil})

	beego.GlobalControllerRouter["app/controllers:CommentController"] = append(beego.GlobalControllerRouter["app/controllers:CommentController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			Params:           nil})

	beego.GlobalControllerRouter["app/controllers:CommentController"] = append(beego.GlobalControllerRouter["app/controllers:CommentController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params:           nil})

	beego.GlobalControllerRouter["app/controllers:CommentController"] = append(beego.GlobalControllerRouter["app/controllers:CommentController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			Params:           nil})

	beego.GlobalControllerRouter["app/controllers:CommentController"] = append(beego.GlobalControllerRouter["app/controllers:CommentController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params:           nil})

	beego.GlobalControllerRouter["app/controllers:CommentController"] = append(beego.GlobalControllerRouter["app/controllers:CommentController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params:           nil})

	beego.GlobalControllerRouter["app/controllers:MediaController"] = append(beego.GlobalControllerRouter["app/controllers:MediaController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			Params:           nil})

	beego.GlobalControllerRouter["app/controllers:MediaController"] = append(beego.GlobalControllerRouter["app/controllers:MediaController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params:           nil})

	beego.GlobalControllerRouter["app/controllers:MediaController"] = append(beego.GlobalControllerRouter["app/controllers:MediaController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			Params:           nil})

	beego.GlobalControllerRouter["app/controllers:MediaController"] = append(beego.GlobalControllerRouter["app/controllers:MediaController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params:           nil})

	beego.GlobalControllerRouter["app/controllers:MediaController"] = append(beego.GlobalControllerRouter["app/controllers:MediaController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params:           nil})

	beego.GlobalControllerRouter["app/controllers:PostController"] = append(beego.GlobalControllerRouter["app/controllers:PostController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			Params:           nil})

	beego.GlobalControllerRouter["app/controllers:PostController"] = append(beego.GlobalControllerRouter["app/controllers:PostController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params:           nil})

	beego.GlobalControllerRouter["app/controllers:PostController"] = append(beego.GlobalControllerRouter["app/controllers:PostController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			Params:           nil})

	beego.GlobalControllerRouter["app/controllers:PostController"] = append(beego.GlobalControllerRouter["app/controllers:PostController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params:           nil})

	beego.GlobalControllerRouter["app/controllers:PostController"] = append(beego.GlobalControllerRouter["app/controllers:PostController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params:           nil})

	beego.GlobalControllerRouter["app/controllers:TagController"] = append(beego.GlobalControllerRouter["app/controllers:TagController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			Params:           nil})

	beego.GlobalControllerRouter["app/controllers:TagController"] = append(beego.GlobalControllerRouter["app/controllers:TagController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params:           nil})

	beego.GlobalControllerRouter["app/controllers:TagController"] = append(beego.GlobalControllerRouter["app/controllers:TagController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			Params:           nil})

	beego.GlobalControllerRouter["app/controllers:TagController"] = append(beego.GlobalControllerRouter["app/controllers:TagController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params:           nil})

	beego.GlobalControllerRouter["app/controllers:TagController"] = append(beego.GlobalControllerRouter["app/controllers:TagController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params:           nil})

	beego.GlobalControllerRouter["app/controllers:UserController"] = append(beego.GlobalControllerRouter["app/controllers:UserController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			Params:           nil})

	beego.GlobalControllerRouter["app/controllers:UserController"] = append(beego.GlobalControllerRouter["app/controllers:UserController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params:           nil})

	beego.GlobalControllerRouter["app/controllers:UserController"] = append(beego.GlobalControllerRouter["app/controllers:UserController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			Params:           nil})

	beego.GlobalControllerRouter["app/controllers:UserController"] = append(beego.GlobalControllerRouter["app/controllers:UserController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params:           nil})

	beego.GlobalControllerRouter["app/controllers:UserController"] = append(beego.GlobalControllerRouter["app/controllers:UserController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params:           nil})

}
