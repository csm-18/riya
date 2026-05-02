import java.io.File

class Compiler {
    fun compile(filepath: String) {
        val file = File(filepath)
        
        if (!file.exists()) {
            println("Error: File not found: $filepath")
            return
        }
        
        println("Reading source: ${file.absolutePath}")
        val source = file.readText()
        
        // TODO: Implement lexer, parser, and code generator
        val tokens = tokenize(source)
        println("Tokens: $tokens")
    }
    
    private fun tokenize(source: String): List<String> {
        // Simple tokenizer - replace with actual implementation
        return source.split(Regex("\\s+")).filter { it.isNotEmpty() }
    }
}
