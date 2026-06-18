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


func lineAndColumn(at targetIndex: Int, in text: String) -> (line: Int, column: Int)? {
    // Validate bounds
    guard targetIndex >= 0 && targetIndex <= text.count else {
        return nil // Out of bounds
    }
    
    var currentLine = 1
    var currentColumn = 1
    var currentIndex = 0
    
    // Iterate through the characters of the string
    for char in text {
        if currentIndex == targetIndex {
            return (currentLine, currentColumn)
        }
        
        if char.isNewline {
            currentLine += 1
            currentColumn = 1
        } else {
            currentColumn += 1
        }
        
        currentIndex += 1
    }
    
    // Handle the edge case where the index is at the very end of the string
    if currentIndex == targetIndex {
        return (currentLine, currentColumn)
    }
    
    return nil
}


func printError(message: String, at index: Int, in text: String, filename: String) -> Never {
    // Call the lineAndColumn function to get line and column
    if let position = lineAndColumn(at: index, in: text) {
        // Standard compiler format: filename:line:col: error: message
        print("\(filename):\(position.line):\(position.column): error: \(message)")
        
        // Extract and print the specific line with a pointing caret (^)
        let lines = text.components(separatedBy: .newlines)
        if position.line - 1 < lines.count {
            let errorLine = lines[position.line - 1]
            print(errorLine)
            let padding = String(repeating: " ", count: max(0, position.column - 1))
            print("\(padding)^")
        }
    } else {
        // Fallback for an out-of-bounds index
        print("\(filename): error: \(message) (at invalid index: \(index))")
    }
    
    exit(0)
}