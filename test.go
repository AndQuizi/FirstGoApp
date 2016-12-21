package firstapp

import (
	"fmt"
	"net/http"
)

func init() {
	http.HandleFunc("/test", testhandler)
}

func testhandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "test page...")
}
