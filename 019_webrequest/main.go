package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "http://txti.es/"

func main() {
	// fmt.Println{"LCO web request"}

	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Response is of type : %T\n", response)

	defer response.Body.Close() // this will close the connection once the request is done

	databytes, err := io.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	content := string(databytes)
	fmt.Println(content)
}
