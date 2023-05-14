package interview

import (
	"context"
	"sort"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"mormixz.com/interview-service/models"
	"mormixz.com/interview-service/store"
)

var (
	INTERVIEW_COLLECTION = "interview"
	USER_COLLECTION      = "users"

	STATUS_TODO       = "To Do"
	STATUS_INPROGRESS = "In Progress"
	STATUS_DONE       = "Done"
)

type Service struct {
	dbStore *store.Store
}

func NewService(dbStore *store.Store) *Service {
	return &Service{
		dbStore,
	}
}

func convertStringtoPrimitiveID(id string) primitive.ObjectID {
	primitiveID, _ := primitive.ObjectIDFromHex(id)
	return primitiveID
}

func (s *Service) GetInterview(id string) (*models.Interview, error) {
	collection := s.dbStore.GetCollection(INTERVIEW_COLLECTION)

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(s.dbStore.ConnectionTimeout)*time.Second)
	defer cancel()

	interview := &models.Interview{}
	query := bson.M{"_id": convertStringtoPrimitiveID(id)}
	if err := collection.FindOne(ctx, query).Decode(&interview); err != nil {
		return nil, err
	}

	interview.CreatedByUser, _ = s.GetUsers(interview.CreatedBy.Hex())

	for index, comment := range interview.Comments {
		interview.Comments[index].CreatedByUser, _ = s.GetUsers(comment.CreatedBy.Hex())
	}

	sort.Slice(interview.Comments, func(i, j int) bool {
		return interview.Comments[i].CreatedAt.After(*interview.Comments[j].CreatedAt)
	})

	return interview, nil
}

func (s *Service) GetInterviewAll(status string, limit int) ([]*models.Interview, error) {
	collection := s.dbStore.GetCollection(INTERVIEW_COLLECTION)

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(s.dbStore.ConnectionTimeout)*time.Second)
	defer cancel()

	var query interface{}
	if status == "all" {
		query = bson.D{}
	} else {
		query = bson.M{"status": status}
	}

	options := options.Find().SetLimit(int64(limit)).SetSort(bson.D{{Key: "created_at", Value: 1}})
	cur, err := collection.Find(ctx, query, options)
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	interviews := []*models.Interview{}
	for cur.Next(ctx) {
		interview := models.Interview{}
		if err := cur.Decode(&interview); err != nil {
			return nil, err
		}

		user, _ := s.GetUsers(interview.CreatedBy.Hex())

		interviews = append(interviews, &models.Interview{
			ID:            interview.ID,
			Description:   interview.Description,
			Status:        interview.Status,
			CreatedByUser: user,
			CreatedAt:     interview.CreatedAt,
		})
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	return interviews, nil
}

func (s *Service) UpdateInterview(id string, updateInterview *models.Interview) error {
	collection := s.dbStore.GetCollection(INTERVIEW_COLLECTION)

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(s.dbStore.ConnectionTimeout)*time.Second)
	defer cancel()

	dateTimeNow := time.Now()
	query := bson.M{"_id": convertStringtoPrimitiveID(id)}
	update := bson.M{"$set": bson.M{
		"description": updateInterview.Description,
		"status":      updateInterview.Status,
		"updated_at":  &dateTimeNow,
	}}

	if _, err := collection.UpdateOne(ctx, query, update); err != nil {
		return err
	}

	return nil
}

func (s *Service) UpdateCommentInterview(id string, comment models.Comment) error {
	collection := s.dbStore.GetCollection(INTERVIEW_COLLECTION)

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(s.dbStore.ConnectionTimeout)*time.Second)
	defer cancel()

	dateTimeNow := time.Now()
	query := bson.M{"_id": convertStringtoPrimitiveID(id)}
	update := bson.M{"$push": bson.M{"comments": models.Comment{
		Message:   comment.Message,
		CreatedBy: comment.CreatedByUser.ID,
		CreatedAt: &dateTimeNow,
	}}}

	if _, err := collection.UpdateOne(ctx, query, update); err != nil {
		return err
	}

	return nil
}

func (s *Service) GetUsers(id string) (models.Users, error) {
	collection := s.dbStore.GetCollection(USER_COLLECTION)

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(s.dbStore.ConnectionTimeout)*time.Second)
	defer cancel()

	user := models.Users{}
	query := bson.M{"_id": convertStringtoPrimitiveID(id)}
	if err := collection.FindOne(ctx, query).Decode(&user); err != nil {
		return user, err
	}

	return user, nil
}
