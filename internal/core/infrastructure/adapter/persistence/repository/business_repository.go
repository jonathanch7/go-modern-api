package repository

import (
	"context"
	"github.com/jonathanch7/go-modern-api/internal/core/domain/business"
	"github.com/jonathanch7/go-modern-api/internal/core/domain/port"
	"github.com/jonathanch7/go-modern-api/internal/core/infrastructure/adapter/persistence/db"
	"github.com/jonathanch7/go-modern-api/internal/core/infrastructure/adapter/persistence/entity"
	"github.com/jonathanch7/go-modern-api/internal/core/infrastructure/adapter/persistence/mapper"
)

type BusinessRepository interface {
	port.BusinessPort
}

type businessesRepository struct {
	db *db.Database
}

func NewBusinessRepository(db *db.Database) BusinessRepository {
	if db == nil {
		panic("Database is nil")
	}
	return &businessesRepository{db}
}

func (r *businessesRepository) SaveBusinesses(ctx context.Context, businessPersist business.Business) (business.Business, error) {
	stmt, err := r.db.DB.Prepare("INSERT INTO business (business_id, name, description, public_code, reg_date) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		return business.Business{}, err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, businessPersist.BusinessId(), businessPersist.Name(), businessPersist.Description(), businessPersist.PublicCode(), businessPersist.RegDate())
	if err != nil {
		return business.Business{}, err
	}
	return businessPersist, nil
}

func (r *businessesRepository) PartialUpdateBusiness(ctx context.Context, business business.Business) (business.Business, error) {
	businessEntity := mapper.Business.ToEntity(business)
	// do
	return mapper.Business.ToDomain(businessEntity), nil
}

func (r *businessesRepository) RemoveBusiness(ctx context.Context, business business.Business) (business.Business, error) {
	businessEntity := mapper.Business.ToEntity(business)
	// do
	return mapper.Business.ToDomain(businessEntity), nil
}

func (r *businessesRepository) SearchAllBusinesses(ctx context.Context) ([]business.Business, error) {
	rows, err := r.db.DB.QueryContext(ctx, "SELECT * FROM business")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []business.Business
	for rows.Next() {
		e := entity.Business{}
		err = rows.Scan(&e.BusinessId, &e.Name, &e.Description, &e.PublicCode, &e.RegDate)
		if err != nil {
			return nil, err
		}
		list = append(list, mapper.Business.ToDomain(e))
	}
	return list, nil
}
