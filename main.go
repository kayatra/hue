package main

import(
  "github.com/home-control/core/plugin"
)

func main() {
  p := plugin.NewPlugin("hue")
  p.Start()
}
