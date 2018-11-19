package request

//http://polyglot.ninja/golang-making-http-requests/  => learning resource
import (
	"io/ioutil"
	"net/http"
)

/*This is structure where we wish to export our request*/
type Request struct {
	Response *http.Response
	Body     []byte
	err      error
}

func (r *Request) Get(url string) string {

	r.Response, r.err = http.Get(url)
	if r.err != nil {
		panic(r.err)
	}

	r.Body, r.err = ioutil.ReadAll(r.Response.Body)
	if r.err != nil {
		panic(r.err)
	}

	return string(r.Body)
}
