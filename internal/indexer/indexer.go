package indexer

import (
	"fmt"

	fileFinder "github.com/SUMUKHA-PK/The-Engine/internal/filefinder"
	"github.com/SUMUKHA-PK/The-Engine/internal/types"
	"github.com/gomodule/redigo/redis"
)

// GetIndex gets the path for the media in the query.
func GetIndex(key types.Runnable) (string, error) {
	path, exists := getIndexedFile(key)
	if !exists {
		var err error
		path, err = indexFile(key)
		if err != nil {
			return "", err
		}
	}
	return path, nil
}

// indexFile indexes a file and returns its path.
func indexFile(key types.Runnable) (string, error) {
	conn, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		fmt.Printf("Failed to connect: %v\n", err)
	}
	defer conn.Close()

	path, err := fileFinder.Find(",", "F")

	_, err = conn.Do("HMSET", "media", key, path)
	if err != nil {
		return "", err
	}
	return "", nil
}

// getIndexedFile gets an already indexed file.
func getIndexedFile(key types.Runnable) (path string, exists bool) {
	conn, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		fmt.Printf("Failed to connect: %v\n", err)
	}
	defer conn.Close()

	switch key.Category {
	case "TV":
		path, err = redis.String(conn.Do("HGET", "media", key))
		if err != nil {
			exists = false
			return
		}
	case "Movies":

	}

	return
}
