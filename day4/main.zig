const std = @import("std");

pub fn main() !void {
    const file = try std.fs.cwd().openFile("input.txt", .{});
    defer file.close();

    var bufReader = std.io.bufferedReader(file.reader());
    var inStream = bufReader.reader();

    var arena = std.heap.ArenaAllocator.init(std.heap.page_allocator);
    defer arena.deinit();
    const ally = arena.allocator();

    var lines = std.ArrayList([]u8).init(ally);

    var buf: [1024]u8 = undefined;
    while (try inStream.readUntilDelimiterOrEof(&buf, '\n')) |line| {
        const copyLine = try ally.alloc(u8, line.len);
        std.mem.copyForwards(u8, copyLine, line);
        try lines.append(copyLine);
    }

    for (lines.items) |line| {
        std.debug.print("{s}\n", .{line});
    }

    const length = 3;

    var count: u32 = 0;
    count += 0;
    for (lines.items, 0..) |line, uy| {
        lineFor: for (line, 0..) |c, ux| {
            if (c != 'X') continue :lineFor;
            const y: isize = @intCast(uy);
            const x: isize = @intCast(ux);
            if (x + length < line.len) {
                const m = line[ux + 1] == 'M';
                const a = line[ux + 2] == 'A';
                const s = line[ux + 3] == 'S';
                if (m and a and s) {
                    count += 1;
                }
            }
            if (0 <= x - length) {
                const m = line[ux - 1] == 'M';
                const a = line[ux - 2] == 'A';
                const s = line[ux - 3] == 'S';
                if (m and a and s) {
                    count += 1;
                }
            }
            if (0 <= y - length) {
                const m = lines.items[uy - 1][ux] == 'M';
                const a = lines.items[uy - 2][ux] == 'A';
                const s = lines.items[uy - 3][ux] == 'S';
                if (m and a and s) {
                    count += 1;
                }
            }
            if (y + length < lines.items.len) {
                const m = lines.items[uy + 1][ux] == 'M';
                const a = lines.items[uy + 2][ux] == 'A';
                const s = lines.items[uy + 3][ux] == 'S';
                if (m and a and s) {
                    count += 1;
                }
            }
            if (y + length < lines.items.len and x + length < line.len) {
                const m = lines.items[uy + 1][ux + 1] == 'M';
                const a = lines.items[uy + 2][ux + 2] == 'A';
                const s = lines.items[uy + 3][ux + 3] == 'S';
                if (m and a and s) {
                    count += 1;
                }
            }
            if (0 <= y - length and x + length < line.len) {
                const m = lines.items[uy - 1][ux + 1] == 'M';
                const a = lines.items[uy - 2][ux + 2] == 'A';
                const s = lines.items[uy - 3][ux + 3] == 'S';
                if (m and a and s) {
                    count += 1;
                }
            }
            if (0 <= y - length and 0 <= x - length) {
                const m = lines.items[uy - 1][ux - 1] == 'M';
                const a = lines.items[uy - 2][ux - 2] == 'A';
                const s = lines.items[uy - 3][ux - 3] == 'S';
                if (m and a and s) {
                    count += 1;
                }
            }
            if (y + length < lines.items.len and 0 <= x - length) {
                const m = lines.items[uy + 1][ux - 1] == 'M';
                const a = lines.items[uy + 2][ux - 2] == 'A';
                const s = lines.items[uy + 3][ux - 3] == 'S';
                if (m and a and s) {
                    count += 1;
                }
            }
        }
    }

    std.debug.print("Example: {}\n", .{count});
}
