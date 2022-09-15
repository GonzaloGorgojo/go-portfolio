package main

import (
	"fmt"
	"log"

	"github.com/GonzaloGorgojo/go-backend-portfolio/src/common"
	"github.com/GonzaloGorgojo/go-backend-portfolio/src/infrastructure/database"
	"github.com/GonzaloGorgojo/go-backend-portfolio/src/infrastructure/router"
	"github.com/GonzaloGorgojo/go-backend-portfolio/src/registry"
	"github.com/labstack/echo"
)

func main() {

	db := database.NewDB()

	r := registry.NewRegistry(db)

	e := echo.New()
	e.Validator = common.NewValidationUtil()
	e = router.NewRouter(e, r.NewAppController())

	fmt.Println("Server listen at http://localhost" + ":" + "8080")
	if err := e.Start(":" + "8080"); err != nil {
		log.Fatalln(err)
	}
}
