package folder_test

import (
	"testing"

	"github.com/georgechieng-sc/interns-2022/folder"
	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/assert"
)

/*
	Testing MoveFolders - cases with no errors
*/
func Test_No_Errors_MoveFolders(t *testing.T) {
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
			name:  "Test - all paths have same parent folder",
			orgID: testUID,  
			folders: []folder.Folder{
				{Name: "alpha", OrgId: testUID, Paths: "alpha"},
				{Name: "bravo", OrgId: testUID, Paths: "alpha.bravo"},
				{Name: "charlie", OrgId: testUID, Paths: "alpha.bravo.charlie"},
				{Name: "delta", OrgId: testUID, Paths: "alpha.delta"},
				{Name: "echo", OrgId: testUID, Paths: "alpha.delta.echo"},
				{Name: "foxtrot", OrgId: otherUID, Paths: "foxtrot"},
			},
			want: []folder.Folder{
				{Name: "alpha", OrgId: testUID, Paths: "alpha"},
				{Name: "bravo", OrgId: testUID, Paths: "alpha.delta.bravo"},
				{Name: "charlie", OrgId: testUID, Paths: "alpha.delta.bravo.charlie"},
				{Name: "delta", OrgId: testUID, Paths: "alpha.delta"},
				{Name: "echo", OrgId: testUID, Paths: "alpha.delta.echo"},
			},
		},
		{
			name:  "Test - paths have different parent folder",
			orgID: testUID,  
			folders: []folder.Folder{
				{Name: "alpha", OrgId: testUID, Paths: "alpha"},
				{Name: "bravo", OrgId: testUID, Paths: "alpha.bravo"},
				{Name: "charlie", OrgId: testUID, Paths: "alpha.bravo.charlie"},
				{Name: "delta", OrgId: testUID, Paths: "foxtrot.delta"},
				{Name: "echo", OrgId: testUID, Paths: "foxtrot.delta.echo"},
				{Name: "foxtrot", OrgId: testUID, Paths: "foxtrot"},
				{Name: "golf", OrgId: testUID, Paths: "golf"},
			},
			want: []folder.Folder{
				{Name: "alpha", OrgId: testUID, Paths: "alpha"},
				{Name: "bravo", OrgId: testUID, Paths: "foxtrot.delta.bravo"},
				{Name: "charlie", OrgId: testUID, Paths: "foxtrot.delta.bravo.charlie"},
				{Name: "delta", OrgId: testUID, Paths: "foxtrot.delta"},
				{Name: "echo", OrgId: testUID, Paths: "foxtrot.delta.echo"},
				{Name: "foxtrot", OrgId: testUID, Paths: "foxtrot"},
				{Name: "golf", OrgId: testUID, Paths: "golf"},
			},
		},
		{
			name:  "Test - src folder is parent in all paths",
			orgID: testUID,  
			folders: []folder.Folder{
				{Name: "alpha", OrgId: testUID, Paths: "alpha"},
				{Name: "bravo", OrgId: testUID, Paths: "bravo"},
				{Name: "charlie", OrgId: testUID, Paths: "bravo.charlie"},
				{Name: "delta", OrgId: testUID, Paths: "alpha.delta"},
				{Name: "echo", OrgId: testUID, Paths: "alpha.delta.echo"},
				{Name: "foxtrot", OrgId: otherUID, Paths: "foxtrot"},
				{Name: "golf", OrgId: testUID, Paths: "golf"},
			},
			want: []folder.Folder{
				{Name: "alpha", OrgId: testUID, Paths: "alpha"},
				{Name: "bravo", OrgId: testUID, Paths: "alpha.delta.bravo"},
				{Name: "charlie", OrgId: testUID, Paths: "alpha.delta.bravo.charlie"},
				{Name: "delta", OrgId: testUID, Paths: "alpha.delta"},
				{Name: "echo", OrgId: testUID, Paths: "alpha.delta.echo"},
				{Name: "golf", OrgId: testUID, Paths: "golf"},
			},
		},
		{
			name:  "Test - dst folder is root folder",
			orgID: testUID,  
			folders: []folder.Folder{
				{Name: "alpha", OrgId: testUID, Paths: "foxtrot.alpha"},
				{Name: "bravo", OrgId: testUID, Paths: "foxtrot.alpha.bravo"},
				{Name: "charlie", OrgId: testUID, Paths: "foxtrot.alpha.bravo.charlie"},
				{Name: "delta", OrgId: testUID, Paths: "delta"},
				{Name: "echo", OrgId: testUID, Paths: "foxtrot.alpha.delta.echo"},
				{Name: "foxtrot", OrgId: testUID, Paths: "foxtrot"},
				{Name: "golf", OrgId: testUID, Paths: "golf"},
			},
			want: []folder.Folder{
				{Name: "alpha", OrgId: testUID, Paths: "foxtrot.alpha"},
				{Name: "bravo", OrgId: testUID, Paths: "delta.bravo"},
				{Name: "charlie", OrgId: testUID, Paths: "delta.bravo.charlie"},
				{Name: "delta", OrgId: testUID, Paths: "delta"},
				{Name: "echo", OrgId: testUID, Paths: "foxtrot.alpha.delta.echo"},
				{Name: "foxtrot", OrgId: testUID, Paths: "foxtrot"},
				{Name: "golf", OrgId: testUID, Paths: "golf"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := folder.NewDriver(tt.folders)
			get, err := f.MoveFolder("bravo", "delta")
			assert.Equal(t, get, tt.want, "They should be equal")
			assert.Nil(t, err)
		})
	}
}

/*
	Testing MoveFolders - cases with errors
*/
func Test_Errors_MoveFolder(t *testing.T) {
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
			name:  "Test - Cannot move a folder to a child of itself",
			orgID: testUID, 
			folders: []folder.Folder{
				{Name: "alpha", OrgId: testUID, Paths: "alpha"},
				{Name: "bravo", OrgId: testUID, Paths: "alpha.bravo"},
				{Name: "charlie", OrgId: testUID, Paths: "alpha.bravo.charlie"},
				{Name: "delta", OrgId: testUID, Paths: "alpha.bravo.delta"},
				{Name: "echo", OrgId: testUID, Paths: "alpha.delta.echo"},
				{Name: "foxtrot", OrgId: otherUID, Paths: "foxtrot"},
			},
			errMsg: "Cannot move a folder to a child of itself",
		},
		{
			name:  "Test - Source folder does not exist",
			orgID: testUID, 
			folders: []folder.Folder{
				{Name: "alpha", OrgId: testUID, Paths: "alpha"},
				{Name: "charlie", OrgId: testUID, Paths: "alpha.charlie"},
				{Name: "delta", OrgId: testUID, Paths: "alpha.delta"},
				{Name: "echo", OrgId: testUID, Paths: "alpha.delta.echo"},
				{Name: "foxtrot", OrgId: otherUID, Paths: "foxtrot"},
			},
			errMsg: "Source folder does not exist",
		},
		{
			name:  "Test - Destination folder does not exist",
			orgID: testUID, 
			folders: []folder.Folder{
				{Name: "alpha", OrgId: testUID, Paths: "alpha"},
				{Name: "bravo", OrgId: testUID, Paths: "alpha.bravo"},
				{Name: "charlie", OrgId: testUID, Paths: "alpha.bravo.charlie"},
				{Name: "echo", OrgId: testUID, Paths: "alpha.echo"},
				{Name: "foxtrot", OrgId: otherUID, Paths: "foxtrot"},
			},
			errMsg: "Destination folder does not exist",
		},
		{
			name:  "Test - Cannot move a folder to a different organization",
			orgID: testUID, 
			folders: []folder.Folder{
				{Name: "alpha", OrgId: testUID, Paths: "alpha"},
				{Name: "bravo", OrgId: testUID, Paths: "alpha.bravo"},
				{Name: "charlie", OrgId: testUID, Paths: "alpha.bravo.charlie"},
				{Name: "delta", OrgId: otherUID, Paths: "alpha.delta"},
				{Name: "echo", OrgId: testUID, Paths: "alpha.delta.echo"},
				{Name: "foxtrot", OrgId: otherUID, Paths: "foxtrot"},
			},
			errMsg: "Cannot move a folder to a different organization",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := folder.NewDriver(tt.folders)
			get, err := f.MoveFolder("bravo", "delta")

			assert.Nil(t, get, "Folder list should be nil when there's an error")
			assert.EqualError(t, err, tt.errMsg)
		})
	}
}

/*
	Testing MoveFolders - special case for error 
*/
func Test_MoveFolder_CannotMoveToItself(t *testing.T) {
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
			name:  "Test - Cannot move a folder to itself",
			orgID: testUID, 
			folders: []folder.Folder{
				{Name: "bravo", OrgId: testUID, Paths: "bravo"},
				{Name: "charlie", OrgId: testUID, Paths: "bravo.charlie"},
				{Name: "delta", OrgId: testUID, Paths: "delta"},
				{Name: "echo", OrgId: otherUID, Paths: "echo"},
				{Name: "foxtrot", OrgId: testUID, Paths: "foxtrot"},
			},
			errMsg: "Cannot move a folder to itself",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := folder.NewDriver(tt.folders)
			get, err := f.MoveFolder("bravo", "bravo")

			assert.Nil(t, get, "Folder list should be nil when there's an error")
			assert.EqualError(t, err, tt.errMsg)
		})
	}
}
