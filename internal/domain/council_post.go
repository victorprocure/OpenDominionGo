package domain

type CouncilPost struct {
	CouncilThread *CouncilThread
	Dominion      *Dominion
	Body          string
	IsDeleted     bool
}

func NewCouncilPost(councilThread *CouncilThread, dominion *Dominion, body string, isDeleted bool) *CouncilPost {
	return &CouncilPost{
		CouncilThread: councilThread,
		Dominion:      dominion,
		Body:          body,
		IsDeleted:     isDeleted,
	}
}