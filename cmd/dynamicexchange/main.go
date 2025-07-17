// cmd/dynamicexchange/main.go
package main

import (
"flag"
"log"
"os"

"dynamicexchange/internal/dynamicexchange"
)

func main() {
verbose := flag.Bool("verbose", false, "Enable verbose logging")
flag.Parse()

app := dynamicexchange.NewApp(*verbose)
if err := app.Run(); err != nil {
log.Fatal(err)
}
}
