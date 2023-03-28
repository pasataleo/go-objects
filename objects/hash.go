package objects

import "hash/fnv"

func hashFromString(object Object) uint64 {
	hash := fnv.New64()
	hash.Write([]byte(object.ToString()))
	return hash.Sum64()
}
