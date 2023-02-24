package Visitor

import "testing"

func TestVisitor(t *testing.T) {
	art := Article{
		title:           "title2",
		numberOfReaders: 20,
	}
	art.accept(&FanVisitor{})
	art.accept(&BloggerVisitor{})
}
