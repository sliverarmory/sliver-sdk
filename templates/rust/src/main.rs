mod parser;

// Use this file for local debugging

const data_source: &str = "CHANGE_ME";

fn main() {
    // read args from file
    let args = std::fs::read(data_source).unwrap();
    // create a new DataParser
    let mut parser = parser::ArgumentParser::from_vec(args);
    // get the first argument
    let arg1 = parser.get_u32().unwrap();
    println!("arg1: {}", arg1);
    // get the second argument
    let arg2 = parser.get_string().unwrap();
    println!("arg2: {}", arg2);
    // get the third argument
    let arg3 = parser.get_wstring().unwrap();
    println!("arg3: {}", arg3);
}
