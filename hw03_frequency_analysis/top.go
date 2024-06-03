package hw03frequencyanalysis

import (
	"fmt"
	"sort"
	"strings"
)

var slice []kv

type kv struct {
	Key   string
	Value int
}

func (k kv) String() string {
	return fmt.Sprintf("%s %d\n", k.Key, k.Value)
}

func Top10(text string) []string {
	var out []string

	wordCounts := countWords(text)
	for k, v := range wordCounts {
		slice = append(slice, kv{k, v})
	}

	sort.Slice(slice, func(i, j int) bool {
		return slice[i].Value > slice[j].Value
	})

	result := make(map[int][]string)
	for i, kv := range slice {
		if i != 10 {
			result[kv.Value] = append(result[kv.Value], kv.Key)
		} else {
			break
		}
	}

	var builder strings.Builder

	// Выводим результат
	for _, value := range result {
		for _, word := range MySorting(value) {
			builder.WriteString(word)
			out = append(out, builder.String())
			builder.Reset()
		}
	}
	return out
}

func countWords(input string) map[string]int {
	words := strings.Fields(input)
	wordCount := make(map[string]int)
	for _, word := range words {
		wordCount[word]++
	}
	return wordCount
}

func MySorting(str []string) []string {
	var result []string
	sort.Strings(str)
	result = append(result, str...)
	return result
}
