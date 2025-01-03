package postgres

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"

	pb "github.com/dilshodforever/nasiya-savdo/genprotos"
	"github.com/google/uuid"
)

type ProductStorage struct {
	db *sql.DB
}

func NewProductStorage(db *sql.DB) *ProductStorage {
	return &ProductStorage{db: db}
}

func (p *ProductStorage) CreateProduct(req *pb.CreateProductRequest) (*pb.ProductResponse, error) {
	id := uuid.NewString()
	query := `
		INSERT INTO products (id, name, color, model, image_url, made_in, date_of_creation, storage_id, created_at, updated_at, deleted_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8,now(), now(), 0)
		
	`

	_, err := p.db.Exec(query, id, req.Name, req.Color, req.Model, req.ImageUrl, req.MadeIn, req.DateOfCreation, req.StorageId)
	if err != nil {
		return nil, err
	}

	return &pb.ProductResponse{
		Message: "Product created successfully",
		Success: true,
	}, nil
}

func (p *ProductStorage) GetProduct(req *pb.ProductIdRequest) (*pb.GetProductResponse, error) {
	var product pb.GetProductResponse
	query := `
		SELECT id, name, color, model, image_url, made_in, date_of_creation, storage_id, created_at, updated_at, deleted_at
		FROM products
		WHERE id = $1 AND deleted_at = 0
	`
	err := p.db.QueryRow(query, req.Id).Scan(
		&product.Id, &product.Name, &product.Color, &product.Model,
		&product.ImageUrl, &product.MadeIn, &product.DateOfCreation,
		&product.StorageId, &product.CreatedAt, &product.UpdatedAt, &product.DeletedAt,
	)
	if err != nil {
		return nil, err
	}

	priceQuery := `
		WITH buy_orders AS (
			SELECT 
				id, 
				product_id, 
				amount, 
				price, 
				created_at
			FROM 
				exchange
			WHERE 
				status = 'buy' 
				AND product_id = $1
				AND deleted_at = 0
			ORDER BY 
				created_at ASC
		),
		sell_total AS (
			SELECT 
				COALESCE(SUM(amount), 0) AS total_sold
			FROM 
				exchange
			WHERE 
				status = 'sell' 
				AND product_id = $1
				AND deleted_at = 0
		),
		remaining_buy_orders AS (
			SELECT 
				id, 
				product_id, 
				amount, 
				price, 
				created_at,
				SUM(amount) OVER (ORDER BY created_at ASC) - (SELECT total_sold FROM sell_total) AS remaining
			FROM 
				buy_orders
		)
		SELECT 
			price
		FROM 
			remaining_buy_orders
		WHERE 
			remaining > 0
		ORDER BY 
			created_at
		LIMIT 1;
	`
	
	var currentPrice sql.NullFloat64
	var totalAmount sql.NullInt32

	err = p.db.QueryRow(priceQuery, req.Id).Scan(&currentPrice)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	
	query = `
		SELECT SUM(
			CASE 
				WHEN status = 'buy' THEN amount 
				WHEN status = 'sell' THEN -amount 
				ELSE 0 
			END
		) 
		FROM exchange 
		WHERE deleted_at = 0 AND product_id = $1
	`
	err = p.db.QueryRow(query, req.Id).Scan(&totalAmount)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	
	if currentPrice.Valid {
		product.Price = currentPrice.Float64
	} else {
		product.Price = 0.0 // agar narx topilmasa
	}
	if totalAmount.Valid {
		product.Amount = totalAmount.Int32
	} else {
		product.Amount = 0 // agar mahsulot miqdori topilmasa
	}

	return &product, nil
}
func (p *ProductStorage) UpdateProduct(req *pb.UpdateProductRequest) (*pb.ProductResponse, error) {
	// Initialize a map to hold the fields that need to be updated
	update := map[string]interface{}{}
	if req.Name != "" {
		update["name"] = req.Name
	}
	if req.Color != "" {
		update["color"] = req.Color
	}
	if req.Model != "" {
		update["model"] = req.Model
	}
	if req.ImageUrl != "" {
		update["image_url"] = req.ImageUrl
	}
	if req.MadeIn != "" {
		update["made_in"] = req.MadeIn
	}
	if req.DateOfCreation != "" {
		update["date_of_creation"] = req.DateOfCreation
	}
	if req.StorageId != "" {
		update["storage_id"] = req.StorageId
	}

	// Check if there's anything to update
	if len(update) == 0 {
		return &pb.ProductResponse{Message: "Nothing to update", Success: false}, nil
	}

	setClauses := []string{}
	args := []interface{}{}
	argCount := 1

	for column, value := range update {
		setClauses = append(setClauses, fmt.Sprintf("%s = $%d", column, argCount))
		args = append(args, value)
		argCount++
	}

	setQuery := strings.Join(setClauses, ", ")

	query := fmt.Sprintf(`
		UPDATE products
		SET %s, updated_at = now()
		WHERE id = $%d AND deleted_at = 0
	`, setQuery, argCount)

	// Append the id as the last argument
	args = append(args, req.Id)

	// Execute the query
	_, err := p.db.Exec(query, args...)
	if err != nil {
		return nil, err
	}

	return &pb.ProductResponse{
		Message: "Product updated successfully",
		Success: true,
	}, nil
}

func (p *ProductStorage) DeleteProduct(req *pb.ProductIdRequest) (*pb.ProductResponse, error) {
	query := `
		UPDATE products
		SET deleted_at = $2
		WHERE id = $1 AND deleted_at = 0
	`
	_, err := p.db.Exec(query, req.Id, time.Now().Unix())
	if err != nil {
		return nil, err
	}

	return &pb.ProductResponse{
		Message: "Product deleted successfully",
		Success: true,
	}, nil
}


func (p *ProductStorage) ListProducts(req *pb.GetAllProductRequest) (*pb.GetAllProductResponse, error) {
	products := pb.GetAllProductResponse{}

	query := `
		SELECT id, name, color, model, image_url, made_in, date_of_creation, storage_id, created_at, updated_at, deleted_at
		FROM products
		WHERE deleted_at = 0
	`

	var args []interface{}
	argCounter := 1

	if req.Name != "" {
		query += fmt.Sprintf(" AND name ILIKE $%d", argCounter)
		args = append(args, "%"+req.Name+"%")
		argCounter++
	}
	if req.Color != "" {
		query += fmt.Sprintf(" AND color ILIKE $%d", argCounter)
		args = append(args, "%"+req.Color+"%")
		argCounter++
	}
	if req.Model != "" {
		query += fmt.Sprintf(" AND model ILIKE $%d", argCounter)
		args = append(args, "%"+req.Model+"%")
		argCounter++
	}

	// Pagination logic based on Page and Limit
	if req.Limit != 0 {
		page := req.Page
		if page < 1 {
			page = 1 // Default to first page if Page is invalid
		}
		offset := (page - 1) * req.Limit
		query += fmt.Sprintf(" LIMIT $%d OFFSET $%d", argCounter, argCounter+1)
		args = append(args, req.Limit, offset)
	}

	rows, err := p.db.Query(query, args...)
	if err != nil {
		return nil, err 
	}
	defer rows.Close()

	for rows.Next() {
		var product pb.GetProductResponse
		err := rows.Scan(
			&product.Id, &product.Name, &product.Color, &product.Model,
			&product.ImageUrl, &product.MadeIn, &product.DateOfCreation,
			&product.StorageId, &product.CreatedAt, &product.UpdatedAt, &product.DeletedAt,
		)
		if err != nil {
			log.Println("Failed to scan product:", err)
			return nil, err
		}
		
		priceQuery := `
			WITH buy_orders AS (
				SELECT 
					id, 
					product_id, 
					amount, 
					price, 
					created_at
				FROM 
					exchange
				WHERE 
					status = 'buy' 
					AND product_id = $1
					AND deleted_at = 0
				ORDER BY 
					created_at ASC
			),
			sell_total AS (
				SELECT 
					COALESCE(SUM(amount), 0) AS total_sold
				FROM 
					exchange
				WHERE 
					status = 'sell' 
					AND product_id = $1
					AND deleted_at = 0
			),
			remaining_buy_orders AS (
				SELECT 
					id, 
					product_id, 
					amount, 
					price, 
					created_at,
					SUM(amount) OVER (ORDER BY created_at ASC) - (SELECT total_sold FROM sell_total) AS remaining
				FROM 
					buy_orders
			)
			SELECT 
				price
			FROM 
				remaining_buy_orders
			WHERE 
				remaining > 0
			ORDER BY 
				created_at
			LIMIT 1;
		`
		
		var currentPrice sql.NullFloat64
		var totalAmount sql.NullInt32

		err = p.db.QueryRow(priceQuery, product.Id).Scan(&currentPrice)
		if err != nil && err != sql.ErrNoRows {
			return nil, err
		}
		
		query = `
			SELECT SUM(
				CASE 
					WHEN status = 'buy' THEN amount 
					WHEN status = 'sell' THEN -amount 
					ELSE 0 
				END
			) 
			FROM exchange 
			WHERE deleted_at = 0 AND product_id = $1
		`
		err = p.db.QueryRow(query, product.Id).Scan(&totalAmount)
		if err != nil && err != sql.ErrNoRows {
			return nil, err
		}
		
		if currentPrice.Valid {
			product.Price = currentPrice.Float64
		} else {
			product.Price = 0.0 // Default if no price is found
		}
		if totalAmount.Valid {
			product.Amount = totalAmount.Int32
		} else {
			product.Amount = 0 // Default if no amount is found
		}

		products.AllProducts = append(products.AllProducts, &product)
	}

	query = `SELECT COUNT(1) FROM products WHERE deleted_at = 0`
	err = p.db.QueryRow(query).Scan(&argCounter)
	if err != nil {
		return nil, err
	}
	products.Count = int32(argCounter)

	if err := rows.Err(); err != nil {
		return nil, err
	}

	products.Message = "Products retrieved successfully"
	return &products, nil
}




// func (p *ProductStorage) ListProducts(req *pb.GetAllProductRequest) (*pb.GetAllProductResponse, error) {
// 	products := pb.GetAllProductResponse{}

// 	// query`
// 	// 	SELECT id, name, color, model, image_url, made_in, date_of_creation, storage_id, created_at, updated_at, deleted_at
// 	// 	FROM products
// 	// 	WHERE deleted_at = 0 and storage_id = $1
// 	// `
// 	query := `
// 		SELECT id, name, color, model, image_url, made_in, date_of_creation, storage_id, created_at, updated_at, deleted_at
// 		FROM products
// 		WHERE deleted_at = 0
// 	`

// 	var args []interface{}
// 	argCounter := 1
// 	// args = append(args, req.StorageId)

// 	if req.Name != "" {
// 		query += fmt.Sprintf(" AND name ILIKE $%d", argCounter)
// 		args = append(args, "%"+req.Name+"%")
// 		argCounter++
// 	}
// 	if req.Color != "" {
// 		query += fmt.Sprintf(" AND color ILIKE $%d", argCounter)
// 		args = append(args, "%"+req.Color+"%")
// 		argCounter++
// 	}
// 	if req.Model != "" {
// 		query += fmt.Sprintf(" AND model ILIKE $%d", argCounter)
// 		args = append(args, "%"+req.Model+"%")
// 		argCounter++
// 	}

// 	if req.Limit != 0 || req.Offset != 0 {
// 		query += fmt.Sprintf(" LIMIT $%d OFFSET $%d", argCounter, argCounter+1)
// 		args = append(args, req.Limit, req.Offset)
// 	}

// 	// if req.StorageId != "" {
// 	// 	query += fmt.Sprintf(" AND storage_id = $%d", argCounter)
// 	// 	args = append(args, req.StorageId)
// 	// 	argCounter++
// 	// }

// 	rows, err := p.db.Query(query, args...)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	for rows.Next() {
// 		var product pb.GetProductResponse
// 		err := rows.Scan(
// 			&product.Id, &product.Name, &product.Color, &product.Model,
// 			&product.ImageUrl, &product.MadeIn, &product.DateOfCreation,
// 			&product.StorageId, &product.CreatedAt, &product.UpdatedAt, &product.DeletedAt,
// 		)
// 		if err != nil {
// 			log.Println("Failed to scan product:", err)
// 			return nil, err
// 		}
		
// 		priceQuery := `
// 			WITH buy_orders AS (
// 				SELECT 
// 					id, 
// 					product_id, 
// 					amount, 
// 					price, 
// 					created_at
// 				FROM 
// 					exchange
// 				WHERE 
// 					status = 'buy' 
// 					AND product_id = $1
// 					AND deleted_at = 0
// 				ORDER BY 
// 					created_at ASC
// 			),
// 			sell_total AS (
// 				SELECT 
// 					COALESCE(SUM(amount), 0) AS total_sold
// 				FROM 
// 					exchange
// 				WHERE 
// 					status = 'sell' 
// 					AND product_id = $1
// 					AND deleted_at = 0
// 			),
// 			remaining_buy_orders AS (
// 				SELECT 
// 					id, 
// 					product_id, 
// 					amount, 
// 					price, 
// 					created_at,
// 					SUM(amount) OVER (ORDER BY created_at ASC) - (SELECT total_sold FROM sell_total) AS remaining
// 				FROM 
// 					buy_orders
// 			)
// 			SELECT 
// 				price
// 			FROM 
// 				remaining_buy_orders
// 			WHERE 
// 				remaining > 0
// 			ORDER BY 
// 				created_at
// 			LIMIT 1;
// 		`
		
// 		var currentPrice sql.NullFloat64
// 		var totalAmount sql.NullInt32

// 		err = p.db.QueryRow(priceQuery, product.Id).Scan(&currentPrice)
// 		if err != nil && err != sql.ErrNoRows {
// 			return nil, err
// 		}
		
// 		query = `
// 			SELECT SUM(
// 				CASE 
// 					WHEN status = 'buy' THEN amount 
// 					WHEN status = 'sell' THEN -amount 
// 					ELSE 0 
// 				END
// 			) 
// 			FROM exchange 
// 			WHERE deleted_at = 0 AND product_id = $1
// 		`
// 		err = p.db.QueryRow(query, product.Id).Scan(&totalAmount)
// 		if err != nil && err != sql.ErrNoRows {
// 			return nil, err
// 		}
		
// 		if currentPrice.Valid {
// 			product.Price = currentPrice.Float64
// 		} else {
// 			product.Price = 0.0 // agar narx topilmasa
// 		}
// 		if totalAmount.Valid {
// 			product.Amount = totalAmount.Int32
// 		} else {
// 			product.Amount = 0 // agar mahsulot miqdori topilmasa
// 		}

// 		products.AllProducts = append(products.AllProducts, &product)
// 	}

// 	query = `SELECT COUNT(1) FROM products`
// 	err = p.db.QueryRow(query).Scan(&argCounter)
// 	if err != nil {
// 		return nil, err
// 	}
// 	products.Count = int32(argCounter)

// 	if err := rows.Err(); err != nil {
// 		return nil, err
// 	}

// 	products.Message = "Products retrieved successfully"
// 	return &products, nil
// }
