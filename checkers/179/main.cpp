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
    int n1 = Out.ReadInt(), n2 = Ans.ReadInt(), n = File.ReadInt();
    if (n1 == n2 && (n1 == 0 || n1 == -1))
    {
        if (Out.HasInput())
        {
            QuitWith(PE, "" Too many input data "");
        }
        QuitWith(AC, "" OK "");
    }
    if ((n1 == -1) ^ (n2 == -1))
    {
        std::stringstream msg;
        msg << "" WA "";
        QuitWith(WA, msg.str());
    }
    if (n1 > n2)
    {
        std::stringstream msg;
        msg << "" WA1 "";
        QuitWith(WA, msg.str());
    }
    std::vector<std::vector<int>> a(n, std::vector<int>(n));
    std::vector<int> path(n1);
    for (int i = 0; i < n; i++)
        for (int j = 0; j < n; j++)
            a[i][j] = File.ReadInt();
    int s = File.ReadInt();
    int t = File.ReadInt();
    for (int i = 0; i < n1 + 1; i++)
    {
        if (i == 0)
        {
            int x = Out.ReadInt(s, s);
            x--;
            path[i] = x;
        }
        else if (i == n1)
        {
            int x = Out.ReadInt(t, t);
            x--;
            path[i] = x;
        }
        else
        {
            int x = Out.ReadInt();
            x--;
            path[i] = x;
        }
    }
    for (int i = 0; i < path.size() - 1; i++)
        if (!(a[path[i]][path[i + 1]]))
        {
            std::stringstream msg;
            msg << "" WA2 "";
            QuitWith(WA, msg.str());
        }
    QuitWith(AC, "" OK "");
}