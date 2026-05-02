import java.io.File

fun main(args: Array<String>) {
    if (args.isEmpty()) {
        println("Kotlin Compiler - Entry point")
        println("Usage: ./run <source-file> [options]")
        return
    }
    
    val inputFile = args[0]
    println("Compiling: $inputFile")
    
    // TODO: Implement compiler logic
    val compiler = Compiler()
    compiler.compile(inputFile)
}
