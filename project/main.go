package main

import (
	"context"

	awssdk "github.com/aws/aws-sdk-go/services/s3"

	"github.com/parsaeisa/S3_try/internal/s3"
)

func main() {
	s3Adaptor := s3.NewS3Adaptor()

	ctx := context.Background()

	s3Adaptor.Svc.GetObjectWithContext(ctx, &awssdk.GetOb)
}
