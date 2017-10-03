package cuckoofilter

import (
	"fmt"
	"math/rand"
)

//MaxNumKicks-
const MaxNumKicks = 500
const bucketSize = 4

type bucket [bucketSize]byte

//CuckooFilter -
type CuckooFilter struct {
	buckets []bucket
	count   int
}

//NewCuckooFilter -
func NewCuckooFilter(cap int) *CuckooFilter {
	buckets := make([]bucket, cap)
	for i := range buckets {
		buckets[i] = [bucketSize]byte{}
	}

	return &CuckooFilter{buckets, 0}
}

func (c *CuckooFilter) fingerprint(data []byte) byte {
	f := byte(hash(data))
	if f == 0 {
		f += 7
	}
	return f
}

//Insert -
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
// // Hashtable is considered full;
// return Failure;
func (c *CuckooFilter) Insert(data []byte) bool {
	f := c.fingerprint(data)
	hashV := byte(hash(data))

	i1 := int(hashV) % len(c.buckets)
	i2 := i1 ^ int(byte(hash([]byte{hashV})))

	fmt.Println(i1, i2)

	if c.insert(i1, f) || c.insert(i2, f) {
		return true
	}

	// address already exist, kick it out
	fp := f
	for k := 0; k < MaxNumKicks; k++ {
		r := rand.Intn(len(c.buckets))
		oldfp := fp
		fp = c.buckets[r][i2]
		c.buckets[r][i2] = oldfp
		if c.insert(r, fp) {
			return true
		}
	}

	return false
}

func (c *CuckooFilter) insert(index int, footprint byte) bool {
	for k, v := range c.buckets[index] {
		if v == 0 {
			c.buckets[index][k] = footprint
			c.count++
			return true
		}
	}
	return false
}

//Lookup -
// f = fingerprint(x);
// i1 = hash(x);
// i2 = i1 ⊕ hash(f);
// if bucket[i1] or bucket[i2] has f then
// return True;
// return False;
func (c *CuckooFilter) Lookup(data []byte) bool {
	f := c.fingerprint(data)
	hashV := byte(hash(data))

	i1 := int(hashV) % len(c.buckets)
	i2 := i1 ^ int(byte(hash([]byte{hashV})))

	b1 := c.buckets[i1]
	b2 := c.buckets[i2]

	for i := 0; i < bucketSize; i++ {
		if b1[i] == f || b2[i] == f {
			return true
		}
	}
	return false
}

//Delete -
// f = fingerprint(x);
// i1 = hash(x);
// i2 = i1 ⊕ hash(f);
// if bucket[i1] or bucket[i2] has f then
// remove a copy of f from this bucket;
// return True;
// return False;
func (c *CuckooFilter) Delete(data []byte) bool {
	f := c.fingerprint(data)
	hashV := byte(hash(data))

	i1 := int(hashV) % len(c.buckets)
	i2 := i1 ^ int(byte(hash([]byte{hashV})))

	b1 := c.buckets[i1]
	b2 := c.buckets[i2]

	for i := 0; i < bucketSize; i++ {
		if b1[i] == f {
			b1[i] = 0
			return true
		}

		if b2[i] == f {
			b2[i] = 0
			return true
		}
	}

	return false
}

//Count -
func (c *CuckooFilter) Count() int {
	return c.count
}
