import com.netflix.hystrix.*;
import com.netflix.hystrix.strategy.concurrency.HystrixRequestContext;

import java.util.concurrent.TimeUnit;

public class HystrixConfigExample {

    /**
     * Demonstrates Hystrix command execution with circuit breaker pattern.
     *
     * Initializes a Hystrix request context and executes 100 instances of {@link CommandExample},
     * printing the result of each command execution. Ensures proper context shutdown after processing.
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
         * @see HystrixCommand
         * @see HystrixCommandGroupKey
         * @see HystrixThreadPoolProperties
         * @see HystrixCommandProperties
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
         * Executes the command logic with a simulated success or failure scenario.
         *
         * @return A success message if the random condition is met
         * @throws Exception If a random failure condition occurs, throwing a RuntimeException
         *
         * @implNote This method randomly determines the command's outcome:
         * - With a 50% probability, it simulates a successful execution by sleeping for 100 milliseconds
         *   and returning a success message with the command's name.
         * - With a 50% probability, it throws a RuntimeException indicating a failure,
         *   which will trigger the fallback mechanism.
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
         * @return A string indicating the fallback message with the command's name
         */
        @Override
        protected String getFallback() {
            return "Fallback: " + name;
        }
    }
}
