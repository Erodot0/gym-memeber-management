package ports

import (
	"time"

	"github.com/Erodot0/gym-memeber-management/internals/app/domains/entities"
	"github.com/gofiber/fiber/v2"
)

// HttpResponses defines methods for handling HTTP responses
type HttpAdapters interface {

	// 200 ok response
	//	Tokens **must** be used only for mobile app
	Success(c *fiber.Ctx, data interface{}, message string, tokens *entities.Tokens) error

	// 400 bad request
	BadRequest(c *fiber.Ctx, message string) error

	// 401 unauthorized
	Unauthorized(c *fiber.Ctx, message string) error

	// 403 forbidden
	Forbidden(c *fiber.Ctx) error

	// 404 not found
	NotFound(c *fiber.Ctx, message string) error

	// 500 internal server error
	InternalServerError(c *fiber.Ctx, message string) error

	// 200 ok response with file
	WithFile(c *fiber.Ctx, pathToFile string) error
}

type CacheAdapters interface {
	// SetCache sets data in Redis based on the provided CachePort data.
	//
	// Parameters:
	//   - data: is Data to be set in Redis using the CachePort interface
	//
	// Returns:
	//   - error: if there was an error setting the data in Redis
	SetCache(data CachePort) error
	// GetCacheKeys retrieves keys from Redis based on the provided CachePort data.
	//
	// Parameters:
	//   - data: the CachePort data used to retrieve keys from Redis
	//
	// Returns:
	//   - []string: a slice of strings representing the keys retrieved from Redis
	//   - error: if there was an error retrieving the keys from Redis
	GetCacheKeys(data CachePort) ([]string, error)
	// GetCacheFromKey retrieves data from Redis based on the provided CachePort data.
	//
	// Parameters:
	//   - key: the key to retrieve data from Redis
	//   - data: the CachePort data used to retrieve the data from Redis
	//
	// Returns:
	//   - error: if there was an error retrieving the data from Redis
	GetCacheFromKey(key string, data CachePort) error
	// GetCacheFromData retrieves data from Redis.
	//
	// Parameters:
	//   - data: data to scan from Redis
	//
	// Returns:
	//   - error: if there was an error retrieving the data from Redis
	GetCacheFromData(data CachePort) error
	// DelCache deletes a key from Redis based on the provided CachePort data.
	//
	// Parameters:
	//   - data: the CachePort data used to retrieve the key from Redis
	//
	// Returns:
	//   - error: if there was an error deleting the key from Redis
	DelCache(data CachePort) error
	// DelCacheMultiple deletes multiple keys from Redis based on the provided CachePort data.
	//
	// Parameters:
	//   - data: the CachePort data used to retrieve the keys from Redis
	//
	// Returns:
	//   - error: if there was an error deleting the keys from Redis
	DelCacheMultiple(data CachePort) error
}

// CachePort defines methods for handling cache
type CachePort interface {

	// SetCacheKey returns the cache key for the cache port.
	SetCacheKey() string

	// SetCacheExpiration returns the expiration time for the cache port.
	SetCacheExpiration() time.Duration

	// GetCacheKey returns the cache key for the cache port.
	GetCacheKey() string
}
