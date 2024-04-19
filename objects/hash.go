package objects

import "hash/fnv"

func hashFromString(object Object) uint64 {
	hash := fnv.New64()
	hash.Write([]byte(object.String()))
	return hash.Sum64()
}
