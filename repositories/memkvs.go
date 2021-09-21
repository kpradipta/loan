package repositories

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/lending/internal/core/domain"
)

type MemKvs struct {
	kvs map[string]string
}

func NewMemKVS() *MemKvs {
	return &MemKvs{kvs: map[string]string{}}
}

func (repo *MemKvs) Get(id string) (domain.Loan, error) {

	if value, ok := repo.kvs[id]; ok {
		loan := domain.Loan{}
		err := json.Unmarshal([]byte(value), &loan)
		if err != nil {
			return domain.Loan{}, errors.New("102: fail to get value from kvs")
		}

		return loan, nil
	}

	return domain.Loan{}, errors.New("101: loan not found in kvs")
}

func (repo *MemKvs) Save(loan domain.Loan) error {
	jsonData, err := json.Marshal(loan)
	fmt.Println(string(jsonData))
	if err != nil {
		return errors.New("500: error save to kvs")
	}
	repo.kvs[loan.Id] = string(jsonData)
	fmt.Println(repo.kvs[loan.Id])
	loans := domain.Loan{}
	err = json.Unmarshal(jsonData, &loan)
	if err != nil {
		
	}
	fmt.Println(loans)
	return nil
}


func (repo *MemKvs) Update(loan domain.Loan) error {
	if _, ok := repo.kvs[loan.Id]; !ok {
		return errors.New("404: Data not found")
	}
	jsonData, err := json.Marshal(loan)
	if err != nil {
		return errors.New("500: error while marshaling json for queueing")
	}
	repo.kvs[loan.Id] = string(jsonData)

	return nil
}
