package main

import "./network"
// требуется только ниже для обработки примера

const CONN_HOST = "localhost"
const CONN_PORT = ":8081"
const CONN_TYPE = "tcp"
const CONN_URL  = CONN_HOST + ":" + CONN_PORT

func main() {
	network.Open()
}