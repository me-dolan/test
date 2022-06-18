package tokens

import (
	"context"
	"fmt"

	"github.com/me-dolan/test/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
)

// func (t *Tokens) creatDb(u domain.User, at AuthTokens) error {
// 	err := t.db.Ping(context.Background(), nil)
// 	if err != nil {
// 		return err
// 	}
// 	collection := t.db.Database("UserSession").Collection("session")

// 	u.RefreshToken, err = HashToken(at.RefreshToken)
// 	fmt.Println(u.RefreshToken)
// 	if err != nil {
// 		return err
// 	}
// 	f := collection.FindOne(context.Background(), bson.D{{"guid", u.Guid}}).Decode(&u)
// 	fmt.Println(f)
// 	if f == mongo.ErrNoDocuments {
// 		_, err = collection.InsertOne(context.Background(), u)
// 		if err != nil {
// 			return err
// 		}
// 		//  else { 
// 		// 	filter := bson.D{{"guid", u.Guid}}
// 		// 	update := bson.D{
// 		// 		{"$set", bson.M{
// 		// 			"session": bson.M{
// 		// 				"refreshtoken": u.Session.RefreshToken,
// 		// 			},
// 		// 		}},
// 		// 	}
// 		// 	_, err := collection.UpdateOne(context.Background(), filter, update)
// 		// 	if err != nil {
// 		// 		return err
// 		// 	}
// 		// }
// 	} else {
// 		filter := bson.D{{"guid", u.Guid}}
// 		update := bson.D{
// 			{"$set", bson.M{
// 				"refreshtoken": u.RefreshToken,
// 			}},
// 		}
// 		_, err := collection.UpdateOne(context.Background(), filter, update)
// 		if err != nil {
// 			return err
// 		}
// 	}
// 	//  else {
// 	// 	_, err = collection.InsertOne(context.Background(), u)
// 	// 	if err != nil {
// 	// 		return err
// 	// 	}
// 	// _, err = collection.InsertOne(context.Background(), session)
// 	// if err != nil {
// 	// 	return err
// 	// }

// 	return nil
// }
func (t *Tokens) creatDb(u domain.User, at AuthTokens) error {
	err := t.db.Ping(context.Background(), nil)
	if err != nil {
		return err
	}
	collection := t.db.Database("UserSession").Collection("session")

	u.RefreshToken, err = HashToken(at.RefreshToken)
	fmt.Println(u.RefreshToken)
	if err != nil {
		return err
	}
	_, err = collection.InsertOne(context.Background(), u)
	if err != nil {
		return err
	}

	return nil
}

func (t *Tokens) checkDb(refreshToken string, guid string) (bool, error) {
	err := t.db.Ping(context.Background(), nil)
	if err != nil {
		return false, err
	}
	collection := t.db.Database("UserSession").Collection("session")
	if err != nil {
		return false, err
	}
	cursor, err := collection.Find(context.Background(), bson.M{"guid": guid})
	if err != nil {
		return false, err
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var episode bson.M
		if err = cursor.Decode(&episode); err != nil {
			return false, err
		}
		str := fmt.Sprintf("%v", episode["refreshtoken"])
		res := CheckTokenHash(refreshToken, str)
		if res {
			// id := episode["_id"]
			// _, err := collection.DeleteOne(context.Background(), bson.M{"_id": id})
			// if err != nil {
			// 	return false, nil
			// }
			return true, nil
		}
	}
	return false, nil
}

func (t *Tokens) refreshDbToken(u domain.User, at AuthTokens) error {
	err := t.db.Ping(context.Background(), nil)
	if err != nil {
		return err
	}
	collection := t.db.Database("UserSession").Collection("session")

	u.RefreshToken, err = HashToken(at.RefreshToken)
	fmt.Println(u.RefreshToken)
	if err != nil {
		return err
	}

	filter := bson.D{{"guid", u.Guid}}
	update := bson.D{
		{"$set", bson.D{
			{"refreshtoken", u.RefreshToken},
		},
		}}
	_, err = collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}

	return nil
}