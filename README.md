# **Recursive Scraper**

This repository contains a web scraper written in Go using the goquery package. Compiles to a command
line tool which takes 2 arguments, a starting URL and a target file location for the json output.

The program concurrently steps through a website form a starting domain, preventing the
same page from being visited twice by cross-referencing after each step, effectively performing a
depth-first search on the provided website.

Only pages under the starting domain extend the search, and the program will continue down each branch
until no new pages are found. once the program has completed all external links are printed to the console
and all pages and resources sitting under the starting domain are written to a json file.

**Build:**
```
go build -o search
```

**Arguments:**
```
-u URL string - the start point for the scrape
-p Path string (optional, defaults to output.json) - the target file to write the json output to
```


**Example:**
```
./scrape -p test.json -u https://jamesmilligan.net 
```


**Further improvements:**

To improve the project beyond this point I suggest the following changes/ additions:
* Further testing
* Addition of an extra flag to prevent the scraper accessing all resources (exclude .png)
