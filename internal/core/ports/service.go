package ports

import "github.com/lending/internal/core/domain"

type LoanService interface {
	Create(loan domain.Loan) (string, error)
	Update(loan domain.Loan) (domain.Loan, error)
	Read(id string) (domain.Loan, error)
	Approve(id string) error
}
