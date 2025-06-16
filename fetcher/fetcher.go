package fetcher

import (
	"context"

	"github.com/goccy/go-json"
	"github.com/redis/go-redis/v9"
)

// Fetcher is a generic interface defining a contract for types that fetch data from a data source.
// It specifies a single method, Fetch, which retrieves tasks of type T and returns them as a slice.
// The generic type T allows implementations to handle any task type specified by the caller without restriction.
// This interface is designed for simplicity, providing direct access to fetched data without partitioning.
type Fetcher[T any] interface {
	// Fetch retrieves a collection of tasks of type T from a data source identified by the provided keys.
	// It returns a slice of T containing all fetched tasks and an error if the operation encounters a failure.
	// The context parameter enables cancellation and timeout management, while keys specify the data source location.
	Fetch(ctx context.Context, keys []string) ([]T, error)
}

// The script extractCommand is a Lua script that interacts with Redis to fetch tasks from a Redis list.
// It uses the LPOP command to pop tasks from the list until a specified maximum number of tasks max_tasks are fetched,
// or the list is empty, whichever comes first. The Lua script ensures efficient retrieval of tasks while respecting the max limit.
var extractCommand = redis.NewScript(`
local key = KEYS[1]
local max_tasks = tonumber(ARGV[1])
local tasks = {}

for i = 1, max_tasks do
	local task = redis.call('LPOP', key)
	if not task then
		break
	end
	table.insert(tasks, task)
end

return tasks
`)

// maxTask defines the maximum number of tasks to be fetched in a single operation.
// This constant ensures that the system will not try to fetch an unreasonably large number of tasks from Redis at once.
const maxTask = 1000

// RedisFetcher is a generic type that represents a structure for fetching data from a Redis storage.
// It includes a Redis client and a logger, allowing it to perform operations like fetching data
// from Redis and logging relevant information during its operations.
type RedisFetcher[T any] struct {
	rdb redis.UniversalClient
}

// New creates a new instance of the RedisFetcher with the provided Redis client and logger.
// The function is generic, allowing it to work with any type `T` that satisfies the `comparable` constraint,
// enabling the RedisFetcher to handle different types of data in a flexible manner.
func New[T any](rdb redis.UniversalClient) *RedisFetcher[T] {
	return &RedisFetcher[T]{rdb: rdb}
}

// Fetch is a method on the RedisFetcher struct that retrieves a list of tasks from Redis based on the provided keys.
// It executes a Lua script using the Redis client to fetch up to a maximum number of tasks from the Redis list.
// The method returns a slice of tasks of type T and an error if any occurred during the operation.
func (f *RedisFetcher[T]) Fetch(ctx context.Context, keys []string) ([]T, error) {
	// Run the Redis Lua script using the provided context, Redis client universal client,
	// and the specified keys, along with the maxTask limit as an argument.
	result, err := extractCommand.Run(ctx, f.rdb, keys, maxTask).Result()
	// Check if an error occurred during the script execution.
	if err != nil {
		return nil, err
	}

	// Create an empty slice tasks of type T using the make function.
	// T is a generic type, so the actual type of tasks will be determined at runtime.
	// The make function initializes the slice with an initial length of 0, meaning it starts empty.
	// The capacity of the slice will be dynamic, growing as elements are appended to it.
	// This slice will store the tasks fetched and unmarshalled from Redis.
	tasks := make([]T, 0)
	// Check if the result from Redis is a slice of empty interfaces and contains elements.
	// This ensures that the result is in the expected format of a list and is not empty.
	if results, ok := result.([]interface{}); ok && len(results) > 0 {
		// Iterate over each task in the results slice.
		for _, task := range results {
			// Create a new instance of type T to hold the unmarshalled task data.
			// The out variable will be populated with the task's data if unmarshalling is successful.
			out := new(T)

			// Check if the task is of type string.
			// The ok variable indicates whether the type assertion was successful.
			// If successful, the task is stored in the value variable for further processing.
			if value, ok := task.(string); ok {
				// Attempt to unmarshal the task (which is in string format) into the out variable of type T.
				// The task is expected to be in JSON format as a string, so json.Unmarshal is used to decode it.
				if err = json.Unmarshal([]byte(value), out); err != nil {
					// If unmarshalling fails, log the error and continue to the next task.
					// This ensures that one failed task does not interrupt the processing of other tasks.
					continue
				}
			}

			// If unmarshalling is successful, append the unmarshalled task to the tasks slice.
			// The task is now an instance of type T, and can be used further in the application.
			tasks = append(tasks, *out)
		}
	}

	// After all tasks are processed, the tasks slice will contain all successfully unmarshalled tasks.
	// If no valid tasks were found or unmarshalled, the tasks slice will be empty, which is valid.
	return tasks, nil
}
