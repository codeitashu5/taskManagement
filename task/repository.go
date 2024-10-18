package task

import (
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
	"taskManagement/envornment"
	"taskManagement/models"
	"taskManagement/util"
	"time"
)

type Repository struct {
	UserCollection *mongo.Collection
	TaskCollection *mongo.Collection
}

func NewRepository(db *mongo.Database) *Repository {
	return &Repository{
		UserCollection: db.Collection(envornment.GetUserCollection()),
		TaskCollection: db.Collection(envornment.GetTaskCollection()),
	}
}

func (r *Repository) registerUser(body RegisterBody) (primitive.ObjectID, error) {
	// register the user and return the registered user
	var (
		user       User
		insertedID primitive.ObjectID
	)

	// check if email already exists or not
	err := r.UserCollection.FindOne(context.Background(), bson.M{"email": body.Email}).Decode(&user)
	if err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
		return insertedID, err
	}

	// check if user and request body have same email
	if user.Email == body.Email {
		return insertedID, errors.New("user already exists")
	}

	encryptedPassword, err := util.EncryptPassword(body.Password)
	if err != nil {
		return insertedID, err
	}

	user = User{
		ID:        primitive.NewObjectID(),
		Firstname: body.Firstname,
		Lastname:  body.Lastname,
		Email:     body.Email,
		Password:  encryptedPassword,
	}

	result, err := r.UserCollection.InsertOne(context.Background(), user)
	if err != nil {
		return insertedID, err
	}

	insertedID = result.InsertedID.(primitive.ObjectID)
	return insertedID, nil
}

func (r *Repository) login(request LoginRequest) (models.User, error) {
	// check if email already exists or not
	user := models.User{}

	// check if the user with email password exits
	err := r.UserCollection.FindOne(context.Background(), bson.M{"email": request.Email}).Decode(&user)
	if err != nil {
		return user, err
	}

	passwordBytes := []byte(request.Password)
	passwordDB := []byte(user.Password)

	err = bcrypt.CompareHashAndPassword(passwordDB, passwordBytes)
	if err != nil {
		return user, err
	}

	// check if the email password is correct
	return user, nil
}

func (r *Repository) createTask(request Request, userId primitive.ObjectID) error {
	task := Task{
		ID:         primitive.NewObjectID(),
		UserID:     userId,
		Task:       request.Task,
		CreatedAt:  time.Now(),
		ArchivedAt: nil,
	}

	// check if the user with email password exits
	_, err := r.TaskCollection.InsertOne(context.Background(), task)
	return err
}

func (r *Repository) getTask(taskID primitive.ObjectID) (*Task, error) {
	var task Task
	err := r.TaskCollection.FindOne(context.Background(), bson.M{"_id": taskID, "archived_at": bson.M{"$eq": nil}}).Decode(&task)
	if err != nil {
		return nil, err
	}
	return &task, nil
}

func (r *Repository) getAllTasks(userId primitive.ObjectID) ([]Task, error) {
	var tasks []Task
	cursor, err := r.TaskCollection.Find(context.Background(), bson.M{"user_id": userId, "archived_at": bson.M{"$eq": nil}})
	if err != nil {
		return nil, err
	}

	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		var task Task
		if err := cursor.Decode(&task); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}

func (r *Repository) deleteTask(taskID primitive.ObjectID) error {
	var task Task
	err := r.TaskCollection.FindOne(context.Background(), bson.M{"_id": taskID}).Decode(&task)
	if err != nil {
		return err
	}

	// Check if the task is already archived (deleted)
	if task.ArchivedAt != nil {
		return errors.New("task is archived")
	}

	// Mark the task as archived by setting ArchivedAt to the current time
	update := bson.M{
		"$set": bson.M{"archived_at": time.Now()},
	}

	_, err = r.TaskCollection.UpdateOne(context.Background(), bson.M{"_id": taskID}, update)
	return err
}

func (r *Repository) updateTask(taskID primitive.ObjectID, taskData Request) error {
	var task Task
	err := r.TaskCollection.FindOne(context.Background(), bson.M{"_id": taskID, "archived_at": bson.M{"$ne": nil}}).Decode(&task)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return fmt.Errorf("task not found or not archived")
		}
		return err
	}

	update := bson.M{
		"$set": bson.M{
			"task": taskData.Task,
		},
	}
	_, err = r.TaskCollection.UpdateOne(context.Background(), bson.M{"_id": taskID, "archived_at": bson.M{"$ne": nil}}, update)
	return err
}
