package interview

import (
	"context"
	"log"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"mormixz.com/interview-service/models"
	"mormixz.com/interview-service/store"
)

func mockService() *Service {
	dbStore, _ := store.NewStore("localhost:27017", "interview", 5)
	return NewService(dbStore)
}

func convertStringtoTime(datetime string) *time.Time {
	resTime, err := time.Parse("2006-01-02T15:04:05.000Z", datetime)
	if err != nil {
		log.Println(err.Error())
	}
	return &resTime
}

func TestGetInterview(t *testing.T) {
	s := mockService()

	// Case Get Success By ID
	caseName := "Case Get Success"
	expect := &models.Interview{
		ID:          convertStringtoPrimitiveID("645f4882bcc2918b5e4605b1"),
		Description: "description interview 002",
		Status:      STATUS_TODO,
		CreatedBy:   convertStringtoPrimitiveID("645f2b47e0931117713e74ca"),
		CreatedByUser: models.Users{
			ID:    convertStringtoPrimitiveID("645f2b47e0931117713e74ca"),
			Name:  "โรบินฮู้ด",
			Email: "user1@robinhood.co.th",
		},
		CreatedAt: convertStringtoTime("2023-01-01T10:00:00.000Z"),
		Comments: []models.Comment{
			{
				Message:   "comment 003",
				CreatedAt: convertStringtoTime("2023-01-02T18:00:00.000Z"),
				CreatedBy: convertStringtoPrimitiveID("645f2b47e0931117713e74cb"),
				CreatedByUser: models.Users{
					ID:    convertStringtoPrimitiveID("645f2b47e0931117713e74cb"),
					Name:  "แบทแมน",
					Email: "user2@robinhood.co.th",
				},
			},
			{
				Message:   "comment 002",
				CreatedAt: convertStringtoTime("2023-01-01T18:00:00.000Z"),
				CreatedBy: convertStringtoPrimitiveID("645f2b47e0931117713e74cc"),
				CreatedByUser: models.Users{
					ID:    convertStringtoPrimitiveID("645f2b47e0931117713e74cc"),
					Name:  "แคทวูแมน",
					Email: "user3@robinhood.co.th",
				},
			},
			{
				Message:   "comment 001",
				CreatedAt: convertStringtoTime("2023-01-01T15:00:00.000Z"),
				CreatedBy: convertStringtoPrimitiveID("645f2b47e0931117713e74cb"),
				CreatedByUser: models.Users{
					ID:    convertStringtoPrimitiveID("645f2b47e0931117713e74cb"),
					Name:  "แบทแมน",
					Email: "user2@robinhood.co.th",
				},
			},
		},
	}

	result, err := s.GetInterview("645f4882bcc2918b5e4605b1")
	if err != nil {
		t.Error(err.Error())
		return
	}
	assert.Equal(t, expect, result, caseName)
	// End Case Get Success By ID

	// Case Not Found ID
	caseName = "Case Not Found"
	result, err = s.GetInterview("645f2b47e0931117713e74ca")
	if err == nil {
		t.Error("Get Interview is should Failed")
	}

	assert.Equal(t, mongo.ErrNoDocuments, err, caseName)
	// End Case Not Found ID
}

func TestGetInterviewAll(t *testing.T) {
	s := mockService()

	// Case Get Status == To Do
	caseName := "Case Get To Do"
	expect := []*models.Interview{
		{
			ID:          convertStringtoPrimitiveID("645f4882bcc2918b5e4605b1"),
			Description: "description interview 002",
			Status:      STATUS_TODO,
			CreatedByUser: models.Users{
				ID:    convertStringtoPrimitiveID("645f2b47e0931117713e74ca"),
				Name:  "โรบินฮู้ด",
				Email: "user1@robinhood.co.th",
			},
			CreatedAt: convertStringtoTime("2023-01-01T10:00:00.000Z"),
		},
		{
			ID:          convertStringtoPrimitiveID("645f4882bcc2918b5e4605b2"),
			Description: "description interview 003",
			Status:      STATUS_TODO,
			CreatedByUser: models.Users{
				ID:    convertStringtoPrimitiveID("645f2b47e0931117713e74ca"),
				Name:  "โรบินฮู้ด",
				Email: "user1@robinhood.co.th",
			},
			CreatedAt: convertStringtoTime("2023-01-01T15:00:00.000Z"),
		},
		{
			ID:          convertStringtoPrimitiveID("645f4882bcc2918b5e4605b3"),
			Description: "description interview 004",
			Status:      STATUS_TODO,
			CreatedByUser: models.Users{
				ID:    convertStringtoPrimitiveID("645f2b47e0931117713e74ca"),
				Name:  "โรบินฮู้ด",
				Email: "user1@robinhood.co.th",
			},
			CreatedAt: convertStringtoTime("2023-01-02T10:00:00.000Z"),
		},
	}

	result, err := s.GetInterviewAll(STATUS_TODO, 3)
	if err != nil {
		t.Error(err.Error())
		return
	}

	assert.Equal(t, expect, result, caseName)
	// End Case Get Status == To Do

	// Case Get Status == In Progress
	caseName = "Case Get In Progress"
	expect = []*models.Interview{
		{
			ID:          convertStringtoPrimitiveID("645f4882bcc2918b5e4605b0"),
			Description: "description interview 001",
			Status:      STATUS_INPROGRESS,
			CreatedByUser: models.Users{
				ID:    convertStringtoPrimitiveID("645f2b47e0931117713e74ca"),
				Name:  "โรบินฮู้ด",
				Email: "user1@robinhood.co.th",
			},
			CreatedAt: convertStringtoTime("2023-01-01T09:00:00.000Z"),
		},
	}

	result, err = s.GetInterviewAll(STATUS_INPROGRESS, 1)
	if err != nil {
		t.Error(err.Error())
		return
	}

	assert.Equal(t, expect, result, caseName)
	// End Case Get Status == In Progress

	// Case Get All
	caseName = "Case Get All"
	expect = []*models.Interview{
		{
			ID:          convertStringtoPrimitiveID("645f4882bcc2918b5e4605af"),
			Description: "description interview 000",
			Status:      STATUS_DONE,
			CreatedByUser: models.Users{
				ID:    convertStringtoPrimitiveID("645f2b47e0931117713e74ca"),
				Name:  "โรบินฮู้ด",
				Email: "user1@robinhood.co.th",
			},
			CreatedAt: convertStringtoTime("2023-01-01T08:00:00.000Z"),
		},
		{
			ID:          convertStringtoPrimitiveID("645f4882bcc2918b5e4605b0"),
			Description: "description interview 001",
			Status:      STATUS_INPROGRESS,
			CreatedByUser: models.Users{
				ID:    convertStringtoPrimitiveID("645f2b47e0931117713e74ca"),
				Name:  "โรบินฮู้ด",
				Email: "user1@robinhood.co.th",
			},
			CreatedAt: convertStringtoTime("2023-01-01T09:00:00.000Z"),
		},
		{
			ID:          convertStringtoPrimitiveID("645f4882bcc2918b5e4605b1"),
			Description: "description interview 002",
			Status:      STATUS_TODO,
			CreatedByUser: models.Users{
				ID:    convertStringtoPrimitiveID("645f2b47e0931117713e74ca"),
				Name:  "โรบินฮู้ด",
				Email: "user1@robinhood.co.th",
			},
			CreatedAt: convertStringtoTime("2023-01-01T10:00:00.000Z"),
		},
		{
			ID:          convertStringtoPrimitiveID("645f4882bcc2918b5e4605b2"),
			Description: "description interview 003",
			Status:      STATUS_TODO,
			CreatedByUser: models.Users{
				ID:    convertStringtoPrimitiveID("645f2b47e0931117713e74ca"),
				Name:  "โรบินฮู้ด",
				Email: "user1@robinhood.co.th",
			},
			CreatedAt: convertStringtoTime("2023-01-01T15:00:00.000Z"),
		},
		{
			ID:          convertStringtoPrimitiveID("645f4882bcc2918b5e4605b3"),
			Description: "description interview 004",
			Status:      STATUS_TODO,
			CreatedByUser: models.Users{
				ID:    convertStringtoPrimitiveID("645f2b47e0931117713e74ca"),
				Name:  "โรบินฮู้ด",
				Email: "user1@robinhood.co.th",
			},
			CreatedAt: convertStringtoTime("2023-01-02T10:00:00.000Z"),
		},
		{
			ID:          convertStringtoPrimitiveID("645f4882bcc2918b5e4605b4"),
			Description: "description interview 005",
			Status:      STATUS_TODO,
			CreatedByUser: models.Users{
				ID:    convertStringtoPrimitiveID("645f2b47e0931117713e74ca"),
				Name:  "โรบินฮู้ด",
				Email: "user1@robinhood.co.th",
			},
			CreatedAt: convertStringtoTime("2023-01-03T10:00:00.000Z"),
		},
		{
			ID:          convertStringtoPrimitiveID("645f4882bcc2918b5e4605b5"),
			Description: "description interview 006",
			Status:      STATUS_TODO,
			CreatedByUser: models.Users{
				ID:    convertStringtoPrimitiveID("645f2b47e0931117713e74ca"),
				Name:  "โรบินฮู้ด",
				Email: "user1@robinhood.co.th",
			},
			CreatedAt: convertStringtoTime("2023-01-04T10:00:00.000Z"),
		},
	}

	result, err = s.GetInterviewAll("all", 0)
	if err != nil {
		t.Error(err.Error())
		return
	}

	assert.Equal(t, expect, result, caseName)
	// End Case Get All
}

func TestUpdateInterview(t *testing.T) {
	s := mockService()

	// Case UpdateInterview Success
	if err := s.UpdateInterview("645f4882bcc2918b5e4605b5", &models.Interview{
		Description: "update description interview 006",
		Status:      STATUS_INPROGRESS,
	}); err != nil {
		t.Error(err.Error())
		return
	}

	result, err := s.GetInterview("645f4882bcc2918b5e4605b5")
	if err != nil {
		t.Error(err.Error())
		return
	}

	expect := &models.Interview{
		ID:          convertStringtoPrimitiveID("645f4882bcc2918b5e4605b5"),
		Description: "update description interview 006",
		Status:      STATUS_INPROGRESS,
		CreatedBy:   convertStringtoPrimitiveID("645f2b47e0931117713e74ca"),
		CreatedByUser: models.Users{
			ID:    convertStringtoPrimitiveID("645f2b47e0931117713e74ca"),
			Name:  "โรบินฮู้ด",
			Email: "user1@robinhood.co.th",
		},
		UpdatedAt: result.UpdatedAt,
		CreatedAt: convertStringtoTime("2023-01-04T10:00:00.000Z"),
		Comments:  []models.Comment{},
	}

	assert.Equal(t, expect, result)
	// End Case UpdateInterview Success

	// Reset Data
	collection := s.dbStore.GetCollection(INTERVIEW_COLLECTION)
	query := bson.M{"_id": convertStringtoPrimitiveID("645f4882bcc2918b5e4605b5")}
	update := bson.M{"$set": bson.M{
		"description": "description interview 006",
		"status":      STATUS_TODO,
		"updated_at":  nil,
	}}

	if _, err := collection.UpdateOne(context.Background(), query, update); err != nil {
		t.Error(err.Error())
	}
}

func TestUpdateCommentInterview(t *testing.T) {
	s := mockService()

	// Case UpdateCommentInterview Success
	if err := s.UpdateCommentInterview("645f4882bcc2918b5e4605b5", models.Comment{
		Message: "add comment",
		CreatedByUser: models.Users{
			ID: convertStringtoPrimitiveID("645f2b47e0931117713e74ca"),
		},
	}); err != nil {
		t.Error(err.Error())
		return
	}

	result, err := s.GetInterview("645f4882bcc2918b5e4605b5")
	if err != nil {
		t.Error(err.Error())
		return
	}

	expect := &models.Interview{
		ID:          convertStringtoPrimitiveID("645f4882bcc2918b5e4605b5"),
		Description: "description interview 006",
		Status:      STATUS_TODO,
		CreatedBy:   convertStringtoPrimitiveID("645f2b47e0931117713e74ca"),
		CreatedByUser: models.Users{
			ID:    convertStringtoPrimitiveID("645f2b47e0931117713e74ca"),
			Name:  "โรบินฮู้ด",
			Email: "user1@robinhood.co.th",
		},
		CreatedAt: convertStringtoTime("2023-01-04T10:00:00.000Z"),
		Comments: []models.Comment{
			{
				Message:   "add comment",
				CreatedAt: result.Comments[0].CreatedAt,
				CreatedBy: convertStringtoPrimitiveID("645f2b47e0931117713e74ca"),
				CreatedByUser: models.Users{
					ID:    convertStringtoPrimitiveID("645f2b47e0931117713e74ca"),
					Name:  "โรบินฮู้ด",
					Email: "user1@robinhood.co.th",
				},
			},
		},
	}

	assert.Equal(t, expect, result)
	//End Case UpdateCommentInterview Success

	// Reset Data
	collection := s.dbStore.GetCollection(INTERVIEW_COLLECTION)
	query := bson.M{"_id": convertStringtoPrimitiveID("645f4882bcc2918b5e4605b5")}
	update := bson.M{"$set": bson.M{"comments": []models.Comment{}}}
	if _, err := collection.UpdateOne(context.Background(), query, update); err != nil {
		t.Error(err.Error())
	}
}

func TestGetUser(t *testing.T) {
	s := mockService()

	// Case Get User Success
	caseName := "Case Get Success"
	expect := models.Users{
		ID:    convertStringtoPrimitiveID("645f2b47e0931117713e74ca"),
		Name:  "โรบินฮู้ด",
		Email: "user1@robinhood.co.th",
	}

	result, err := s.GetUsers("645f2b47e0931117713e74ca")
	if err != nil {
		t.Error(err.Error())
		return
	}
	assert.Equal(t, expect, result, caseName)
	//End Case Get User Success

	// Case Get Not Found User
	caseName = "Case Not Found"
	_, err = s.GetUsers("")
	if err == nil {
		t.Error("Get Users is should Failed")
	}

	assert.Equal(t, mongo.ErrNoDocuments, err, caseName)
	// End Case Get Not Found User
}
