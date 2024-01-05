mod {{.ExtensionName}};
mod parser;

const SUCCESS: u32 = 0;
const ERROR: u32 = 1;

#[no_mangle]
pub extern "C" fn start(
    data: *mut u8,
    size: u64,
    result_callback: extern "C" fn(*mut u8, u64),
) -> u32 {
    // create a new ArgumentParser
    let mut dp = parser::ArgumentParser::from_ptr(data, size);
    // create a new OutputBuffer
    let mut ob = parser::OutputBuffer::from_callback(result_callback);
    let firt_arg = match dp.get_u32() {
        Ok(v) => v,
        Err(e) => {
            ob.send_error(&format!("failed to get first argument: {}", e));
            ob.flush();
            return ERROR;
        }
    };
    let second_arg = match dp.get_string() {
        Ok(v) => v,
        Err(e) => {
            ob.send_error(&format!("failed to get second argument: {}", e));
            ob.flush();
            return ERROR;
        }
    };
    // Append data to output buffer
    ob.send_output(format!("first arg: {}", firt_arg).as_str());
    ob.send_output(format!("second arg: {}", second_arg).as_str());

    // call do_stuff from myextension.rs
    let result = match {{.ExtensionName}}::do_stuff(firt_arg, second_arg.as_str()) {
        Ok(v) => v,
        Err(e) => {
            ob.send_error(&format!("failed to call do_stuff: {}", e));
            ob.flush();
            return ERROR;
        }
    };
    // Append data to output buffer
    ob.send_output(format!("result: {}", result).as_str());
    // Send data back to implant
    ob.flush();
    // return success
    SUCCESS
}
