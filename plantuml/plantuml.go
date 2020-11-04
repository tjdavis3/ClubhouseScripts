package plantuml

import (
	"encoding/base64"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	b64chars  = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
	pumlchars = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-_"
)

// PlantUML represents a PlantUML Server
type PlantUML interface {

	// GetDiagram takes a plantuml diagram image and returns the image bytes
	GetDiagramImage(diagram []byte) (image []byte, err error)
}

type server struct {
	URL string
}

func NewServer(URL string) *server {
	if URL == "" {
		URL = "https://plantuml.com/plantuml/png/"
	}
	return &server{URL: URL}
}

func (s *server) GetDiagramImage(diagram []byte) (imageData []byte, err error) {
	encodedDiagram, err := EncodeDiagram(diagram, 0)
	if err != nil {
		return nil, err
	}
	fmt.Println("calling ", s.URL+encodedDiagram)
	resp, err := http.Get(s.URL + encodedDiagram)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		fmt.Println("Error calling plantuml")
		return nil, errors.New(resp.Status)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

// EncodeDiagram takes a plantuml diagram description and returns the encoded string
// that can be used to call a PlantUML server
func EncodeDiagram(diagram []byte, level int) (encodedDiagram string, err error) {
	// var buf bytes.Buffer
	// def, err := flate.NewWriter(&buf, flate.BestCompression)
	// if err != nil {
	// 	return "", err
	// }
	// defer def.Close()
	// fmt.Println(string(diagram))
	// if _, err = def.Write(diagram); err != nil {
	// 	return "", err
	// }
	// def.Flush()
	// deflated := buf.Bytes()
	// fmt.Printf("Len deflated: %d\n", len(deflated))
	b64encoded := base64.RawStdEncoding.EncodeToString(diagram)
	encoded := string(reEncodeB64([]byte(b64encoded)))
	if level == 0 {
		encoded = "~h" + encoded
	}
	return encoded, nil
}

func reEncodeB64(input []byte) (retval []byte) {
	var loop int
	for x := 0; x < len(input); x++ {
		for loop = 0; loop < len(b64chars); loop++ {
			if input[x] == b64chars[loop] {
				input[x] = pumlchars[loop]
				break
			}
		}
	}
	return input
}
