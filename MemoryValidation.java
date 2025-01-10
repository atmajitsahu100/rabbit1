import java.io.FileReader;
import java.util.Map;
import org.yaml.snakeyaml.Yaml;

public class MemoryValidation {
    /**
     * Validates memory configuration by comparing memory settings from a YAML file with Java arguments.
     *
     * This method reads memory configuration from 'app.yaml' and compares it with the maximum heap size
     * specified in the Java virtual machine arguments. It performs a case-insensitive comparison and
     * prints the result of the validation.
     *
     * @throws Exception if there are issues reading the YAML file or processing its contents
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
     * Retrieves the maximum heap size (-Xmx) from Java Virtual Machine arguments.
     *
     * @return A string representing the memory value specified by the -Xmx flag,
     *         or null if no -Xmx argument is found in the JVM arguments
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
