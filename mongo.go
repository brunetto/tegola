package tegola

import (
	"gopkg.in/mgo.v2"
)

func (b *Bot) Collection( collectionName string ) (*mgo.Collection, error) {

	var (
		session *mgo.Session
		err error
		collection *mgo.Collection
	)

	b.MongoSession, err = mgo.DialWithInfo(&(b.MongoAuth))
	if err != nil {
		return collection, err
	}

	collection = session.DB(b.MongoAuth.Database).C(collectionName)

	return collection, err
}

