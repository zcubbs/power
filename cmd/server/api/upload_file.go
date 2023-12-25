package api

import (
	"fmt"
	"time"
)

func (s *Server) UploadFileWithPreSignedUrl(bucketName, objectName, filePath string) (string, error) {
	// Upload the file to the bucket
	_, err := s.s3Client.UploadFile(s.cfg.S3.BucketName, objectName, filePath)
	if err != nil {
		return "", fmt.Errorf("failed to upload project to S3 bucket: %v", err)
	}

	// Generate a download URL for the uploaded file
	downloadUrl, err := s.s3Client.GetDownloadURL(s.cfg.S3.BucketName, objectName, 10*time.Minute, nil)
	if err != nil {
		return "", fmt.Errorf("failed to get download url: %v", err)
	}

	return downloadUrl.String(), nil
}
