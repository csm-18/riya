#include <iostream>
#include <string>
#include <vector>

std::vector<std::string> getArgsFrom(int argc, char* argv[]);


const std::string RIYA_VERSION = "riya 0.1.0\n";
const std::string RIYA_ABOUT = "A high-level programming language\nFor help:\n riya help\n";
const std::string RIYA_HELP = "riya commands:\n"
        " 1. riya                -> print about info\n"
        " 2. riya version        -> print riya version\n"
        " 3. riya help           -> print riya commands\n";

int main(int argc, char* argv[]) {
    //cli args
    std::vector<std::string> args = getArgsFrom(argc, argv);


    if (args.empty()){
        std::cout << RIYA_VERSION << RIYA_ABOUT;
    }else if (args.size() == 1){
        if (args[0] == "version"){
            std::cout << RIYA_VERSION;
        }else if (args[0] == "help"){
            std::cout << RIYA_HELP;
        }
    }
    return 0;
}

// convert cli args into vector<string> without program name
std::vector<std::string> getArgsFrom(int argc, char* argv[]) {
    std::vector<std::string> args;
    args.reserve(argc-1);
    for (int i = 1; i < argc; ++i) {
        args.emplace_back(argv[i]);
    }
    return args;
}