package utils

import (
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/disintegration/imaging"
	"github.com/renaldyhidayatt/twittersqlc/dto/response"
)

func UploadFileImage(w http.ResponseWriter, r *http.Request) response.FileResponse {

	err := r.ParseMultipartForm(32 << 20) // 32 MB
	if err != nil {
		http.Error(w, fmt.Sprintf("file err : %s", err.Error()), http.StatusBadRequest)

	}

	fileImage, headerImage, err := r.FormFile("profileImage")
	if err != nil {
		response.ResponseError(w, http.StatusBadRequest, err)

	}
	defer fileImage.Close()

	fileCover, headerCover, err := r.FormFile("profileCover")
	if err != nil {
		response.ResponseError(w, http.StatusBadRequest, err)

	}
	defer fileCover.Close()

	folderImage := "profileImage"
	folderCover := "profileCover"

	folderPathImage := fmt.Sprintf("public/%s", folderImage)
	folderPathCover := fmt.Sprintf("public/%s", folderCover)
	if _, err := os.Stat(folderPathImage); os.IsNotExist(err) {
		err = os.MkdirAll(folderPathImage, 0755)
		if err != nil {
			response.ResponseError(w, http.StatusBadRequest, err)

		}
	}
	if _, err := os.Stat(folderPathCover); os.IsNotExist(err) {
		err = os.MkdirAll(folderPathCover, 0755)
		if err != nil {
			response.ResponseError(w, http.StatusBadRequest, err)

		}
	}

	fileExtImage := filepath.Ext(headerImage.Filename)
	newFileNameImage := fmt.Sprintf("%d%s", time.Now().UnixNano(), fileExtImage)
	// filePathImage := fmt.Sprintf("http://localhost:8000/images/%s/%s", folderImage, newFileNameImage)
	dstImage := fmt.Sprintf("%s/%s", folderPathImage, newFileNameImage)
	if _, err := saveFile(dstImage, headerImage); err != nil {
		http.Error(w, fmt.Sprintf("error : %s", err.Error()), http.StatusInternalServerError)

	}

	fileExtCover := filepath.Ext(headerCover.Filename)
	newFileNameCover := fmt.Sprintf("%d%s", time.Now().UnixNano(), fileExtCover)
	filePathCover := fmt.Sprintf("http://localhost:8000/images/%s/%s", folderCover, newFileNameCover)
	dstCover := fmt.Sprintf("%s/%s", folderPathCover, newFileNameCover)
	if _, err := saveFile(dstCover, headerCover); err != nil {
		http.Error(w, fmt.Sprintf("error : %s", err.Error()), http.StatusInternalServerError)

	}

	img, err := imaging.Open(dstImage)
	if err != nil {
		http.Error(w, fmt.Sprintf("error : %s", err.Error()), http.StatusInternalServerError)

	}
	resized := imaging.Resize(img, 300, 0, imaging.Lanczos)
	newFileNameResized := fmt.Sprintf("%d%s", time.Now().UnixNano(), fileExtImage)
	dstResized := fmt.Sprintf("%s/%s", folderPathImage, newFileNameResized)
	err = imaging.Save(resized, dstResized)
	if err != nil {
		http.Error(w, fmt.Sprintf("error : %s", err.Error()), http.StatusInternalServerError)

	}

	return response.FileResponse{
		ProfileImage: fmt.Sprintf("http://localhost:8000/images/%s/%s", folderImage, newFileNameResized),
		ProfileCover: filePathCover,
	}

}

func saveFile(filePath string, fileHeader *multipart.FileHeader) (bool, error) {
	src, err := fileHeader.Open()
	if err != nil {
		return false, err
	}
	defer src.Close()

	dst, err := os.Create(filePath)
	if err != nil {
		return false, err
	}
	defer dst.Close()

	if _, err = io.Copy(dst, src); err != nil {
		return false, err
	}

	return true, nil
}
