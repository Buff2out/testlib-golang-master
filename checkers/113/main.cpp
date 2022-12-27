#include <checkers/testlib.h>
#include <cmath>
#include <iostream>
#include <vector>
#include <algorithm>
#include <string>

using namespace NTestlib;

int main(int argc, char *argv[])
{
    InitChecker(argc, argv);
    int n = File.ReadInt();
    std::vector<std::vector<int>> t(3);
    for (int i = 1; i <= n; i++)
        t[0].push_back(n - i + 1);
    for (int i = 0; i < pow(2, n) - 1; i++)
    {
        int a = Out.ReadInt(1, n), b = Out.ReadInt(1, 3), c = Out.ReadInt(1, 3);
        if (t[b - 1].back() != a || !t[c - 1].empty())
            if (t[c - 1].back() < a)
                QuitWith(WA, "Wrong");
        t[c - 1].push_back(a);
        t[b - 1].pop_back();
    }
    if (t[0].size() != 0 || t[1].size() != 0 || t[2].size() != n)
        QuitWith(WA, "Wrong");
    QuitWith(AC, "OK");
}