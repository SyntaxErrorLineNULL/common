package buffer

import (
	"testing"
)

// Benchmarks the Write method of ByteBuffer using a small payload of 10 bytes.
// This benchmark verifies the efficiency and performance of the Write operation when
// handling frequent, small data writes. It ensures that the buffer can process these
// operations efficiently, providing insight into the performance impact of Write when
// used repeatedly with small inputs.
// BenchmarkWriteSmall tests the Write method with small data sizes (e.g., 10 bytes).
// BenchmarkWriteSmall-8   	170206342	         7.188 ns/op
func BenchmarkWriteSmall(b *testing.B) {
	// Create a new instance of ByteBuffer for repeated Write operations in this benchmark.
	// This instance will be the target for each Write, allowing measurement of how well the buffer
	// manages high volumes of data and whether it scales effectively under load.
	buf := &ByteBuffer{}

	// Define a small byte slice to be written to the buffer.
	// The slice contains the string "small data" which is 10 bytes, providing a consistent small payload size.
	data := []byte("small data")

	// Reset the timer to ignore any time taken in setup and begin measurement only from the Write calls.
	// This ensures accurate timing of the Write operation itself, isolating it from setup overhead.
	b.ResetTimer()

	// Execute the Write operation `b.N` times, which represents the number of iterations defined by the benchmark framework.
	// In each iteration, the data slice is written to the buffer to measure the performance impact of repeated small writes.
	for i := 0; i < b.N; i++ {
		// Perform the Write operation, writing the small data slice to the buffer.
		// The underscore ignores the return value, as we only measure Write's performance in this benchmark.
		_, _ = buf.Write(data)
	}
}

// Benchmarks the Write method of ByteBuffer with a medium-sized payload of 1,000 bytes.
// This benchmark is designed to measure the performance and efficiency of the Write operation
// when handling a moderately sized data buffer. By using 1,000 bytes of patterned data,
// this test simulates real-world usage, where data in memory is rarely zero-filled and often follows
// certain byte patterns. The benchmark results provide insight into how the Write method scales
// with medium data sizes and its ability to handle typical buffer loads without significant overhead.
// BenchmarkWriteMedium-8   	 2416652	      2122 ns/op
func BenchmarkWriteMedium(b *testing.B) {
	// Create a new instance of ByteBuffer for repeated Write operations in this benchmark.
	// This instance will be the target for each Write, allowing measurement of how well the buffer
	// manages high volumes of data and whether it scales effectively under load.
	buf := &ByteBuffer{}

	// Create a large byte slice of 1,000 bytes to simulate a high-volume write.
	// This slice represents the data that will be written to the ByteBuffer in each iteration,
	// helping us test the buffer’s efficiency in handling large amounts of data.
	data := make([]byte, 1000)

	// Fill the data slice with a repeating pattern, providing meaningful non-zero content.
	// This step iterates over each byte in `data`, assigning values to make the content non-uniform,
	// which simulates real-world data more closely than a zero-initialized buffer.
	for i := range data {
		// Assign a value based on the index to each byte, with the value cycling every 256 bytes.
		// This patterned data makes the buffer contents more realistic for benchmark purposes.
		data[i] = byte(i % 256)
	}

	// Reset the benchmark timer to focus on the Write call, excluding setup operations.
	// This ensures that timing measurements only include the actual Write operation, isolating it for accuracy.
	b.ResetTimer()

	// Execute the benchmark loop, repeating Write calls `b.N` times for performance measurement.
	// The number of iterations (`b.N`) is determined by the testing framework to reach a stable
	// average duration, allowing for accurate benchmarking results.
	for i := 0; i < b.N; i++ {
		// Write the large data slice to the ByteBuffer. The Write method’s result is ignored here,
		// as the focus is purely on benchmarking the time it takes to execute this Write operation.
		_, _ = buf.Write(data)
	}
}

// Benchmarks the Write method of ByteBuffer with a large payload of 100,000 bytes.
// This test evaluates ByteBuffer’s handling of large data sizes, which is essential for applications
// needing efficient buffer management under high data loads. By filling the data buffer
// with a patterned, non-zero content, the benchmark replicates real-world conditions,
// offering insight into performance during intensive write operations.
// BenchmarkWriteLarge-8   	   29431	     44777 ns/op
func BenchmarkWriteLarge(b *testing.B) {
	// Create a new instance of ByteBuffer for repeated Write operations in this benchmark.
	// This instance will be the target for each Write, allowing measurement of how well the buffer
	// manages high volumes of data and whether it scales effectively under load.
	buf := &ByteBuffer{}

	// Create a large byte slice of 100,000 bytes to simulate a high-volume write.
	// This slice represents the data that will be written to the ByteBuffer in each iteration,
	// helping us test the buffer’s efficiency in handling large amounts of data.
	data := make([]byte, 100000)

	// Fill the data slice with a repeating pattern, providing meaningful non-zero content.
	// This step iterates over each byte in `data`, assigning values to make the content non-uniform,
	// which simulates real-world data more closely than a zero-initialized buffer.
	for i := range data {
		// Assign a value based on the index to each byte, with the value cycling every 256 bytes.
		// This patterned data makes the buffer contents more realistic for benchmark purposes.
		data[i] = byte(i % 256)
	}

	// Reset the benchmark timer to focus on the Write call, excluding setup operations.
	// This ensures that timing measurements only include the actual Write operation, isolating it for accuracy.
	b.ResetTimer()

	// Execute the benchmark loop, repeating Write calls `b.N` times for performance measurement.
	// The number of iterations (`b.N`) is determined by the testing framework to reach a stable
	// average duration, allowing for accurate benchmarking results.
	for i := 0; i < b.N; i++ {
		// Write the large data slice to the ByteBuffer. The Write method’s result is ignored here,
		// as the focus is purely on benchmarking the time it takes to execute this Write operation.
		_, _ = buf.Write(data)
	}
}

// Benchmarks the parallel writing capability of ByteBuffer with a small data payload.
// This benchmark simulates a high-concurrency scenario where multiple goroutines
// perform write operations concurrently, each writing a small slice of bytes to ByteBuffer.
// It assesses the performance and thread safety of ByteBuffer’s Write method under
// parallel conditions, specifically with frequent, lightweight writes. The test
// helps evaluate how well ByteBuffer performs in environments with high levels of
// concurrency, ensuring it can handle simultaneous access without data races or errors.
// BenchmarkByteBufferWriteParallelSmall-8   	234693730	        13.83 ns/op
func BenchmarkByteBufferWriteParallelSmall(b *testing.B) {
	// Define a small byte slice with sample data to be written to ByteBuffer.
	// This small payload is designed to simulate frequent but lightweight write operations,
	// making it suitable for testing concurrent small writes to the buffer.
	data := []byte("hello, world!")

	// Run the benchmark in parallel, where multiple goroutines perform writes concurrently.
	// The testing framework coordinates parallel execution, allowing each goroutine to repeatedly
	// write to ByteBuffer until the benchmark completion conditions are met.
	b.RunParallel(func(pb *testing.PB) {
		// Initialize a new instance of ByteBuffer for each parallel goroutine.
		// This buffer instance is isolated to each goroutine, ensuring no data races occur.
		var buf ByteBuffer

		// Enter the main loop where each iteration represents a write operation.
		// The pb.Next() method controls loop progression based on the benchmark’s configuration,
		// allowing each goroutine to perform write operations until the benchmark concludes.
		for pb.Next() {
			// Perform a Write operation to add the data slice to ByteBuffer.
			// This line writes the predefined small data to ByteBuffer, simulating a concurrent write.
			// The result is ignored here as the focus is solely on performance.
			_, err := buf.Write(data)

			// Check for any errors that occurred during the Write operation.
			// If an error is encountered, it is logged as a fatal error, stopping the benchmark,
			// as it indicates an unexpected failure in the buffer’s write handling.
			if err != nil {
				b.Fatal(err)
			}
		}
	})
}

// Benchmarks the parallel writing of a medium-sized byte slice (1,000 bytes) to ByteBuffer.
// This test is designed to evaluate the performance and correctness of the ByteBuffer's
// Write method under a high-concurrency workload. The benchmark runs multiple goroutines
// simultaneously, each performing write operations. The objective is to measure how well
// ByteBuffer handles concurrent writes, ensuring that the implementation is efficient and
// thread-safe without encountering race conditions or performance bottlenecks.
// BenchmarkByteBufferWriteParallelMedium-8   	 4056788	       939.3 ns/op
func BenchmarkByteBufferWriteParallelMedium(b *testing.B) {
	// Initialize a shared instance of ByteBuffer that will be used by all parallel goroutines.
	// The ByteBuffer struct will be written to by multiple goroutines, so it needs to be shared.
	// This setup ensures that each iteration of the benchmark uses the same buffer instance.
	buf := &ByteBuffer{}

	// Create a byte slice of 1,000 bytes to serve as the data written to the buffer.
	// This size is considered medium-sized and will help simulate a moderately sized workload.
	// The data will be reused by all parallel goroutines to benchmark the Write method with the same input.
	data := make([]byte, 1000)

	// Fill the byte slice with a repeating pattern of bytes. This creates a predictable and varied
	// sequence of data that will be written to the buffer. The pattern ensures the data isn't just
	// random and can be verified during testing, should any validation be necessary.
	for i := range data {
		// Set each byte in the slice to a value based on its index mod 256.
		// This creates a simple repeating pattern of values from 0 to 255.
		// The pattern ensures that the data is diverse, avoiding the use of single repeated byte values.
		data[i] = byte(i % 256) // Creates a byte pattern from 0 to 255, looping back after 255.
	}

	// Run the benchmark in parallel mode, which allows multiple goroutines to execute the benchmark code
	// simultaneously. This simulates the environment where multiple threads are performing write operations.
	// It helps to test the performance and concurrency characteristics of the Write method in ByteBuffer.
	// The function 'b.RunParallel' accepts a function that will be run in parallel by multiple goroutines.
	b.RunParallel(func(pb *testing.PB) {
		// Each goroutine will execute this loop, repeatedly writing the same data to the buffer.
		// The 'pb.Next()' method is used to indicate the next iteration in the parallel benchmark.
		// The loop runs until the benchmark has completed the desired number of iterations for the parallel test.
		for pb.Next() {
			// Perform the write operation, which writes the data slice to the buffer.
			// This line invokes the 'Write' method on the ByteBuffer instance, passing the data slice.
			// The result of the write operation is discarded, but we check for errors to ensure that it succeeds.
			_, err := buf.Write(data)

			// If an error occurs during the write operation, the benchmark is immediately halted.
			// The error is passed to 'b.Fatal', which logs the error and terminates the benchmark early.
			// This ensures that any failure is immediately detected, and the benchmark doesn't continue with incorrect behavior.
			if err != nil {
				// If an error occurs, it indicates a failure in the write operation.
				// The benchmark will stop and report the error.
				b.Fatal(err)
			}
		}
	})
}

// Benchmarks the parallel writing of a large byte slice (100,000 bytes) to ByteBuffer.
// This test is designed to evaluate the performance and correctness of the ByteBuffer's
// Write method under a high-concurrency workload. The benchmark runs multiple goroutines
// simultaneously, each performing write operations. The objective is to measure how well
// ByteBuffer handles concurrent writes, ensuring that the implementation is efficient and
// thread-safe without encountering race conditions or performance bottlenecks.
// BenchmarkByteBufferWriteParallelLarge-8   	   51974	    266938 ns/op
func BenchmarkByteBufferWriteParallelLarge(b *testing.B) {
	// Initialize a shared instance of ByteBuffer that will be used by all parallel goroutines.
	// The ByteBuffer struct will be written to by multiple goroutines, so it needs to be shared.
	// This setup ensures that each iteration of the benchmark uses the same buffer instance.
	buf := &ByteBuffer{}

	// Create a byte slice of 100,000 bytes to be written to the buffer.
	// This data size is considered large and will help simulate a real-world scenario where large data
	// needs to be written to a buffer. This large data size allows us to test ByteBuffer's handling
	// of substantial data in a multi-threaded environment.
	data := make([]byte, 100000)

	// Fill the byte slice with a repeating pattern of values between 0 and 255. This creates a
	// predictable pattern in the data, which ensures that the test data is meaningful for both
	// performance and testing purposes. Repeating patterns help us confirm that data is being handled
	// consistently across parallel operations.
	for i := range data {
		// Set each byte to a value based on its index modulo 256, creating a repeating pattern of 0-255.
		// This allows the data to vary while staying within a reasonable byte range.
		// The pattern is useful for ensuring that the data written to the buffer is not static and varies.
		data[i] = byte(i % 256) // Fill with pattern data for meaningful content.
	}

	// Run the benchmark in parallel mode. The 'b.RunParallel' function executes the provided closure
	// across multiple goroutines, which helps simulate a concurrent write load on the ByteBuffer.
	// The goal of running the benchmark in parallel is to measure how well the buffer handles simultaneous
	// writes from multiple threads.
	b.RunParallel(func(pb *testing.PB) {
		// Each goroutine will execute this loop, performing a write operation on the buffer during each iteration.
		// The 'pb.Next()' method is used to signal the next iteration of the benchmark for the goroutine.
		// The loop continues running for the desired number of benchmark iterations as defined by the framework.
		for pb.Next() {
			// Perform the write operation for each iteration, writing the large data slice to the buffer.
			// The result of the write operation is ignored, but we check for errors to ensure the operation succeeds.
			_, err := buf.Write(data)

			// If the write operation encounters an error, the benchmark is immediately halted.
			// The error is passed to 'b.Fatal', which will log the error and stop the benchmark.
			// This ensures that any failures in writing the data will cause the benchmark to fail immediately.
			if err != nil {
				// If an error occurs during writing, the benchmark is stopped and the error is logged.
				// This prevents any invalid benchmarks from continuing, ensuring only successful operations are measured.
				b.Fatal(err)
			}
		}
	})
}
