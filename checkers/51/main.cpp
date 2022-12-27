#include <checkers/testlib.h>
#include <iostream>
#include <string>
#include <vector>
#include <algorithm>

using namespace NTestlib;

int main(int argc, char *argv[])
{
    InitChecker(argc, argv);
    std::string out = "", answer = "";
    // File - programm input (from .in file)
    // Out - User output
    // Ans - Correct output (from .out file)
    out = Out.ReadLine();
    answer = Ans.ReadLine();
    while (out[out.size() - 1] == 32)
        out = out.substr(0, out.size() - 1);
    while (answer[answer.size() - 1] == 32)
        answer = answer.substr(0, answer.size() - 1);
    while (out[0] == 32)
        out = out.substr(1, out.size());
    while (answer[0] == 32)
        answer = answer.substr(1, answer.size());
    if (out == answer)
        QuitWith(AC, "Full solution");
    std::stringstream msg;
    msg << "Out " << out.c_str() << ", Correct answer: " << answer.c_str();
    QuitWith(WA, msg.str());
}
