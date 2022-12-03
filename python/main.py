#!/usr/bin/env python3

import time
import hashlib


if __name__ == "__main__":
    with open("../test.jpg", "rb") as f:
        contents = f.read()
    start = time.time()
    print(f"Sum: 0x{hashlib.sha3_256(contents).hexdigest()}")
    print(f"Took: {(time.time() - start)*1000:.8f}ms")
    print(f"Input Size: {len(contents)}B", )
