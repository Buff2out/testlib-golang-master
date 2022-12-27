#include <checkers/testlib.h>
#include <iostream>
#include <vector>
#include <set>
#include <algorithm>
#include <map>
#include <set>
#include <string>
#include <sstream>
#include <regex>

using namespace NTestlib;

int main(int argc, char *argv[])
{
    InitChecker(argc, argv);
    std::vector<bool> a;
    std::vector<std::string> b;
    while (File.HasInput())
        b.push_back(File.ReadLine());
    int cnt1 = 0;
    while (Ans.HasInput())
    {
        int x = Ans.ReadInt();
        if (x)
            cnt1++;
        a.push_back((x));
    }
    int cnt = 0;
    while (Out.HasInput())
    {
        int x = Out.ReadInt(1, a.size());
        int y = Out.ReadInt(1, b[x - 1].length() - 13);
        if (!a[x - 1])
            QuitWith(WA, "Merlin was silent");
        if (b[x - 1].substr(y - 1, 14) != "Avada-ke-davra")
            QuitWith(WA, "Merlin didn't say that");
        cnt++;
    }
    if (cnt == cnt1)
        QuitWith(AC, "OK");
    else
        QuitWith(WA, "Merlin said more");
}