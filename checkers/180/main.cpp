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
    int n1 = Ans.ReadInt(), n2 = Out.ReadInt();
    if (n1 != n2)
        QuitWith(WA, "Wrong answer. Incorrect first value");
    for (int i = 0; i < n1; i++)
    {
        int nn1 = Ans.ReadInt(), nn2 = Out.ReadInt();
        if (nn1 != nn2)
        {
            QuitWith(WA, "Wrong answer. Incorrect value");
        }
        std::set<int> s;
        for (int j = 0; j < nn1; j++)
        {
            s.insert(Ans.ReadInt());
        }
        for (int j = 0; j < nn1; j++)
        {
            if (s.find(Out.ReadInt()) == s.end())
                QuitWith(WA, "Wrong answer. Incorrect value");
        }
    }
    QuitWith(AC, "Full solution");
}