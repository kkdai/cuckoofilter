package cukoofilter

import (
	"fmt"
	"math/rand"

	"github.com/dgryski/go-metro"
)

//MaxNumKicks-
const MaxNumKicks = 500

type bucket [4]byte

//CukooFilter -
type CukooFilter struct {
	buckets []bucket
}

//NewCukooFilter -
func NewCukooFilter(cap int) *CukooFilter {
	return new(CukooFilter)
}

func (c *CukooFilter) fingerprint(data []byte) byte {
	f := byte(metro.Hash64(data, 1337))
	if f == 0 {
		f += 7
	}
	return f
}

//Insert -
func (c *CukooFilter) Insert(data []byte) bool {
	f := c.fingerprint(data)
	hash := byte(metro.Hash64(data, 1337))

	i1 := int(hash) % len(c.buckets)
	i2 := i1 ^ int(byte(metro.Hash64([]byte{hash}, 1337)))

	fmt.Println(i1, i2)

	// if c.insert(i1, f) || c.insert(i2, f) {
	// 	return true
	// }

	fp := f
	for k := 0; k < MaxNumKicks; k++ {
		r := rand.Intn(len(c.buckets))
		oldfp := fp
		fp = c.buckets[r][i2]
		c.buckets[r][i2] = oldfp
	}
	// c.insertCollision(i1, f)
	// f = fingerprint(x);
	// i1 = hash(x);
	// i2 = i1 ⊕ hash(f);
	// if bucket[i1] or bucket[i2] has an empty entry then
	// add f to that bucket;
	// return Done;
	// // must relocate existing items;
	// i = randomly pick i1 or i2;
	// for n = 0; n < MaxNumKicks; n++ do
	// randomly select an entry e from bucket[i];
	// swap f and the fingerprint stored in entry e;
	// i = i ⊕ hash(f);
	// if bucket[i] has an empty entry then
	// add f to bucket[i];
	// return Done;
	// // Hashtable is considered ful

	return false
}

//Lookup -
func (c *CukooFilter) Lookup(data []byte) bool {
	return false
}

//Delete -
func (c *CukooFilter) Delete(data []byte) bool {
	return false
}

//Count -
func (c *CukooFilter) Count() int {
	return 0
}
