package repositories

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"sync"

	"go-code-challenge/internal/actions"
	"go-code-challenge/internal/users"
)

type JSONRepository struct {
	users    []users.User
	actions  []actions.Action
	initOnce sync.Once
}

const fileBase = "../../datas/"

func NewJSONRepository() *JSONRepository {
	return &JSONRepository{}
}

func (r *JSONRepository) initialize() error {
	usersFile := fmt.Sprintf("%susers.json", fileBase)
	actionsFile := fmt.Sprintf("%sactions.json", fileBase)
	var err error

	r.initOnce.Do(func() {
		userBytes, readErr := ioutil.ReadFile(usersFile)
		if readErr != nil {
			err = fmt.Errorf("error reading users.json: %v", readErr)
			return
		}

		if unmarshalErr := json.Unmarshal(userBytes, &r.users); unmarshalErr != nil {
			err = fmt.Errorf("error parsing users.json: %v", unmarshalErr)
			return
		}

		actionBytes, readErr := ioutil.ReadFile(actionsFile)
		if readErr != nil {
			err = fmt.Errorf("error reading actions.json: %v", readErr)
			return
		}

		if unmarshalErr := json.Unmarshal(actionBytes, &r.actions); unmarshalErr != nil {
			err = fmt.Errorf("error parsing actions.json: %v", unmarshalErr)
			return
		}
	})

	return err
}

func (r *JSONRepository) GetUserByID(id int) (*users.User, error) {
	if err := r.initialize(); err != nil {
		return nil, fmt.Errorf("error initializing repository: %v", err)
	}

	for _, user := range r.users {
		if user.ID == id {
			return &user, nil
		}
	}
	return nil, fmt.Errorf("user not found")
}

func (r *JSONRepository) GetActionsByUserID(userID int) ([]actions.Action, error) {
	if err := r.initialize(); err != nil {
		return nil, err
	}

	var userActions []actions.Action
	for _, action := range r.actions {
		if action.UserID == userID {
			userActions = append(userActions, action)
		}
	}
	return userActions, nil
}

func (r *JSONRepository) GetAllActions() ([]actions.Action, error) {
	if err := r.initialize(); err != nil {
		return nil, fmt.Errorf("error initializing JSON repository: %v", err)
	}
	return r.actions, nil
}
