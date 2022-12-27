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
    if (n1 == 0 && n2 == 0)
        QuitWith(AC, "OK");
    if (n1 > n2)
        QuitWith(WA, "WA");
    if (n1 == -1 && n2 == -1)
        QuitWith(AC, "OK");
    int n = 8;
    int m = 8;
    int x = File.ReadInt();
    int y = File.ReadInt();
    int x5 = File.ReadInt();
    int y5 = File.ReadInt();
    std::vector<int> x1, y1, x2, y2;
    x1.push_back(x);
    y1.push_back(y);
    x2.push_back(x5);
    y2.push_back(y5);
    for (int i = 0; i < n1; i++)
    {
        int x3 = Out.ReadInt();
        int y3 = Out.ReadInt();
        int x4 = Out.ReadInt();
        int y4 = Out.ReadInt();
        x1.push_back(x3);
        y1.push_back(y3);
        x2.push_back(x4);
        y2.push_back(y4);
    }
    if (x1.back() != x2.back() || y1.back() != y2.back())
    {
        QuitWith(WA, "WA1");
    }
    for (int i = 0; i < x1.size() - 1; i++)
    {
        int dx1 = x1[i + 1] - x1[i];
        int dy1 = y1[i + 1] - y1[i];
        int dx2 = x2[i + 1] - x2[i];
        int dy2 = y2[i + 1] - y2[i];
        int dx[8] = {1, 1, -1, -1, 2, 2, -2, -2};
        int dy[8] = {2, -2, 2, -2, 1, -1, 1, -1};
        bool f1 = 0;
        bool f2 = 0;
        for (int j = 0; j < 8; j++)
        {
            if (dx1 == dx[j] && dy1 == dy[j])
                f1 = 1;
            if (dx2 == dx[j] && dy2 == dy[j])
                f2 = 1;
        }
        if (!f1 || !f2)
            QuitWith(WA, "WA2");
    }
    QuitWith(AC, "OK");
}