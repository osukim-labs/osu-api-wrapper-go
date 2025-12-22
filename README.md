# Osu API Wrapper - GoLang

A Go wrapper for interacting with osu! APIs, designed to be simple, extensible, and practical for real-world use. This library supports the official osu! API and multiple mirror services, with the ability to plug in your own custom mirror providers.

> ⚠️ **Status:** Early development / hobby project. Breaking changes may occur.

---

## ✨ Features (Current & Planned)

* [ ] Wrapper foundation & API structure
* [ ] Rate limit handling
* [ ] osu!(lazer) API support
* [ ] Mirror beatmap download support
* [ ] Legacy v1 API support
* [ ] **Custom Mirror API Support**
  * Allows users to define their own mirror API provider
  * Must follow **osu!direct-compatible API formatting**
  * Flexible for self-hosted mirrors or alternative public APIs

---

## 🚀 Installation

```bash
go get github.com/osukim-labs/osu-api-wrapper-go
```

---

## 🧩 Basic Usage

```go
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/osukim-labs/osu-api-wrapper-go/v2"
)

func main() {
	// Using Official API
	api, err := v2.NewOsuV2API(12345, "abcdefghijklmnopqrstuvwxyz", 10*time.Second)
	if err != nil {
		log.Fatal(err)
	}

	beatmap, err := api.GetBeatmap("123")
	if err != nil {
		panic(err)
	}

	beatmapPretty, _ := json.MarshalIndent(beatmap, "", "  ")
	fmt.Println(string(beatmapPretty))

	// Using Built-in Mirror API
	mirror, err := v2.NewOsuV2Mirror("osu!Direct", 10*time.Second)
	if err != nil {
		log.Fatal(err)
	}

	beatmap2, err := mirror.GetBeatmap("123")
	if err != nil {
		panic(err)
	}

	beatmap2Pretty, _ := json.MarshalIndent(beatmap2, "", "  ")
	fmt.Println(string(beatmap2Pretty))

	// Using Custom Mirror API
	custom := v2.NewOsuV2MirrorWithCustomHost(
		"https://custom-mirror.com",
		"/api/v2",
		10*time.Second,
	)

	beatmap3, err := custom.GetBeatmap("123")
	if err != nil {
		panic(err)
	}

	beatmap3Pretty, _ := json.MarshalIndent(beatmap3, "", "  ")
	fmt.Println(string(beatmap3Pretty))
}
```

---

## 🧰 Custom Mirror API (osu!direct Compatible)

If the built-in mirror APIs are not what you need, you can register your own mirror endpoint - as long as it follows the osu!direct-compatible response structure.

### ✔️ Goals

* Easily swap mirror sources
* Support self-hosted or private mirrors
* Maintain a reliable and predictable API format

### 🧩 Custom Mirror API Requirements

Your mirror API must:
* Return data in osu!direct compatible JSON format
* Follow the same routes/endpoints as osu!direct

---

### 🛠 Example Custom Mirror API Usage

```go
func main() {
	api := v2.NewOsuV2MirrorWithCustomHost(
		"https://custom-mirror.com",
		"/api/v2",
		10*time.Second,
	)

	// now everything uses your custom provider
} 
```

---

## 📌 Roadmap

* [ ] Structured rate limiting
* [ ] Full osu!lazer API endpoint coverage
* [ ] Full osu!v1/v2 API endpoint coverage
* [ ] Built-in multiple mirror providers
* [ ] Better error typing & logging
* [ ] Documentation & examples

---

## 📝 Important Notes

This is a hobby project and development may be slow, so please don't expect active maintenance or frequent updates.  
Many parts of this project were built with the help of AI tools (ChatGPT / Claude), and its main purpose is to support my own projects that rely on osu! APIs and mirror services. Expect occasional rough edges or experimental code.

---

## 🤝 Contributing

Contributions, ideas, and issue reports are welcome!

1. Fork the repo
2. Create a feature branch
3. Submit a PR

---

## 📜 License

