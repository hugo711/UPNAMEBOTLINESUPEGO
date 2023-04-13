สวัสดี <img src="https://media.giphy.com/media/M3nwJpDEUxkCzVftCi/giphy.gif" width="25px">
📊 **this week i spent my time on:**
<!--START_SECTION:waka-->
```text
Python   26 297,4383 saat       █████████████████████████████████████   100 % 
CSS          600 dakika         ██░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░   7.92 % 
JSON         8 765,81277 saat   ████████████████████████████░░░░░░░░░   65.2 % 
HTML         730,484398 saat    █████████░░░░░░░░░░░░░░░░░░░░░░░░░░░░   40.5 % 
PHP         4320 dakika         ████░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░   17.2 % 
```
<!--END_SECTION:waka-->
<!---
if you like what i do, maybe consider buying me a coffee/tea 🥺👉👈

<a href="https://www.buymeacoffee.com/cybertkr" target="_blank"><img src="https://cdn.buymeacoffee.com/buttons/v2/default-red.png" alt="Buy Me A Coffee" width="150" ></a>
-->
🚧 **my todoist stats:**
<!-- TODO-IST:START -->
🏆  7,982  Points           
🌸  Completed 0 tasks today           
✅  Completed 669 tasks so far           
⏳  Longest streak is 10 days
<!-- TODO-IST:END -->




# namegenerator

A random name generator (for projects, servers, cluster nodes, etc ...)
implementation in Golang.

## Badges

[![License][License-Image]][License-URL]
[![CircleCI Status][CircleCI-Image]][CircleCI-URL]
[![Coverage Report][Coverage-Image]][Coverage-URL]
[![Go Report Card][GoReportCard-Image]][GoReportCard-URL]
[![CII Best Practices][CII-Image]][CII-URL]
[![GoDoc][GoDoc-Image]][GoDoc-URL]

## Install

```bash
go get github.com/goombaio/namegenerator
```

You can also update an already installed version:

```bash
go get -u github.com/goombaio/namegenerator
```

## Example of use

```go
package main

import (
    "github.com/goombaio/namegenerator"
)

func main() {
    seed := time.Now().UTC().UnixNano()
    nameGenerator := namegenerator.NewNameGenerator(seed)

    name := nameGenerator.Generate()

    fmt.Println(name)
}
```
