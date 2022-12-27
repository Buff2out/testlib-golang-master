
#include <checkers/testlib.h>
#include <iostream>
#include <vector>
#include <set>
#include <algorithm>
#include <map>
#include <set>
#include <string>
#include <sstream>

using namespace NTestlib;

int main(int argc, char *argv[])
{
    InitChecker(argc, argv);
    int n1 = 0, n2 = 0, x;
    std::set<int> s1, s2;
    while (Ans.HasInput())
    {
        x = Ans.ReadInt();
        n1++;
        s1.insert(x);
    }
    while (Out.HasInput())
    {
        x = Out.ReadInt();
        n2++;
        s2.insert(x);
    }
    if (n1 != n2)
    {
        QuitWith(WA, "WA");
    }
    for (std::set<int>::iterator i = s1.begin(); i != s1.end(); i++)
    {
        if (s2.find(*i) == s2.end())
            QuitWith(WA, "WA");
    }
    QuitWith(AC, "Full solution");
}