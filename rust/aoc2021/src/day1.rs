/// Day 1
pub fn day1(strs: Vec<String>) -> u16 {
    // Convert to integers
    let numbers: Vec<u16> = strs.iter().flat_map(|x| x.parse()).collect();

    let mut prev: u16 = 0;
    let mut ics: u16 = 0;

    for (i, v) in numbers.iter().enumerate() {
	if i == 0 {
	    prev = *v;
	    continue;
	}

	if v > &prev {
	    ics += 1;
	}
	prev = *v
    }
    ics
}
