package config

var Host string = "0.0.0.0"
var Port int = 7379
var KeysLimit int = 100

// Will evict EvictionRatio of keys whenever eviction runs
var EvictionRatio float64 = 0.4
var EvictionStrategy string = "allkeys-random"
// var EvictionStrategy string = "simple-first"

var AOFFile string = "./goredis-master.aof"
