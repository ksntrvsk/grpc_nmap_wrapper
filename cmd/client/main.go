package main

import (
	"context"
	"log"

	"github.com/ksntrvsk/grpc_nmap_wrapper/api/pb"
	"google.golang.org/grpc"
)

func main() {

	connection, err := grpc.Dial(":8000", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	if err != nil {
		log.Fatalf("could not connect: %s", err)
	}
	defer connection.Close()

	client := pb.NewNetVulnServiceClient(connection)
	request := pb.CheckVulnRequest{
		// Targets: []string{"scanme.nmap.org"},
		Targets: []string{"challenge02.root-me.org"},
		TcpPort: []int32{2222},
	}

	response, err := client.CheckVuln(context.Background(), &request)
	if err != nil {
		log.Fatalf("Error when calling CheckVuln: %s", err)
	}
	for index, result := range response.Results {
		log.Printf("Response from the server: \n")
		log.Printf("target: %s\n", response.Results[index].Target)
		for _, vuln := range result.Services[0].Vulns {
			log.Printf("vuln: %s\n", vuln)
		}
	}
}
