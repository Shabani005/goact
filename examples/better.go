package main

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/shabani005/goact"
)

const page = `
<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8" />
<title>{{.Title}}</title>

<style>
:root {
  --bg: #020617;
  --panel: #020617;
  --text: #e5e7eb;
  --muted: #94a3b8;
  --accent: #a5b4fc;
  --border: rgba(148, 163, 184, 0.15);
}

* {
  box-sizing: border-box;
  font-family: system-ui, -apple-system, BlinkMacSystemFont,
    "Segoe UI", sans-serif;
}

body {
  margin: 0;
  min-height: 100vh;
  background:
    radial-gradient(circle at 40% 0%, #020617, transparent 45%),
    var(--bg);
  color: var(--text);
}

main {
  max-width: 680px;
  margin: 0 auto;
  padding: 4rem 2rem 3rem;
}

h1 {
  margin: 0;
  font-size: 3rem;
  letter-spacing: -0.04em;
}

.subtitle {
  margin-top: 0.5rem;
  color: var(--muted);
  font-size: 1.1rem;
}

.card {
  margin-top: 2.5rem;
  padding: 1.75rem;
  border: 1px solid var(--border);
  border-radius: 1.25rem;
  background: linear-gradient(
    180deg,
    rgba(2, 6, 23, 0.85),
    rgba(2, 6, 23, 1)
  );
}

.row {
  display: flex;
  justify-content: space-between;
  gap: 1rem;
  padding: 0.6rem 0;
  font-size: 0.95rem;
}

.row span {
  color: var(--muted);
}

footer {
  margin-top: 3rem;
  color: var(--muted);
  font-size: 0.8rem;
  text-align: center;
}
</style>
</head>

<body>
<main>
  <h1>{{.Greeting}}</h1>
  <div class="subtitle">{{.Today}}</div>

  <section class="card">
    <div class="row">
      <span>Local time</span>
      <div>{{.Clock}}</div>
    </div>

    <div class="row">
      <span>Timezone</span>
      <div>{{.Timezone}}</div>
    </div>

    <div class="row">
      <span>ISO timestamp</span>
      <div>{{.ISOTime}}</div>
    </div>

    <div class="row">
      <span>Locale</span>
      <div>{{.Locale}}</div>
    </div>

    {{if .Referrer}}
    <div class="row">
      <span>Referrer</span>
      <div>{{.Referrer}}</div>
    </div>
    {{end}}
  </section>

  <footer>
    Rendered at {{.RenderedAt}} · powered by goact
  </footer>
</main>
</body>
</html>
`

type PageData struct {
	Title string

	Greeting string
	Today    string
	Clock    string
	Timezone string
	ISOTime  string

	Locale string
	Referrer string

	RenderedAt string
}

func greeting(now time.Time) string {
	h := now.Hour()
	switch {
	case h < 12:
		return "Good morning."
	case h < 18:
		return "Good afternoon."
	default:
		return "Good evening."
	}
}

func main() {
	mux := http.NewServeMux()
	tmpl := goact.MustTemplate("today", page)

	goact.Handle(mux, "/", tmpl, func(r *http.Request) any {
		now := time.Now()

		locale := "en-US"
		if al := r.Header.Get("Accept-Language"); al != "" {
			locale = strings.Split(al, ",")[0]
		}

		return PageData{
			Title: "Today",

			Greeting: greeting(now),
			Today:    now.Format("Monday, 02 January 2006"),
			Clock:    now.Format("15:04:05"),
			Timezone: now.Format("MST"),
			ISOTime:  now.Format(time.RFC3339),

			Locale:   locale,
			Referrer: r.Referer(),

			RenderedAt: now.Format(time.RFC1123Z),
		}
	})

	fmt.Println("Listening on http://localhost:8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		panic(err)
	}
}
