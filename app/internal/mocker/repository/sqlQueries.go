package repository

const (
	queryGetTableNames = `
	SELECT table_schema, table_name AS full_table_name
	FROM information_schema.tables
	WHERE table_schema NOT IN ('information_schema', 'pg_catalog');
	  `

	// queryGetTableNames = `SELECT table_name AS full_table_name
	//  FROM information_schema.tables
	//  WHERE table_schema NOT IN ('information_schema', 'pg_catalog');`

	queryGetColumns = `
  SELECT column_name, data_type FROM information_schema.columns WHERE table_name=$1
	`

	queryGetRowsNum = `
	SELECT COUNT(1) FROM %s
	`
)
