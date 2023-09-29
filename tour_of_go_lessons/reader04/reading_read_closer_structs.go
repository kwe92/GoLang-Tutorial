package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"time"
)

func main() {

	url := "https://vast-puce-mite-fez.cyclic.app/animeme"

	client := http.Client{Timeout: 10 * time.Second}

	response0, err := client.Get(url)

	checkError(err)

	//? Read content from io.ReadCloser With Custom Implementation

	fmt.Printf("\nresponse0:\n\n%+v\n\n", response0)

	fmt.Printf("\nresponse0 body:\n\n%+v\n\n", response0.Body)

	response0Bytes := make([]byte, response0.ContentLength)

	ReadAndClose(response0.Body, response0Bytes)

	fmt.Printf("\nresponse0 body as Slice of bytes:\n\n%+v\n\n", response0Bytes)

	fmt.Printf("\nresponse0 body as string:\n\n%+v\n\n", string(response0Bytes))

	//? Read content from io.ReadCloser With io Package

	response1, err := client.Get(url)

	response1Bytes, err := io.ReadAll(response1.Body)

	checkError(err)

	response1StringRep := string(response1Bytes)

	fmt.Printf("\nresponse1 string representation:\n\n%+v\n\n", response1StringRep)

	//? Read content from io.ReadCloser with bytes package

	response2, err := client.Get(url)

	buffer := bytes.NewBuffer(make([]byte, response2.ContentLength))

	_, err = buffer.ReadFrom(response2.Body)

	checkError(err)

	response2StringRep := buffer.String()

	fmt.Printf("\nresponse2 string representation:\n\n%+v\n\n", response2StringRep)

}

func ReadAndClose(r io.ReadCloser, buf []byte) (n int, err error) {

	for len(buf) > 0 {

		var nr int

		if err == io.EOF {
			return
		}

		nr, err = r.Read(buf)

		n += nr

		buf = buf[:nr]

		fmt.Println("\nbuffer bytes from ReadAndClose: ", buf[:n])

		fmt.Println(string(buf))

	}
	r.Close()
	return
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

// http.Response

//   - a struct defined in the http package that represents an HTTP response0 received
//     for a request

// response0.Body | io.ReadCloser

//   - the content of an HTTP response0 is an io.ReadCloser implementation
//   - the content can not be represented as a string until it is unmarshalled or decoded
//   - there are several ways to unmarshal an io.ReadCloser

// Least Efficient Ways From Worst to Best:

// Implement ReadAndClose function

//   - which takes an io.ReadCloser and a Slice of bytes as an argument
//   - and continuously reads into the Slice of bytes until the end of the file reached
//   - the Slice of bytes can then be converted to a string with type conversion

// Use io package

//   - pass the io.ReadCloser as an argument to io.ReadAll
//     which returns a Slice of bytes the length of the content and any errors encountered
//   - the Slice of bytes can then be converted to a string with type conversion

// Use bytes package to Write Contents into in-memory Buffer
