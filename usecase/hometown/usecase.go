package hometown

import (
	"context"
	"net/http"

	"go-template/model"
	"go-template/repository/hometown"
	"go-template/util"
)

//Viet noi dung ham va lay cac ham tuong tac ben db

type usecase struct {
	// cardRepo card.Repository
	hometownRepo hometown.Repository
}

type UploadResponse struct {
	Url string `json:"url,omitempty"`
}

func New(hometownRepo hometown.Repository) IUsecase {
	return &usecase{hometownRepo}
}

func (e *usecase) ShowHometownList(ctx context.Context) (*[]model.Hometown, error) {

	hometownlist, err := e.hometownRepo.GetHometownList(ctx)
	//err la khong tra ve ket qua mong muon
	if err != nil {
		return nil, util.NewError(err, http.StatusInternalServerError, 1010, "failed get hometown list")
	}

	//account da tao trong db
	return hometownlist, nil
}

// func (e *usecase) Upload(form *multipart.Form) (*UploadResponse, error) {

// 	files := form.File["files"]

// 	for _, file := range files {
// 		// Source
// 		src, err := file.Open()
// 		if err != nil {
// 			return nil, err
// 		}
// 		defer src.Close()

// 		// Destination
// 		dst, err := os.Create(file.Filename)
// 		if err != nil {
// 			return nil, err
// 		}
// 		defer dst.Close()

// 		// Copy
// 		if _, err = io.Copy(dst, src); err != nil {
// 			return nil, err
// 		}

// 	}

// 	response := UploadResponse{
// 		Url: "http://localhost:3000/static/" + file.Filename,
// 	}

// 	return &response, nil
// }
