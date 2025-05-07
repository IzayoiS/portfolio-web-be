package utils

import (
	"context"
	"mime/multipart"
	"os"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

func UploadImageToCloudinary(file multipart.File, fileHeader *multipart.FileHeader) (string, error) {
	cloudname := os.Getenv("CLOUDINARY_CLOUD_NAME")
	apiKey := os.Getenv("CLOUDINARY_API_KEY")
	apiSecret := os.Getenv("CLOUDINARY_API_SECRET")

	cld, err := cloudinary.NewFromParams(cloudname,apiKey,apiSecret)
	if err != nil {
		return "", err
	}

	uploadParam, err := cld.Upload.Upload(context.Background(),file,uploader.UploadParams {
		PublicID: fileHeader.Filename,
		Folder: "portfolio",
	})

	if err != nil {
		return "", err
	}
	return uploadParam.SecureURL, nil
}

func DeleteImage(publicID string) error {
	cloudname := os.Getenv("CLOUDINARY_CLOUD_NAME")
	apiKey := os.Getenv("CLOUDINARY_API_KEY")
	apiSecret := os.Getenv("CLOUDINARY_API_SECRET")
	
	cld, err := cloudinary.NewFromParams(cloudname,apiKey,apiSecret)
	if err != nil {
		return err
	}

	_, err = cld.Upload.Destroy(context.Background(), uploader.DestroyParams{
		PublicID: publicID,
	})
	return err
}