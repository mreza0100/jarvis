package runners

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"github.com/pkg/errors"

	"github.com/mreza0100/jarvis/internal/models"
	runnerport "github.com/mreza0100/jarvis/internal/ports/runnerport"
)

type pgsRunner struct {
	conn *sql.DB
}

type PgsRunnerReq struct {
	Configs *models.PostgresConnConfig
}

func NewPgsRunner(req *PgsRunnerReq) runnerport.PgsRunner {
	runner := &pgsRunner{}
	runner.initPgsConnection(req.Configs)
	return runner
}

func (r *pgsRunner) ExecScript(req *models.PgsRunnerRequest) (*models.PgsRunnerResponse, error) {
	resultSets := &models.PgsRunnerResponse{
		QueryResponses: make([]*models.QueryResult, 0, 5),
	}

	rows, err := r.conn.Query(req.Query)
	if err != nil {
		resultSets.Err = err
		return resultSets, nil
	}

	isFirst := true
	for isFirst || rows.NextResultSet() {
		isFirst = false
		result, err := r.scanResult(rows)
		if err != nil {
			return &models.PgsRunnerResponse{
				Err: err,
			}, nil
		}
		resultSets.QueryResponses = append(resultSets.QueryResponses, result)
	}

	return resultSets, nil
}

func (r *pgsRunner) scanResult(rows *sql.Rows) (result *models.QueryResult, err error) {
	result = new(models.QueryResult)

	fmt.Println(rows.Columns())
	result.Columns, err = rows.Columns()
	if err != nil {
		return nil, err
	}

	columnsType, err := rows.ColumnTypes()
	if err != nil {
		return nil, err
	}

	for _, v := range columnsType {
		columnType := &models.ColumnType{
			Name:             v.Name(),
			DatabaseTypeName: v.DatabaseTypeName(),
			Length:           nil,
			Nullable:         nil,
		}

		len, ok := v.Length()
		if ok {
			columnType.Length = &len
		}

		isNullable, ok := v.Nullable()
		if ok {
			columnType.Nullable = &isNullable
		}

		result.ColumnsType = append(result.ColumnsType, columnType)
	}

	if err := rows.Err(); err != nil {
		result.Err = errors.Wrap(err, "Error iterating over rows")
	}
	for rows.Next() {
		columnValues := make([]any, len(result.Columns))
		for i := range columnValues {
			columnValues[i] = new(any)
		}

		if err := rows.Scan(columnValues...); err != nil {
			return nil, err
		}

		result.ColumnValues = append(result.ColumnValues, columnValues...)
	}

	return result, nil
}

func (pgs *pgsRunner) initPgsConnection(config *models.PostgresConnConfig) {
	// TODO: try connecting with linux socket domain
	connString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.Host, config.Port, config.Username, config.Password, config.Database)

	db, err := sql.Open("postgres", connString)
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	pgs.conn = db
}
