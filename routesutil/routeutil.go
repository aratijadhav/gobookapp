package routesutil

import (
	"github.com/aratijadhav1610/bookapp/rest"
	"github.com/aratijadhav1610/bookapp/routes"
)

func CreateRoute() *HandlerStruct {

	var hs HandlerStruct

	URL := routes.API + routes.GetBook
	hs.Url = URL
	hs.HandlerName = rest.GetBook
	hs.Methods = []string{"GET"}

	return &hs

}
