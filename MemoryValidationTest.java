import java.io.FileReader;
import java.util.Properties;

public class MemoryValidationTest {
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
