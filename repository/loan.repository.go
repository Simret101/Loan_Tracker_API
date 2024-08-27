package repository

import (
	"context"
	"errors"

	"Loan_Tracker_API/database"
	"Loan_Tracker_API/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type LoanRepository struct {
	collection database.CollectionInterface
}

// FilterLoanDocument implements domain.Loan_Repository_interface.
func (lr *LoanRepository) FilterLoanDocument(filters map[string]interface{}) ([]domain.Loan, error) {
	panic("unimplemented")
}

func NewLoanRepository(collection database.CollectionInterface) *LoanRepository {
	return &LoanRepository{
		collection: collection,
	}
}

func (lr *LoanRepository) CreateLoanDocument(loan domain.Loan) (domain.Loan, error) {
	_, err := lr.collection.InsertOne(context.TODO(), loan)
	return loan, err
}

func (lr *LoanRepository) GetOneLoanDocument(id string) (domain.Loan, error) {
	obId, _ := primitive.ObjectIDFromHex(id)
	var loan domain.Loan
	query := bson.M{"_id": obId}

	err := lr.collection.FindOne(context.TODO(), query).Decode(&loan)
	if err != nil {
		return domain.Loan{}, err
	}

	return loan, nil
}

func (lr *LoanRepository) GetLoanDocuments(page, limit int) ([]domain.Loan, error) {
	var loans []domain.Loan

	options := options.Find()
	options.SetSkip(int64((page - 1) * limit))
	options.SetLimit(int64(limit))

	cursor, err := lr.collection.Find(context.TODO(), bson.M{}, options)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var loan domain.Loan
		if err := cursor.Decode(&loan); err != nil {
			return nil, err
		}
		loans = append(loans, loan)
	}

	return loans, nil
}

func (lr *LoanRepository) UpdateLoanDocument(id string, loan domain.Loan) (domain.Loan, error) {
	obId, _ := primitive.ObjectIDFromHex(id)
	_, err := lr.collection.UpdateOne(context.TODO(), bson.M{"_id": obId}, bson.M{"$set": loan})
	return loan, err
}

func (lr *LoanRepository) DeleteLoanDocument(id string) error {
	obID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	query := bson.M{"_id": obID}

	res, err := lr.collection.DeleteOne(context.TODO(), query)
	if err != nil {
		return err
	}

	if res.DeletedCount() == 0 {
		return errors.New("no loan with this ID found")
	}

	return nil
}

func (lr *LoanRepository) FilterLoanDocuments(filter map[string]interface{}) ([]domain.Loan, error) {
	var loans []domain.Loan
	query := bson.M{}

	for key, value := range filter {
		switch v := value.(type) {
		case string:
			query[key] = bson.M{"$regex": v, "$options": "i"}
		case []string:
			query[key] = bson.M{"$in": v}
		}
	}

	cursor, err := lr.collection.Find(context.TODO(), query)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var loan domain.Loan
		if err := cursor.Decode(&loan); err != nil {
			return nil, err
		}
		loans = append(loans, loan)
	}

	return loans, nil
}

func (lr *LoanRepository) UpdateLoanStatus(id string, status string) error {
	obId, _ := primitive.ObjectIDFromHex(id)
	_, err := lr.collection.UpdateOne(context.TODO(), bson.M{"_id": obId}, bson.M{"$set": bson.M{"status": status}})
	return err
}

func (lr *LoanRepository) GetSystemLogs() ([]domain.ResponseLoan, error) {
	// Placeholder for actual log retrieval logic
	return []domain.ResponseLoan{}, nil
}
