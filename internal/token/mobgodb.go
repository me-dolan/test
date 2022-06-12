package token

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"github.com/me-dolan/test/internal/user"
)

func (t *Tokens) creatDb(u user.User, at AuthTokens) error {
	err := t.db.Ping(context.Background(), nil)
	if err != nil {
		return err
	}

	collection := t.db.Database("UserSession").Collection("session")
	var session Session
	session.User = u
	session.RefreshToken, err = HashToken(at.RefreshToken)
	fmt.Println(session.RefreshToken)
	if err != nil {
		return err
	}

	_, err = collection.InsertOne(context.Background(), session)
	if err != nil {
		return err
	}

	return nil
}

func (t *Tokens) checkDb(refreshToken, guid string) (bool, error) {
	err := t.db.Ping(context.Background(), nil)
	if err != nil {
		return false, err
	}
	collection := t.db.Database("UserSession").Collection("Session")
	if err != nil {
		return false, err
	}
	cursor, err := collection.Find(context.Background(), bson.M{"user": bson.M{"guid": guid}})
	if err != nil {
		return false, err
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var epispde bson.M
		if err = cursor.Decode(&epispde); err != nil {
			return false, nil
		}
		str := fmt.Sprintf("%v", epispde["refreshtoken"])
		res := CheckTokenHash(refreshToken, str)
		if res {
			id := epispde["_id"]
			_, err := collection.DeleteOne(context.Background(), bson.M{"_id": id})
			if err != nil {
				return false, nil
			}
			return true, nil
		}

	}
	return false, nil
}
