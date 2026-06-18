import Foundation

public extension String {
    
    // 1. Direct indexing: text[3] -> Returns "s" (String)
    // Note: Matches standard array behavior, will crash if index is out of bounds
    subscript(index: Int) -> String {
        let charIndex = self.index(startIndex, offsetBy: index)
        return String(self[charIndex])
    }

    // 2. Labeled safe indexing: text[at: 3] -> Returns Optional("s") (String?)
    subscript(at index: Int) -> String? {
        guard index >= 0 && index < count else { return nil }
        let charIndex = self.index(startIndex, offsetBy: index)
        return String(self[charIndex])
    }
    
    // 3. Safe method: text.at(3) -> Returns Optional("s") (String?)
    func at(_ index: Int) -> String? {
        return self[at: index]
    }

    // 4. Closed Range Slicing: text[1...3] -> Returns "wim" (String)
    subscript(range: ClosedRange<Int>) -> String {
        let start = index(startIndex, offsetBy: max(0, range.lowerBound))
        let end = index(startIndex, offsetBy: min(count, range.upperBound))
        guard start <= end else { return "" }
        return String(self[start...end])
    }
    
    // 5. Partial Range Slicing: text[2...] -> Returns "ift" (String)
    subscript(range: PartialRangeFrom<Int>) -> String {
        let start = index(startIndex, offsetBy: max(0, range.lowerBound))
        return String(self[start...])
    }
}