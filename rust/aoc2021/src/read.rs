use std::fs::File;
use std::io::{prelude::*, BufReader};
use std::path::Path;

pub fn from_file(filename: impl AsRef<Path>) -> Vec<String> {
    let file = File::open(filename).expect("no such file");
    let buf = BufReader::new(file);

    // Return the vector of strings
    buf.lines()
	.map(|l| l.expect("could not parse line"))
	.collect()
}
