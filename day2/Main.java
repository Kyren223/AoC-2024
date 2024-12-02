import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Path;
import java.util.List;

public class Main {
    public static void main(String[] args) throws IOException {
        System.out.println("hello world");
        List<String> reports = Files.readAllLines(Path.of("example_input1.txt"));
        System.out.println("Example reports: " + getValids(reports));
        reports = Files.readAllLines(Path.of("input1.txt"));
        System.out.println("Reports: " + getValids(reports));

    }

    public static int getValids(List<String> reports) {
        int valids = 0;

        for (String report : reports) {
            String[] levels = report.split(" ");
            boolean increasing = Integer.parseInt(levels[0]) < Integer.parseInt(levels[1]);

            boolean isValid = true;
            for (int i = 0; i < levels.length - 1; i++) {
                int diff = Integer.parseInt(levels[i]) - Integer.parseInt(levels[i + 1]);
                if (increasing && (diff < -3 || diff >= 0)) {
                    isValid = false;
                    break;
                } else if (!increasing && (diff > 3 || diff <= 0)) {
                    isValid = false;
                    break;
                }
            }
            if (isValid) {
                valids += 1;
            }
        }

        return valids;
    }
}
