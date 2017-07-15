package downloader

import (
	"io"
	"net/http"
	"os"
)

func downloadFile(remoteUrl, localFile string) (*os.File, error) {
	resp, err := http.Get(remoteUrl)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	lFile, err := os.Create(localFile)
	if err != nil {
		return nil, err
	}
	defer lFile.Close()

	var buff []byte = make([]byte, 1<<10)
	for {
		n, err := resp.Body.Read(buff)
		if err != nil && err != io.EOF {
			return nil, err
		}
		if _, err := lFile.Write(buff[:n]); err != nil {
			return nil, err
		}
		if n == 0 || err == io.EOF {
			break
		}
	}

	return lFile, nil
}
