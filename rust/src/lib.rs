mod nist;

use nist::sp800108::{stuff, CounterLength};

pub fn do_stuff() {
    stuff(CounterLength::U8, "".to_string())
}
