package load

import "net/http"
import "fmt"
import "io/ioutil"

//Fetch ...............
func Fetch(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(resp.Status)
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}
