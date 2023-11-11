package s3

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

var sessions map[string]*session.Session

type S3Adaptor struct {
	svc *s3.S3
}

func NewS3Adaptor() *S3Adaptor {

	sessions = make(map[string]*session.Session)

	cp := x509.NewCertPool()
	// don't know what is it
	if cfg.CAPath != "" {
		data, err := ioutil.ReadFile(cfg.CAPath)
		if err != nil {
			return nil, re.New(op, err, "could not read the certificate")
		}
		cp.AppendCertsFromPEM(data)
	}

	newSession,ok := sessions["test-name"]
	if !ok {
		newSession, err := session.NewSession(&aws.Config{
			Region: ,
			Endpoint: ,
			S3ForcePathStyle: ,
			HTTPClient: &http.Client{
				Transport: &http.Transport{
					TLSClientConfig: &tls.Config{
						RootCAs: cp,
					},
				},
			},
		})

		sessions["test-name"] = newSession
	}

	return &S3Adaptor{
		svc: s3.New(newSession, &aws.Config{

		}),
	}
}

func (s S3Adaptor) GetObject(ctx context.Context) (io.ReadCloser, error){
	resp, err := s.svc.GetObjectWithContext(ctx, &s3.GetObjectInput{})
	if err != nil {
		return nil, err
	}

	return resp.Body, nil
}

// Other methods
/*
- _, err := s.svc.PutObjectWithContext(c, &s3.PutObjectInput{
- _, err := s.svc.DeleteObjectWithContext(c, &s3.DeleteObjectInput{
- resp, err := s.svc.ListObjects(params)
 */