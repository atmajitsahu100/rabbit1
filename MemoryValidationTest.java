import java.io.FileReader;
import java.util.Properties;

public class MemoryValidationTest {
    /**
     * Validates memory configuration by comparing settings from app.yaml and Java system properties.
     *
     * @param args Command-line arguments (not used in this method)
     * @throws Exception If there are issues reading the YAML file or retrieving system properties
     * @throws RuntimeException If memory configuration in app.yaml does not match Java runtime memory settings
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
