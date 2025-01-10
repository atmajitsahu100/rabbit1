import java.io.FileReader;
import java.util.Map;
import org.yaml.snakeyaml.Yaml;

public class MemoryValidation {
    /**
     * Main method to validate memory configuration from a YAML file against Java runtime arguments.
     *
     * This method reads memory configuration from 'app.yaml', retrieves the current Java VM memory settings,
     * and compares them for consistency. It handles potential file reading and parsing exceptions.
     *
     * @throws Exception if there are issues reading the YAML file or accessing system properties
     */
    public static void main(String[] args) {
        String yamlFile = "app.yaml";
        Yaml yaml = new Yaml();
        try {
            Map<String, Object> yamlData = yaml.load(new FileReader(yamlFile));
            String definedMemory = (String) yamlData.get("memory");
            String javaArgsMemory = getJavaArgsMemory();

            if (!definedMemory.equalsIgnoreCase(javaArgsMemory)) {
                System.out.println("Mismatch: YAML memory: " + definedMemory + " vs Java Args memory: " + javaArgsMemory);
            } else {
                System.out.println("Memory configuration matches");
            }
        } catch (Exception e) {
            e.printStackTrace();
        }
    }

    /**
     * Retrieves the maximum heap size specified in the Java virtual machine arguments.
     *
     * @return A string representing the maximum heap size (e.g., "2g", "512m"), 
     *         or {@code null} if no maximum heap size is specified
     */
    private static String getJavaArgsMemory() {
        String javaArgs = System.getProperty("java.vm.args");
        String memoryValue = null;
        if (javaArgs.contains("-Xmx")) {
            int startIndex = javaArgs.indexOf("-Xmx") + 4;
            int endIndex = javaArgs.indexOf(" ", startIndex);
            memoryValue = endIndex == -1 ? javaArgs.substring(startIndex) : javaArgs.substring(startIndex, endIndex);
        }
        return memoryValue;
    }
}
