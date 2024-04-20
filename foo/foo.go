package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	err := run()
	if err != nil {
		fmt.Println("Failed:", err)
		os.Exit(1)
	}
}

func run() error {
	data, err := os.ReadFile("peq.txt")
	if err != nil {
		return fmt.Errorf("os.Open: %w", err)
	}

	type entry struct {
		name       string
		spellid    string
		merchantid string
	}

	entries := []entry{}

	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		records := strings.Split(line, "\t")
		merchant := records[0]
		merchant = strings.ReplaceAll(merchant, "0", "")
		merchantid := records[1]

		spellid := records[3]

		entries = append(entries, entry{name: merchant, spellid: spellid, merchantid: merchantid})
	}

	data, err = os.ReadFile("live.txt")
	if err != nil {
		return fmt.Errorf("os.Open: %w", err)
	}

	onLive := []entry{}

	for _, line := range strings.Split(string(data), "\n") {
		records := strings.Split(line, "\t")
		merchant := records[2]
		spellid := records[3]

		merchant = strings.ReplaceAll(merchant, "0", "")

		isFound := false
		merchantid := ""
		for i, entry := range entries {
			if entry.name == merchant {
				merchantid = entry.merchantid
				if entry.spellid == spellid {
					isFound = true
					entries = append(entries[:i], entries[i+1:]...)
					break
				}
			}
		}
		if !isFound {
			onLive = append(onLive, entry{name: merchant, spellid: spellid, merchantid: merchantid})
		}
	}

	for _, entry := range onLive {
		fmt.Println("missing", entry.name, entry.spellid, entry.merchantid)
	}

	for _, entry := range entries {
		fmt.Println("extra", entry.name, entry.spellid, entry.merchantid)
	}

	return nil
}
