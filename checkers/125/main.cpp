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
    std::vector<int> a;
    for (int i = 0; i < n; i++)
    {
        a.push_back(File.ReadInt());
    }
    int m = File.ReadInt();
    for (int i = 0; i < m; i++)
    {
        int index = Out.ReadInt(), target = File.ReadInt(), answ = Ans.ReadInt();
        if (answ == index)
            continue;
        if (index < 0 || index > n)
        {
            QuitWith(WA, "Incorrect user output");
        }
        if (a[index - 1] != target)
        {
            QuitWith(WA, "Incorrect user output");
        }
    }
    QuitWith(AC, "OK");
}