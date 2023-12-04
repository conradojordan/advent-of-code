use std::fs;

const MAX_RED: u32 = 12;
const MAX_GREEN: u32 = 13;
const MAX_BLUE: u32 = 14;

#[derive(Debug)]
struct Game<'a> {
    id: u32,
    sets: Vec<Vec<&'a str>>,
}

impl<'a> Game<'a> {
    // "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green"
    fn from_string(input: &str) -> Game {
        let splitted_game: Vec<&str> = input.split(":").collect();
        let sets: Vec<Vec<&str>> = splitted_game
            .last()
            .unwrap()
            .split(";")
            .map(|set| set.split(",").map(|set| set.trim()).collect())
            .collect();
        Game {
            id: splitted_game
                .first()
                .unwrap()
                .split(" ")
                .last()
                .unwrap()
                .parse()
                .unwrap(),
            sets: sets,
        }
    }
}

fn solve_day_2(input: &str) -> (u32, u32) {
    let mut result_part_1: u32 = 0;
    let mut result_part_2: u32 = 0;

    for line in input.lines() {
        let game = Game::from_string(line);
        let mut valid_game = true;
        let mut curr_max_red = 0;
        let mut curr_max_green = 0;
        let mut curr_max_blue = 0;

        for set in game.sets {
            for play in set {
                match play.split(" ").last().unwrap() {
                    "red" => {
                        let red_value = play.split(" ").collect::<Vec<&str>>()[0]
                            .parse::<u32>()
                            .unwrap();
                        if red_value > curr_max_red {
                            curr_max_red = red_value;
                        }
                        if red_value > MAX_RED {
                            valid_game = false;
                        }
                    }
                    "green" => {
                        let green_value = play.split(" ").collect::<Vec<&str>>()[0]
                            .parse::<u32>()
                            .unwrap();
                        if green_value > curr_max_green {
                            curr_max_green = green_value;
                        }
                        if green_value > MAX_GREEN {
                            valid_game = false;
                        }
                    }
                    "blue" => {
                        let blue_value = play.split(" ").collect::<Vec<&str>>()[0]
                            .parse::<u32>()
                            .unwrap();
                        if blue_value > curr_max_blue {
                            curr_max_blue = blue_value;
                        }
                        if blue_value > MAX_BLUE {
                            valid_game = false;
                        }
                    }
                    _ => (),
                }
            }
        }
        if valid_game {
            result_part_1 += game.id
        }
        result_part_2 += curr_max_red * curr_max_green * curr_max_blue;
    }
    (result_part_1, result_part_2)
}

fn main() {
    println!("{:-^30}", " Day 2 ");
    let input_path = "../../inputs/day2";
    let file_content = fs::read_to_string(input_path).unwrap();
    let input = file_content.trim_end();

    let (part_1_result, part_2_result) = solve_day_2(input);
    println!("Result of part 1 is {}", part_1_result);
    println!("Result of part 2 is {}", part_2_result);
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_solve_day_2() {
        assert_eq!(
            solve_day_2(
                "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green"
            ),
            (8, 2286)
        );
    }
}
