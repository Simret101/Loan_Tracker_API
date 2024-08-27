package usecase

import (
	"Loan_Tracker_API/domain"
	//passwordservice "Loan_Tracker_API/infrastructure/password_service"
)

type UserUseCase struct {
	UserRepo domain.User_Repository_interface
}

// UpdatePassword implements domain.User_Usecase_interface.
func (usecase *UserUseCase) UpdatePassword(id string, updated_user domain.UpdatePassword) (domain.ResponseUser, error) {
	panic("unimplemented")
}

func NewUserUseCase(repo domain.User_Repository_interface) *UserUseCase {
	return &UserUseCase{UserRepo: repo}
}

func (usecase *UserUseCase) GetOneUser(id string) (domain.ResponseUser, error) {
	user, err := usecase.UserRepo.GetUserDocumentByID(id)
	if err != nil {
		return domain.ResponseUser{}, err
	}
	response_user := domain.CreateResponseUser(user)
	return response_user, nil
}

func (usecase *UserUseCase) GetUsers() ([]domain.ResponseUser, error) {
	users, err := usecase.UserRepo.GetUserDocuments()
	if err != nil {
		return []domain.ResponseUser{}, err
	}
	responses_users := []domain.ResponseUser{}

	for _, user := range users {
		responses_users = append(responses_users, domain.CreateResponseUser(user))
	}
	return responses_users, nil
}

func (usecase *UserUseCase) UpdateUser(id string, user domain.UpdateUser) (domain.ResponseUser, error) {
	new_user, err := usecase.UserRepo.UpdateUserDocument(id, user)
	if err != nil {
		return domain.ResponseUser{}, err
	}
	return domain.CreateResponseUser(new_user), nil
}

func (usecase *UserUseCase) DeleteUser(id string) error {
	return usecase.UserRepo.DeleteUserDocument(id)
}

func (usecase *UserUseCase) FilterUser(filter map[string]string) ([]domain.ResponseUser, error) {
	users, err := usecase.UserRepo.FilterUserDocument(filter)
	if err != nil {
		return []domain.ResponseUser{}, err
	}

	response_users := []domain.ResponseUser{}

	for _, user := range users {
		response_users = append(response_users, domain.CreateResponseUser(user))
	}

	return response_users, nil
}
