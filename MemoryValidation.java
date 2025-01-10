import java.io.FileReader;
import java.util.Map;
import org.yaml.snakeyaml.Yaml;

public class MemoryValidation {
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
