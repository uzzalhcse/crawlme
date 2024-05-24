#!/bin/bash

# Run Playwright test
echo "Running Playwright tests..."
cd playwright
go run main.go > ../data/playwright_results.txt

# Run Rod test
echo "Running Rod tests..."
cd ../rod
go run main.go > ../data/rod_results.txt

echo "Tests completed. Check the data directory for results."
