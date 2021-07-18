package utils

import "fmt"

// PrintKey prints a desired key and its values based off of the inferred type
func PrintKey(id string, m interface{}) {
	switch m.(type) {
	case string:
		fmt.Printf("%s: \n\n", id)
		fmt.Printf("\t%s\n", m)
	case []string:
		fmt.Printf("%s: \n\n", id)
		for _, value := range m.([]string) {
			fmt.Printf("\t%s\n", value)
		}
	case map[string]string:
		fmt.Printf("%s: \n\n", id)
		for key, value := range m.(map[string]string) {
			fmt.Printf("\t%s: %s\n", key, value)
		}
	}
	fmt.Println()
}