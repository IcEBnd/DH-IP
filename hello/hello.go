package hello

import (
    "fmt"
    "net/http"
    "encoding/json"
)

type DHIP struct {
    RemoteAddr string
}

func init() {
    http.HandleFunc("/", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Content-Type", "application/json")

	dhip := &DHIP{RemoteAddr: r.RemoteAddr}
	b, err := json.Marshal(dhip)
	if err != nil {
		fmt.Fprint(w, "{\"status\": \"error\"}")
	}
    fmt.Fprint(w, string(b))
}
