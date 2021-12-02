// **Challenge**

// Write an in-memory, key-value value store that can "time travel."

// **Part 1**

// You should be able to get and set values for arbitrary keys. In Golang, this might look something like:

//	kv.set('foo', 'bar')
//  kv.get('foo')
//  =>  "bar"

// **Part 2**

// If a timestamp is provided for a given key, fetch the value for that key at that particular time.
// If no timestamp is supplied, fetch the most recently set value for that key. In Golang, this might look like:

//	now := Time.Now.Unix()
//  kv.set('foo', 'bar')
//	time.Sleep(time.Second * 2)
//  kv.set('foo', 'bar2')

//  # Fetch the key 'foo' with the 'now' timestamp
//  kv.get('foo', now)
//  => "bar"

//  # Fetch the key 'foo' without a timestamp
//  kv.get('foo', nil)
//  => "bar2" # returns the last set value

// **Part 3**

// Support 'fuzzy' matching on a timestamp.

//  now := Time.Now()
//  kv.set('foo', 'bar')
//  Time.Sleep(Time.Second * 3)
//  kv.set('foo', 'bar2')

//  # Fetch the key 'foo' with the 'now' timestamp, plus 2 seconds
// 	newTime := now.Add(time.Duration(+2) * time.Second).UnixNano())
// 	kv.Get("foo", newTime)
//  => "bar" # returns the closest set value to that timestamp, but always in the past

/*
# Understand The Problem
- Goal: key value store that can time travel

# Assumptions

# Constraints

# Find the Obvious / Key Insights

# Questions

# Scenarions & Data Flow
## Edge Cases

# Go Brute Force

# Iterate

# Can We Do Better?

*/

package main

import (
	"fmt"
	"sort"
	"time"
)

func main() {
	fmt.Println("starting")
	kv := MakeKeyValueStore()

	fmt.Printf("now: %v\n", time.Now().Unix())
	fmt.Printf("now: %v\n", time.Now().Unix())

	kv.Put("Messi", "Barcelona")
	time.Sleep(time.Second * 2)

	kv.Put("Messi", "PSG")
	fmt.Printf("Get Latest: %v\n", *kv.Get("Messi", nil))

	fmt.Printf("kv: %+v\n", kv)
	now := time.Now()
	unixTime := now.UnixNano()
	_ = unixTime
	fourSecondsAgo := now.Add(time.Duration(-1) * time.Second).UnixNano()

	// fmt.Printf("Get with no timestamp: %v\n", *kv.Get("Messi", nil))
	// fmt.Printf("Get latest: %v\n", *kv.Get("Messi", &unixTime)) // check value not nil
	fmt.Printf("Get in the past: %v\n", *kv.Get("Messi", &fourSecondsAgo))
}

type keyTimeCompositeKey struct {
	k string
	t int64
}

func makeCompositeKey(k string, timeStamp int64) *keyTimeCompositeKey {
	return &keyTimeCompositeKey{
		k: k,
		t: timeStamp,
	}
}

type KeyValueStore struct {
	keyToTimeStamps          map[string][]int64
	keyToTimeStampToValueMap map[string]map[int64]string
	keyToValue               map[string]string
}

func MakeKeyValueStore() *KeyValueStore {
	return &KeyValueStore{
		keyToTimeStamps:          make(map[string][]int64),
		keyToTimeStampToValueMap: make(map[string]map[int64]string),
		keyToValue:               make(map[string]string),
	}
}

func (k *KeyValueStore) SortTimeStamps(key string) {
	sort.SliceStable(k.keyToTimeStamps[key], func(i, j int) bool {
		return k.keyToTimeStamps[key][i] < k.keyToTimeStamps[key][j]
	})
}

func (k *KeyValueStore) Put(key, value string) int64 {
	now := time.Now().UnixNano()

	// see if current key has a map having timelineToValue map, if not create it
	_, ok := k.keyToTimeStampToValueMap[key]
	if !ok {
		newMap := make(map[int64]string)
		newMap[now] = value
		k.keyToTimeStampToValueMap[key] = newMap
		fmt.Println(k.keyToTimeStampToValueMap[key])
	}

	// now that we know the value exists, fetch it, and add to the map the new timeline
	mapKey := k.keyToTimeStampToValueMap[key]
	mapKey[now] = value

	k.keyToTimeStamps[key] = append(k.keyToTimeStamps[key], now)

	k.keyToValue[key] = value

	return now
}

func (k *KeyValueStore) Get(key string, timeStamp *int64) *string {
	if timeStamp == nil {
		k, ok := k.keyToValue[key]
		if !ok {
			return nil
		}
		return &k
	}
	_, ok := k.keyToTimeStampToValueMap[key]
	if !ok {
		fmt.Println("found no results")
	}

	closestTimeStamp := k.GetClosestTimeStamp(key, *timeStamp)
	if closestTimeStamp == nil {
		return nil
	}
	timeStampMap, ok := k.keyToTimeStampToValueMap[key]
	fmt.Println("fetched from keyToTImeStamp", timeStampMap, ok, key)
	if !ok {
		return nil
	}

	stamp := timeStampMap[*closestTimeStamp]
	return &stamp
}

func (k *KeyValueStore) GetClosestTimeStamp(key string, time int64) *int64 {
	timeStampList, ok := k.keyToTimeStamps[key]
	if !ok {
		fmt.Println("No value found in GetClosestTimeSamp")
		return nil
	}

	if len(timeStampList) == 1 {
		if timeStampList[0] < time {
			return &timeStampList[0]
		}
		return nil
	}

	left := 0
	right := len(timeStampList) - 1
	for left < right {
		fmt.Printf("left: %v, right: %v\n", left, right)
		mid := left + (right-left)/2

		fmt.Println("mid:", mid)
		fmt.Printf("timeStampList[mid]: %v, time: %v\n", timeStampList[mid], time)
		if timeStampList[mid] < time {
			left = mid + 1
		}
		right = mid
	}
	result := timeStampList[left-1]
	return &result
}
