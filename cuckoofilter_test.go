package cuckoofilter

import "testing"

func TestInsert(t *testing.T) {
	ck := NewCukooFilter(1000)
	if ret := ck.Insert([]byte("abc")); !ret {
		t.Error("Insert1: failed")
	}
	if ret := ck.Insert([]byte("bcd")); !ret {
		t.Error("Insert2: failed")
	}

	if ret := ck.Insert([]byte("bcd")); !ret {
		t.Error("Insert2: failed")
	}
	if ret := ck.Insert([]byte("bcd")); !ret {
		t.Error("Insert2: failed")
	}
	if count := ck.Count(); count != 4 {
		t.Error("Couting errer:", count)
	}
}
