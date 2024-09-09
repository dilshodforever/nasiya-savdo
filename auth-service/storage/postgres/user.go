package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	pb "github.com/dilshodforever/nasiya-savdo/genprotos"

	"github.com/google/uuid"
)

type UserStorage struct {
	db *sql.DB
}

func NewUserStorage(db *sql.DB) *UserStorage {
	return &UserStorage{db: db}
}

func (p *UserStorage) Register(user *pb.UserReq) (*pb.Void, error) {
	tr, err := p.db.Begin()
	if err != nil {
		return nil, err
	}
	defer tr.Commit()

	u_id := uuid.NewString()
	query := `
		INSERT INTO users (id, full_name, email, address, phone_number, username, password_hash, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
	`
	_, err = tr.ExecContext(context.Background(), query, u_id, user.FullName, user.Email, user.Address, user.PhoneNumber, user.Username, user.Password, time.Now(), time.Now())
	if err != nil {
		tr.Rollback()
		return nil, err
	}

	id := uuid.NewString()
	query = `
		INSERT INTO storage (id, name, user_id, created_at)
		VALUES ($1, $2, $3, now())
		RETURNING id
	`

	_, err = tr.Exec(query, id, user.FullName+"'s storage", u_id)
	if err != nil {
		tr.Rollback()
		return nil, err
	}

	return &pb.Void{}, nil
}

func (p *UserStorage) GetById(id *pb.ById) (*pb.User, error) {
	query := `
		SELECT id, full_name, email, address, phone_number, username, password_hash, created_at, updated_at FROM users 
		WHERE id = $1 AND deleted_at = 0
	`
	row := p.db.QueryRowContext(context.Background(), query, id.Id)

	var user pb.User

	err := row.Scan(
		&user.Id,
		&user.FullName,
		&user.Email,
		&user.Address,
		&user.PhoneNumber,
		&user.Username,
		&user.PasswordHash,
		&user.CreatedAt,
		&user.UpdatedAt,
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
	query := `SELECT id, full_name, email, address, phone_number, username FROM users WHERE deleted_at = 0`

	var params []interface{}
	count := 1

	if filter.FullName != "" {
		query += fmt.Sprintf(" AND full_name ILIKE $%d", count)
		params = append(params, "%"+filter.FullName+"%")
		count++
	}
	if filter.Email != "" {
		query += fmt.Sprintf(" AND email ILIKE $%d", count)
		params = append(params, "%"+filter.Email+"%")
		count++
	}
	if filter.Address != "" {
		query += fmt.Sprintf(" AND address ILIKE $%d", count)
		params = append(params, "%"+filter.Address+"%")
		count++
	}

	// Apply limit and offset
	fmt.Println(filter.Limit, filter.Offset)
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
			&user.FullName,
			&user.Email,
			&user.Address,
			&user.PhoneNumber,
			&user.Username,
		)
		if err != nil {
			return nil, err
		}
		users.Users = append(users.Users, &user)
	}

	query = "SELECT COUNT(1) FROM users WHERE deleted_at = 0"
	err = p.db.QueryRowContext(context.Background(), query).Scan(&count)
	if err != nil {
		return nil, err
	}
	users.Count = int32(count)

	return users, nil
}

func (p *UserStorage) Update(user *pb.User) (*pb.Void, error) {
	query := ` 
		UPDATE users
		SET full_name = $2, email = $3, address = $4, phone_number = $5, username = $6, updated_at = $7
		WHERE id = $1 AND deleted_at = 0
	`
	_, err := p.db.ExecContext(context.Background(), query, user.Id, user.FullName, user.Email, user.Address, user.PhoneNumber, user.Username, time.Now())
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

func (p *UserStorage) Login(login *pb.UserLogin) (*pb.UserLoginRes, error) {
	query := `
		SELECT u.id, u.full_name, u.email, u.address, u.phone_number, u.username, s.id
		FROM users u JOIN storage s ON u.id = s.user_id 
		WHERE username = $1 AND password_hash = $2 AND deleted_at = 0
	`
	row := p.db.QueryRowContext(context.Background(), query, login.Username, login.Password)

	var user pb.UserLoginRes

	err := row.Scan(
		&user.Id,
		&user.FullName,
		&user.Email,
		&user.Address,
		&user.PhoneNumber,
		&user.Username,
		&user.StorageId,
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
		SET password_hash = $3, updated_at = $4
		WHERE id = $1 AND password_hash = $2 AND deleted_at = 0
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
		return nil, fmt.Errorf("invalid current password")
	}

	return &pb.Void{}, nil
}

func (p *UserStorage) ForgotPassword(forgotPass *pb.ForgotPass) (*pb.Void, error) {
	// Implementation here
	return &pb.Void{}, nil
}

func (p *UserStorage) ResetPassword(resetPass *pb.ResetPass) (*pb.Void, error) {
	query := `
		UPDATE users
		SET password_hash = $2, updated_at = $3
		WHERE id = $1 AND deleted_at = 0
	`

	_, err := p.db.ExecContext(context.Background(), query, resetPass.Id, resetPass.NewPassword, time.Now())
	if err != nil {
		return nil, err
	}

	return &pb.Void{}, nil
}

// func (p *UserStorage) SaveToken(token *pb.Token) (*pb.Void, error) {
// 	id := uuid.NewString()
// 	query := `
// 		INSERT INTO user_token (id, access, refresh, user_id)
// 		VALUES ($1, $2, $3, $4)
// 	`

// 	_, err := p.db.ExecContext(context.Background(), query, id, token.Access, token.Refresh, token.UserId)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &pb.Void{}, nil
// }

// func (p *UserStorage) GetToken(id *pb.ById) (*pb.Token, error) {
// 	query := `
// 		SELECT id, access, refresh, user_id FROM user_token WHERE user_id = $1
// 	`

// 	var token pb.Token
// 	err := p.db.QueryRowContext(context.Background(), query, id.Id).Scan(&token.Id, &token.Access, &token.Refresh, &token.UserId)
// 	if err == sql.ErrNoRows {
// 		return nil, fmt.Errorf("token not found")
// 	}
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &token, nil
// }

// func (p *UserStorage) UpdateToken(token *pb.Token) (*pb.Void, error) {
// 	query := `
// 		UPDATE user_token
// 		SET access = $2, refresh = $3
// 		WHERE user_id = $1
// 	`

// 	_, err := p.db.ExecContext(context.Background(), query, token.UserId, token.Access, token.Refresh)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &pb.Void{}, nil
// }
