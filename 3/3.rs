use std::env;
use std::fs::File;
use std::io::{self, BufRead, BufReader};

fn main() -> io::Result<()> {
    let filename = env::args().nth(1).expect("Usage: <program> <filename>");

    let file = File::open(&filename)?;
    let reader = BufReader::new(file);

    let mut lines: Vec<String> = Vec::new();
    for line in reader.lines() {
        let line = line?;
        lines.push(line);
    }

    let mut result = 0;
    for (line_index, line) in lines.iter().enumerate() {
        for (index, char) in line.chars().enumerate() {
            if char != '.' && !char.is_numeric() {
                // symbol found!
                // check up, down, left, right, and 4 diagonals
                // if any is numeric, get integer by left and right bounds ('.', '')
                // add integer to result 
            }
        }
    }


    Ok(())
}
