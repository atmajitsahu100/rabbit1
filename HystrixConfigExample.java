import com.netflix.hystrix.*;
import com.netflix.hystrix.strategy.concurrency.HystrixRequestContext;

import java.util.concurrent.TimeUnit;

public class HystrixConfigExample {

    /**
     * Demonstrates the execution of Hystrix commands with circuit breaker pattern.
     *
     * This method initializes a Hystrix request context and executes 100 instances of {@link CommandExample},
     * printing the result of each command execution. The request context is properly managed and shut down
     * in a finally block to ensure resource cleanup.
     *
     * @param args Command-line arguments (not used in this implementation)
     */
    public static void main(String[] args) {
        HystrixRequestContext context = HystrixRequestContext.initializeContext();
        try {
            for (int i = 0; i < 100; i++) {
                System.out.println(new CommandExample("Task-" + i).execute());
            }
        } finally {
            context.shutdown();
        }
    }

    static class CommandExample extends HystrixCommand<String> {

        private final String name;

        /**
         * Constructs a CommandExample with Hystrix configuration for circuit breaker and thread pool management.
         *
         * @param name A unique identifier for the command instance
         *
         * @implNote This constructor configures a Hystrix command with:
         * - Group key: "ExampleGroup"
         * - Command key: "ExampleCommand"
         * - Thread pool key: "ExampleThreadPool"
         * - Thread pool settings:
         *   - Core size: 10 threads
         *   - Maximum size: 15 threads
         *   - Dynamic size adjustment enabled
         *   - Unbounded queue (-1)
         *   - Queue rejection threshold: 100
         * - Circuit breaker properties:
         *   - Enabled with request volume threshold of 50
         *   - Sleep window of 5 seconds after circuit opens
         *   - Execution timeout of 2 seconds
         *   - Fallback mechanism enabled
         */
        protected CommandExample(String name) {
            super(Setter.withGroupKey(HystrixCommandGroupKey.Factory.asKey("ExampleGroup"))
                    .andCommandKey(HystrixCommandKey.Factory.asKey("ExampleCommand"))
                    .andThreadPoolKey(HystrixThreadPoolKey.Factory.asKey("ExampleThreadPool"))
                    .andThreadPoolPropertiesDefaults(HystrixThreadPoolProperties.Setter()
                            .withCoreSize(10)
                            .withMaximumSize(15)
                            .withAllowMaximumSizeToDivergeFromCoreSize(true)
                            .withMaxQueueSize(-1)
                            .withQueueSizeRejectionThreshold(100))
                    .andCommandPropertiesDefaults(HystrixCommandProperties.Setter()
                            .withCircuitBreakerEnabled(true)
                            .withCircuitBreakerRequestVolumeThreshold(50)
                            .withCircuitBreakerSleepWindowInMilliseconds(5000)
                            .withExecutionTimeoutInMilliseconds(2000)
                            .withFallbackEnabled(true)));
            this.name = name;
        }

        /**
         * Executes the main logic of the Hystrix command with simulated success or failure.
         *
         * @return A success message if the random condition is met
         * @throws Exception If a random failure condition occurs, throwing a RuntimeException
         * 
         * This method demonstrates a probabilistic execution scenario:
         * - 50% chance of successful execution with a 100ms delay
         * - 50% chance of throwing a RuntimeException
         * 
         * The method uses Math.random() to determine the outcome and simulates 
         * potential variability in command execution.
         */
        @Override
        protected String run() throws Exception {
            if (Math.random() > 0.5) {
                TimeUnit.MILLISECONDS.sleep(100);
                return "Success: " + name;
            } else {
                throw new RuntimeException("Failure: " + name);
            }
        }

        /**
         * Provides a fallback response when the Hystrix command execution fails.
         *
         * This method is called automatically by Hystrix when the primary command execution
         * encounters an error, throws an exception, or is rejected. It returns a predefined
         * fallback message that includes the command's name.
         *
         * @return A string representing the fallback response, prefixed with "Fallback: "
         *         and appended with the command's name
         */
        @Override
        protected String getFallback() {
            return "Fallback: " + name;
        }
    }
}
