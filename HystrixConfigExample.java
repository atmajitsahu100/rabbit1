import com.netflix.hystrix.*;
import com.netflix.hystrix.strategy.concurrency.HystrixRequestContext;

import java.util.concurrent.TimeUnit;

public class HystrixConfigExample {

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

        @Override
        protected String run() throws Exception {
            if (Math.random() > 0.5) {
                TimeUnit.MILLISECONDS.sleep(100);
                return "Success: " + name;
            } else {
                throw new RuntimeException("Failure: " + name);
            }
        }

        @Override
        protected String getFallback() {
            return "Fallback: " + name;
        }
    }
}
