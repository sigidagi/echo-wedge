package main

import "log"

func main() {
	givenID := []string{"netId", "devId", "valId", "stateId"}
	// wantedID := []string{"netId", "devId"}
	// log.Println(len(wantedID))

	all, _ := findIDs(givenID, []string{"netId", "devId"})
	one, _ := findExact("netId", all)

	log.Println("Exact devID : ", one)
}

func findIDs(slice []string, desired []string) ([]string, error) {
	var returnedIds []string

	for _, v := range slice {
		for _, dID := range desired {
			if v == dID {
				returnedIds = append(returnedIds, dID)
				break
			}
		}
	}
	return returnedIds, nil
}

func findExact(id string, slice []string) (string, error) {
	var exactID string
	for _, v := range slice {
		if v == id {
			exactID = id
			break
		}
	}
	return exactID, nil
}