#include <checkers/testlib.h>
#include <iostream>
#include <vector>
#include <algorithm>
#include <string>

using namespace NTestlib;

int main(int argc, char *argv[])
{
    InitChecker(argc, argv);
    int n = File.ReadInt();
    std::vector<std::string> a;
    for (int i = 0; i < n; i++)
    {
        a.push_back(File.ReadWord());
    }
    int m = File.ReadInt();
    for (int i = 0; i < m; i++)
    {
        std::string target = File.ReadWord();
        int answ = Out.ReadInt();
        if ((answ < 1 || answ > n) && (answ != -1))
            QuitWith(PE, "Presentation error. Incorrect index");
        if ((answ == -1) && (std::find(a.begin(), a.end(), target) != a.end()))
            QuitWith(PE, "Presentation error. Can't find word with this index");
        if (answ == -1)
            continue;
        if (a[answ - 1] != target)
            QuitWith(WA, "Incorrect user output");
    }
    QuitWith(AC, "OK");
}
