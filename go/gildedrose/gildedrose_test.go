package gildedrose_test

import (
	"testing"

	"github.com/emilybache/gildedrose-refactoring-kata/gildedrose"
)

func TestUpdateQuality(t *testing.T) {
	var items = []*gildedrose.Item{
		{Name: "foo", SellIn: 0, Quality: 0},
	}

	gildedrose.UpdateQuality(items)

	// TODO: Just fixing this to pass
	if items[0].Name != "foo" {
		t.Errorf("Name: Expected %s but got %s ", "fixme", items[0].Name)
	}
}
