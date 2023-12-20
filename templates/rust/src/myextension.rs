use std::error::Error;

pub fn do_stuff(int_arg: u32, str_arg: &str) -> Result<String, Box<dyn Error>> {
    Ok("Hello World!".to_string())
}
