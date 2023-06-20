package server

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Ullaakut/nmap/v3"
	"github.com/ksntrvsk/grpc_nmap_wrapper/api/pb"
	log "github.com/sirupsen/logrus"
)

type GRPCServer struct {
	pb.UnimplementedNetVulnServiceServer
}

func (server *GRPCServer) CheckVuln(ctx context.Context, req *pb.CheckVulnRequest) (*pb.CheckVulnResponse, error) {

	// Constants
	const (
		scriptName = "vulners"
		flagName   = "-sV"
	)

	// Variables
	checkVulnResponse := pb.CheckVulnResponse{
		Results: make([]*pb.TargetResult, 0),
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	// Create nmap scanner
	scanner, err := nmap.NewScanner(
		ctx,
		nmap.WithTargets(targets(req.Targets)...),
		nmap.WithPorts(ports(req.TcpPort)...),
		nmap.WithScripts(scriptName),
		nmap.WithCustomArguments(flagName),
	)
	if err != nil {
		return nil, fmt.Errorf("unable to create a scanner: %v", err)
	}

	// Run nmap scanner
	result, warnings, err := scanner.Run()
	if warnings != nil {
		log.Printf("run finished with warnings: %s\n", *warnings)
	}
	if err != nil {
		log.Println(err)
		log.Fatalf("unable to run nmap scan: %v", err)
	}

	// Parse the result and forming a response
	for _, host := range result.Hosts {
		if len(host.Ports) == 0 || len(host.Addresses) == 0 {
			continue
		}

		targetResult := pb.TargetResult{
			Target:   host.Addresses[0].Addr,
			Services: make([]*pb.Service, 0),
		}

		for _, port := range host.Ports {
			service := pb.Service{
				Name:    port.Service.Name,
				Version: port.Service.Version,
				TcpPort: int32(port.ID),
				Vulns:   make([]*pb.Vulnerability, 0),
			}

			for _, script := range port.Scripts {
				if script.ID == scriptName {
					vulnerabilities := strings.Split(script.Output, "\n")
					for _, vulnerability := range vulnerabilities {
						vuln := strings.Fields(vulnerability)
						if len(vuln) == 0 || len(vuln) == 1 {
							continue
						}
						cvss, err := strconv.ParseFloat(vuln[1], 32)
						if err != nil {
							return nil, fmt.Errorf("uable parse float value from string: %v", err)
						}
						identifier := vuln[0]
						service.Vulns = append(service.Vulns, &pb.Vulnerability{
							Identifier: identifier,
							CvssScore:  float32(cvss),
						})
					}
				}
			}
			targetResult.Services = append(targetResult.Services, &service)
		}
		checkVulnResponse.Results = append(checkVulnResponse.Results, &targetResult)
	}
	return &checkVulnResponse, nil
}

func targets(trgt []string) []string {
	var targets []string

	if len(trgt) != 0 {
		targets = trgt
	}
	return targets
}

func ports(port []int32) []string {
	ports := make([]string, len(port))

	if len(port) != 0 {
		for i := 0; i < len(port); i++ {
			ports[i] = strconv.Itoa(int(port[i]))
		}
	}
	return ports
}
