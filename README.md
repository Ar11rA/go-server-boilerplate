# Structuring go web servers!

## Why and What?

There is no one way of creating a web server in go. So many libraries, so many design patterns...
So as a curious developer, just had to explore!

I will be covering some design patterns we can follow in go.

## Installation

Install `dep` on your machine and run `dep ensure` 

Install postgres and run the following:

```
createdb temp;
create table users (id serial, username text, name text, email text, phone text);
```

To get the server up and running:
go run main.go