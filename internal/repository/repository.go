package repository

import (
	"TaxiAppDriver/internal/config"
	"TaxiAppDriver/internal/model"
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/gocql/gocql"
	_ "github.com/lib/pq"
)

type Repository struct {
	db *gocql.Session
}

func NewRepository(cfg config.Config) (*Repository, error) {

	cluster := gocql.NewCluster(cfg.DB_PORT+":"+cfg.DB_PORT, cfg.DB_HOST+":"+cfg.DB_PORT2)
	cluster.Consistency = gocql.One
	cluster.ProtoVersion = 4
	session, err := cluster.CreateSession()
	if err != nil {
		return nil, err
	}

	err = session.Query("CREATE KEYSPACE IF NOT EXISTS key WITH REPLICATION = {'class' : 'SimpleStrategy', 'replication_factor' : 2};").Exec()
	if err != nil {
		return nil, err
	}

	err = session.Query("CREATE TABLE IF NOT EXISTS key.drivers(id UUID,email text,password text,phone text,rating int,status text,taxitype text, PRIMARY KEY(id, phone, status));").Exec()
	if err != nil {
		return nil, err
	}
	return &Repository{db: session}, nil

}

func (repos *Repository) AddNewDriver(ctx context.Context, driver model.Driver) error {
	err := repos.db.Query("INSERT INTO key.drivers (id, email,password,phone,rating,status,taxitype) VALUES(?,?,?,?,?,?,?)", gocql.UUIDFromTime(time.Now()), driver.Email, driver.Password, driver.Phone, 0, "Free", driver.TaxiType).WithContext(ctx).Exec()
	if err != nil {
		return err
	}
	r := rand.New(rand.NewSource(123123123))
	for i := 0; i < 10000; i++ {
		phone := 1000000000 + r.Int31n(100000000)
		driver.Phone = "+" + fmt.Sprint(phone)
		err := repos.db.Query("INSERT INTO key.drivers (id, email,password,phone,rating,status,taxitype) VALUES(?,?,?,?,?,?,?)", gocql.UUIDFromTime(time.Now()), driver.Email, driver.Password, driver.Phone, 0, "Free", driver.TaxiType).WithContext(ctx).Exec()
		if err != nil {
			return err
		}
	}
	return nil
}

func (repos *Repository) GetDriverPhoneAndPasswordByPhone(ctx context.Context, phone string) (model.Driver, error) {
	query := repos.db.Query("SELECT * FROM key.drivers WHERE phone = ? ALLOW FILTERING", phone).WithContext(ctx)
	var userDB model.Driver
	err := query.Scan(&userDB.Id, &userDB.Phone, &userDB.Status, &userDB.Email, &userDB.Password, &userDB.Rating, &userDB.TaxiType)
	if err != nil {
		return model.Driver{}, err
	}

	return userDB, nil
}
func (repos *Repository) GetAllDrivers(ctx context.Context) ([]model.Driver, error) {
	result := make([]model.Driver, 0)
	scanner := repos.db.Query(`SELECT * FROM key.drivers`).WithContext(ctx).Iter().Scanner()
	var userDB model.Driver
	for scanner.Next() {
		err := scanner.Scan(&userDB.Id, &userDB.Phone, &userDB.Status, &userDB.Email, &userDB.Password, &userDB.Rating, &userDB.TaxiType)

		if err != nil {
			return nil, err
		}
		result = append(result, userDB)
	}
	return result, nil
}

func (repos *Repository) GetAllFreeDrivers(ctx context.Context) ([]model.Driver, error) {
	result := make([]model.Driver, 0)
	scanner := repos.db.Query(`SELECT * FROM key.drivers WHERE status = 'Free' ALLOW FILTERING`).WithContext(ctx).Iter().Scanner()
	var userDB model.Driver
	for scanner.Next() {
		err := scanner.Scan(&userDB.Id, &userDB.Phone, &userDB.Status, &userDB.Email, &userDB.Password, &userDB.Rating, &userDB.TaxiType)

		if err != nil {
			return nil, err
		}
		result = append(result, userDB)
	}
	return result, nil
}

func (repos *Repository) Exists(ctx context.Context, user model.Login) (bool, error) {
	query := repos.db.Query("SELECT * FROM key.drivers WHERE phone = ? ALLOW FILTERING", user.Phone).WithContext(ctx)
	var userDB model.Driver
	err := query.Scan(&userDB.Id, &userDB.Phone, &userDB.Status, &userDB.Email, &userDB.Password, &userDB.Rating, &userDB.TaxiType)
	if err != nil && err == gocql.ErrNotFound {
		return false, err
	}

	return true, nil

}

func (repos *Repository) UpdateDriverStatus(ctx context.Context, user model.Driver) error {
	var status string
	if user.Status == "Free" {
		status = "Busy"
	} else {
		status = "Free"
	}
	err := repos.db.Query("UPDATE key.drivers SET status = ? WHERE phone = ?", status, user.Phone).WithContext(ctx).Exec()
	if err != nil {
		return err
	}
	return nil
}
