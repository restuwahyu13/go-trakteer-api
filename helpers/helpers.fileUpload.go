package helpers

import (
	"net/http"

	"github.com/restuwahyu13/go-trakteer-api/providers"
)

func FileUpload(r *http.Request, field ...string) (map[string]string, error) {

	if err := r.ParseMultipartForm(3145728); err != nil {
		return nil, err
	}

	if err := r.ParseForm(); err != nil {
		return nil, err
	}

	mapping := make(map[string]string)

	for _, v := range field {
		fileRes, _, err := r.FormFile(v)
		if err != nil {
			return nil, err
		}

		uploadRes, uploadErr := providers.Cloudinary(r.Context(), fileRes)
		if uploadErr != nil {
			return nil, uploadErr
		}

		mapping[v] = uploadRes
	}

	return mapping, nil
}
