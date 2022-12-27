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
    int g[100][100];
    int in[100][100];
    for (int i = 0; i < n; i++)
        for (int j = 0; j < n; j++)
            in[i][j] = File.ReadInt();
    while (Out.HasInput())
    {
        int a = Out.ReadInt(1, n);
        int b = Out.ReadInt(1, n);
        a--;
        b--;
        g[a][b] = 1;
        g[b][a] = 1;
    }
    for (int i = 0; i < n; i++)
        for (int j = 0; j < n; j++)
            if (g[i][j] != in[i][j])
            {
                std::stringstream msg;
                msg << "WA wrong cell " << i << " " << j;
                QuitWith(WA, msg.str());
            }
    QuitWith(AC, "OK");
}
