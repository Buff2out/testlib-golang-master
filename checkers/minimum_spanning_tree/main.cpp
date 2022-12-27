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

void dfs(std::vector<std::vector<int>> &a, std::vector<int> &u, int s, int &n)
{
    for (int i = 0; i < n; i++)
        if (a[s][i] && !u[i])
        {
            u[i]++;
            dfs(a, u, i, n);
        }
}

int main(int argc, char *argv[])
{
    InitChecker(argc, argv);
    int n1 = Ans.ReadInt(), n2 = Out.ReadInt();
    if (n1 > n2)
    {
        QuitWith(WA, "Wrong answer. User output greater than correct");
    }
    int n = File.ReadInt();
    std::vector<std::vector<int>> a(n, std::vector<int>(n));
    int s = 0;
    for (int i = 0; i < n; i++)
        for (int j = 0; j < n; j++)
        {
            a[i][j] = Out.ReadInt();
            s += a[i][j];
        }
    if (n2 != s / 2)
        QuitWith(WA, "Wrong answer");
    std::vector<int> u(n, 0);
    u[0]++;
    dfs(a, u, 0, n);
    int count = 0;
    for (int i = 0; i < n; i++)
        if (u[i])
            count++;
    if (count != n)
        QuitWith(WA, "Wrong answer");
    QuitWith(AC, "Full solution");
}