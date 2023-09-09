package postgres

import (
	_ "github.com/lib/pq"
)

// type postgresRunner struct {
// 	conn *sql.DB
// }

// type PostgresRunnerReq struct {
// 	Configs *runnerport.PostgresConfig
// }

// func NewPostgresRunner(req *PostgresRunnerReq) runnerport.Runner {
// 	conn, err := getPostgresConnection(req.Configs)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	return &postgresRunner{
// 		conn: conn,
// 	}
// }

// func (r *postgresRunner) ExecuteScript(request runnerport.RunnerReq) (runnerport.RunnerRes, error) {
// 	fmt.Printf("script request%+v\n", request)

// 	return &models.RunnerOSResponse{}, nil
// }

// func getPostgresConnection(config *runnerport.PostgresConfig) (*sql.DB, error) {
// 	connString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
// 		config.Host, config.Port, config.Username, config.Password, config.Database)
// 	fmt.Println(connString)

// 	db, err := sql.Open("postgres", connString)
// 	if err != nil {
// 		return nil, err
// 	}

// 	// Ping the database to ensure a connection can be made
// 	err = db.Ping()
// 	if err != nil {
// 		return nil, err
// 	}

// 	return db, nil
// }
