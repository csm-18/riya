data class Token(val type: String, val value: String, val line: Int, val column: Int)

class Lexer(private val source: String) {
    private var position = 0
    private var line = 1
    private var column = 1
    
    fun tokenize(): List<Token> {
        val tokens = mutableListOf<Token>()
        
        // TODO: Implement lexical analysis
        // For now, return empty list
        
        return tokens
    }
    
    private fun currentChar(): Char? = 
        if (position < source.length) source[position] else null
}
