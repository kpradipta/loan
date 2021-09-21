package ports

import "github.com/lending/internal/core/domain"

type LoanRepository interface {
	Get(id string) (domain.Loan, error)
	Save(loan domain.Loan) error
	Update(loan domain.Loan) error
}