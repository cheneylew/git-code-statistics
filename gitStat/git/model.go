package git

import "time"

type Commit struct {
	Branch string
	Id string
	Author string
	Email string
	Date time.Time
	Message string
	IsMerge bool
	IsRevert bool
	MergeIds []string
	AddLines int
	RemoveLines int
	TotalLines int
	WorkDir string
}
