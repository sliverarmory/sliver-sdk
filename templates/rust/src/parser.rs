use std::io::Write;
use std::os::raw::c_void;
use std::ptr;

pub struct DataParser {
    original: Vec<u8>,
    n: usize,
}

impl DataParser {
    pub fn get_data_length(&self) -> usize {
        self.original.len() - self.n
    }

    pub fn get_next_u16(&mut self) -> Result<u16, &'static str> {
        if self.get_data_length() < 2 {
            return Err("no more data to return");
        }
        let r = u16::from_le_bytes([self.original[self.n], self.original[self.n + 1]]);
        self.n += 2;
        Ok(r)
    }

    pub fn get_string(&mut self, length: usize) -> Result<String, &'static str> {
        if self.get_data_length() < length {
            return Err("no more data to return");
        }
        let result = str::from_utf8(&self.original[self.n..self.n + length]);
        match result {
            Ok(v) => {
                self.n += length;
                Ok(v.to_string())
            }
            Err(_e) => Err("failed to convert bytes to string"),
        }
    }

    pub fn get_data(&mut self, length: usize) -> Result<Vec<u8>, &'static str> {
        if self.get_data_length() < length {
            return Err("no more data to return");
        }
        let result = self.original[self.n..self.n + length].to_vec();
        self.n += length;
        Ok(result)
    }
}

pub struct OutputBuffer {
    b: Vec<u8>,
    done: bool,
    callback: extern "C" fn(*mut u8, u64),
}

impl OutputBuffer {
    pub fn new(callback: extern "C" fn(*mut u8, u64)) -> OutputBuffer {
        OutputBuffer {
            b: Vec::new(),
            done: false,
            callback,
        }
    }

    pub fn send_output(&mut self, data: &str) {
        self.b.write_all(data.as_bytes()).unwrap();
        self.b.write_all(b"\n").unwrap();
    }

    pub fn send_error(&mut self, err: &str) {
        self.send_output(&format!("error: {}", err));
    }

    pub fn flush(&mut self) {
        if self.done {
            return;
        }
        (self.callback)(self.b.as_mut_ptr(), self.b.len() as u64);
        self.done = true;
    }
}
