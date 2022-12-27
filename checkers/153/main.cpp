#include <checkers/testlib.h>
#include <iostream>
#include <vector>
#include <algorithm>
#include <string>

using namespace NTestlib;

int main(int argc, char *argv[])
{
    InitChecker(argc, argv);
    int n = File.ReadInt(), w = File.ReadInt();
    std::vector<int> a;
    std::vector<int> b;
    std::vector<bool> c(n, false);
    for (int i = 0; i < n; i++)
    {
        a.push_back(File.ReadInt());
    }
    for (int i = 0; i < n; i++)
    {
        b.push_back(File.ReadInt());
    }
    int rescnt = 0;
    while (Ans.HasInput())
    {
        int x = Ans.ReadInt(1, n);
        rescnt += a[x - 1];
    }
    int oufcnt = 0, oufw = 0;
    while (Out.HasInput())
    {
        int x = Out.ReadInt(1, n);
        if (c[x - 1])
        {
            QuitWith(WA, "Same ingot");
        }
        c[x - 1] = true;
        oufcnt += a[x - 1];
        oufw += b[x - 1];
        if (oufw > w)
        {
            QuitWith(WA, "Rucksack overflow");
        }
    }
    if (oufcnt < rescnt)
    {
        QuitWith(WA, "Wrong answer");
    }
    if (oufcnt == rescnt)
    {
        QuitWith(AC, "Full solution");
    }
    if (oufcnt > rescnt)
    {
        QuitWith(EF, "Contestant output better than jury's");
    }
}