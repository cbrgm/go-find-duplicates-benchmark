package main

func findDuplicatesTypeToBool(strings []string) []string {
	// use a map type -> bool
	seenStrings := make(map[string]bool)
	var duplicates []string
	for _, str := range strings {
		if seenStrings[str] {
			duplicates = append(duplicates, str)
		} else {
			seenStrings[str] = true
		}
	}
	return duplicates
}

func findDuplicatesTypeToInterface(strings []string) []string {
	seenStrings := make(map[string]interface{})
	var duplicates []string
	for _, str := range strings {
		if _, ok := seenStrings[str]; ok {
			duplicates = append(duplicates, str)
		} else {
			seenStrings[str] = true
		}
	}
	return duplicates
}

func findDuplicatesTypeToStruct(strings []string) []string {
	seenStrings := make(map[string]struct{})
	var duplicates []string
	for _, str := range strings {
		if _, ok := seenStrings[str]; ok {
			duplicates = append(duplicates, str)
		} else {
			seenStrings[str] = struct{}{}
		}
	}
	return duplicates
}
