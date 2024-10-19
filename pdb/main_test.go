package bankAccount

import (
    "context"
    "fmt"
    "log"
    "os"
    "testing"

    "github.com/jackc/pgx/v5"
)

const (
    dbDriver = "postgres"
    dbSource = "postgresql://bank_ag_db_user:123456@bank-postgres:5432/bank_ag_db?sslmode=disable"
)

var testQueries *Queries

func TestMain(m *testing.M) {
    conn, err := pgx.Connect(context.Background(), dbSource)
    if err != nil {
        log.Fatal("cannot connect to database: ", err)
    }
    testQueries = New(conn)
    fmt.Printf("connected successfully ")
    os.Exit(m.Run())
}

func TestQueryMigrations(t *testing.T) {
    // Replace this with a real query for your database
    ctx := context.Background()
    rows, err := testQueries.db.Query(ctx, "SELECT * FROM schema_migrations")
    if err != nil {
        t.Fatal("Failed to execute query:", err)
    }
    defer rows.Close()
    rowNum := 0
    for rows.Next() {
        rowNum++
        var version int
        var dirty bool
        if err := rows.Scan(&version, &dirty); err != nil {
            t.Fatalf("Failed to scan row: %v", err)
        }
        fmt.Printf("Row %d: version=%d, dirty=%t\n", rowNum, version, dirty)
    }
    t.Logf("run query")
}