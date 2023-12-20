use std::error::Error;
use std::io::Write;

pub struct ArgumentParser {
    pub original: Vec<u8>,
    pub n: usize,
}

impl ArgumentParser {
    pub fn from_vec(v: Vec<u8>) -> ArgumentParser {
        // Start at 4 to skip the length of the whole buffer
        ArgumentParser { original: v, n: 4 }
    }

    pub fn from_ptr(data: *mut u8, size: u64) -> ArgumentParser {
        let v = unsafe { std::slice::from_raw_parts(data, size as usize) };
        ArgumentParser::from_vec(v.to_vec())
    }

    pub fn get_data_length(&self) -> usize {
        self.original.len() - self.n
    }

    pub fn get_u32(&mut self) -> Result<u32, &'static str> {
        if self.get_data_length() < 4 {
            return Err("no more data to return");
        }
        let r = u32::from_le_bytes([
            self.original[self.n],
            self.original[self.n + 1],
            self.original[self.n + 2],
            self.original[self.n + 3],
        ]);
        self.n += 4;
        Ok(r)
    }

    pub fn get_u16(&mut self) -> Result<u16, &'static str> {
        if self.get_data_length() < 2 {
            return Err("no more data to return");
        }
        let r = u16::from_le_bytes([self.original[self.n], self.original[self.n + 1]]);
        self.n += 2;
        Ok(r)
    }

    pub fn get_string(&mut self) -> Result<String, Box<dyn Error>> {
        // get_data
        let out_str = self.get_data()?;
        let decoded = String::from_utf8(out_str)?;
        Ok(decoded)
    }

    pub fn get_wstring(&mut self) -> Result<String, Box<dyn Error>> {
        let rb = self.get_data()?;
        // convert to [u16]
        let decoded: Vec<u16> = rb
            .chunks_exact(2)
            .into_iter()
            .map(|a| u16::from_ne_bytes([a[0], a[1]]))
            .collect();
        let decoded = decoded.as_slice();
        let decoded = String::from_utf16_lossy(decoded);
        Ok(decoded)
    }

    pub fn get_data(&mut self) -> Result<Vec<u8>, Box<dyn Error>> {
        if self.get_data_length() < 4 {
            return Err("no more data to return".into());
        }
        //extract the length
        let l = u32::from_le_bytes(self.original[self.n..self.n + 4].try_into()?);
        //increment n
        self.n += 4;
        //copy to a new buffer to avoid mutating the underlying state
        if self.get_data_length() < l as usize {
            return Err("no more data to return".into());
        }
        let mut rb = vec![0; l as usize];
        rb.copy_from_slice(&self.original[self.n..self.n + l as usize]);
        //increment n
        self.n += l as usize;

        Ok(rb)
    }
}

pub struct OutputBuffer {
    b: Vec<u8>,
    done: bool,
    callback: extern "C" fn(*mut u8, u64),
}

impl OutputBuffer {
    pub fn from_callback(callback: extern "C" fn(*mut u8, u64)) -> OutputBuffer {
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
