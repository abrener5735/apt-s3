package main

import (
	"flag"
	"fmt"
	"github.com/abrener5735/apt-s3/downloader"
	"log"
	"os"
	"runtime"

	"github.com/zendesk/apt-s3/method"
)

var (
	downloadUri  = flag.String("download", "", "S3 URI for downloading a single file")
	downloadPath = flag.String("path", "", "Path to download file to")
	versionFlag  = flag.Bool("version", false, "Show version")
	Version      = "master"
)

func main() {
	m := method.New()
	programName := os.Args[0]

	flag.Parse()

	if *versionFlag {
		fmt.Printf("%s %s (Go version: %s)\n", programName, Version, runtime.Version())
		os.Exit(0)
		// Called outside of apt to download a file
	} else if *downloadUri != "" {
		if !downloader.URIRegex.MatchString(*downloadUri) {
			log.Fatalf("Incorrect bucket format.\nExpected: s3://<bucket>.s3-<region>.amazonaws.com/path/to/file\n")
		}

		filename, err := m.Downloader.DownloadFile(*downloadUri, *downloadPath)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Downloaded %s\n", filename)
		os.Exit(0)
	} else {
		m.Start()
	}
}
