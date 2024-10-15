package main

import (
	"fmt"

	"github.com/georgechieng-sc/interns-2022/folder"
	"github.com/gofrs/uuid"
)

func main() {
	orgID := uuid.FromStringOrNil(folder.DefaultOrgID)

	res := folder.GetAllFolders()

	// example usage
	folderDriver := folder.NewDriver(res)
	orgFolder, err := folderDriver.MoveFolder("bravo", "delta")

	// folder.PrettyPrint(res)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("\n Folders for orgID: %s", orgID)
		folder.PrettyPrint(orgFolder)
	}
}
