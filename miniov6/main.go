package main

import (
	"fmt"

	"github.com/minio/minio-go/v6"
)

func main() {
	endpoint := "cloudchain.hechuan.rockontrol.com:20443"
	accessKeyID := "cloudchian"
	secretAccessKey := "Urc3hL4jy86yPEV3"
	useSSL := true

	// Initialize minio client object.
	minioClient, _ := minio.New(endpoint, accessKeyID, secretAccessKey, useSSL)

	// fmt.Println("success:")
	// log.Printf("%#v\n", minioClient) // minioClient is now setup
	buckets, err := minioClient.ListBuckets()
	if err != nil {
		panic(err)
	}
	for _, b := range buckets {
		fmt.Println(b.Name)
	}
}
