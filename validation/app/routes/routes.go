// GENERATED CODE - DO NOT EDIT
package routes

import "github.com/revel/revel"


type tApplication struct {}
var Application tApplication


func (_ tApplication) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Application.Index", args).Url
}


type tSample2 struct {}
var Sample2 tSample2


func (_ tSample2) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Sample2.Index", args).Url
}

func (_ tSample2) HandleSubmit(
		username string,
		firstname string,
		lastname string,
		age int,
		password string,
		passwordConfirm string,
		email string,
		emailConfirm string,
		termsOfUse bool,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "username", username)
	revel.Unbind(args, "firstname", firstname)
	revel.Unbind(args, "lastname", lastname)
	revel.Unbind(args, "age", age)
	revel.Unbind(args, "password", password)
	revel.Unbind(args, "passwordConfirm", passwordConfirm)
	revel.Unbind(args, "email", email)
	revel.Unbind(args, "emailConfirm", emailConfirm)
	revel.Unbind(args, "termsOfUse", termsOfUse)
	return revel.MainRouter.Reverse("Sample2.HandleSubmit", args).Url
}


type tSample3 struct {}
var Sample3 tSample3


func (_ tSample3) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Sample3.Index", args).Url
}

func (_ tSample3) HandleSubmit(
		user interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "user", user)
	return revel.MainRouter.Reverse("Sample3.HandleSubmit", args).Url
}


type tSample4 struct {}
var Sample4 tSample4


func (_ tSample4) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Sample4.Index", args).Url
}

func (_ tSample4) HandleSubmit(
		user interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "user", user)
	return revel.MainRouter.Reverse("Sample4.HandleSubmit", args).Url
}


type tSample1 struct {}
var Sample1 tSample1


func (_ tSample1) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Sample1.Index", args).Url
}

func (_ tSample1) HandleSubmit(
		username string,
		firstname string,
		lastname string,
		age int,
		password string,
		passwordConfirm string,
		email string,
		emailConfirm string,
		termsOfUse bool,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "username", username)
	revel.Unbind(args, "firstname", firstname)
	revel.Unbind(args, "lastname", lastname)
	revel.Unbind(args, "age", age)
	revel.Unbind(args, "password", password)
	revel.Unbind(args, "passwordConfirm", passwordConfirm)
	revel.Unbind(args, "email", email)
	revel.Unbind(args, "emailConfirm", emailConfirm)
	revel.Unbind(args, "termsOfUse", termsOfUse)
	return revel.MainRouter.Reverse("Sample1.HandleSubmit", args).Url
}


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


