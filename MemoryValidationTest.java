import java.io.FileReader;
import java.util.Properties;

public class MemoryValidationTest {
    /**
     * Validates memory configuration by comparing settings from app.yaml and Java runtime arguments.
     *
     * <p>This method performs the following steps:
     * 1. Loads memory configuration from the 'app.yaml' file
     * 2. Retrieves the memory setting from Java system properties
     * 3. Compares both memory configurations in a case-insensitive manner
     *
     * @param args Command-line arguments (not used in this method)
     * @throws Exception If file reading fails or memory configurations do not match
     * @throws RuntimeException If memory settings in app.yaml and Java arguments differ
     */
    public static void main(String[] args) throws Exception {
        Properties yamlProps = new Properties();
        yamlProps.load(new FileReader("app.yaml"));

        String yamlMemory = yamlProps.getProperty("memory").toLowerCase();
        String javaArgsMemory = System.getProperty("Xmx").toLowerCase();

        if (!yamlMemory.equals(javaArgsMemory)) {
            throw new RuntimeException("Memory mismatch: app.yaml=" + yamlMemory + ", Java Args=" + javaArgsMemory);
        }

        System.out.println("Memory configuration is valid.");
    }
}
