package folder_test

import (
	"testing"

	"github.com/georgechieng-sc/interns-2022/folder"
	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/assert"
)

// orgId's used for testing
const testUIDStr = "c1556e17-b7c0-45a3-a6ae-9546248fb17a"
const otherUIDStr = "38b9879b-f73b-4b0e-b9d9-4fc4c23643a7"

/*
	Testing GetFoldersByOrgID
*/
func Test_folder_GetFoldersByOrgID(t *testing.T) {
	t.Parallel()
	testUID := uuid.FromStringOrNil(testUIDStr)
	otherUID := uuid.FromStringOrNil(otherUIDStr)
	tests := [...]struct {
		name    string
		orgID   uuid.UUID
		folders []folder.Folder
		want    []folder.Folder
	}{
		{
			name:  "Test - all folders are the same org",
			orgID: testUID,  // Replace with a real UUID
			folders: []folder.Folder{
				{Name: "alpha", OrgId: testUID, Paths: "alpha"},
				{Name: "bravo", OrgId: testUID, Paths: "alpha.bravo"},
				{Name: "charlie", OrgId: testUID, Paths: "alpha.bravo.charlie"},
				{Name: "delta", OrgId: testUID, Paths: "alpha.delta"},
				{Name: "echo", OrgId: testUID, Paths: "echo"},
				{Name: "foxtrot", OrgId: testUID, Paths: "foxtrot"},
			},
			want: []folder.Folder{
				{Name: "alpha", OrgId: testUID, Paths: "alpha"},
				{Name: "bravo", OrgId: testUID, Paths: "alpha.bravo"},
				{Name: "charlie", OrgId: testUID, Paths: "alpha.bravo.charlie"},
				{Name: "delta", OrgId: testUID, Paths: "alpha.delta"},
				{Name: "echo", OrgId: testUID, Paths: "echo"},
				{Name: "foxtrot", OrgId: testUID, Paths: "foxtrot"},
			},
		},
		{
			name:  "Test - different orgs",
			orgID: testUID,  
			folders: []folder.Folder{
				{Name: "alpha", OrgId: testUID, Paths: "alpha"},
				{Name: "bravo", OrgId: otherUID, Paths: "alpha.bravo"},
				{Name: "charlie", OrgId: otherUID, Paths: "alpha.bravo.charlie"},
				{Name: "delta", OrgId: testUID, Paths: "alpha.delta"},
				{Name: "echo", OrgId: testUID, Paths: "echo"},
				{Name: "foxtrot", OrgId: otherUID, Paths: "foxtrot"},
			},
			want: []folder.Folder{
				{Name: "alpha", OrgId: testUID, Paths: "alpha"},
				{Name: "delta", OrgId: testUID, Paths: "alpha.delta"},
				{Name: "echo", OrgId: testUID, Paths: "echo"},
			},
		},
		{
			name:  "Test - no org/empty output",
			orgID: testUID,  
			folders: []folder.Folder{
				{Name: "alpha", OrgId: otherUID, Paths: "alpha"},
				{Name: "bravo", OrgId: otherUID, Paths: "alpha.bravo"},
				{Name: "charlie", OrgId: otherUID, Paths: "alpha.bravo.charlie"},
				{Name: "foxtrot", OrgId: otherUID, Paths: "foxtrot"},
			},
			want: []folder.Folder{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := folder.NewDriver(tt.folders)
			get := f.GetFoldersByOrgID(testUID)
			assert.Equal(t, get, tt.want, "They should be equal")
		})
	}
}

/*
	Testing GetAllChildFolders - cases with no errors
*/
func Test_No_Errors_GetAllChildFolders(t *testing.T) {
	t.Parallel()
	testUID := uuid.FromStringOrNil(testUIDStr)
	otherUID := uuid.FromStringOrNil(otherUIDStr)
	tests := [...]struct {
		name    string
		orgID   uuid.UUID
		folders []folder.Folder
		want    []folder.Folder
	}{
		{
			name:  "Test - all folders are the same org",
			orgID: testUID,  
			folders: []folder.Folder{
				{Name: "alpha", OrgId: testUID, Paths: "alpha"},
				{Name: "bravo", OrgId: testUID, Paths: "alpha.bravo"},
				{Name: "charlie", OrgId: testUID, Paths: "alpha.bravo.charlie"},
				{Name: "delta", OrgId: testUID, Paths: "alpha.delta"},
				{Name: "echo", OrgId: testUID, Paths: "echo"},
				{Name: "foxtrot", OrgId: testUID, Paths: "foxtrot"},
			},
			want: []folder.Folder{
				{Name: "bravo", OrgId: testUID, Paths: "alpha.bravo"},
				{Name: "charlie", OrgId: testUID, Paths: "alpha.bravo.charlie"},
				{Name: "delta", OrgId: testUID, Paths: "alpha.delta"},
			},
		},
		{
			name:  "Test - different orgs",
			orgID: testUID, 
			folders: []folder.Folder{
				{Name: "alpha", OrgId: testUID, Paths: "alpha"},
				{Name: "bravo", OrgId: testUID, Paths: "alpha.bravo"},
				{Name: "charlie", OrgId: otherUID, Paths: "alpha.bravo.charlie"},
				{Name: "delta", OrgId: testUID, Paths: "alpha.delta"},
				{Name: "echo", OrgId: testUID, Paths: "echo"},
				{Name: "foxtrot", OrgId: testUID, Paths: "foxtrot"},
			},
			want: []folder.Folder{
				{Name: "bravo", OrgId: testUID, Paths: "alpha.bravo"},
				{Name: "delta", OrgId: testUID, Paths: "alpha.delta"},
			},
		},
		{
			name:  "Test - child folder within nested folder",
			orgID: testUID, 
			folders: []folder.Folder{
				{Name: "alpha", OrgId: testUID, Paths: "golf.alpha"},
				{Name: "bravo", OrgId: testUID, Paths: "golf.alpha.bravo"},
				{Name: "charlie", OrgId: otherUID, Paths: "golf.alpha.bravo.charlie"},
				{Name: "delta", OrgId: testUID, Paths: "golf.alpha.delta"},
				{Name: "echo", OrgId: testUID, Paths: "echo"},
				{Name: "foxtrot", OrgId: testUID, Paths: "foxtrot"},
				{Name: "golf", OrgId: testUID, Paths: "golf"},
			},
			want: []folder.Folder{
				{Name: "bravo", OrgId: testUID, Paths: "golf.alpha.bravo"},
				{Name: "delta", OrgId: testUID, Paths: "golf.alpha.delta"},
			},
		},
		{
			name:  "Test - no child folders/empty output",
			orgID: testUID, 
			folders: []folder.Folder{
				{Name: "alpha", OrgId: testUID, Paths: "delta.alpha"},
				{Name: "bravo", OrgId: testUID, Paths: "bravo"},
				{Name: "charlie", OrgId: testUID, Paths: "charlie"},
				{Name: "delta", OrgId: testUID, Paths: "delta"},
				{Name: "echo", OrgId: testUID, Paths: "echo"},
				{Name: "foxtrot", OrgId: testUID, Paths: "foxtrot"},
			},
			want: []folder.Folder{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := folder.NewDriver(tt.folders)
			get, err := f.GetAllChildFolders(testUID, "alpha")
			assert.Equal(t, get, tt.want, "They should be equal")
			assert.Nil(t, err)
		})
	}
}

/*
	Testing GetAllChildFolders - cases with errors
*/
func Test_Errors_GetAllChildFolders(t *testing.T) {
	t.Parallel()
	testUID := uuid.FromStringOrNil(testUIDStr)
	otherUID := uuid.FromStringOrNil(otherUIDStr)
	tests := [...]struct {
		name    string
		orgID   uuid.UUID
		folders []folder.Folder
		errMsg  string
	}{
		{
			name:  "Test - folder does not exist",
			orgID: testUID, 
			folders: []folder.Folder{
				{Name: "bravo", OrgId: testUID, Paths: "bravo"},
				{Name: "charlie", OrgId: testUID, Paths: "bravo.charlie"},
				{Name: "delta", OrgId: testUID, Paths: "delta"},
				{Name: "echo", OrgId: testUID, Paths: "echo"},
				{Name: "foxtrot", OrgId: testUID, Paths: "foxtrot"},
			},
			errMsg: "Folder does not exist",
		},
		{
			name:  "Test - folder does not exist in org",
			orgID: testUID,  
			folders: []folder.Folder{
				{Name: "alpha", OrgId: otherUID, Paths: "delta.alpha"},
				{Name: "bravo", OrgId: testUID, Paths: "bravo"},
				{Name: "charlie", OrgId: testUID, Paths: "bravo.charlie"},
				{Name: "delta", OrgId: testUID, Paths: "delta"},
				{Name: "echo", OrgId: testUID, Paths: "echo"},
				{Name: "foxtrot", OrgId: testUID, Paths: "foxtrot"},
			},
			errMsg: "Folder does not exist in the specified organization",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := folder.NewDriver(tt.folders)
			get, err := f.GetAllChildFolders(testUID, "alpha")

			assert.Nil(t, get, "Folder list should be nil when there's an error")
			assert.EqualError(t, err, tt.errMsg)
		})
	}
}
