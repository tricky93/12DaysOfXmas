package christmas

var gifts = map[int]string{1: "And a partridge in a pear tree.", 2: "Two turtle doves,", 3: "Three french hens,", 4: "Four calling birds,", 5: "Five golden rings,", 6: "Six geese-a-laying,", 7: "Six geese-a-laying,", 8: "Eight maids-a-milking,", 9: "Nine ladies dancing,", 10: "Ten lords-a-leaping,", 11: "Eleven pipers piping", 12: "Twelve drummers drumming,"}

var days = map[int]string{1: "first", 2: "second", 3: "third", 4: "fouth", 5: "fifth", 6: "sixth", 7: "seventh", 8: "eighth", 9: "ninth", 10: "tenth", 11: "eleventh", 12: "twelfth"}

var song []string

func GetGift(day int) string {
	return gifts[day]
}

func GetOrdinalNumber(day int) string {
	return days[day]
}

func CreateOpeningLines(ordinal string) (openingLines string) {

	lineOne := "On the " + ordinal + " day of Christmas,\n"
	lineTwo := "My true love gave to me,\n"

	return lineOne + lineTwo

}

func AppendGift(gift string) (giftList []string) {

	return append(giftList, gift)
}
