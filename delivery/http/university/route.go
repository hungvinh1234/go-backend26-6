package university

import (
	"github.com/labstack/echo/v4"

	"go-template/usecase/university"
	"go-template/util"
)

type Route struct {
	universityUsecase university.IUsecase
}

func Init(group *echo.Group, universityUsecase university.IUsecase) {
	r := &Route{universityUsecase}
	group.GET("/university", r.ShowUniversityList)

}

func (r *Route) ShowUniversityList(c echo.Context) error {
	ctx := &util.CustomEchoContext{c}

	universities, err := r.universityUsecase.ShowUniversityList(ctx)

	if err != nil {
		util.Response.Error(c, err.(util.MyError))
		return nil
	}

	util.Response.Success(c, universities)
	return nil

}
