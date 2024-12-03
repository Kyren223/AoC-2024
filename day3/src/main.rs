use std::fs;

fn main() {
    println!("Hello, world!");
    let example = "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))";
    println!("Example: {}", part1(&example));
    
    let input = fs::read_to_string("input.txt").unwrap();
    println!("Part1: {}", part1(input.as_str()));
}

fn part1(input: &str) -> u32 {
    let mut sum = 0;
    for (i, c) in input.char_indices() {
        if c == 'm' {
            let result = try_parse(&input[i..]);
            if let Some(v) = result {
                // println!("Result: {}", v);
                sum += v
            }
        }
    }
    sum
}

fn try_parse(input: &str) -> Option<u32> {
    // println!("Input: {}", input);
    if !input.starts_with("mul(") {
        return None;
    }

    let (result, length) = parse_three(&input[4..]);
    if length == 0 {
        return None;
    }
    let len: usize = 5_usize + length as usize;
    let sep = input.chars().nth(len - 1)?;
    if sep != ',' {
        return None;
    }
    let (result2, length2) = parse_three(&input[len..]);
    if length2 == 0 {
        return None;
    }
    if input.chars().nth(len + length2 as usize)? != ')' {
        return None;
    }

    return Some(result.unwrap() * result2.unwrap());
}

fn parse_three(input: &str) -> (Option<u32>, i32) {
    let potential = match input.chars().nth(0) {
        Some(it) => it,
        None => return (None, 0),
    };
    if !potential.is_digit(10) {
        return (None, 0);
    }
    let mut digit = potential.to_digit(10).unwrap();

    let potential = match input.chars().nth(1) {
        Some(it) => it,
        None => return (Some(digit), 1),
    };
    if !potential.is_digit(10) {
        return (Some(digit), 1);
    }
    digit *= 10;
    digit += potential.to_digit(10).unwrap();

    let potential = match input.chars().nth(2) {
        Some(it) => it,
        None => return (Some(digit), 2),
    };
    if !potential.is_digit(10) {
        return (Some(digit), 2);
    }
    digit *= 10;
    digit += potential.to_digit(10).unwrap();
    return (Some(digit), 3);
}
