import Foundation

public func readFile(path: String) -> String {
    do {
        return try String(contentsOfFile: path, encoding: .utf8)
    } catch {
        print("Error reading file: \(path)")
        exit(0)
    }
}