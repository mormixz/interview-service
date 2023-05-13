package interview

import (
	"context"
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
	options := options.FindOne().SetSort(bson.D{primitive.E{Key: "comment.created_at", Value: 1}})
	if err := collection.FindOne(ctx, query, options).Decode(&interview); err != nil {
		return nil, err
	}

	interview.CreatedByUser, _ = s.GetUsers(interview.CreatedBy.Hex())

	for index, comment := range interview.Comments {
		interview.Comments[index].CreatedByUser, _ = s.GetUsers(comment.CreatedBy.Hex())
	}

	return interview, nil
}

func (s *Service) GetInterviewAll(status string, limit int) ([]*models.Interview, error) {
	collection := s.dbStore.GetCollection(INTERVIEW_COLLECTION)

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(s.dbStore.ConnectionTimeout)*time.Second)
	defer cancel()

	query := bson.M{"status": status}
	options := options.Find().SetLimit(int64(limit)).SetSort(bson.D{primitive.E{Key: "created_at", Value: 1}})
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

	query := bson.M{"_id": convertStringtoPrimitiveID(id)}
	update := bson.M{"$set": bson.M{
		"description": updateInterview.Description,
		"status":      updateInterview.Status,
		"updated_at":  updateInterview.UpdatedAt,
	}}

	if _, err := collection.UpdateOne(ctx, query, update); err != nil {
		return err
	}

	return nil
}

func (s *Service) FlushCommentInterview(id string) error {
	collection := s.dbStore.GetCollection(INTERVIEW_COLLECTION)

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(s.dbStore.ConnectionTimeout)*time.Second)
	defer cancel()

	query := bson.M{"_id": convertStringtoPrimitiveID(id)}
	update := bson.M{"$set": bson.M{"comments": []models.Comment{}}}

	if _, err := collection.UpdateOne(ctx, query, update); err != nil {
		return err
	}

	return nil
}

func (s *Service) UpdateCommentInterview(id string, comment models.Comment) error {
	collection := s.dbStore.GetCollection(INTERVIEW_COLLECTION)

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(s.dbStore.ConnectionTimeout)*time.Second)
	defer cancel()

	comment.CreatedBy = comment.CreatedByUser.ID

	query := bson.M{"_id": convertStringtoPrimitiveID(id)}
	update := bson.M{"$push": bson.M{"comments": comment}}

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
