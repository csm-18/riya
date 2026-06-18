import Foundation

public func readFile(path: String) -> String {
    do {
        return try String(contentsOfFile: path, encoding: .utf8)
    } catch {
        print("Error reading file: \(path)")
        exit(0)
    }
}

func printTokens(_ tokens: [Token], title: String = "Lexer Output") {
    let separator = String(repeating: "-", count: 40)
    
    print("\n\(separator)")
    print("  \(title.uppercased()) (\(tokens.count) found)")
    print(separator)
    
    for token in tokens {
        let paddedIndex = String(format: "%3d", token.index)
        
        // This converts the enum case to a String without needing rawValue
        let typeString = String(describing: token.type).uppercased()
        let paddedType = typeString.padding(toLength: 15, withPad: " ", startingAt: 0)
        
        print("[\(paddedIndex)] \(paddedType) -> \"\(token.value)\"")
    }
    
    print("\(separator)\n")
}