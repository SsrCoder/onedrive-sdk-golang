package microsoft

import "time"

type BaseResp struct {
	OdataContext string `json:"@odata.context"`
}

type FileInfo struct {
	MicrosoftGraphDownloadURL *string          `json:"@microsoft.graph.downloadUrl"`
	CreatedDateTime           time.Time        `json:"createdDateTime"`
	ETag                      string           `json:"eTag"`
	ID                        string           `json:"id"`
	LastModifiedDateTime      time.Time        `json:"lastModifiedDateTime"`
	Name                      string           `json:"name"`
	WebURL                    string           `json:"webUrl"`
	CTag                      string           `json:"cTag"`
	Size                      int              `json:"size"`
	CreatedBy                 *EditBy          `json:"createdBy,omitempty"`
	LastModifiedBy            *EditBy          `json:"lastModifiedBy,omitempty"`
	ParentReference           *ParentReference `json:"parentReference"`
	File                      *File            `json:"file"`
	FileSystemInfo            *FileSystemInfo  `json:"fileSystemInfo"`
	Folder                    *Folder          `json:"folder"`
}

type User struct {
	Email       string `json:"email"`
	ID          string `json:"id"`
	DisplayName string `json:"displayName"`
}

type EditBy struct {
	User        *User        `json:"user"`
	Application *Application `json:"application"`
}

type Application struct {
	ID          string `json:"id"`
	DisplayName string `json:"displayName"`
}

type ParentReference struct {
	DriveID   string `json:"driveId"`
	DriveType string `json:"driveType"`
	ID        string `json:"id"`
	Path      string `json:"path"`
}

type File struct {
	MimeType string    `json:"mimeType"`
	Hashes   *FileHash `json:"hashes"`
}

type FileHash struct {
	QuickXorHash string `json:"quickXorHash"`
}

type FileSystemInfo struct {
	CreatedDateTime      time.Time `json:"createdDateTime"`
	LastModifiedDateTime time.Time `json:"lastModifiedDateTime"`
}

type DirChildrenInfoResp struct {
	BaseResp
	Value []FileInfo `json:"value"`
}

type FileInfoResp struct {
	BaseResp
	FileInfo
}

type Folder struct {
	ChildCount int `json:"child_count"`
}

type ProfileResp struct {
	BaseResp
	Profile
}

type Profile struct {
	BusinessPhones    []interface{} `json:"businessPhones"`
	DisplayName       string        `json:"displayName"`
	GivenName         interface{}   `json:"givenName"`
	JobTitle          interface{}   `json:"jobTitle"`
	Mail              string        `json:"mail"`
	MobilePhone       interface{}   `json:"mobilePhone"`
	OfficeLocation    interface{}   `json:"officeLocation"`
	PreferredLanguage interface{}   `json:"preferredLanguage"`
	Surname           interface{}   `json:"surname"`
	UserPrincipalName string        `json:"userPrincipalName"`
	ID                string        `json:"id"`
}
