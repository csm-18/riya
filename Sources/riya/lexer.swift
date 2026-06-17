import Foundation

enum TokenType {
    case string_literal
    case keyword
    case identifier
    case number_literal
    case left_paren
    case right_paren
    case newline
}

struct Token {
    var type: TokenType
    var value: String
    var index: Int
}


func lex(filename: String,text : String) -> [Token] {
    var tokens: [Token] = []

    var x = 0
    while x < text.count {
        
        x+=1
    }
    return tokens
}