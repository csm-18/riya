#include <iostream>
#include <string>
#include <vector>

std::vector<std::string> getArgsFrom(int argc, char* argv[]);

int main(int argc, char* argv[]) {
    //cli args
    std::vector<std::string> args = getArgsFrom(argc, argv);
    for (const std::string &arg : args) {
        std::cout << arg << std::endl;
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