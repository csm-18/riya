import Foundation

let RIYA_VERSION = "Riya 0.1.0"

let args = Array(CommandLine.arguments.dropFirst())

if args.isEmpty {
    print(RIYA_VERSION)
    print("")
    print("A simple programming language that compiles to JS.")
    print("")
    print("For help: ")
    print("  riya help")        
}else if args.count == 1 {
    switch args[0] {
    case "version":
        print(RIYA_VERSION)
    case "help":
        print("Riya commands:")
        print("  riya                  ->  Print about message")
        print("  riya <filename.riya>  ->  Compile .riya file to .js file")
        print("  riya version          ->  Print riya version")
        print("  riya help             ->  Print riya commands list")
    default:
        print("Error: Unknown command!")
        exit(0)
    }
}else{
    print("Error: Unknown command!")
    exit(0)
}
