package main

import (
	"app/gen/ent"
	"app/quick"
	"context"
	"fmt"
	"sync"
)

var plugins quick.Plugins

func main() {
	plugins = *quick.GetPlugins()

	characters := "abcdefghijklmnopqrstuvwxyz"
	charList := []rune(characters)
	length := 7 // Length of the combinations

	// Channel to collect words
	wordChan := make(chan string, 5000) // Buffered channel

	// WaitGroup for synchronization
	var wg sync.WaitGroup

	// Start generating combinations in a separate goroutine
	wg.Add(1)
	go func() {
		defer wg.Done()
		generateCombinations(charList, "", length, wordChan)
		close(wordChan) // Close channel after generation is done
	}()

	// Start batch insertion in a separate goroutine
	wg.Add(1)
	go func() {
		defer wg.Done()
		batchInsertWords(wordChan)
	}()

	// Wait for all operations to complete
	wg.Wait()
}

// generateCombinations generates and sends all combinations of the given length using the characters in charList to a channel
func generateCombinations(charList []rune, prefix string, length int, wordChan chan<- string) {
	if length == 0 {
		fmt.Println(prefix)
		wordChan <- prefix
		return
	}

	for i := 0; i < len(charList); i++ {
		newPrefix := prefix + string(charList[i])
		generateCombinations(charList, newPrefix, length-1, wordChan)
	}
}

// batchInsertWords inserts the words into the database in batches from a channel
func batchInsertWords(wordChan <-chan string) {
	ctx := context.Background()
	batchSize := 1000 // Define your batch size here
	batch := make([]string, 0, batchSize)

	for word := range wordChan {
		batch = append(batch, word)
		if len(batch) >= batchSize {
			insertBatch(ctx, batch)
			batch = batch[:0] // Clear the batch
		}
	}

	// Insert any remaining words
	if len(batch) > 0 {
		insertBatch(ctx, batch)
	}
}

// insertBatch inserts a batch of words into the database
func insertBatch(ctx context.Context, batch []string) {
	bulk := make([]*ent.WordCreate, len(batch))
	for i, word := range batch {
		bulk[i] = plugins.EntDB.Client().Word.Create().SetName(word)
	}
	_, err := plugins.EntDB.Client().Word.CreateBulk(bulk...).Save(ctx)
	if err != nil {
		fmt.Printf("Failed to insert batch: %v\n", err)
	}
}
