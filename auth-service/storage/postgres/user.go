package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	pb "gitlab.com/lingualeap/auth/genprotos/users"

	"github.com/google/uuid"
)

type UserStorage struct {
	db *sql.DB
}

func NewUserStorage(db *sql.DB) *UserStorage {
	return &UserStorage{db: db}
}

func (p *UserStorage) Register(user *pb.UserReq) (*pb.Void, error) {
	id := uuid.NewString()
	query := `
		INSERT INTO users (id, user_name, email, password, full_name, native_language, created_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`
	_, err := p.db.ExecContext(context.Background(), query, id, user.UserName, user.Email, user.Password, user.FullName, user.NativeLanguage, time.Now())
	if err != nil {
		return nil, err
	}
	return &pb.Void{}, nil
}

func (p *UserStorage) GetById(id *pb.ById) (*pb.User, error) {
	query := `
		SELECT id, user_name, email, full_name, native_language FROM users 
		WHERE id = $1 AND deleted_at = 0
	`
	row := p.db.QueryRowContext(context.Background(), query, id.Id)

	var user pb.User

	err := row.Scan(
		&user.Id,
		&user.UserName,
		&user.Email,
		&user.FullName,
		&user.NativeLanguage,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user not found")
		}
		return nil, err
	}

	return &user, nil
}

func (p *UserStorage) GetAll(filter *pb.UserFilter) (*pb.AllUsers, error) {
	users := &pb.AllUsers{}
	query := `SELECT id, user_name, email, full_name, native_language FROM users WHERE deleted_at = 0`

	var params []interface{}
	count := 1

	if filter.UserName != "" {
		query += fmt.Sprintf(" AND user_name ILIKE $%d", count)
		params = append(params, "%"+filter.UserName+"%")
		count++
	}
	if filter.Email != "" {
		query += fmt.Sprintf(" AND email ILIKE $%d", count)
		params = append(params, "%"+filter.Email+"%")
		count++
	}

	query += fmt.Sprintf(" LIMIT $%d OFFSET $%d", count, count+1)
	params = append(params, filter.Limit, filter.Offset)

	rows, err := p.db.QueryContext(context.Background(), query, params...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user pb.User
		err := rows.Scan(
			&user.Id,
			&user.UserName,
			&user.Email,
			&user.FullName,
			&user.NativeLanguage,
		)
		if err != nil {
			return nil, err
		}
		users.Users = append(users.Users, &user)
	}

	return users, nil
}

func (p *UserStorage) Update(user *pb.User) (*pb.Void, error) {
	query := `
		UPDATE users
		SET user_name = $2, email = $3, full_name = $4, native_language = $5, updated_at = $6
		WHERE id = $1 AND deleted_at = 0
	`
	_, err := p.db.ExecContext(context.Background(), query, user.Id, user.UserName, user.Email, user.FullName, user.NativeLanguage, time.Now())
	if err != nil {
		return nil, err
	}
	return &pb.Void{}, nil
}

func (p *UserStorage) Delete(id *pb.ById) (*pb.Void, error) {
	query := `
		UPDATE users
		SET deleted_at = $2
		WHERE id = $1 AND deleted_at = 0
	`
	_, err := p.db.ExecContext(context.Background(), query, id.Id, time.Now().Unix())
	if err != nil {
		return nil, err
	}
	return &pb.Void{}, nil
}

func (p *UserStorage) Login(login *pb.UserLogin) (*pb.User, error) {
	query := `
		SELECT id, user_name, email, full_name, native_language, role FROM users 
		WHERE user_name = $1 AND password = $2 AND deleted_at = 0
	`
	row := p.db.QueryRowContext(context.Background(), query, login.UserName, login.Password)

	var user pb.User

	err := row.Scan(
		&user.Id,
		&user.UserName,
		&user.Email,
		&user.FullName,
		&user.NativeLanguage,
		&user.Role,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("invalid username or password")
		}
		return nil, err
	}

	return &user, nil
}

func (p *UserStorage) ChangePassword(changePass *pb.ChangePass) (*pb.Void, error) {
	query := `
		UPDATE users
		SET password = $3, updated_at = $4
		WHERE id = $1 AND password = $2 AND deleted_at = 0
	`
	r, err := p.db.ExecContext(context.Background(), query, changePass.Id, changePass.CurrentPassword, changePass.NewPassword, time.Now())
	if err != nil {
		return nil, err
	}
	l, err := r.RowsAffected()
	if err != nil {
		return nil, err
	}
	if l == 0 {
		return nil, fmt.Errorf("error while db: invalid current password")
	}
	return &pb.Void{}, nil
}

func (p *UserStorage) ForgotPassword(forgotPass *pb.ForgotPass) (*pb.Void, error) {
	return &pb.Void{}, nil
}

func (p *UserStorage) ResetPassword(resetPass *pb.ResetPass) (*pb.Void, error) {
	query := `
		UPDATE users
		SET password = $2, updated_at = $3
		WHERE id = $1 AND deleted_at = 0
	`
	_, err := p.db.ExecContext(context.Background(), query, resetPass.Id, resetPass.NewPassword, time.Now())
	if err != nil {
		return nil, err
	}
	return &pb.Void{}, nil
}
