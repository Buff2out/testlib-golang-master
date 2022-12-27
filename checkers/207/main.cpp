#include <checkers/testlib.h>
#include <iostream>
#include <vector>
#include <set>
#include <algorithm>
#include <map>
#include <set>
#include <string>
#include <sstream>
#include <regex>

using namespace NTestlib;

int main(int argc, char *argv[])
{
    InitChecker(argc, argv);
    int a[100][100];
    int cnt = 0;
    std::cout << 0;
    while (Ans.HasInput())
    {
        int x = Ans.ReadInt();
        int y = Ans.ReadInt();
        if (x != y)
        {
            a[x][y] = 1;
            a[y][x] = 1;
            cnt++;
        }
    }
    long double cnt1 = 0;
    while (Out.HasInput())
    {
        int x = Out.ReadInt(1, 100);
        int y = Out.ReadInt(1, 100);
        x--;
        y--;
        if (a[x][y] == 1)
            cnt1 += 1;
        else
            cnt1 -= 0.5;
    }
    if ((cnt1 * 1L) / (cnt * 1L) > 0.5)
    {
        std::stringstream msg;
        msg << "OK " << cnt << " " << cnt1;
        QuitWith(AC, msg.str());
    }
    else
    {
        std::stringstream msg;
        msg << "WA " << int((cnt1 * 100) / cnt);
        QuitWith(WA, msg.str());
    }
}