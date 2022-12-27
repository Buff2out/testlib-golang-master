#include <checkers/testlib.h>
#include <iostream>
#include <vector>
#include <algorithm>
#include <string>
using namespace NTestlib;

int a[10000][10000];

int main(int argc, char *argv[])
{
    InitChecker(argc, argv);
    int n = File.ReadInt();
    int m = File.ReadInt();
    int px1 = File.ReadInt();
    int py1 = File.ReadInt();
    int px2 = File.ReadInt();
    int py2 = File.ReadInt();
    for (int i = 0; i < n; i++)
        for (int j = 0; j < m; j++)
            a[i][j] = File.ReadInt();
    int dx, dy;
    int ln = Out.ReadInt();
    dx = Out.ReadInt();
    dy = Out.ReadInt();
    if (dx != px1 || dy != py1)
        QuitWith(WA, "Wrong start position");
    for (int i = 0; i < ln - 1; i++)
    {
        dx = Out.ReadInt();
        dy = Out.ReadInt();
        if ((dx <= m && dx >= 0 && dy <= n && dy >= 0) && a[dy - 1][dx - 1] == 0 && ((dy == py1 + 1 && dx == px1) || (dy == py1 - 1 && dx == px1) || (dy == py1 && dx == px1 + 1) || (dy == py1 && dx == px1 - 1)))
        {
            px1 = dx;
            py1 = dy;
        }
        else
        {
            std::stringstream msg;
            msg << "Impossible movement " << px1 << ", " << py1 << ", " << dx << ", " << dy;
            QuitWith(WA, msg.str());
        }
    }
    if (px1 != px2 || py1 != py2)
        QuitWith(WA, "Wrong end position");
    int opCount = Ans.ReadInt();
    int eps = 3;
    if (ln - opCount > eps)
        QuitWith(WA, "Too long way");
    QuitWith(AC, "Full solution");
}