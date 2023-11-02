package repository

const (
	queryGetTableNames = `
SELECT table_schema || '.' || table_name AS full_table_name
FROM information_schema.tables
WHERE table_schema NOT IN ('information_schema', 'pg_catalog');
  `
)
