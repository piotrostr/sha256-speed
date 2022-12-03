use std::{fs::File, io::Read, time::Instant};

use sha3::{Digest, Sha3_256};

fn main() {
    let mut f = File::open("../test.jpg").unwrap();
    let mut buf = Vec::new();
    let num_read = f.read_to_end(&mut buf).unwrap();

    let mut hasher = Sha3_256::new();

    // create a SHA3-256 object
    let now = Instant::now();

    // write input message
    hasher.update(buf);
    println!("Sum: 0x{:x}", hasher.finalize());
    println!("Elapsed: {}ms", now.elapsed().as_millis());
    println!("Number of bytes: {}", num_read);
}
