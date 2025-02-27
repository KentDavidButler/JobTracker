package receiver

import (
	"fmt"
	"io"
	"net/http"
)

func Receiver(url string) {
	// goal of the receiver is to receive data from the url
	// that will fill in the data of the model.JobPosting struct
	// while trying to keep both items uncoupled
	resp, err := http.Get(url)
	if err != nil {
		fmt.Print("Error Occured: ", err)
	}

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Print("Error Occured: ", err)
	}

	fmt.Println("Here's the response")
	fmt.Println(resp)
	fmt.Println(string(body))
}
