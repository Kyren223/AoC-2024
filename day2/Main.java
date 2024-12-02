import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Path;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

public class Main {
    public static void main(String[] args) throws IOException {
        List<String> reports = Files.readAllLines(Path.of("example_input1.txt"));
        System.out.println("Example reports: " + getValids(reports));
        reports = Files.readAllLines(Path.of("input1.txt"));
        System.out.println("Reports: " + getValids(reports));

        reports = Files.readAllLines(Path.of("example_input1.txt"));
        System.out.println("Example reports part2: " + getValids2(reports));
        reports = Files.readAllLines(Path.of("input1.txt"));
        System.out.println("Reports part2: " + getValids2(reports));
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

    public static int getValids2(List<String> reports) {
        int valids = 0;

        for (String report : reports) {
            int oldValid = valids;
            List<String> levels = new ArrayList<>(Arrays.asList(report.split(" ")));
            for (int ignore = 0; ignore < levels.size(); ignore++) {
                List<String> copy = new ArrayList<>(levels);
                copy.remove(ignore);
                boolean increasing = Integer.parseInt(copy.get(0)) < Integer.parseInt(copy.get(1));
                boolean isValid = true;
                for (int i = 0; i < copy.size() - 1; i++) {
                    int diff = Integer.parseInt(copy.get(i)) - Integer.parseInt(copy.get(i + 1));
                    // System.out.println("Diff: " + diff);
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
                    break;
                } else {
                    // System.out.println(copy + ""+ isValid);
                }
            }

            if (valids == oldValid) {
                // System.out.println(report);
            }
        }

        return valids;
    }
}
