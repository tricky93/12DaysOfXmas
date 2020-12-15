package christmas_test

import (
	christmas "12DaysOfXmas"
	"testing"
)

func TestGetGift(t *testing.T) {
	t.Run("returns correct verse when passed a day", func(t *testing.T) {
		got := christmas.GetGift(1)
		want := "And a partridge in a pear tree."

		if got != want {
			t.Errorf("got %s, want %s", got, want)
		}

		got = christmas.GetGift(10)
		want = "Ten lords-a-leaping,"

		if got != want {
			t.Errorf("got %s, want %s", got, want)
		}
	})
}

func TestGetOrdinalNumber(t *testing.T) {
	t.Run("returns correct ordinal number when passed a day", func(t *testing.T) {
		got := christmas.GetOrdinalNumber(1)
		want := "first"

		if got != want {
			t.Errorf("got %s, want %s", got, want)
		}

		got = christmas.GetOrdinalNumber(10)
		want = "tenth"

		if got != want {
			t.Errorf("got %s, want %s", got, want)
		}
	})
}

func TestCreateVerse(t *testing.T) {
	t.Run("returns the opening line of the verse when passed a day", func(t *testing.T) {
		got := christmas.CreateOpeningLines("first")
		want := "On the first day of Christmas,\nMy true love gave to me,\n"

		if got != want {
			t.Errorf("got %s, want %s", got, want)
		}

		got = christmas.CreateOpeningLines("fourth")
		want = "On the fourth day of Christmas,\nMy true love gave to me,\n"

		if got != want {
			t.Errorf("got %s, want %s", got, want)
		}
	})
}

func TestAppendGift(t *testing.T) {
	t.Run("appends the lastest gift to the global slice", func(t *testing.T) {
		got := christmas.AppendGift("Two turtle doves,")
		want := []string{"Two turtle doves,"}

		if len(got) != len(want) {
			t.Errorf("got %d, want %d", len(got), len(want))
		}
	})
}
