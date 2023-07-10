package mongo

import (
	"context"
	"fmt"
	"log"
	"time"

	"GoNews/pkg/storage"

	"go.mongodb.org/mongo-driver/bson"
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

	collection := client.Database("testdb").Collection("posts")

	return &Store{
		client:     client,
		collection: collection,
	}, nil
}

//вывести список всех постов
func (s *Store) Posts() ([]storage.Post, error) {

	cursor, err := s.collection.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, fmt.Errorf("Не удалось выполнить поиск %w", err)
	}
	defer cursor.Close(context.Background())

	var results []storage.Post

	for cursor.Next(context.TODO()) {

		var elem storage.Post
		err := cursor.Decode(&elem)
		if err != nil {
			return nil, fmt.Errorf("Ошибка декодинга %w", err)
		}

		results = append(results, elem)
	}

	if err := cursor.Err(); err != nil {
		log.Fatal(err)
	}

	cursor.Close(context.TODO())
	return results, nil

}

// добавление поста (отредактированно)
func (s *Store) AddPost(post storage.Post) error {

	post.CreatedAt = time.Now().Unix()
	post.PublishedAt = time.Now().Unix()

	_, err := s.collection.InsertOne(context.TODO(), post)
	if err != nil {
		return fmt.Errorf("Ошибка добавления записи: %w", err)
	}

	return nil
}

//обновление поста
func (s *Store) UpdatePost(post storage.Post) error {
	filter := bson.M{"id": post.ID}
	update := bson.M{
		"$set": bson.M{
			"Title":       post.Title,
			"Content":     post.Content,
			"AuthorID":    post.AuthorID,
			"AuthorName":  post.AuthorName,
			"PublishedAt": time.Now().Unix(),
		},
	}

	_, err := s.collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return fmt.Errorf("failed to execute MongoDB update query: %w", err)
	}

	return nil
}

func (s *Store) DeletePost(post storage.Post) error {

	filter := bson.D{{"id", post.ID}}

	_, err := s.collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return fmt.Errorf("failed to execute MongoDB delete query: %w", err)
	}

	return nil
}

func (s *Store) GetPost(id int) (storage.Post, error) {
	var post storage.Post
	filter := bson.M{"id": id}
	err := s.collection.FindOne(context.Background(), filter).Decode(&post)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return post, storage.ErrPostNotFound
		}
		return post, fmt.Errorf("failed to execute MongoDB find query: %w", err)
	}
	return post, nil
}

func (s *Store) DeleteAllPosts() error {
	_, err := s.collection.DeleteMany(context.Background(), bson.M{})
	if err != nil {
		return fmt.Errorf("failed to execute MongoDB delete query: %w", err)
	}
	return nil
}
