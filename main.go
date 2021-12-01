package main

import "fmt"

func CompareSlicesB(got, want []string) error {
	var (
		unexpectedEntries = []string{}
		notFoundEntries   = []string{}
	)

Got:
	for _, entry := range got {
		for _, wanted := range want {
			if entry == wanted {
				continue Got
			}
		}
		unexpectedEntries = append(unexpectedEntries, entry)
	}

Want:
	for _, wanted := range want {
		for _, entry := range got {
			if entry == wanted {
				continue Want
			}
		}
		notFoundEntries = append(notFoundEntries, wanted)
	}

	if len(unexpectedEntries) != 0 || len(notFoundEntries) != 0 {
		return fmt.Errorf("got %d unexpected entries, and %d not found entries", len(unexpectedEntries), len(notFoundEntries))
	}

	return nil

}

func CompareSlicesA(got, want []string) error {
	w := make(map[string]bool, len(want))
	var (
		unexpectedEntries = []string{}
		notFoundEntries   = []string{}
	)

	for _, entry := range want {
		w[entry] = false
	}

	for _, entry := range got {
		if _, ok := w[entry]; !ok {
			unexpectedEntries = append(unexpectedEntries, entry)
		}
		w[entry] = true
	}

	for entry, ok := range w {
		if !ok {
			notFoundEntries = append(notFoundEntries, entry)
		}
	}

	if len(unexpectedEntries) != 0 || len(notFoundEntries) != 0 {
		return fmt.Errorf("got %d unexpected entries, and %d not found entries", len(unexpectedEntries), len(notFoundEntries))
	}

	return nil
}
