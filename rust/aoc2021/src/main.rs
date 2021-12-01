mod day1;
mod read;

fn main() {
    println!("{}", day1::day1(read::from_file("../../data/day1.txt")));
}
