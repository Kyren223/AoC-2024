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

    // for (lines.items) |line| {
    // std.debug.print("{s}\n", .{line});
    // }

    var count: u32 = 0;
    for (lines.items, 0..) |line, uy| {
        lineFor: for (line, 0..) |c, ux| {
            if (c != 'A') continue :lineFor;
            if (@as(isize, @intCast(uy)) - 1 < 0 or uy + 1 >= lines.items.len) {
                continue :lineFor;
            }
            if (@as(isize, @intCast(ux)) - 1 < 0 or ux + 1 >= lines.items[uy].len) {
                continue :lineFor;
            }
            const tl = lines.items[uy - 1][ux - 1];
            const tr = lines.items[uy - 1][ux + 1];
            const bl = lines.items[uy + 1][ux - 1];
            const br = lines.items[uy + 1][ux + 1];

            const d1 = (tl == 'M' and br == 'S') or (tl == 'S' and br == 'M');
            const d2 = (tr == 'M' and bl == 'S') or (tr == 'S' and bl == 'M');
            if (d1 and d2) {
                count += 1;
            }
        }
    }

    std.debug.print("Example: {}\n", .{count});
}
