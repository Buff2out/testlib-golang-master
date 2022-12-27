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

int g[100][100];
int main(int argc, char *argv[])
{
    InitChecker(argc, argv);
    int n = File.ReadInt();
    int m = 0;
    int a[100][100];
    for (int i = 0; i < n; i++)
        for (int j = 0; j < n; j++)
        {
            a[i][j] = File.ReadInt();
            if (a[i][j])
                m++;
        }
    m /= 2;
    int p = Out.ReadInt();
    int anss = Ans.ReadInt();
    if ((anss == -1 && p != -1))
        QuitWith(WA, "-1");
    if (anss == -1 && p == -1)
        QuitWith(AC, "OK");
    p--;
    int f = p;
    for (int i = 0; i < m; i++)
    {
        int r = Out.ReadInt();
        r--;
        if (a[p][r] == 0)
            QuitWith(WA, "WA1");
        if (g[r][p])
            QuitWith(WA, "WA3");
        g[r][p] = 1;
        g[p][r] = 1;
        p = r;
    }
    if (f != p)
        QuitWith(WA, "WA2");
    QuitWith(AC, "OK");
}