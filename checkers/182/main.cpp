#include <checkers/testlib.h>
#include <iostream>
#include <vector>
#include <set>
#include <algorithm>
#include <map>
#include <string>
#include <sstream>

using namespace NTestlib;

int main(int argc, char *argv[])
{
    InitChecker(argc, argv);
    int n1 = Out.ReadInt();
    int n2 = Ans.ReadInt();
    if (n1 > n2)
    {
        QuitWith(WA, "WA");
    }
    if (n1 == -1)
        if (n2 == -1)
            QuitWith(AC, "OK");
        else
            QuitWith(WA, "WA");
    int n = File.ReadInt();
    int m = File.ReadInt();
    int x1 = File.ReadInt();
    int y1 = File.ReadInt();
    int x2 = File.ReadInt();
    int y2 = File.ReadInt();
    std::vector<int> x, y;
    for (int i = 0; i < n1 + 1; i++)
    {
        int x3 = Out.ReadInt();
        int y3 = Out.ReadInt();
        x.push_back(x3);
        y.push_back(y3);
    }
    if (x[0] != x1 || y[0] != y1 || x[x.size() - 1] != x2 || y[y.size() - 1] != y2)
        QuitWith(WA, "WA1");
    for (int i = 0; i < x.size() - 1; i++)
    {
        int dx = x[i + 1] - x[i], dy = y[i + 1] - y[i];
        if ((!(dx == 2 && dy == 1) && !(dx == 2 && dy == -1) && !(dx == 1 && dy == 2) && !(dx == 1 && dy == -2) &&
             !(dx == -1 && dy == 2) && !(dx == -1 && dy == -2) && !(dx == -2 && dy == -1) && !(dx == -2 && dy == 1)) ||
            x[i] > n || x[i] == 0 || x[i + 1] > n || x[i + 1] == 0 || y[i] > m || y[i] == 0 || y[i + 1] > m || y[i + 1] == 0)
            QuitWith(WA, "WA2");
    }
    QuitWith(AC, "OK");
}