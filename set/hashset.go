package set

const mod = 997

type MyHashSet struct {
	set [mod][]int
}

func Constructor() MyHashSet {
	return MyHashSet{}
}

func (this *MyHashSet) Add(key int) {
	if this.Contains(key) {
		return
	} // no duplicates
	this.set[key%mod] = append(this.set[key%mod], key)
}

func (this *MyHashSet) Remove(key int) {
	list := &this.set[key%mod]
	for i, k := range *list {
		if k != key {
			continue
		}
		(*list)[i] = (*list)[len(*list)-1]
		(*list) = (*list)[:len(*list)-1]
		return
	}
}

func (this *MyHashSet) Contains(key int) bool {
	for _, k := range this.set[key%mod] {
		if k == key {
			return true
		}
	}
	return false
}

func (this *MyHashSet) Get(key int) interface{} {
	for _, k := range this.set[key%mod] {
		if k == key {
			return k
		}
	}
	return nil
}
