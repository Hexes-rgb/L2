package main

import (
	"fmt"
	"sort"
	"strings"
)

// функция для преобразования слова в ключ анаграммы
func createAnagramKey(word string) string {
	chars := strings.Split(word, "")
	sort.Strings(chars)
	return strings.Join(chars, "")
}

// функция для поиска всех множеств анаграмм по словарю.
func findAnagrams(dict []string) map[string][]string {
	anagramSets := make(map[string][]string)

	for _, word := range dict {
		wordLower := strings.ToLower(word)
		key := createAnagramKey(wordLower)

		anagramSets[key] = append(anagramSets[key], wordLower)
	}

	// удаляем множества из одного элемента.
	for key, set := range anagramSets {
		if len(set) < 2 {
			delete(anagramSets, key)
		} else {
			sort.Strings(set)
		}
	}

	// формируем итоговую мапу с ключом из первого слова множества.
	finalSets := make(map[string][]string)
	for _, set := range anagramSets {
		key := set[0]
		finalSets[key] = set
	}

	return finalSets
}

func main() {
	dictionary := []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик", "слово"}
	anagrams := findAnagrams(dictionary)

	for key, set := range anagrams {
		fmt.Printf("Key: '%s', Set: %v\n", key, set)
	}
}
