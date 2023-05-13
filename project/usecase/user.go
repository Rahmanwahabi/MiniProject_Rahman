package usecase

import (
	"MiniProject_Rahman/project/lib/database"
	"MiniProject_Rahman/project/middlewares"
	"MiniProject_Rahman/project/models"
	"fmt"
)

type UserUsecase interface {
	LoginUser(user *models.User) (err error)
	CreateUser(user *models.User) error
	GetUser(id uint) (user models.User, err error)
	GetListUsers() (users []models.User, err error)
	UpdateUser(user *models.User) (err error)
	DeleteUser(id uint) (err error)
}

type userUsecase struct {
	userRepository  database.UserRepository
	boardRepository database.BoardRepository
	taskRepository  database.TaskRepository
}

func NewUserUsecase(
	userRepo database.UserRepository,
	boardRepo database.BoardRepository,
	taskRepo database.TaskRepository,
) *userUsecase {
	return &userUsecase{
		userRepository:  userRepo,
		boardRepository: boardRepo,
		taskRepository:  taskRepo,
	}
}

func (u *userUsecase) LoginUser(user *models.User) (err error) {
	// check to db email and password
	err = u.userRepository.LoginUser(user)
	if err != nil {
		fmt.Println("GetUser: Error getting user from database")
		return
	}
	// generate jwt
	token, err := middlewares.CreateToken(int(user.ID))
	if err != nil {
		fmt.Println("GetUser: Error Generate token")
		return
	}
	user.Token = token
	return
}

func (u *userUsecase) CreateUser(user *models.User) error {

	err := u.userRepository.CreateUser(user)
	if err != nil {
		return err
	}

	return nil
}

func (u *userUsecase) GetUser(id uint) (user models.User, err error) {
	user.ID = id
	err = u.userRepository.GetUser(&user)
	if err != nil {
		fmt.Println("GetUser: Error getting user from database")
		return
	}
	board, err := u.boardRepository.GetBoardsByUserId(id)
	if err != nil {
		fmt.Println("GetUser: Error getting user from database")
		return
	}
	user.Boards = append(user.Boards, board...)

	task, err := u.taskRepository.GetTasksByUserId(id)
	if err != nil {
		fmt.Println("GetUser: Error getting user from database")
		return
	}
	user.Tasks = append(user.Tasks, task...)

	return
}

func (u *userUsecase) GetListUsers() (users []models.User, err error) {
	users, err = u.userRepository.GetUsers()
	if err != nil {
		fmt.Println("GetListUsers: Error getting users from database")
		return
	}
	return
}

func (u *userUsecase) UpdateUser(user *models.User) (err error) {
	err = u.userRepository.UpdateUser(user)
	if err != nil {
		fmt.Println("UpdateUser : Error updating user, err: ", err)
		return
	}

	return
}

func (u *userUsecase) DeleteUser(id uint) (err error) {
	user := models.User{}
	user.ID = id
	err = u.userRepository.DeleteUser(&user)
	if err != nil {
		fmt.Println("DeleteUser : error deleting user, err: ", err)
		return
	}

	return
}
