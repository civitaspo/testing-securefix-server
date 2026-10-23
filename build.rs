use std::process::Command;

fn main() {
    let result = Command::new("git").args(["rev-parse", "HEAD"]).output().unwrap();
    assert!(result.status.success());
    println!("cargo:rustc-env=SOURCE_SHA={}", String::from_utf8(result.stdout).unwrap().trim());
}
