#include <iostream>
#include <string>
#include <vector>

std::vector<std::string> getArgsFrom(int argc, char* argv[]);


const std::string RIYA_VERSION = "riya 0.1.0\n";
const std::string RIYA_ABOUT = "A high-level programming language\nFor help:\n riya help\n";

int main(int argc, char* argv[]) {
    //cli args
    std::vector<std::string> args = getArgsFrom(argc, argv);


    if (args.empty()){
        std::cout << RIYA_VERSION<< RIYA_ABOUT;
    }else if (args.size() == 1){
        if (args[0] == "version"){
            std::cout << RIYA_VERSION;
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