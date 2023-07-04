package mongo

import (
	"context"
	"fmt"

	"GoNews/pkg/storage"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Хранилище данных в MongoDB.
type Store struct {
	client     *mongo.Client
	collection *mongo.Collection
}

// Конструктор объекта хранилища.
func New(connectionString string) (*Store, error) {
	clientOptions := options.Client().ApplyURI(connectionString)

	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MongoDB: %w", err)
	}

	dbName := "mongodb"             // Замените на имя вашей базы данных
	collectionName := "collections" // Замените на имя вашей коллекции

	db := client.Database(dbName)
	collection := db.Collection(collectionName)

	return &Store{
		client:     client,
		collection: collection,
	}, nil
}
func (s *Store) Posts() ([]storage.Post, error) {
	// Здесь нужно выполнить запрос к коллекции с постами в MongoDB
	// и преобразовать полученные данные в структуры storage.Post.

	// Пример кода для выполнения запроса:
	cursor, err := s.collection.Find(context.Background(), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to execute MongoDB find query: %w", err)
	}
	defer cursor.Close(context.Background())

	var posts []storage.Post
	for cursor.Next(context.Background()) {
		var post storage.Post
		err := cursor.Decode(&post)
		if err != nil {
			return nil, fmt.Errorf("failed to decode document: %w", err)
		}
		posts = append(posts, post)
	}

	return posts, nil
}

func (s *Store) AddPost(post storage.Post) error {
	// Здесь нужно выполнить INSERT операцию для добавления поста в коллекцию MongoDB.

	_, err := s.collection.InsertOne(context.Background(), post)
	if err != nil {
		return fmt.Errorf("failed to execute MongoDB insert query: %w", err)
	}

	return nil
}

func (s *Store) UpdatePost(post storage.Post) error {
	// Здесь нужно выполнить UPDATE операцию для обновления поста в коллекции MongoDB.

	filter := primitive.M{"_id": post.ID}
	update := primitive.M{"$set": bson.M{"title": post.Title, "content": post.Content}}

	_, err := s.collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return fmt.Errorf("failed to execute MongoDB update query: %w", err)
	}

	return nil
}

func (s *Store) DeletePost(post storage.Post) error {
	// Здесь нужно выполнить DELETE операцию для удаления поста из коллекции MongoDB.

	filter := primitive.M{"_id": post.ID}

	_, err := s.collection.DeleteOne(context.Background(), filter)
	if err != nil {
		return fmt.Errorf("failed to execute MongoDB delete query: %w", err)
	}

	return nil
}
