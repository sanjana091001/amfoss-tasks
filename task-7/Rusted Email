extern crate regex;
use std::io;
fn main()
{
let mut input = String::new();
println!("{}","Enter your mail-id");
io::stdin().read_line(&mut input).ok();
let word=input.trim();         //trimming the string to get rid of whitespaces

use regex::Regex;           //using external crate regex to match with given input 
let formatted=format!("{}","email validification failed");
assert!(Regex::new(r"(?x)
            ^(?P<text>[^@\s]+)@
            ([[:word:]]+\.)*
            [[:word:]]+\.([a-z|A-Z])+$
            ").unwrap().is_match(word),formatted);          //matching regex expression with input

}
