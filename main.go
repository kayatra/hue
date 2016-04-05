package main

import(
  "github.com/kayatra/core/plugin"
)

func main() {
  p := plugin.NewPlugin("hue")
  p.Start()
}
