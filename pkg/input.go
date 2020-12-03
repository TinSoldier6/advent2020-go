package input

import (
	"io/ioutil"
	"net/http"
)

// Get returns the text of the problem input file.
// Input: A URL containing the name of the input file.
// Returns: A string representing the input file.
func Get(url string) (string, error) {
	res, err := http.Get(url)
	if err != nil {
		return "", err
	}

	file, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return "", err
	}

	return string(file), nil
}
