package queries

import (
	"context"
	"net/http"
	"time"
)

func GetHardOpQuery(url string) (bool, int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	isRequestDone := make(chan *http.Response)
	isRequestError := make(chan error)

	go func() {
		resp, err := http.Get(url)
		if err != nil {
			isRequestError <- err
		}
		isRequestDone <- resp
	}()

	select {
	case resp := <-isRequestDone:
		return true, resp.StatusCode, nil
	case err := <-isRequestError:
		return false, -1, err
	case <-ctx.Done():
		return false, -1, nil
	}
}
