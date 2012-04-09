package main

import "log"
var main_log *log.Logger

func main() {
  // Create a logger
  main_log = create_logger("MAIN")
  server_log = create_logger("PROXY")
  upstream_log = create_logger("UPSTREAM")

  // Run the server loop
  server(8888)
}

