package handlers

import (
	"bytes"
	"errors"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/coreos/discovery.etcd.io/pkg/lockstring"
)

var (
	currentLeader lockstring.LockString
)

func init() {
	// etcd cluster is passed in via commandline option -etcd
	//currentLeader.Set("127.0.0.1:4001")
}

func proxyRequest(r *http.Request) (*http.Response, error) {
	body, _ := ioutil.ReadAll(r.Body)

	machines := strings.Split((*etcdCluster), ",")

	for i := 0; i <= 10; i++ {

		if i == 0 {
			log.Printf("etcd initial leader: %v", machines[i])
			currentLeader.Set(machines[i])
		}

		leader := currentLeader.String();
		lurl, err := url.Parse(leader)
		if err != nil {
			return nil, err
		}

		u := url.URL{
			Scheme: lurl.Scheme,
			Host: lurl.Host,
			Path: path.Join("v2", "keys", "_etcd", "registry", r.URL.Path),
			RawQuery: r.URL.RawQuery,
		}

		buf := bytes.NewBuffer(body)
		outreq, err := http.NewRequest(r.Method, u.String(), buf)
		if err != nil {
			return nil, err
		}

		copyHeader(outreq.Header, r.Header)

		client := http.Client{}
		resp, err := client.Do(outreq)
		if err != nil {
			if i < len(machines) {
				log.Printf("etcd leader: %v", machines[i])
				currentLeader.Set(machines[i])
				continue
			}
			return nil, err
		}

		// Try again on the next host
		if resp.StatusCode == 307 &&
			(r.Method == "PUT" || r.Method == "DELETE") {
			u, err := resp.Location()
			if err != nil {
				return nil, err
			}
			currentLeader.Set(u.Host)
			continue
		}

		return resp, nil
	}

	return nil, errors.New("All attempts at proxying to etcd failed")
}

// copyHeader copies all of the headers from dst to src.
func copyHeader(dst, src http.Header) {
	for k, v := range src {
		for _, q := range v {
			dst.Add(k, q)
		}
	}
}

func TokenHandler(w http.ResponseWriter, r *http.Request) {
	resp, err := proxyRequest(r)
	if err != nil {
		log.Printf("Error making request: %v", err)
		http.Error(w, "", 500)
		return
	}

	copyHeader(w.Header(), resp.Header)
	w.WriteHeader(resp.StatusCode)
	io.Copy(w, resp.Body)
}
