// GENERATED CODE - DO NOT EDIT
package routes

import "github.com/revel/revel"


type tApp struct {}
var App tApp


func (_ tApp) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("App.Index", args).Url
}


type tCommonController struct {}
var CommonController tCommonController



type tStatic struct {}
var Static tStatic


func (_ tStatic) Serve(
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.Serve", args).Url
}

func (_ tStatic) ServeModule(
		moduleName string,
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "moduleName", moduleName)
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.ServeModule", args).Url
}


type tTestRunner struct {}
var TestRunner tTestRunner


func (_ tTestRunner) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("TestRunner.Index", args).Url
}

func (_ tTestRunner) Suite(
		suite string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "suite", suite)
	return revel.MainRouter.Reverse("TestRunner.Suite", args).Url
}

func (_ tTestRunner) Run(
		suite string,
		test string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "suite", suite)
	revel.Unbind(args, "test", test)
	return revel.MainRouter.Reverse("TestRunner.Run", args).Url
}

func (_ tTestRunner) List(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("TestRunner.List", args).Url
}


type tUserController struct {}
var UserController tUserController


func (_ tUserController) Init(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("UserController.Init", args).Url
}

func (_ tUserController) List(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("UserController.List", args).Url
}

func (_ tUserController) Get(
		user_id int64,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "user_id", user_id)
	return revel.MainRouter.Reverse("UserController.Get", args).Url
}

func (_ tUserController) Create(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("UserController.Create", args).Url
}

func (_ tUserController) Update(
		user_id int64,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "user_id", user_id)
	return revel.MainRouter.Reverse("UserController.Update", args).Url
}

func (_ tUserController) Authenticate(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("UserController.Authenticate", args).Url
}


type tBookController struct {}
var BookController tBookController


func (_ tBookController) Init(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("BookController.Init", args).Url
}

func (_ tBookController) List(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("BookController.List", args).Url
}

func (_ tBookController) PublicBooks(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("BookController.PublicBooks", args).Url
}

func (_ tBookController) MyBooks(
		user_id int64,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "user_id", user_id)
	return revel.MainRouter.Reverse("BookController.MyBooks", args).Url
}

func (_ tBookController) Get(
		book_id int64,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "book_id", book_id)
	return revel.MainRouter.Reverse("BookController.Get", args).Url
}

func (_ tBookController) Create(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("BookController.Create", args).Url
}

func (_ tBookController) Update(
		book_id int64,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "book_id", book_id)
	return revel.MainRouter.Reverse("BookController.Update", args).Url
}


