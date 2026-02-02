package usage

import (
	"net/http"
	"time"
	"github.com/shabani005/goact"
)

const htmlTemplate = `
<!DOCTYPE html>
<html>
  <body>
    <h1>{{.Title}}</h1>
    <p>{{.Time}}</p>
  </body>
</html>
`

type PageData struct {
	Title string
	Time  string
}

func main() {
	mux := http.NewServeMux()

	tmpl := goact.MustTemplate("page", htmlTemplate)

	goact.Handle(mux, "/", tmpl, func(r *http.Request) any {
		return PageData{
			Title: "goact",
			Time:  time.Now().Format(time.RFC1123),
		}
	})

	http.ListenAndServe(":8080", mux)
}
