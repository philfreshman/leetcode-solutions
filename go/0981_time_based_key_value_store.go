package main

// time: O(log n)
// space: O(1)

type TimeMap struct {
	store map[string][]ValStamp // key : list of [val, timestamp]
}

type ValStamp struct {
	Val  string
	Time int
}

func Constructor() TimeMap {
	return TimeMap{store: make(map[string][]ValStamp)}
}

// time: O(1)
// space: O(n)

func (this *TimeMap) Set(key string, value string, timestamp int) {
	this.store[key] = append(this.store[key], ValStamp{value, timestamp})
}

// time: O(log n)
// space: O(1)

func (this *TimeMap) Get(key string, timestamp int) string {
	var res string
	var values []ValStamp
	if _, ok := this.store[key]; ok {
		values = this.store[key]
	}
	l, r := 0, len(values)-1
	for l <= r {
		mid := l + (r-l)/2
		if values[mid].Time <= timestamp {
			res = values[mid].Val
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return res
}
