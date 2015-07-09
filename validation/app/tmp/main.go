// GENERATED CODE - DO NOT EDIT
package main

import (
	"flag"
	"reflect"
	"github.com/revel/revel"
	controllers0 "github.com/revel/modules/static/app/controllers"
	_ "github.com/revel/modules/testrunner/app"
	controllers1 "github.com/revel/modules/testrunner/app/controllers"
	models "github.com/revel/samples/validation/app/models"
	controllers "validation/app/controllers"
	tests "validation/tests"
	"github.com/revel/revel/testing"
)

var (
	runMode    *string = flag.String("runMode", "", "Run mode.")
	port       *int    = flag.Int("port", 0, "By default, read from app.conf")
	importPath *string = flag.String("importPath", "", "Go Import Path for the app.")
	srcPath    *string = flag.String("srcPath", "", "Path to the source root.")

	// So compiler won't complain if the generated code doesn't reference reflect package...
	_ = reflect.Invalid
)

func main() {
	flag.Parse()
	revel.Init(*runMode, *importPath, *srcPath)
	revel.INFO.Println("Running revel server")
	
	revel.RegisterController((*controllers.Application)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Index",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
					10: []string{ 
					},
				},
			},
			
		})
	
	revel.RegisterController((*controllers.Sample2)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Index",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
					12: []string{ 
					},
				},
			},
			&revel.MethodType{
				Name: "HandleSubmit",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "username", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "firstname", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "lastname", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "age", Type: reflect.TypeOf((*int)(nil)) },
					&revel.MethodArg{Name: "password", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "passwordConfirm", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "email", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "emailConfirm", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "termsOfUse", Type: reflect.TypeOf((*bool)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					46: []string{ 
						"username",
						"firstname",
						"lastname",
						"age",
						"password",
						"email",
					},
				},
			},
			
		})
	
	revel.RegisterController((*controllers.Sample3)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Index",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
					13: []string{ 
					},
				},
			},
			&revel.MethodType{
				Name: "HandleSubmit",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "user", Type: reflect.TypeOf((**models.User)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					27: []string{ 
						"user",
					},
				},
			},
			
		})
	
	revel.RegisterController((*controllers.Sample4)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Index",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
					13: []string{ 
					},
				},
			},
			&revel.MethodType{
				Name: "HandleSubmit",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "user", Type: reflect.TypeOf((**models.User)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					27: []string{ 
						"user",
					},
				},
			},
			
		})
	
	revel.RegisterController((*controllers.Sample1)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Index",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
					12: []string{ 
					},
				},
			},
			&revel.MethodType{
				Name: "HandleSubmit",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "username", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "firstname", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "lastname", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "age", Type: reflect.TypeOf((*int)(nil)) },
					&revel.MethodArg{Name: "password", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "passwordConfirm", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "email", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "emailConfirm", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "termsOfUse", Type: reflect.TypeOf((*bool)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					46: []string{ 
						"username",
						"firstname",
						"lastname",
						"age",
						"password",
						"email",
					},
				},
			},
			
		})
	
	revel.RegisterController((*controllers0.Static)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Serve",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "prefix", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "filepath", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "ServeModule",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "moduleName", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "prefix", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "filepath", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			
		})
	
	revel.RegisterController((*controllers1.TestRunner)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Index",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
					70: []string{ 
						"testSuites",
					},
				},
			},
			&revel.MethodType{
				Name: "Run",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "suite", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "test", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					107: []string{ 
					},
				},
			},
			&revel.MethodType{
				Name: "List",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			
		})
	
	revel.DefaultValidationKeys = map[string]map[int]string{ 
		"validation/app/controllers.Sample1.HandleSubmit": { 
			22: "username",
			23: "username",
			24: "firstname",
			25: "lastname",
			26: "age",
			27: "age",
			28: "password",
			29: "password",
			30: "passwordConfirm",
			31: "passwordConfirm",
			32: "email",
			33: "email",
			34: "emailConfirm",
			35: "emailConfirm",
			36: "termsOfUse",
		},
		"validation/app/controllers.Sample2.HandleSubmit": { 
			22: "username",
			23: "username",
			24: "firstname",
			25: "lastname",
			26: "age",
			27: "age",
			28: "password",
			29: "password",
			30: "passwordConfirm",
			31: "passwordConfirm",
			32: "email",
			33: "email",
			34: "emailConfirm",
			35: "emailConfirm",
			36: "termsOfUse",
		},
		"validation/app/models.(*User).Validate": { 
			18: "user.Username",
			19: "user.Username",
			20: "user.FirstName",
			21: "user.LastName",
			22: "user.Age",
			23: "user.Age",
			24: "user.Password",
			25: "user.Password",
			26: "user.PasswordConfirm",
			27: "user.PasswordConfirm",
			29: "user.Email",
			30: "user.Email",
			31: "user.EmailConfirm",
			32: "user.EmailConfirm",
			34: "user.TermsOfUse",
		},
	}
	testing.TestSuites = []interface{}{ 
		(*tests.ApplicationTest)(nil),
	}

	revel.Run(*port)
}
