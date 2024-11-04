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

type ExchangeStorage struct {
	db *sql.DB
}

func NewExchangeStorage(db *sql.DB) *ExchangeStorage {
	return &ExchangeStorage{db: db}
}

func (p *ExchangeStorage) CreateExchange(req *pb.CreateExchangeRequest) (*pb.ExchangeResponse, error) {
	id := uuid.NewString()
	if req.Status == `sell` {
		query := `
		SELECT SUM(
			CASE 
				WHEN status = 'buy' THEN amount 
				WHEN status = 'sell' THEN -amount 
				ELSE 0 
			END
		) 
		FROM exchange 
		WHERE deleted_at = 0 AND product_id = $1`
		var amount int
		err := p.db.QueryRow(query, req.ProductId).Scan(&amount)
		if err != nil {
			return nil, err
		}
		if amount < int(req.Amount) {
			return nil, fmt.Errorf("insufficient product quantity available")
		}
		query = `
			INSERT INTO exchange (id, product_id, amount, price, status, contract_id, created_at, deleted_at)
			VALUES ($1, $2, $3, $4, $5, $6, now(), 0)
		`
		_, err = p.db.Exec(query, id, req.ProductId, req.Amount, req.Price, req.Status, req.ContractId)
		if err != nil {
			return nil, err
		}
	} else {
		query := `
			INSERT INTO exchange (id, product_id, amount, price, status,  created_at, deleted_at)
			VALUES ($1, $2, $3, $4, $5, now(), 0)
		`
		_, err := p.db.Exec(query, id, req.ProductId, req.Amount, req.Price, req.Status)
		if err != nil {
			return nil, err
		}
	}
	return &pb.ExchangeResponse{
		Message: "Exchange created successfully ",
		Success: true,
	}, nil
}

func (p *ExchangeStorage) GetExchange(req *pb.ExchangeIdRequest) (*pb.GetExchangeResponse, error) {
	query := `
		SELECT id, product_id, amount, price, status, contract_id, created_at, updated_at, deleted_at
		FROM exchange
		WHERE id = $1 AND deleted_at = 0
	`
	var exchange pb.GetExchangeResponse
	var contractId sql.NullString
	err := p.db.QueryRow(query, req.Id).Scan(
		&exchange.Id, &exchange.ProductId, &exchange.Amount, &exchange.Price, &exchange.Status,
		&contractId, &exchange.CreatedAt, &exchange.UpdatedAt, &exchange.DeletedAt,
	)
	// Handle the nullable field
	if contractId.Valid {
		exchange.ContractId = contractId.String
	} else {
		exchange.ContractId = "" // Or set it to a default value
	}
	if err != nil {
		return nil, err
	}

	return &exchange, nil
}

func (p *ExchangeStorage) UpdateExchange(req *pb.UpdateExchangeRequest) (*pb.ExchangeResponse, error) {

	update := map[string]interface{}{}
	if req.ProductId != "" {
		update["product_id"] = req.ProductId
	}
	if req.Amount != 0 {
		update["amount"] = req.Amount
	}
	if req.Price != 0 {
		update["price"] = req.Price
	}

	if req.ContractId != "" {
		update["contract_id"] = req.ContractId
	}

	// Check if there's anything to update
	if len(update) == 0 {
		return &pb.ExchangeResponse{Message: "Nothing to update", Success: false}, nil
	}

	// Build the dynamic SQL query based on the fields to be updated
	setClauses := []string{}
	args := []interface{}{}
	argCount := 1

	for column, value := range update {
		setClauses = append(setClauses, fmt.Sprintf("%s = $%d", column, argCount))
		args = append(args, value)
		argCount++
	}

	// Join the SET clauses
	setQuery := strings.Join(setClauses, ", ")

	// Final update query
	query := fmt.Sprintf(`
		UPDATE exchange
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

	return &pb.ExchangeResponse{
		Message: "Exchange successfully updated",
		Success: true,
	}, nil
}

func (p *ExchangeStorage) DeleteExchange(req *pb.ExchangeIdRequest) (*pb.ExchangeResponse, error) {
	query := `
		UPDATE exchange
		SET deleted_at = $2
		WHERE id = $1 AND deleted_at = 0
	`
	_, err := p.db.Exec(query, req.Id, time.Now().Unix())
	if err != nil {
		return nil, err
	}

	return &pb.ExchangeResponse{
		Message: "Exchange successfully deleted",
		Success: true,
	}, nil
}


func (p *ExchangeStorage) ListExchanges(req *pb.GetAllExchangeRequest) (*pb.GetAllExchangeResponse, error) {
	exchanges := pb.GetAllExchangeResponse{
		AllExchanges: []*pb.GetExchangeResponse{}, // Explicitly initialize to an empty slice
	}

	// SQL query with corrected JOIN condition
	query := `
		SELECT e.id, e.product_id, e.amount, e.price, e.status, e.contract_id, e.created_at, e.updated_at, e.deleted_at
		FROM exchange as e
		JOIN products as p ON p.id = e.product_id
		JOIN storage as s ON s.id = p.storage_id
		WHERE e.deleted_at = 0
	`

	var args []interface{}
	count := 1

	// Adding filtering conditions based on the request
	if req.ProductId != "" {
		query += fmt.Sprintf(" AND e.product_id = $%d", count)
		args = append(args, req.ProductId)
		count++
	}
	if req.Status != "" {
		query += fmt.Sprintf(" AND e.status = $%d", count)
		args = append(args, req.Status)
		count++
	}

	// Calculate limit and offset based on page
	if req.Limit != 0 {
		page := req.Page
		if page < 1 {
			page = 1 // Default to the first page if an invalid page is given
		}
		offset := (page - 1) * req.Limit
		query += fmt.Sprintf(" LIMIT $%d OFFSET $%d", count, count+1)
		args = append(args, req.Limit, offset)
	}

	// Execute the query
	rows, err := p.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Scan through the result set
	for rows.Next() {
		var exchange pb.GetExchangeResponse
		var contractId sql.NullString

		// Scanning data into fields
		err := rows.Scan(
			&exchange.Id, &exchange.ProductId, &exchange.Amount, &exchange.Price, &exchange.Status,
			&contractId, &exchange.CreatedAt, &exchange.UpdatedAt, &exchange.DeletedAt,
		)

		if err != nil {
			log.Println(err)
			return nil, err
		}

		// Handle the nullable contract ID
		if contractId.Valid {
			exchange.ContractId = contractId.String
		} else {
			exchange.ContractId = ""
		}

		// Append to the list of exchanges
		exchanges.AllExchanges = append(exchanges.AllExchanges, &exchange)
	}

	// Query the total count of exchanges for pagination
	countQuery := `SELECT COUNT(1) FROM exchange where deleted_at=0`
	err = p.db.QueryRow(countQuery).Scan(&count)
	if err != nil {
		return nil, err
	}
	exchanges.Count = int32(count)

	// Check for errors after iterating through rows
	if err = rows.Err(); err != nil {
		return nil, err
	}

	// Return the list of exchanges
	return &exchanges, nil
}


func (p *ExchangeStorage) GetMonthlyStatistics(req *pb.ExchangeStatisticsRequest) (*pb.ExchangeStatisticsResponse, error) {
	// Define start and end dates for the given month
	fmt.Println(req.Month,"\n", req.Year)
	startDate := time.Date(int(req.Year), time.Month(req.Month), 1, 0, 0, 0, 0, time.UTC)
	endDate := startDate.AddDate(0, 1, 0) // First day of the next month (exclusive)

	// Initialize variables for tracking statistics
	var totalBought, totalSold int
	var totalSpend, totalRevenue float64

	// Log start and end dates for debugging
	log.Printf("Querying statistics from %s to %s", startDate.Format(time.RFC3339), endDate.Format(time.RFC3339))

	// Query for total 'buy' statistics
	buyQuery := `
		SELECT COALESCE(SUM(amount), 0) AS total_bought, COALESCE(SUM(amount * price), 0) AS total_spend
		FROM exchange 
		WHERE status = 'buy' AND created_at >= $1 AND created_at < $2 AND deleted_at = 0
	`
	err := p.db.QueryRow(buyQuery, startDate, endDate).Scan(&totalBought, &totalSpend)
	if err != nil {
		log.Printf("Error fetching buy statistics: %v", err)
		return nil, fmt.Errorf("error fetching buy statistics: %v", err)
	}
	fmt.Println("TotalBought:", totalBought, "\nTotalSpend:", totalSpend)
	// Query for total 'sell' statistics
	sellQuery := `
		SELECT COALESCE(SUM(amount), 0) AS total_sold, COALESCE(SUM(amount * price), 0) AS total_revenue
		FROM exchange 
		WHERE status = 'sell' AND created_at >= $1 AND created_at < $2 AND deleted_at = 0
	`
	err = p.db.QueryRow(sellQuery, startDate, endDate).Scan(&totalSold, &totalRevenue)
	if err != nil {
		log.Printf("Error fetching sell statistics: %v", err)
		return nil, fmt.Errorf("error fetching sell statistics: %v", err)
	}
	fmt.Println("TotalSold:", totalSold, "\nTotalRevenue:", totalRevenue)
	// Calculate net amount and profit
	netAmount := totalBought - totalSold
	netProfit := totalRevenue - totalSpend

	// Create and return the response using protobuf types
	return &pb.ExchangeStatisticsResponse{
		TotalBought:  int32(totalBought),
		TotalSold:    int32(totalSold),
		TotalSpend:   totalSpend,
		TotalRevenue: totalRevenue,
		NetAmount:    int32(netAmount),
		NetProfit:    netProfit,
	}, nil
}




func (p *ExchangeStorage) GetExchangeGetbyProductId(req *pb.GetExchangeGetbyProductIdRequest) (*pb.GetExchangeGetbyProductIdResponse, error) {
	query := `
		SELECT id, amount, price, created_at
		FROM exchange
		WHERE product_id = $1 AND deleted_at = 0
	`
	args := []interface{}{req.ProductId}
	count := 2

	// Calculate limit and offset based on page
	if req.Limit != 0 {
		page := req.Page
		if page < 1 {
			page = 1 // Default to the first page if an invalid page is given
		}
		offset := (page - 1) * req.Limit
		query += fmt.Sprintf(" LIMIT $%d OFFSET $%d", count, count+1)
		args = append(args, req.Limit, offset)
	}

	rows, err := p.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var response pb.GetExchangeGetbyProductIdResponse
	for rows.Next() {
		var exchange pb.GetExchangeGetbyProductIdList
		err = rows.Scan(&exchange.Id,&exchange.Amount, &exchange.Price, &exchange.CreatedAt)
		if err != nil {
			return nil, err
		}
		response.Exchange = append(response.Exchange, &exchange)
	}

	// Count query for pagination
	countQuery := `SELECT COUNT(*) FROM exchange WHERE product_id = $1 AND deleted_at = 0`
	err = p.db.QueryRow(countQuery, req.ProductId).Scan(&response.Count)
	if err != nil {
		return nil, err
	}

	return &response, nil
}


// func (p *ExchangeStorage) ListExchanges(req *pb.GetAllExchangeRequest) (*pb.GetAllExchangeResponse, error) {
// 	exchanges := pb.GetAllExchangeResponse{
// 		AllExchanges: []*pb.GetExchangeResponse{}, // Explicitly initialize to an empty slice
// 	}

// 	// SQL query with corrected JOIN condition
// 	query := `
// 		SELECT e.id, e.product_id, e.amount, e.price, e.status, e.contract_id, e.created_at, e.updated_at, e.deleted_at
// 		FROM exchange as e
// 		JOIN products as p ON p.id = e.product_id
// 		JOIN storage as s ON s.id = p.storage_id
// 		WHERE e.deleted_at = 0
// 	`

// 	var args []interface{}
// 	count := 1

// 	// Adding filtering conditions based on the request
// 	if req.ProductId != "" {
// 		query += fmt.Sprintf(" AND e.product_id = $%d", count)
// 		args = append(args, req.ProductId)
// 		count++
// 	}
// 	if req.Status != "" {
// 		query += fmt.Sprintf(" AND e.status = $%d", count)
// 		args = append(args, req.Status)
// 		count++
// 	}

// 	if req.Limit != 0 || req.Offset != 0 {
// 		query += fmt.Sprintf(" LIMIT $%d OFFSET $%d", count, count+1)
// 		args = append(args, req.Limit, req.Offset)
// 	}

// 	// Execute the query
// 	rows, err := p.db.Query(query, args...)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	// Scan through the result set
// 	for rows.Next() {
// 		var exchange pb.GetExchangeResponse
// 		var contractId sql.NullString

// 		// Scanning data into fields
// 		err := rows.Scan(
// 			&exchange.Id, &exchange.ProductId, &exchange.Amount, &exchange.Price, &exchange.Status,
// 			&contractId, &exchange.CreatedAt, &exchange.UpdatedAt, &exchange.DeletedAt,
// 		)

// 		if err != nil {
// 			log.Println(err)
// 			return nil, err
// 		}

// 		// Handle the nullable contract ID
// 		if contractId.Valid {
// 			exchange.ContractId = contractId.String
// 		} else {
// 			exchange.ContractId = ""
// 		}

// 		// Append to the list of exchanges
// 		exchanges.AllExchanges = append(exchanges.AllExchanges, &exchange)
// 	}

// 	query = `SELECT COUNT(1) FROM exchange`
// 	err = p.db.QueryRow(query).Scan(&count)
// 	if err != nil {
// 		return nil, err
// 	}
// 	exchanges.Count = int32(count)

// 	// Check for errors after iterating through rows
// 	if err = rows.Err(); err != nil {
// 		return nil, err
// 	}

// 	// Return the list of exchanges
// 	return &exchanges, nil
// }


