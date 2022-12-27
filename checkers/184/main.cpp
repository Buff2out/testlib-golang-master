#include <checkers/testlib.h>
#include <iostream>
#include <vector>
#include <set>
#include <algorithm>
#include <map>
#include <string>
#include <sstream>

using namespace NTestlib;

struct node
{
    int x1, y1, x2, y2;
    bool l;
};

int main(int argc, char *argv[])
{
    InitChecker(argc, argv);
    int n1 = Out.ReadInt();
    int n2 = Ans.ReadInt();
    if (n1 > n2)
        QuitWith(WA, "WA");
    if (n1 == -1 && n2 == -1)
        QuitWith(AC, "OK");
    int n = 8;
    int m = 8;
    int x1 = File.ReadInt();
    int y1 = File.ReadInt();
    int x2 = File.ReadInt();
    int y2 = File.ReadInt();
    node now = {x1, y1, x2, y2, 1};
    bool f = 0;
    for (int i = 0; i < n1; i++)
    {
        int l3 = Out.ReadInt();
        int x3 = Out.ReadInt();
        int y3 = Out.ReadInt();
        l3--;
        if (l3 == 1)
        {
            if (now.l == l3)
                QuitWith(WA, "WA1");
            int dx[8] = {2, 2, -2, -2, 1, 1, -1, -1};
            int dy[8] = {-1, 1, -1, 1, -2, 2, -2, 2};
            bool ff = 0;
            for (int i = 0; i < 8; i++)
                if (now.x2 - x3 == dx[i] && now.y2 - y3 == dy[i])
                    ff = 1;
            if (ff = 0)
                QuitWith(WA, "WA2");
        }
        else
        {
            if (now.l == l3)
                QuitWith(WA, "WA3");
            int dx[8] = {2, 2, -2, -2, 1, 1, -1, -1};
            int dy[8] = {-1, 1, -1, 1, -2, 2, -2, 2};
            bool ff = 0;
            for (int i = 0; i < 8; i++)
                if (now.x1 - x3 == dx[i] && now.y1 - y3 == dy[i])
                    ff = 1;
            if (ff = 0)
                QuitWith(WA, "WA4");
        }
        if (l3 == 1)
        {
            now.x2 = x3;
            now.y2 = y3;
        }
        else
        {
            now.x1 = x3;
            now.y1 = y3;
        }
        now.l = !now.l;
    }
    if (f == 1)
        QuitWith(WA, "WA5");
    if (now.x1 != x2 || now.y1 != y2 || now.x2 != x1 || now.y2 != y1)
        QuitWith(WA, "WA6");
    QuitWith(AC, "OK");
}