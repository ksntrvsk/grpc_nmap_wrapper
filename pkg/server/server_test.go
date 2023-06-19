package server

import (
	"context"
	"fmt"
	"log"
	"testing"

	"github.com/ksntrvsk/grpc_nmap_wrapper/api/pb"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
)

func TestCheckVuln(test *testing.T) {

	checkVulnResponse := pb.CheckVulnResponse{
		Results: make([]*pb.TargetResult, 0),
	}

	targetResult := pb.TargetResult{
		Target:   "212.129.38.224",
		Services: make([]*pb.Service, 0),
	}

	service := pb.Service{
		Name:    "ssh",
		Version: "7.6p1 Ubuntu 4ubuntu0.6",
		TcpPort: 2222,
		Vulns:   make([]*pb.Vulnerability, 0),
	}

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
		Targets: []string{"challenge02.root-me.org"},
		TcpPort: []int32{2222},
	}

	vulns := []pb.Vulnerability{
		{
			Identifier: "EXPLOITPACK:98FE96309F9524B8C84C508837551A19",
			CvssScore:  5.8,
		},
		{
			Identifier: "EXPLOITPACK:5330EA02EBDE345BFC9D6DDDD97F9E97",
			CvssScore:  5.8,
		},
		{
			Identifier: "EDB-ID:46516",
			CvssScore:  5.8,
		},
		{
			Identifier: "EDB-ID:46193",
			CvssScore:  5.8,
		},
		{
			Identifier: "CVE-2019-6111",
			CvssScore:  5.8,
		},
		{
			Identifier: "1337DAY-ID-32328",
			CvssScore:  5.8,
		},
		{
			Identifier: "1337DAY-ID-32009",
			CvssScore:  5.8,
		},
		{
			Identifier: "SSH_ENUM",
			CvssScore:  5.0,
		},
		{
			Identifier: "PACKETSTORM:150621",
			CvssScore:  5.0,
		},
		{
			Identifier: "EXPLOITPACK:F957D7E8A0CC1E23C3C649B764E13FB0",
			CvssScore:  5.0,
		},
		{
			Identifier: "EXPLOITPACK:EBDBC5685E3276D648B4D14B75563283",
			CvssScore:  5.0,
		},
		{
			Identifier: "EDB-ID:45939",
			CvssScore:  5.0,
		},
		{
			Identifier: "EDB-ID:45233",
			CvssScore:  5.0,
		},
		{
			Identifier: "CVE-2018-15919",
			CvssScore:  5.0,
		},
		{
			Identifier: "CVE-2018-15473",
			CvssScore:  5.0,
		},
		{
			Identifier: "1337DAY-ID-31730",
			CvssScore:  5.0,
		},
		{
			Identifier: "CVE-2021-41617",
			CvssScore:  4.4,
		},
		{
			Identifier: "CVE-2020-14145",
			CvssScore:  4.3,
		},
		{
			Identifier: "CVE-2019-6110",
			CvssScore:  4.0,
		},
		{
			Identifier: "CVE-2019-6109",
			CvssScore:  4.0,
		},
		{
			Identifier: "CVE-2018-20685",
			CvssScore:  2.6,
		},
		{
			Identifier: "PACKETSTORM:151227",
			CvssScore:  0.0,
		},
		{
			Identifier: "MSF:AUXILIARY-SCANNER-SSH-SSH_ENUMUSERS-",
			CvssScore:  0.0,
		},
		{
			Identifier: "1337DAY-ID-30937",
			CvssScore:  0.0,
		},
	}

	for index := range vulns {
		service.Vulns = append(service.Vulns, &vulns[index])
	}
	targetResult.Services = append(targetResult.Services, &service)
	checkVulnResponse.Results = append(checkVulnResponse.Results, &targetResult)

	expected := &checkVulnResponse

	response, err := client.CheckVuln(context.Background(), &request)
	if err != nil {
		log.Fatalf("Error when calling CheckVuln: %s", err)
	}

	assert.Equal(test, expected, response, fmt.Sprintf("Incorrect result. Expect %s, got %s", expected, response))
}
