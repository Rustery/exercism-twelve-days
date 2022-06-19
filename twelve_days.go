package twelve

import (
	"fmt"
	"strings"
)

var positions = []string{
	"first",
	"second",
	"third",
	"fourth",
	"fifth",
	"sixth",
	"seventh",
	"eighth",
	"ninth",
	"tenth",
	"eleventh",
	"twelfth",
}

var objects = []string{
	"a Partridge in a Pear Tree",
	"two Turtle Doves, and",
	"three French Hens,",
	"four Calling Birds,",
	"five Gold Rings,",
	"six Geese-a-Laying,",
	"seven Swans-a-Swimming,",
	"eight Maids-a-Milking,",
	"nine Ladies Dancing,",
	"ten Lords-a-Leaping,",
	"eleven Pipers Piping,",
	"twelve Drummers Drumming,",
}

func Verse(i int) string {
	list := []string{}
	for j := i - 1; j >= 0; j-- {
		list = append(list, objects[j])
	}
	return fmt.Sprintf("On the %s day of Christmas my true love gave to me: %s.", positions[i-1], strings.Join(list, " "))
}

func Song() string {
	verses := []string{}
	for i := 1; i <= len(positions); i++ {
		verses = append(verses, Verse(i))
	}
	return strings.Join(verses, "\n")
}
