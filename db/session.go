package db

import "gopkg.in/mgo.v2"

// Session - Session Struct
type Session struct {
	session *mgo.Session
}

// NewSession - sessions
func NewSession() *Session {
	ses := new(Session)
	ses.session = connect().Copy()
	return ses
}

// Close - close the session
func (s *Session) Close() {
	if s.session != nil {
		s.session.Close()
	}
}

// Collection - get the specified collection with model
func (s *Session) Collection(collectionName string, model interface{}) *Collection {
	if s.session != nil {
		mgoCol := s.session.DB("").C(collectionName)
		c := new(Collection)
		c.col = mgoCol
		c.model = model
		return c
	}
	panic("Attempt to get Collection when session is nil")
}
