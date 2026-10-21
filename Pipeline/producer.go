package pipeline

import (
	"encoding/csv"
	"sync"
	// "fmt"
	"os"
)

// import "os"

func LoadRecipient(filepath string, queue chan Emails, wg *sync.WaitGroup) error {
	defer wg.Done()
	defer close(queue)
	data, err := os.Open(filepath)

	if err != nil {
		return err
	}

	defer data.Close()

	r := csv.NewReader(data)
	records, err := r.ReadAll()
	if err != nil {
		return err
	}

	for _, record := range records[1:] {
		queue <- Emails{
			Name:  record[0],
			Email: record[1],
		}

	}
	return nil
}
