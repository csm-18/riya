import Foundation

let RIYA_VERSION = "riya 0.1.0"

let args = Array(CommandLine.arguments.dropFirst())

if args.isEmpty {
    print(RIYA_VERSION)
    print("")
    print("A simple programming language that compiles to JS.")
    print("")
    print("For help: ")
    print("  riya help")        
}else if args.count == 1 {
    if args[0] == "version" {
        print(RIYA_VERSION)
    }
}
