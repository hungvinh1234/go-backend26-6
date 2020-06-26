package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go-template/client/postgres"
	"go-template/config"
	serviceHttp "go-template/delivery/http"
	"go-template/repository"
	"go-template/usecase"

	"github.com/labstack/echo/v4"
)

func upload(c echo.Context) error {
	// Read form fields
	name := c.FormValue("name")
	email := c.FormValue("email")

	//------------
	// Read files
	//------------

	// Multipart form
	form, err := c.MultipartForm()
	if err != nil {
		return err
	}
	files := form.File["files"]

	for _, file := range files {
		// Source
		src, err := file.Open()
		if err != nil {
			return err
		}
		defer src.Close()

		// Destination
		dst, err := os.Create(file.Filename)
		if err != nil {
			return err
		}
		defer dst.Close()

		// Copy
		if _, err = io.Copy(dst, src); err != nil {
			return err
		}

	}

	return c.HTML(http.StatusOK, fmt.Sprintf("<p>Uploaded successfully %d files with fields name=%s and email=%s.</p>", len(files), name, email))
}

func main() {
	cfg := config.GetConfig()

	defer postgres.Disconnect()

	// setup locale
	{
		loc, err := time.LoadLocation("Asia/Bangkok")
		if err != nil {
			fmt.Println("error", err)
			os.Exit(1)
		}
		time.Local = loc
	}

	repo := repository.New(postgres.GetClient)

	ucase := usecase.New(repo)

	errs := make(chan error)

	// http
	{
		h := serviceHttp.NewHTTPHandler(ucase)
		go func() {
			errs <- h.Start(":" + cfg.Port)
		}()
	}

	// graceful
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	log.Println("exit", <-errs)

	// e := echo.New()

	// e.POST("/upload", upload)

}
