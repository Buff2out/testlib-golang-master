#include <checkers/testlib.h>
#include <iostream>
#include <vector>
#include <algorithm>
#include <string>

using namespace NTestlib;

int main(int argc, char *argv[])
{
    InitChecker(argc, argv);
    int n = File.ReadInt();
    //
    std::vector<std::vector<int>> 
    g(n, std::vector<int>(n, 0)); // одна строка
    for (int i = 0; i < n; i++)
        for (int j = 0; j < n; j++)
            g[i][j] = File.ReadInt();
    std::string res = Ans.ReadWord();
    std::string ans = Out.ReadWord();
    if (res != ans)
        QuitWith(WA, "NOPE");
    if (res == "NO")
        QuitWith(AC, "OK");
    int cnt1 = Out.ReadInt(2, 2 * n);
    int x = Out.ReadInt(1, n);
    int cnt = 0;
    for (int i = 0; i < cnt1 - 1; i++)
    {
        int y = Out.ReadInt(1, n);
        if (g[x - 1][y - 1] == 100000)
        {
            std::stringstream msg;
            //
            msg << 
            "WA two unconnected vertices " << 
            x << " " << y; // одна строка
            QuitWith(WA, msg.str());
        }
        cnt += g[x - 1][y - 1];
        x = y;
    }
    if (cnt >= 0)
        QuitWith(WA, "ABOVE ZERO");
    QuitWith(AC, "OK");
}