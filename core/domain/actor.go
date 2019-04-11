package domain

type Actor struct {
	ID_    UserID
	Token_ string
}

func (a Actor) UserID() UserID {
	return a.ID_
}

func (a Actor) Token() string {
	return a.Token_
}

func NewTestActor(id int) Actor {
	actor := Actor{}
	actor.ID_ = UserID{Value: id}
	actor.Token_ = "dummy"
	return actor
}
