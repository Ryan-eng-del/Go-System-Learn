package rpc_server

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"rpc/proto"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr = flag.String("addr", "127.0.0.1", "Address")
	http_port = flag.Int("http_port", 5000, "Port")
	grpc_addr = flag.String("grpc_addr", "127.0.0.1:5051", "GrpcServer")
)

func HTTPServer () {
	flag.Parse()
	log.Println(*addr, *http_port)
	service := http.NewServeMux()
	service.HandleFunc("/orders", func(w http.ResponseWriter, request *http.Request) {
		conn, err := grpc.Dial(*grpc_addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			log.Fatal(err)
		}

		defer conn.Close()
		client := proto.NewProductClient(conn)
		ctx, cancel := context.WithTimeout(context.Background(), 3 * time.Second)
		defer cancel()

		response, err := client.ProductInfo(ctx, &proto.ProductRequest{
			Id: 1,
		})

		if err != nil {
			log.Fatal(err)
		}
		
		log.Printf("response ID is %d Name is %s", response.Id, response.Name)


		data := struct {
			ID int64 `json:"id"`
			Quantity int `json:"quantity"`
			Products []*proto.ProductResponse `json:"products"`
		} {
			ID: 1,
			Quantity: 100,
			Products: []*proto.ProductResponse{
				response,
			},
		}

		dataJson, err := json.Marshal(data)


		if err != nil {
			log.Fatalln(err)
		}

		w.Header().Set("Content-Type", "application/json")
		if _, err := fmt.Fprint(w, string(dataJson)); err != nil {
			log.Fatalln(err)
		}
	})


	address := fmt.Sprintf("%s:%d", *addr, *http_port)
	fmt.Printf("Order Service is listening on %s\n", address)
	log.Fatalln(http.ListenAndServe(address, service))
}
