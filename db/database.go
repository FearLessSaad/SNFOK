package db

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/FearLessSaad/SNFOK/tooling/logger"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

// dbInstance holds the singleton default database connection and managed connections.
var (
	db          *bun.DB            // Default singleton connection
	once        sync.Once          // For default connection initialization
	connections map[string]*bun.DB // Cache of managed connections by DSN
	connMutex   sync.RWMutex       // Mutex for thread-safe connection map access
	connOnce    sync.Once          // For initializing connections map
)

// Config holds database configuration.
type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

// initDB initializes the default database connection.
func initDB() error {
	// Load configuration from environment variables
	config := Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSL_MODE"),
	}

	// Set defaults if not provided
	if config.Host == "" {
		config.Host = "localhost"
	}
	if config.Port == "" {
		config.Port = "5432"
	}
	if config.User == "" {
		config.User = "postgres"
	}
	if config.DBName == "" {
		config.DBName = "pos"
	}
	if config.SSLMode == "" {
		config.SSLMode = "disable"
	}

	// Create connection
	bunDB, err := connectToDB(config)
	if err != nil {
		return err
	}

	db = bunDB
	return nil
}

// connectToDB creates a new database connection for the given config.
func connectToDB(config Config) (*bun.DB, error) {
	// Create connection string
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		config.User, config.Password, config.Host, config.Port, config.DBName, config.SSLMode,
	)

	// Open SQL connection
	sqlDB := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))

	// Configure connection pool
	sqlDB.SetMaxOpenConns(25)
	sqlDB.SetMaxIdleConns(25)
	sqlDB.SetConnMaxLifetime(5 * time.Minute)

	// Create Bun DB instance
	bunDB := bun.NewDB(sqlDB, pgdialect.New())

	// Verify connection
	ctx := context.Background()
	if err := bunDB.PingContext(ctx); err != nil {
		logger.Log("error", "Failed to connect to database", logger.Field{Key: "error", Value: err.Error()})
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	logger.Log("", "Database connection established successfully",
		logger.Field{Key: "host", Value: config.Host},
		logger.Field{Key: "port", Value: config.Port},
		logger.Field{Key: "db_name", Value: config.DBName},
	)

	return bunDB, nil
}

// Connect establishes a connection to a specific database and caches it.
func Connect(config Config) (*bun.DB, error) {
	// Initialize connections map
	connOnce.Do(func() {
		connections = make(map[string]*bun.DB)
	})

	// Create DSN for cache key
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		config.User, config.Password, config.Host, config.Port, config.DBName, config.SSLMode,
	)

	// Check if connection already exists
	connMutex.RLock()
	if bunDB, exists := connections[dsn]; exists {
		connMutex.RUnlock()
		return bunDB, nil
	}
	connMutex.RUnlock()

	// Create new connection
	bunDB, err := connectToDB(config)
	if err != nil {
		return nil, err
	}

	// Cache the connection
	connMutex.Lock()
	connections[dsn] = bunDB
	connMutex.Unlock()

	return bunDB, nil
}

// GetOrConnect returns an existing connection or creates a new one.
func GetOrConnect(config Config) (*bun.DB, error) {
	return Connect(config)
}

// GetDB returns the singleton default Bun database instance.
func GetDB() *bun.DB {
	once.Do(func() {
		if err := initDB(); err != nil {
			logger.Log("error", "Failed to initialize database", logger.Field{Key: "error", Value: err.Error()})
			panic(err)
		}
	})
	return db
}

// Close closes the default singleton database connection.
func Close() error {
	if db != nil {
		logger.Log("", "Closing default database connection")
		err := db.Close()
		db = nil
		return err
	}
	return nil
}

// CloseDB closes a specific database connection and removes it from the cache.
func CloseDB(bunDB *bun.DB) error {
	if bunDB == nil {
		return nil
	}

	connMutex.Lock()
	defer connMutex.Unlock()

	// Find and remove the connection from the cache
	for dsn, conn := range connections {
		if conn == bunDB {
			delete(connections, dsn)
			logger.Log("", "Closing specific database connection", logger.Field{Key: "dsn", Value: dsn})
			return bunDB.Close()
		}
	}

	// If not in cache, close anyway
	logger.Log("", "Closing uncached database connection")
	return bunDB.Close()
}

// CloseAll closes all managed database connections.
func CloseAll() error {
	connMutex.Lock()
	defer connMutex.Unlock()

	var lastErr error
	for dsn, bunDB := range connections {
		logger.Log("", "Closing database connection", logger.Field{Key: "dsn", Value: dsn})
		if err := bunDB.Close(); err != nil {
			lastErr = err
		}
		delete(connections, dsn)
	}

	// Close default connection
	if db != nil {
		logger.Log("", "Closing default database connection")
		if err := db.Close(); err != nil {
			lastErr = err
		}
		db = nil
	}

	return lastErr
}
