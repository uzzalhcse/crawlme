
# Crawling Efficiency Comparison

This repository contains code to compare the efficiency of web crawling using two different Go packages: [Playwright](https://github.com/playwright-community/playwright-go) and [Rod](https://github.com/go-rod/rod).

## Introduction

This project aims to evaluate the performance and efficiency of two Go-based web crawling packages by extracting pricing information from a specific web page. The results will help determine which package is more efficient for this type of task.

### Installation

1.  **Clone the repository:**

    `git clone https://github.com/uzzalhcse/crawlme.git
    cd crawlme`

2.  **Install dependencies for Playwright:**

    `cd playwright
    go mod tidy`

3.  **Install dependencies for Rod:**

    `cd ../rod
    go mod tidy`


## Running the Tests

To run the tests for both Playwright and Rod, use the provided script. This script will execute both implementations and save the results in the `data/` directory.


`./scripts/run_tests.sh`

### Explanation of the Script

-   The script navigates to the `playwright` directory and runs the Playwright implementation, saving the output to `data/playwright_results.txt`.
-   It then navigates to the `rod` directory and runs the Rod implementation, saving the output to `data/rod_results.txt`.

## Results

After running the tests, compare the results in the `data/` directory:

-   `data/playwright_results.txt`
-   `data/rod_results.txt`

These files contain the output from each implementation, showing the extracted pricing information and the number of times the price was empty.
