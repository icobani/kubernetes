/*
	B1 Yönetim Sistemleri Yazılım ve Danışmanlık Ltd. Şti.
	User        : ICI
	Name        : Ibrahim COBANI
	Date        : 15.07.2025
	Time        : 19:07
	Notes       :

Bunu yapmak lazım.
	kubectl gport-forward -n milvus svc/my-milvus 19530:19530


*/

package main

import (
	"context"
	"fmt"
	"log"

	"github.com/milvus-io/milvus-sdk-go/v2/client"
)

func main() {
	cfg := client.Config{
		Address:  "milvus.local:32080",
		Username: "ici",
		Password: "Zp8mR2xL9vBt4qNh",
	}

	cli := mustConnect(context.Background(), cfg)

	defer cli.Close()

	// Version bilgisi alalım
	version, err := cli.GetVersion(context.Background())
	if err != nil {
		log.Fatalf("get version failed: %v", err)
	}
	fmt.Println("Milvus Version:", version)
}

func mustConnect(ctx context.Context, cfg client.Config) client.Client {

	c, err := client.NewClient(ctx, cfg)
	if err != nil {
		log.Fatalf("connect to database failed, %+v", err)
	} else {
		fmt.Println("Milvus bağlantısı başarılı!")
	}
	return c
}
