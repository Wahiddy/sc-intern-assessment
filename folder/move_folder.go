package folder

import (
	"github.com/gofrs/uuid"
	"strings"
	"errors"
)

// Get the path of the dst, and wherever you see the '.name' in any path replace everything before it with the path of dst
func (f *driver) MoveFolder(name string, dst string) ([]Folder, error) {
	// Your code here...
	if name == dst {
		return nil, errors.New("Cannot move a folder to itself")
	}

	folders := f.folders

	// Find the orgid and extract path
	var src_id uuid.UUID = uuid.Nil
	var dst_id uuid.UUID = uuid.Nil

	var src_path string
	var dst_path string

	for _, f := range folders {
		if f.Name == name {
			src_id = f.OrgId
			src_path = f.Paths
		} else if f.Name == dst {
			dst_id = f.OrgId
			dst_path = f.Paths
		}
	}	

	if strings.HasPrefix(dst_path, src_path + ".") {
		return nil, errors.New("Cannot move a folder to a child of itself")
	}

	if src_id == uuid.Nil {
		return nil, errors.New("Source folder does not exist")
	}

	if dst_id == uuid.Nil {
		return nil, errors.New("Destination folder does not exist")
	}

	if src_id != dst_id {
		return nil, errors.New("Cannot move a folder to a different organization")
	}

	res := []Folder{}
	for _, f := range folders {
		if src_id == f.OrgId {
			if strings.Contains(f.Paths, name) {
				newPath := strings.Replace(f.Paths, src_path, dst_path + "." + name, 1)
				res = append(res, Folder{
					Name:  f.Name,
					OrgId: f.OrgId,
					Paths: newPath,
				})
			} else {
				res = append(res, f)
			}
		}
	}

	return res, nil
}
