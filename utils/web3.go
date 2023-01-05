package utils

import (
	"context"
	"fmt"
	"github.com/ipfs/go-cid"
	"github.com/web3-storage/go-w3s-client"
	"io/fs"
	"path"

	"os"
)

func UploadFile(fpath string, fileName string) (error, string, string) {
	token := os.Getenv("WEB3_STORAGE_TOKEN")
	/*if !ok {
		fmt.Fprintln(os.Stderr, "No API token - set the WEB3_STORAGE_TOKEN environment var and try again.")
		return fmt.Errorf("no API token - set the WEB3_STORAGE_TOKEN environment var and try again")
	}*/

	f, err := os.OpenFile(fpath, 0, 0777)
	defer f.Close()

	client, err := w3s.NewClient(w3s.WithToken(token))
	if err != nil {
		return err, "", ""
	}
	baseName := path.Base(fpath)
	fmt.Printf("Storing %s ...\n", baseName)
	cid, err := client.Put(context.Background(), f)
	if err != nil {
		return err, "", ""
	}
	gatewayURL := fmt.Sprintf("https://%s.ipfs.dweb.link/%s\n", cid.String(), baseName)
	fmt.Printf("Stored %s with web3.storage! View it at: %s\n", baseName, gatewayURL)
	cidString := cid.String()
	return nil, gatewayURL, cidString
}

func RetrieveFiles(client w3s.Client, cidString string) (fs.File, fs.FS, error) {
	c, err := cid.Parse(cidString)
	if err != nil {
		return nil, nil, err
	}

	res, err := client.Get(context.Background(), c)
	if err != nil {
		return nil, nil, err
	}

	if res.StatusCode != 200 {
		return nil, nil,
			fmt.Errorf("request for %s was unsuccessful: [%d]: %s",
				cidString, res.StatusCode, res.Status)
	}

	// res is an http.Response with an extra method for reading IPFS UnixFS files
	return res.Files()
}
