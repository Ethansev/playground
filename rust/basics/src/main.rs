fn main() {
    let mut x = 5;
    println!("The value of x is {x}");
    x = 6;
    println!("The value of new x is now {x}");

    const HELLO: u32 = 6 * 60 * 60;
    println!("HELLO_WORLD {HELLO}");

    // Arrays
    let array_test = [1, 2, 3, 4, 5];

    let another_array: [u8; 5] = [1, 2, 3, 4, 5];

    new_function(5);
    let new_num = return_five();

    println!("{new_num}");

    testing_if(9);

    // look into expressions vs statements in rust

    // rust has 3 loops: loop, while, for
    // loops go forever until we use break. there's also continue

    for num in array_test {
        println!("{num}")
    }
}

fn new_function(num: i32) {
    println!("{num}")
}

fn return_five() -> i32 {
    5
}

fn testing_if(num: i32) {
    if (num < 10) {
        println!("Num is less than 10")
    } else if (num > 10) {
        println!("Num is greater than 10")
    } else {
        println!("who knows")
    }
}

fn testing_assignment_if() {
    let bool = true;
    let number = if bool { 10 } else { 5 };
    println!("{number}")
}

fn print_in_reverse() {
    // printing 1-4 inclusively in reverse
    for num in (1..4).rev() {
        println!("{num}")
    }
}
