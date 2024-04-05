package repository

import (
	"AudioTranscription/serve/models"
	"go.mongodb.org/mongo-driver/mongo"
)

const UsersCollection = "users"

type UsersRepository interface {
	Save(user *models.User) error
	Update(user *models.User) error
	GetById(id string) (*models.User, error)
	GetByEmail(email string) (*models.User, error)
	GetAll() ([]*models.User, error)
	Delete(id string) error
}

type usersRepository struct {
	collection *mongo.Collection
}

//func NewUsersRepository(conn db.Connection) UsersRepository {
//	return &usersRepository{collection: conn.DB().Collection(UsersCollection)}
//}
//
//func (r *usersRepository) Save(user *models.User) error {
//	_, err := r.collection.InsertOne(context.Background(), user)
//	return err
//}
//
//func (r *usersRepository) Update(user *models.User) error {
//	filter := bson.M{"_id": user.Id}
//	update := bson.M{"$set": user}
//	_, err := r.collection.UpdateOne(context.Background(), filter, update)
//	return err
//}
//
//func (r *usersRepository) GetById(id string) (*models.User, error) {
//	var user models.User
//	filter := bson.M{"_id": id}
//	err := r.collection.FindOne(context.Background(), filter).Decode(&user)
//	if err != nil {
//		return nil, err
//	}
//	return &user, nil
//}
//
//func (r *usersRepository) GetByEmail(email string) (*models.User, error) {
//	var user models.User
//	filter := bson.M{"email": email}
//	err := r.collection.FindOne(context.Background(), filter).Decode(&user)
//	if err != nil {
//		return nil, err
//	}
//	return &user, nil
//}
//
//func (r *usersRepository) GetAll() ([]*models.User, error) {
//	var users []*models.User
//	cur, err := r.collection.Find(context.Background(), bson.M{})
//	if err != nil {
//		return nil, err
//	}
//	defer cur.Close(context.Background())
//
//	for cur.Next(context.Background()) {
//		var user models.User
//		err := cur.Decode(&user)
//		if err != nil {
//			return nil, err
//		}
//		users = append(users, &user)
//	}
//
//	if err := cur.Err(); err != nil {
//		return nil, err
//	}
//
//	return users, nil
//}
//
//func (r *usersRepository) Delete(id string) error {
//	filter := bson.M{"_id": id}
//	_, err := r.collection.DeleteOne(context.Background(), filter)
//	return err
//}
