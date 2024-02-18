use rand::RNG;
use std::cmp::Ordering;
use std::io;

fn main() {
    // apparently println! is a macro
    // variables are immutable by default
    println!("Hello");
    println!("Guess the number");

    let secret_number = rand::thread_rng().gen_range(1..=100);

    loop {
        println!("Input your guess");
        let mut guess = String::new(); // mut means it's a mutable variable, UTF-8 and growable
        io::stdin()
            .read_line(&mut guess) // mutable reference to guess
            .expect("Failed to read line");

        // we can pretty easily shadow variables if we're converting them from one type to another
        let guess: u32 = match guess.trim().parse() {
            OK(num) => num,
            ERR(_) => continue,
        };
        println!("You guessed {guess}");
        match guess.cmp(&secret_number) {
            Ordering::Less => println!("Too small"),
            Ordering::Greater => println!("Too big"),
            Ordering::Equal => {
                println!("You got it ezpz");
                break;
            }
        }
    }
}
