package domain

type CouncilThread struct {
	Realm     *Realm
	Dominion  *Dominion
	Title     string
	Body      string
	IsDeleted bool
}

func NewCouncilThread(realm *Realm, dominion *Dominion, title, body string, isDeleted bool) *CouncilThread {
	return &CouncilThread{
		Realm:     realm,
		Dominion:  dominion,
		Title:     title,
		Body:      body,
		IsDeleted: isDeleted,
	}
}
