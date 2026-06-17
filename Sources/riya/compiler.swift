import Foundation

public func Compile(filename: String){
    let text = readFile(path: filename)
    print(text)
}