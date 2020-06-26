package hometown

import (
	"github.com/labstack/echo/v4"

	"go-template/usecase/hometown"
	"go-template/util"
)

type Route struct {
	hometownUsecase hometown.IUsecase
}

func Init(group *echo.Group, hometownUsecase hometown.IUsecase) {
	r := &Route{hometownUsecase}
	group.GET("/hometown", r.ShowHometownList)

	// group.POST("/upload", r.upload)

}

func (r *Route) ShowHometownList(c echo.Context) error {
	ctx := &util.CustomEchoContext{c}

	hometowns, err := r.hometownUsecase.ShowHometownList(ctx)

	if err != nil {
		util.Response.Error(c, err.(util.MyError))
		return nil
	}

	util.Response.Success(c, hometowns)
	return nil

}

// func (r *Route) Upload(c echo.Context) error {

// 	// Read form fields
// 	// name := c.FormValue("name")
// 	// email := c.FormValue("email")

// 	//------------
// 	// Read files
// 	//------------
// 	// Multipart form
// 	form, err := c.MultipartForm()
// 	if err != nil {
// 		return err
// 	}

// 	uploads, err := r.hometownUsecase.Upload(form)
// 	if err != nil {
// 		util.Response.Error(c, err.(util.MyError))
// 		return nil
// 	}

// 	util.Response.Success(c, uploads)
// 	return nil

// }
