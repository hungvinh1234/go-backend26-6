package university

import (
	"context"
	"net/http"

	"go-template/model"
	"go-template/repository/university"
	"go-template/util"
)

//Viet noi dung ham va lay cac ham tuong tac ben db

type usecase struct {
	// cardRepo card.Repository
	universityRepo university.Repository
}

func New(universityRepo university.Repository) IUsecase {
	return &usecase{universityRepo}
}

func (e *usecase) ShowUniversityList(ctx context.Context) (*[]model.University, error) {

	universitylist, err := e.universityRepo.GetUniversityList(ctx)
	//err la khong tra ve ket qua mong muon
	if err != nil {
		return nil, util.NewError(err, http.StatusInternalServerError, 1010, "failed get university list")
	}

	return universitylist, nil
}
