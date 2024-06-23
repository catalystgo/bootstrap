package main

import (
	"bytes"
	"context"

	minio_go "github.com/minio/minio-go/v7"

	"github.com/catalystgo/bootstrap/minio"
)

func main() {
	ctx := context.Background()

	// Create a new Minio Client
	client, err := minio.NewClient(`localhost:9000`, `access_key_minio`, `secret_key_minio`, minio.WithSSL(false))
	if err != nil {
		panic(err)
	}

	exist, err := client.BucketExists(ctx, "mybucket")
	if err != nil {
		panic(err)
	}

	if !exist {
		err = client.MakeBucket(ctx, "mybucket", minio_go.MakeBucketOptions{})
		if err != nil {
			panic(err)
		}
	}

	content := []byte("Hello, World!")
	_, err = client.PutObject(ctx, "mybucket", "myobject", bytes.NewReader(content), -1, minio_go.PutObjectOptions{})
	if err != nil {
		panic(err)
	}

	obj, err := client.GetObject(ctx, "mybucket", "myobject", minio_go.GetObjectOptions{})
	if err != nil {
		panic(err)
	}

	defer obj.Close()

	buf := new(bytes.Buffer)
	buf.ReadFrom(obj)
	println(buf.String(), "✅")
}
