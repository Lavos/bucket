package bucket

type Bucket struct {
	slice []interface{}
}

func NewBucket(size uint) *Bucket {
	return &Bucket {
		slice: make([]interface{}, 0, size),
	}
}

func (b *Bucket) Bump(item interface{}) {
	if len(b.slice) == cap(b.slice) {
		t := make([]interface{}, cap(b.slice)-1, cap(b.slice));
		copy(t, b.slice[1:])
		b.slice = t
	}

	b.slice = append(b.slice, item)
}

func (b *Bucket) Get() []interface{} {
	return b.slice
}
