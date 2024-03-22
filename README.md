# Intro:

This simple program is designed to merge contact records based on certain fields such as phone number, email, and state ID. Overall, the program aims to consolidate duplicate contact records by merging them based on the shared attributes, providing a cleaner and more organized representation of contact information.


# How to Use:
Here's a more detailed guide of what the program does:
1. The program defines a Contact struct to represent an individual contact record. Each contact has fields defined as Phone, Email, StateID, and Platform (the CRM system).
2. The mergeContacts() function combines contacts with the same phone number, email, and state ID into one. It arranges the contacts by platform and then goes through them, merging those that share these three attributes.
3. During merging, the program looks at each contact's details like phone number, email, and state ID. If it finds a contact with the same combination of these details that's already been merged before, it updates that merged contact with any new information from the current contact.
4. The mergeValues function combines two contacts by checking for empty values in each attribute. If an attribute is empty in the first contact, it's replaced with the corresponding attribute from the second contact.
5. The main func prints the merged contacts with a simple fmt.Println.


# Sample Data:

This data could come from an API, a CSV file, a database, etc. In this instance I hardcoded into the program the data given for simplicity, keeping the same three sets below.

Reference:

### **Set A:**

{Phone: "2135550100", Email: "", StateID: "", Platform: "CRM 1"},

{Phone: "4155550199", Email: "test@alpineiq.com", StateID: "", Platform: "POS1"},

{Phone: "2135550100", Email: "", StateID: "", Platform: "CRM 1"},

### **Set B:**

{Phone: "2135550100", Email: "test@alpineiq.com", StateID: "A0001", Platform: "CRM 1"},

{Phone: "4155550199", Email: "test@alpineiq.com", StateID: "A0001", Platform: "CRM 2"},

{Phone: "2135550100", Email: "test@alpineiq.com", StateID: "A0001", Platform: "CRM 3"},

### **Set C:**

{Phone: "", Email: "", StateID: "A0001", Platform: "CRM 1"},

{Phone: "4155550199", Email: "", StateID: "", Platform: "CRM 2"},

{Phone: "2135550100", Email: "", StateID: "", Platform: "CRM 3"},

# Expectations:

### **Set A:**
{Phone: "2135550100", Email: "", StateID: "", Platform: "CRM 1"},

{Phone: "4155550199", Email: "test@alpineiq.com", StateID: "", Platform: "POS1"},

We can get rid of one entry of this data set, because the two can be combined into one entry based on the criteria for merging. The first and third entry looked like a duplicate, therefore it was merged into one. 

### **Set B:**
{Phone: "2135550100", Email: "test@alpineiq.com", StateID: "A0001", Platform: "CRM 1"},

{Phone: "4155550199", Email: "test@alpineiq.com", StateID: "A0001", Platform: "CRM 2"},

Based on the criteria for merging, we can elimate the third entry because it comes in the same from Platform: CRM1 as it does CRM 3. The CRM 3 merges into the CRM 1 entry. 

### **Set C:**
{Phone: "", Email: "", StateID: "A0001", Platform: "CRM 1"},

{Phone: "4155550199", Email: "", StateID: "", Platform: "CRM 2"},

{Phone: "2135550100", Email: "", StateID: "", Platform: "CRM 3"},

There is not enough criteria to merge any of these entries, so it remains unaffected.

# Limitations:

1. I understood this assignment to look at the three sample data sets as independent of each other. While the program works for all of these data sets, I assumed that if a contact was merged with two different Platforms values, then the returned merged contact would have the Platform value of the first contact, not both.
2. This program takes the assumption that an empty value means the value is missing and will look for it to be merged if there is a match. However, in real world scenarios, there could be an alternatve meaning for an empty value (for example, an empty email could mean they do not have one at all and will not find one in any of the CRM systems).
3. The sample data is given in a correct format for the program to consume. In real world scenarios, there needs to be a handling of invalid input data, bad data, messy data, etc.
4. Further test cases can be included to ensure all edge cases.

# Test/Test Results
The test compares the actual merged contacts with the expected merged contacts (expectedMerged). It uses the reflect.DeepEqual function to check if the two slices of contacts are equal. It verifies that the actual output matches the expected output to pass. If they are not equal, it reports an error, indicating that the test has failed.


# Thank you!
