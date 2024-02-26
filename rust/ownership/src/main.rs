fn main() {
    // ownership is how rust manages memory
    // main point of ownership is to manage the heap data

    // Basically every value in Rust has an owner and there can only be one owner at a time. When
    // the owner goes out of scope,the value is dropped

    // this variable goes into the heap since we want to be able to expand it and allocate more
    // memory
    let mut s = String::from("hello");
    s.push_str(" world");
    println!("{}", s)
}
