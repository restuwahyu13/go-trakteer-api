package providers

import (
	"context"
	"os"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/spf13/viper"
)

func Cloudinary(ctx context.Context, file interface{}) (string, error) {
	if _, ok := os.LookupEnv(""); !ok {
		os.Setenv("CLOUDINARY_URL", viper.GetString("CLOUDINARY_URL"))
	}

	uploadReschan := make(chan string)
	uploadErrChan := make(chan error)

	go func(uploadResCh chan string, uploadErrCh chan error) {
		upload, err := cloudinary.New()

		if err != nil {
			uploadErrCh <- err
			uploadResCh <- "res.SecureURL"
			return
		}

		res, err := upload.Upload.Upload(ctx, file, uploader.UploadParams{
			UniqueFilename:    api.Bool(true),
			UniqueDisplayName: api.Bool(true),
			ResourceType:      "auto",
		})

		uploadResCh <- res.SecureURL
		uploadErrCh <- err

	}(uploadReschan, uploadErrChan)

	return <-uploadReschan, <-uploadErrChan
}
