package task

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"taskManagement/envornment"
	"testing"
)

func cleanupCollections(db *mongo.Database) {
	db.Collection(envornment.GetUserCollection()).Drop(context.Background())
	db.Collection(envornment.GetTaskCollection()).Drop(context.Background())
}

func newMongoClient() *mongo.Client {
	mongoTestClient, err := mongo.Connect(context.Background(), options.Client().ApplyURI(envornment.GetMongoURI()))
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Print("test client connection successful")
	err = mongoTestClient.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Print("test client connection successful")

	return mongoTestClient
}

func TestMongoOperations(t *testing.T) {
	mongoClient := newMongoClient()
	defer func() {
		err := mongoClient.Disconnect(context.Background())
		if err != nil {
			log.Printf("error while disconnecting mongo : %v", err)
		}
	}()

	// Connect to the test database
	db := mongoClient.Database("testDb")
	cleanupCollections(db)

	repo := NewRepository(db)

	t.Run("Register User", func(t *testing.T) {
		userBody := RegisterBody{
			Firstname: "Ashutosh",
			Lastname:  "Pandey",
			Email:     "ashutosh.pandey@gmail.com",
			Password:  "password123",
		}

		insertedID, err := repo.registerUser(userBody)
		require.NoError(t, err)
		require.NotEqual(t, primitive.NilObjectID, insertedID, "User should be registered successfully")
	})

	t.Run("Login User", func(t *testing.T) {
		loginRequest := LoginRequest{
			Email:    "ashutosh.pandey@gmail.com",
			Password: "password123",
		}

		user, err := repo.login(loginRequest)
		require.NoError(t, err)
		assert.Equal(t, "Ashutosh", user.Firstname, "Login should return correct user")
	})

	t.Run("Create Task", func(t *testing.T) {
		// Get the user from the login to extract the ID
		loginRequest := LoginRequest{
			Email:    "ashutosh.pandey@gmail.com",
			Password: "password123",
		}
		user, err := repo.login(loginRequest)
		require.NoError(t, err)

		taskRequest := Request{
			Task: "Complete unit testing",
		}

		err = repo.createTask(taskRequest, user.ID)
		require.NoError(t, err, "Task should be created successfully")
	})

	t.Run("Get All Tasks", func(t *testing.T) {
		// Get the user ID from the previous login
		loginRequest := LoginRequest{
			Email:    "ashutosh.pandey@gmail.com",
			Password: "password123",
		}
		user, err := repo.login(loginRequest)
		require.NoError(t, err)

		tasks, err := repo.getAllTasks(user.ID)
		require.NoError(t, err)
		assert.Len(t, tasks, 1, "There should be one active task")
		assert.Equal(t, "Complete unit testing", tasks[0].Task, "Task content should match")
	})

	t.Run("Delete Task", func(t *testing.T) {
		loginRequest := LoginRequest{
			Email:    "ashutosh.pandey@gmail.com",
			Password: "password123",
		}
		user, err := repo.login(loginRequest)
		require.NoError(t, err)

		// Get all tasks, then delete the first one
		tasks, err := repo.getAllTasks(user.ID)
		require.NoError(t, err)

		err = repo.deleteTask(tasks[0].ID)
		require.NoError(t, err, "Task should be archived successfully")

		// Verify the task is archived
		tasks, err = repo.getAllTasks(user.ID)
		require.NoError(t, err)
		assert.Len(t, tasks, 0, "No active tasks should remain after deletion")
	})
}
