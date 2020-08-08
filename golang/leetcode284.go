package main

func main() {

}

/*   Below is the interface for Iterator, which is already defined for you.
 *
 *   type Iterator struct {
 *
 *   }
 *
 *   func (this *Iterator) hasNext() bool {
 *		// Returns true if the iteration has more elements.
 *   }
 *
 *   func (this *Iterator) next() int {
 *		// Returns the next element in the iteration.
 *   }
 */

ttype PeekingIterator struct {
	iters *Iterator
	peeks int
	end   bool
}

func Constructor(iter *Iterator) *PeekingIterator {
	var peekIterator PeekingIterator
	peekIterator.end = iter.hasNext()
	peekIterator.iters = iter
	peekIterator.peeks = iter.next()

	return &peekIterator
}

func (this *PeekingIterator) hasNext() bool {
	return this.end
}

func (this *PeekingIterator) next() int {
	if this.end {
		tmp := this.peeks
		this.end = this.iters.hasNext()
		this.peeks = this.iters.next()
		return tmp
	} else {
		return this.peeks
	}

}

func (this *PeekingIterator) peek() int {
	return this.peeks
}