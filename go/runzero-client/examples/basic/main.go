package main

import (
	"context"
	"log"
	"net/http"

	rc "github.com/runZeroInc/runzero-api/go/runzero-client"
)

func main() {
	var res *http.Response
	var err error

	ctx := context.Background()
	client := rc.NewClient(&rc.Config{})

	log.Printf("runZero Components")
	log.Printf("=================")
	areq := client.PublicAPI.GetLatestAgentVersion(ctx)
	aver, res, err := areq.Execute()
	if err != nil {
		log.Fatalf("failed to get agent version: %s", err)
	}
	log.Printf("     Agent: %s", aver.GetVersion())

	sreq := client.PublicAPI.GetLatestScannerVersion(ctx)
	sver, res, err := sreq.Execute()
	if err != nil {
		log.Fatalf("failed to get scanner version: %s", err)
	}
	log.Printf("   Scanner: %s", sver.GetVersion())

	preq := client.PublicAPI.GetLatestPlatformVersion(ctx)
	pver, res, err := preq.Execute()
	if err != nil {
		log.Fatalf("failed to get platform version: %s", err)
	}
	log.Printf("  Platform: %s", pver.GetVersion())

	cnf := client.GetConfig()
	if len(cnf.DefaultHeader) == 0 {
		log.Printf("Set the RUMBLE_API_KEY environment variable to test authenticated APIs")
		return
	}
	oreq := client.OrganizationAPI.GetOrganization(ctx)
	org, res, err := oreq.Execute()
	if err != nil {
		log.Fatalf("failed to get organization %s", err)
	}
	log.Printf("")
	log.Printf("Organization %s has %d assets", org.GetName(), int64(org.GetAssetCount()))

	_ = res
}
