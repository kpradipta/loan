package services

import (
	json2 "encoding/json"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/lending/internal/core/domain"
	"github.com/lending/internal/core/ports"
	"github.com/lending/internal/fsm"
	"strings"
)

type service struct {
	loanRepository  ports.LoanRepository
	messageQueueing ports.Queueing
}

func New(loanRepository ports.LoanRepository, messageQueueing ports.Queueing) *service {
	return &service{
		loanRepository: loanRepository,
		messageQueueing: messageQueueing,
	}
}

func (srv *service) Read(id string) (domain.Loan, error){
	fmt.Println("id on service : ", id)
	loan, err := srv.loanRepository.Get(id)
	if err != nil {
		return domain.Loan{}, err
	}
	return loan, nil
}

func (srv *service) Create(param domain.Loan) (string, error) {
	uuidWithHyphen := uuid.New()
	fmt.Println(uuidWithHyphen)
	id := strings.Replace(uuidWithHyphen.String(), "-", "", -1)

	_, err := srv.loanRepository.Get(id)
	errCode := strings.Split(err.Error(), ":")
	if err != nil && errCode[0] != "101" {
		return "", err
	}

	param.State = fsm.CREATED.Name()
	param.Id = id
	if err = srv.loanRepository.Save(param); err != nil {
		return "", errors.New("500: error while saving to storage")
	}


	json, err := json2.Marshal(param)

	if err != nil {
		fmt.Println(err)
		return "", errors.New("500: error while marshaling json for queueing")
	}

	go srv.messageQueueing.Publish("loan_created", json)

	return id, nil
}

func (srv *service) Update(param domain.Loan) (domain.Loan, error) {
	loan, err := srv.loanRepository.Get(param.Id)

	if err != nil {
		return domain.Loan{}, err
	}

	if loan.Id == "" {
		return domain.Loan{}, errors.New("404: Loan not found")
	}

	if loan.State != fsm.CREATED.Name() {
		return domain.Loan{}, errors.New("406: Can't update loan, Loan has been approved or there something wrong with loan data")
	}

	loan.Amount = param.Amount
	loan.Tenor = param.Tenor

	err = srv.loanRepository.Update(loan)
	if err != nil {
		return domain.Loan{}, errors.New("500: Something went wrong")
	}


	return param, nil
}


func (srv *service) Approve(id string) error {
	loan, err := srv.loanRepository.Get(id)

	if err != nil {
		return err
	}


	if loan.State != fsm.CREATED.Name() {
		return errors.New("406: Can't update loan, Loan has been approved or there something wrong with loan data")
	}

	loan.State = fsm.APPROVED.Name()
	if err := srv.loanRepository.Update(loan); err != nil {
		return errors.New("500: error while approve loan")
	}

	json, err := json2.Marshal(loan)

	if err != nil {
		fmt.Println(err)
		return errors.New("500: error while marshaling json for queueing")
	}

	go srv.messageQueueing.Publish("loan_approved", json)

	return nil
}

