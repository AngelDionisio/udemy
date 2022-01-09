package main

import (
	"fmt"
	"time"
)

const (
	EmptyString = ""
)

type versionedKeyValueStore struct {
	keyToValue            map[string]string
	keyToTimestampToValue map[string]map[int64]string
}

func MakeNewKeyValueStore() *versionedKeyValueStore {
	return &versionedKeyValueStore{
		keyToValue:            make(map[string]string),
		keyToTimestampToValue: make(map[string]map[int64]string),
	}
}

func (kv *versionedKeyValueStore) Set(key, value string) int64 {
	now := time.Now().UnixNano()

	//check if there is a value in keyToTimeStampToValue, this allows us to know if we need to add a new timestamp
	_, ok := kv.keyToTimestampToValue[key]
	if !ok {
		m := make(map[int64]string)
		m[now] = value
		kv.keyToTimestampToValue[key] = m
	}

	// at this point we know that there is a map for the key in question, fetch it and set the value
	// this will cover the case where value already existed
	fetchedMap := kv.keyToTimestampToValue[key]
	fetchedMap[now] = value
	// set the latest value map
	kv.keyToValue[key] = value

	return now
}

func (kv *versionedKeyValueStore) Get(key string, time *time.Time) string {
	if time == nil {
		value, ok := kv.keyToValue[key]
		if !ok {
			fmt.Printf("no value found for provided key: %s\n", key)
			return EmptyString
		}
		return value
	}
	return ""
}

// func (kv *versionedKeyValueStore) findClosestTimeStamp(key string, targetTimestamp int64) *int64 {
// 	left, right := 0, len(timestampList)-1
// 	for left < right {
// 		mid := (left + right) / 2
// 		switch {
// 		case timestampList[mid] > target:
// 			right = mid - 1
// 		case timestampList[mid] < target:
// 			left = mid + 1
// 		}
// 	}
// 	result := timestampList[left]
// 	return &result
// }

func main() {
	kv := MakeNewKeyValueStore()
	kv.Set("foo", "bar")
	time.Sleep(time.Second * 2)
	kv.Set("foo", "bar2")

	fmt.Println("kv:", kv)
	fmt.Println(kv.Get("foo", nil))

	fmt.Println("about to exit")
}
