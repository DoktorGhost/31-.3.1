package mongo_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"GoNews/pkg/storage"
	"GoNews/pkg/storage/mongo"
)

func TestMongoDB_Posts(t *testing.T) {

	connectionString := "mongodb://localhost:27017/testdb"
	store, err := mongo.New(connectionString)
	if err != nil {
		t.Fatalf("ошибка подключения к MongoDB: %v", err)
	}

	// Создаем тестовую коллекцию записей
	testPosts := []storage.Post{
		{
			ID:         1,
			Title:      "Новый пост",
			Content:    "Это новый пост",
			AuthorID:   1,
			AuthorName: "Автор 1",
		},
		{
			ID:         2,
			Title:      "Новый пост 2",
			Content:    "Это второй новый пост",
			AuthorID:   2,
			AuthorName: "Автор 2",
		},
		{
			ID:         3,
			Title:      "Новый пост 3",
			Content:    "Совершенно новый пост под номером 3",
			AuthorID:   3,
			AuthorName: "Автор 3",
		},
	}

	for _, post := range testPosts {
		err := store.AddPost(post)
		if err != nil {
			t.Fatalf("Ошибка добаления записи: %v", err)
		}
	}

	//Получаем список записей
	posts, err := store.Posts()
	if err != nil {
		t.Fatalf("не удалось получить запись: %v", err)
	}

	for i, expectedPost := range testPosts {
		actualPost := posts[i]

		if expectedPost.ID != actualPost.ID {
			t.Errorf("ID не совпадают %d и %d", expectedPost.ID, actualPost.ID)
		}

		if expectedPost.Title != actualPost.Title {
			t.Error("Название постов не совпадает")
		}

		if expectedPost.Content != actualPost.Content {
			t.Error("Текст постов не совпадает")
		}

	}

}

func TestMongoDB_UpdatePost(t *testing.T) {
	connString := "mongodb://localhost:27017/testdb"
	store, err := mongo.New(connString)
	if err != nil {
		t.Fatalf("ошибка подключения к MongoDB: %v", err)
	}

	// Создаем новую запись
	post := storage.Post{
		ID:         5,
		Title:      "Тестовая запись 5",
		Content:    "Эту запись нужно срочно изменить!",
		AuthorID:   100,
		AuthorName: "Автор 100",
	}

	// Добавляем запись в хранилище
	err = store.AddPost(post)
	if err != nil {
		t.Fatalf("ошибка добавления записи: %v", err)
	}

	// Обновляем содержимое записи
	post.Content = "Это измененная запись!"
	err = store.UpdatePost(post)
	if err != nil {
		t.Fatalf("Ошибка обновления записи: %v", err)
	}

	// Получаем обновленную запись для проверки
	updatedPost, err := store.GetPost(5)
	if err != nil {
		t.Fatalf("не удалось получить обновленную запись: %v", err)
	}

	assert.Equal(t, updatedPost.Content, "Это измененная запись!", "содержимое обновленной записи должно соответствовать ожидаемому значению")
}

func TestMongoDB_DeletePost(t *testing.T) {
	connString := "mongodb://localhost:27017/testdb"
	store, err := mongo.New(connString)
	if err != nil {
		t.Fatalf("ошибка подключения к MongoDB: %v", err)
	}

	err = store.DeleteAllPosts()
	if err != nil {
		t.Fatalf("ошибка удаления всех постов: %v", err)
	}

	// Создаем новую запись
	post := storage.Post{
		ID:         6,
		Title:      "Пост для удаления",
		Content:    "Этот пост будет удален!",
		AuthorID:   666,
		AuthorName: "Автор демон",
	}

	// Добавляем запись в хранилище
	err = store.AddPost(post)
	if err != nil {
		t.Fatalf("ошибка добавления записи: %v", err)
	}

	// Удаляем запись
	err = store.DeletePost(post)
	if err != nil {
		t.Fatalf("ошибка удаления записи: %v", err)
	}

	// Пытаемся получить удаленную запись для проверки
	_, err = store.GetPost(post.ID)
	assert.Equal(t, storage.ErrPostNotFound, err, "удаленный пост не должен быть найден")
}
