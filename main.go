package main

import (
	"fmt"
	"sort"
)

// Contract struct with attributes of interest.
type Contact struct {
	Phone    string
	Email    string
	StateID  string
	Platform string
}

func main() {

	// Example contact sets from README.
	set1 := []Contact{
		{Phone: "2135550100", Email: "", StateID: "", Platform: "CRM 1"},
		{Phone: "4155550199", Email: "test@alpineiq.com", StateID: "", Platform: "POS1"},
		{Phone: "2135550100", Email: "", StateID: "", Platform: "CRM 1"},
	}

	// set2 := []Contact{
	// 	{Phone: "2135550100", Email: "test@alpineiq.com", StateID: "A0001", Platform: "CRM 1"},
	// 	{Phone: "4155550199", Email: "test@alpineiq.com", StateID: "A0001", Platform: "CRM 2"},
	// 	{Phone: "2135550100", Email: "test@alpineiq.com", StateID: "A0001", Platform: "CRM 3"},
	// }

	// set3 := []Contact{
	// 	{Phone: "", Email: "", StateID: "A0001", Platform: "CRM 1"},
	// 	{Phone: "4155550199", Email: "", StateID: "", Platform: "CRM 2"},
	// 	{Phone: "2135550100", Email: "", StateID: "", Platform: "CRM 3"},
	// }

	merged := mergeContacts(set1)

	//Print the merged contacts in a readable format as the output.
	for _, contact := range merged {
		fmt.Printf("  Phone: %s\n", contact.Phone)
		fmt.Printf("  Email: %s\n", contact.Email)
		fmt.Printf("  StateID: %s\n", contact.StateID)
		fmt.Printf("  Platform: %s\n", contact.Platform)
		fmt.Println("----------")
	}
}

// mergeContacts merges contacts with duplicate data in entries
func mergeContacts(contacts []Contact) []Contact {
	//sort contacts by the platform field
	sort.Slice(contacts, func(i, j int) bool {
		return contacts[i].Platform < contacts[j].Platform
	})

	merged := make(map[string]Contact)
	mergedList := make([]Contact, 0)

	//merge contacts
	for _, contact := range contacts {
		key := contact.Phone + contact.Email + contact.StateID
		if _, ok := merged[key]; !ok {
			merged[key] = contact
		} else {
			merged[key] = mergeValues(merged[key], contact)
		}
	}

	//converts the map of merged contacts into a slice
	for _, contact := range merged {
		mergedList = append(mergedList, contact)
	}

	//return slice
	return mergedList
}

// mergeValues takes in the two Contact structs that are to be merged, checks for empty values.
// Returns one Contact struct as a result.
func mergeValues(contact1, contact2 Contact) Contact {
	if contact1.Phone == "" {
		contact1.Phone = contact2.Phone
	}
	if contact1.Email == "" {
		contact1.Email = contact2.Email
	}
	if contact1.StateID == "" {
		contact1.StateID = contact2.StateID
	}
	return contact1
}
