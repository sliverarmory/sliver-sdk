use std::error::Error;

pub fn do_stuff(int_arg: u32, str_arg: &str) -> Result<String, Box<dyn Error>> {
    let out = format!("[do_stuff] got {} and {}", int_arg, str_arg);
    Ok(out)
}
