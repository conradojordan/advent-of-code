use std::fs;

fn first_and_last_digit(line: &str) -> u32 {
    let mut digits: [char; 2] = ['a'; 2];
    let mut first = true;

    for c in line.chars() {
        if c.is_digit(10) {
            if first {
                digits[0] = c;
                first = false;
            } else {
                digits[1] = c;
            }
        }
    }
    if digits[1] == 'a' {
        digits[1] = digits[0];
    }

    format!("{}{}", digits[0], digits[1]).parse().unwrap()
}

fn first_and_last_digit_part_two(line: &str) -> u32 {
    let mut parsed_digits: Vec<char> = Vec::new();
    let line_length: usize = line.len();

    for (index, ch) in line.char_indices() {
        match ch {
            '0'..='9' => {
                parsed_digits.push(ch);
            }
            'e' => {
                if line_length - index >= 5 {
                    if &line[index..index + 5] == "eight" {
                        parsed_digits.push('8')
                    }
                }
            }
            'f' => {
                if line_length - index >= 4 {
                    if &line[index..index + 4] == "five" {
                        parsed_digits.push('5')
                    }
                    if &line[index..index + 4] == "four" {
                        parsed_digits.push('4')
                    }
                }
            }
            'n' => {
                if line_length - index >= 4 {
                    if &line[index..index + 4] == "nine" {
                        parsed_digits.push('9')
                    }
                }
            }
            'o' => {
                if line_length - index >= 3 {
                    if &line[index..index + 3] == "one" {
                        parsed_digits.push('1')
                    }
                }
            }
            's' => {
                if line_length - index >= 3 {
                    if &line[index..index + 3] == "six" {
                        parsed_digits.push('6')
                    }
                }
                if line_length - index >= 5 {
                    if &line[index..index + 5] == "seven" {
                        parsed_digits.push('7')
                    }
                }
            }
            't' => {
                if line_length - index >= 3 {
                    if &line[index..index + 3] == "two" {
                        parsed_digits.push('2')
                    }
                }
                if line_length - index >= 5 {
                    if &line[index..index + 5] == "three" {
                        parsed_digits.push('3')
                    }
                }
            }
            _ => (),
        }
    }

    format!(
        "{}{}",
        parsed_digits[0],
        parsed_digits[parsed_digits.len() - 1]
    )
    .parse()
    .unwrap()
}

fn main() {
    println!("{:-^30}", " Day 1 ");
    let input_path = "../../inputs/day1";
    let file_content = fs::read_to_string(input_path).unwrap();
    let input = file_content.trim_end();

    let mut result: u32 = 0;
    for line in input.split("\n") {
        result += first_and_last_digit(line);
    }
    println!("Result of part 1 is {}", result);

    let mut result: u32 = 0;
    for line in input.split("\n") {
        result += first_and_last_digit_part_two(line);
    }
    println!("Result of part 2 is {}", result);
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_first_and_last_digit() {
        assert_eq!(first_and_last_digit("1abc2"), 12);
        assert_eq!(first_and_last_digit("pqr3stu8vwx"), 38);
        assert_eq!(first_and_last_digit("a1b2c3d4e5f"), 15);
        assert_eq!(first_and_last_digit("treb7uchet"), 77);
    }

    #[test]
    fn test_first_and_last_digit_part_two() {
        assert_eq!(first_and_last_digit_part_two("two1nine"), 29);
        assert_eq!(first_and_last_digit_part_two("eightwothree"), 83);
        assert_eq!(first_and_last_digit_part_two("abcone2threexyz"), 13);
        assert_eq!(first_and_last_digit_part_two("xtwone3four"), 24);
        assert_eq!(first_and_last_digit_part_two("4nineeightseven2"), 42);
        assert_eq!(first_and_last_digit_part_two("zoneight234"), 14);
        assert_eq!(first_and_last_digit_part_two("7pqrstsixteen"), 76);
        assert_eq!(first_and_last_digit_part_two("1bvjgdjlll"), 11);
        assert_eq!(first_and_last_digit_part_two("cf8"), 88);
        assert_eq!(first_and_last_digit_part_two("dzrt197twonine"), 19);
        assert_eq!(first_and_last_digit_part_two("7two8sevencvfjhqmdtfone"), 71);
    }
}
