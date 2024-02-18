package main

import (
	hash "goRedis/hash"
	_ "goRedis/redis"
)


func main() {
		
	hash.GetSetter()
}
