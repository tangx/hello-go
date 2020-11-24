package main

import (
	"context"
	"fmt"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func main() {
	// endpoint := "cloudchain.hechuan.rockontrol.com:19000"
	endpoint := "cloudchain.hechuan.rockontrol.com:20443"
	accessKeyID := "cloudchian"
	secretAccessKey := "Urc3hL4jy86yPEV3"
	useSSL := true

	// Initialize minio client object.
	minioClient, _ := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// log.Printf("%#v\n", minioClient) // minioClient is now set up
	buckets, err := minioClient.ListBuckets(context.Background())
	if err != nil {
		panic(err)
	}

	for _, b := range buckets {
		fmt.Println(b.Name)
	}
}
