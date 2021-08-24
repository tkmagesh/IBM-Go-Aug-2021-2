package main

import "strings"

func main() {
	str := "Dolore duis velit magna eu sunt et excepteur cupidatat ullamco non ex Aliquip amet consequat enim enim occaecat quis fugiat officia quis aliquip pariatur Consequat cillum eiusmod proident culpa duis dolor incididunt id occaecat ex consectetur id dolor Mollit fugiat sit irure do cupidatat officia deserunt laboris ametElit ut dolore incididunt irure elit consectetur Quis exercitation reprehenderit nostrud occaecat nisi ipsum nulla cillum quis labore tempor minim magna ullamco Et nostrud nostrud enim sunt esse excepteur exercitation ad officia sint aliqua exercitation doCillum Lorem reprehenderit minim minim Voluptate magna ipsum incididunt voluptate cillum enim Aute cupidatat fugiat tempor sint labore dolore dolore consectetur id anim ea voluptate Occaecat dolore do aute nulla reprehenderit Minim magna aliquip magna commodo dolore culpa sint nostrud sit non fugiat Eu sunt incididunt deserunt enim adipisicing nulla quis nostrud culpa dolore Mollit pariatur id velit qui duis irureId ea minim commodo labore ullamco proident laborum ad adipisicing Elit ullamco officia culpa magna amet voluptate dolor Occaecat pariatur ea consectetur sunt anim et Lorem in tempor labore velit pariatur velit Minim qui aliquip ad aliquip id nulla labore dolore nulla culpa culpa Irure laborum mollit veniam dolore Lorem Commodo aliqua esse incididunt sit in aliquip incididunt deserunt"

	words := strings.Split(str, " ")
	wordsBySize := getWordsBySize(words)
	maxSize, maxOccurance := getMaxOccuranceBySize(wordsBySize)
	println("Max size:", maxSize, "Max occurance:", maxOccurance)
}

func getWordsBySize(words []string) map[int]int {
	wordsBySize := map[int]int{}
	for _, word := range words {
		size := len(word)
		wordsBySize[size]++
	}
	return wordsBySize
}

func getMaxOccuranceBySize(wordsBySize map[int]int) (int, int) {
	maxOccurance := 0
	maxSize := 0
	for size, occurance := range wordsBySize {
		if occurance > maxOccurance {
			maxOccurance = occurance
			maxSize = size
		}
	}
	return maxSize, maxOccurance
}
