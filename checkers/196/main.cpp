#include <checkers/testlib.h>
#include <iostream>
#include <string>
#include <vector>
#include <algorithm>

using namespace NTestlib;

int main(int argc, char *argv[])
{
    InitChecker(argc, argv);
    int n = File.ReadInt();
    int start = File.ReadInt();
    std::vector<std::vector<int>> g(n, std::vector<int>(n, 0));
    for (int i = 0; i < n; i++)
        for (int j = 0; j < n; j++)
            g[i][j] = File.ReadInt();
    int x = Out.ReadInt(start, start);
    int y;
    std::vector<int> visited(n, 0);
    for (int i = 0; i < n; i++)
    {
        y = Out.ReadInt(1, n);
        if (!g[x - 1][y - 1])
        {
            std::stringstream msg;
            msg << "WA two unconnected vertices " << x << " " << y;
            QuitWith(WA, msg.str());
        }
        visited[y - 1]++;
        x = y;
    }
    if (y != start)
        QuitWith(WA, "WA last != start");
    for (int i = 0; i < n; i++)
        if (visited[i] != 1)
            QuitWith(WA, "vertice visited not 1 time");
    QuitWith(AC, "OK");
}