#include <checkers/testlib.h>
#include <iostream>
#include <vector>
#include <string>
#include <set>
#include <algorithm>
#include <map>
#include <string>
#include <sstream>

using namespace NTestlib;

int w[1000][1000], m[1000][1000];
int main(int argc, char *argv[])
{
    InitChecker(argc, argv);
    int n = File.ReadInt();
    int so = Out.ReadInt();
    int sa = Ans.ReadInt();
    for (int i = 0; i < n; i++)
        for (int j = 0; j < n; j++)
        {
            w[i][j] = File.ReadInt();
        }
    for (int i = 0; i < n; i++)
        for (int j = 0; j < n; j++)
        {
            m[i][j] = File.ReadInt();
        }
    int s = 0;
    std::vector<int> uw(n, 0);
    for (int i = 0; i < n; i++)
    {
        int k = Out.ReadInt(1, n);
        k--;
        if (uw[k])
            QuitWith(WA, "WA2");
        uw[k]++;
        s += w[i][k];
        s += m[k][i];
    }
    if (s != so)
        QuitWith(WA, "WA3");
    if (so < sa)
        QuitWith(WA, "WA1");
    QuitWith(AC, "OK");
}