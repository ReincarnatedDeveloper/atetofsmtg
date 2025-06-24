#!/bin/bash
ssserver -c config.json -p 8388 -k your-secure-password -m aes-256-gcm --reuse-port -d start
