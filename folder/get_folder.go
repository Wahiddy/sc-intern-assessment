package folder

import (
	"github.com/gofrs/uuid"
	"strings"
	"errors"
)

func GetAllFolders() []Folder {
	return GetSampleData()
}

func (f *driver) GetFoldersByOrgID(orgID uuid.UUID) []Folder {
	folders := f.folders

	res := []Folder{}
	for _, f := range folders {
		if f.OrgId == orgID {
			res = append(res, f)
		}
	}

	return res

}

func (f *driver) GetAllChildFolders(orgID uuid.UUID, name string) ([]Folder, error) {
	// Your code here...
	tempFolders := f.folders

	// Existence of folder is checked in a single loop
	folderExists := false
	folderOrgExists := false
	folders := []Folder{}
	for _, t := range tempFolders {
		if t.Name == name {
			folderExists = true
		}

		if t.OrgId == orgID {
			folders = append(folders, t)
			if t.Name == name {
				folderOrgExists = true
			}
		}
	}
	if !folderExists {
		return nil, errors.New("Folder does not exist")
	}
	if !folderOrgExists {
		return nil, errors.New("Folder does not exist in the specified organization")
	}

	res := []Folder{}
	for _, fo := range folders {
		if fo.Name==name {
			continue
		}
		nameIndex := strings.Index(fo.Paths, name)

		if nameIndex != -1{
			reqIndex := strings.Index(fo.Paths, fo.Name)
			if nameIndex < reqIndex {
				res = append(res, fo)
			}
		} 
	}

	return res, nil
}
