package usecase

import (
	"Loan_Tracker_API/domain"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type LoanUseCase struct {
	LoanRepo domain.Loan_Repository_interface
}

// FilterLoan implements domain.Loan_Usecase_interface.
func (uc *LoanUseCase) FilterLoan(filters map[string]interface{}) ([]domain.Loan, error) {
	panic("unimplemented")
}

// GetUniqueLoan implements domain.Loan_Usecase_interface.
func (uc *LoanUseCase) GetUniqueLoan(filter map[string]interface{}, posts *[]domain.Loan) error {
	panic("unimplemented")
}

func NewLoanUseCase(repo domain.Loan_Repository_interface) *LoanUseCase {
	return &LoanUseCase{
		LoanRepo: repo,
	}
}

func (uc *LoanUseCase) CreateLoan(inputLoan domain.CreateLoan) (domain.Loan, error) {
	loan := domain.Loan{
		Amount:      inputLoan.Amount,
		Description: inputLoan.Description,
		Status:      "pending",
	}
	if loan.ID.IsZero() {
		loan.ID = primitive.NewObjectID()
	}
	createdLoan, err := uc.LoanRepo.CreateLoanDocument(loan)
	if err != nil {
		return domain.Loan{}, err
	}
	return createdLoan, nil
}

func (uc *LoanUseCase) GetLoans(limit int, pageNumber int) ([]domain.Loan, error) {
	loans, err := uc.LoanRepo.GetLoanDocuments(pageNumber, limit)
	if err != nil {
		return []domain.Loan{}, err
	}
	return loans, nil
}

func (uc *LoanUseCase) GetAllLoans(page int, limit int, status string, order string) ([]domain.Loan, error) {
	// For simplicity, status and order are not used in this example
	loans, err := uc.LoanRepo.GetLoanDocuments(page, limit)
	if err != nil {
		return []domain.Loan{}, err
	}
	return loans, nil
}

func (uc *LoanUseCase) GetOneLoan(id string) (domain.Loan, error) {
	loan, err := uc.LoanRepo.GetOneLoanDocument(id)
	if err != nil {
		return domain.Loan{}, err
	}
	return loan, nil
}

func (uc *LoanUseCase) UpdateLoan(id string, updatedLoan domain.Loan) (domain.Loan, error) {
	loan, err := uc.LoanRepo.UpdateLoanDocument(id, updatedLoan)
	if err != nil {
		return domain.Loan{}, err
	}
	return loan, nil
}

func (uc *LoanUseCase) UpdateLoanStatus(id string, status string) error {
	return uc.LoanRepo.UpdateLoanStatus(id, status)
}

func (uc *LoanUseCase) DeleteLoan(id string) error {
	return uc.LoanRepo.DeleteLoanDocument(id)
}

func (uc *LoanUseCase) FilterLoans(filters map[string]interface{}) ([]domain.Loan, error) {
	loans, err := uc.LoanRepo.FilterLoanDocument(filters)
	if err != nil {
		return []domain.Loan{}, err
	}
	return loans, nil
}

func (uc *LoanUseCase) GetSystemLogs() ([]domain.ResponseLoan, error) {
	return uc.LoanRepo.GetSystemLogs()
}
