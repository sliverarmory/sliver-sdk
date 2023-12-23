mod parser;
mod {{.ExtensionName}};

// Use this file for local debugging
// cargo run -- PATH_TO_FILE

fn main() {
    let cmd_args = std::env::args().collect::<Vec<String>>();
    let file_path = cmd_args[1].clone();
    // read args from file
    let args = std::fs::read(file_path).unwrap();
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
    let result = {{.ExtensionName}}::do_stuff(arg1, &arg2).unwrap();
    println!("result: {}", result);
}
