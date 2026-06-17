import Foundation

let RIYA_VERSION = "riya 0.1.0"

let args = CommandLine.arguments[1...]

if args.isEmpty {
    print(RIYA_VERSION)
    print("")
    print("A simple programming language that compiles to JS.")
    print("")
    print("For help: ")
    print("  riya help")        
}
