import std/[os, strutils]


const RIYA_VERSION = "riya 0.1.0"

proc main() =
    let args = commandLineParams() # Get command line arguments

    if args.len == 0:
        echo RIYA_VERSION
        echo ""
        echo "Usage:"
        echo " riya <filename.riya>        - compile .riya file"
        echo " riya run <filename.riya>    - compile and run .riya file"
    elif args.len == 1:
        if args[0].len() > 5 and args[0].endsWith(".riya"):
            echo "Compiling ", args[0]
    elif args.len == 2 and args[0] == "run":
        if args[1].len() > 5 and args[1].endsWith(".riya"):
            echo "Compiling and running ", args[1]
           
            
        


if isMainModule:
    main()
    