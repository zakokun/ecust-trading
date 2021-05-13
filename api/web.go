package api

import "net/http"

func DefaultServer() {
	http.HandleFunc("/wallet/info", getWallet)
}

func getWallet(w http.ResponseWriter, r *http.Request) {
	r.GetBody

}