package core

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/subhande/goredis/config"
)

// TODO: Support Expiration
// TODO: Support non-kv data structures
// TODO: Support sync write
func dumpKey(fp *os.File, key string, obj *Obj) {
	cmd := fmt.Sprintf("SET %s %s", key, obj.Value)
	log.Println("dumping", cmd)
	tokens := strings.Split(cmd, " ")
	log.Println("tokens", tokens)
	log.Println("encoded tokens", Encode(tokens, false))
	fp.Write(Encode(tokens, false))
}

// TODO: To to new and switch
func DumpAllAOF() {
	fp, err := os.OpenFile(config.AOFFile, os.O_CREATE|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		fmt.Print("error", err)
		return
	}
	log.Println("rewriting AOF file at", config.AOFFile)
	for k, obj := range store {
		dumpKey(fp, k, obj)
	}
	log.Println("AOF file rewrite complete")
}
