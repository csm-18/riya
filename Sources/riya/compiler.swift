import Foundation

public func Compile(filename: String){
    let text = readFile(path: filename)
    let tokens = lex(filename: filename, text: text)
    print(tokens)
    
}