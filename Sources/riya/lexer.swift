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
        if text[x] == "(" {
            let token = Token(type: TokenType.left_paren,value: "(",index: x)
            tokens.append(token)
        }else if text[x] == ")" {
            let token = Token(type: TokenType.right_paren, value: ")", index: x)
            tokens.append(token)            
        }else if text[x] == "\""{
            var end_quote = false
            var y = x+1
            while y < text.count {
                if text[y] == "\"" {
                    end_quote = true
                    break
                }
                y+=1
            }

            if end_quote {
                let token = Token(type: TokenType.string_literal, value: text[x+1...y-1], index: x)
                tokens.append(token)
                x = y+1
                continue
            }else{
                printError(message: "Unclosed string literal!", at: x, in: text, filename: filename)
            }
        }
        x+=1
    }
    return tokens
}