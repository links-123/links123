package mongo

import (
	"github.com/globalsign/mgo"
	"github.com/pkg/errors"
)

func InitConnection(mongoURL, mongoDB string) (*mgo.Database, error) {
	if mongoURL == "" {
		return nil, errors.New("mongo URL is empty")
	}

	session, err := mgo.Dial(mongoURL)
	if err != nil {
		return nil, errors.Wrap(err, "unable to establish connection to Mongo DB")
	}

	return session.DB(mongoDB), nil
}
