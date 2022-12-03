# sha256-speed

Comparison of same SHA speed in Go, Rust and Python

## Rules

* SHA-3 standard, 256-bytes hash (easy to check if the implementation is exactly the same as hashes would differ lol)
* Built-in libraries, no boilerplate
* Benchmark the hashing itself as well as the time of the execution using [hyperfine](github.com/sharkdp/hyperfine)

Limitations for now is that I am only going to be running locally on an M1 chip, which is aarch64. On amd64 there could be slightly different results.

The test image used is in the repository, it's [test.jpg](https://github.com/piotrostr/sha256-speed/blob/main/test.jpg).

It's hash is `0xf99e4c9afa9903b4399d044e0252520e68446d0036a2dfe54a692fb8e7c3d0e8`.

## Results
* Go:
<img width="680" alt="image" src="https://user-images.githubusercontent.com/63755291/205464846-4740adb3-fa67-49a3-a998-4aa09dd19d8e.png">
