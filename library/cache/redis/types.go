package redis

const (
	redisKeyPrefix = "redis."

	InstantDefault Instant = "default"
	InstantQueue   Instant = "queue"
)

type Instant string

type RDBConfig struct {
	Host     string
	Port     int64
	Pass     string
	Database int
}
