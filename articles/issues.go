package articles

type Issue struct{
	URL string
	FileName string
	Name string
}

type Issues struct{
	Prefix string
	List []Issue
}

type CurrentIssue struct{
	Name string
	CurrentPast string
}
