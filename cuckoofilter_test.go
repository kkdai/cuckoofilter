package cuckoofilter

import "testing"

func TestInsert(t *testing.T) {
	ck := NewCuckooFilter(1000)
	if ret := ck.Insert([]byte("abc")); !ret {
		t.Error("Insert1: failed")
	}
	if ret := ck.Insert([]byte("bcd")); !ret {
		t.Error("Insert2: failed")
	}

	if ret := ck.Delete([]byte("bcd")); !ret {
		t.Error("Delete: failed")
	}

	if ret := ck.Lookup([]byte("bcd")); ret {
		t.Error("Lookup: failed")
	}

	if ret := ck.Insert([]byte("bcd")); !ret {
		t.Error("Insert2: failed")
	}

	if count := ck.Count(); count != 3 {
		t.Error("Couting errer:", count)
	}
}
