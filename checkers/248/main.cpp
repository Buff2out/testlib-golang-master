#include <checkers/testlib.h>
#define _CRT_SECURE_NO_DEPRECATE
#define _USE_MATH_DEFINES
#include <iostream>
#include <vector>
#include <algorithm>
#define eps 1e-6
using namespace NTestlib;
#define mp std::make_pair
typedef long double ld;

using namespace NTestlib;

int main(int argc, char *argv[])
{
    InitChecker(argc, argv);
    ld x1 = File.ReadDouble();
    ld y1 = File.ReadDouble();
    ld x2 = File.ReadDouble();
    ld y2 = File.ReadDouble();
    std::pair<ld, ld> ans11 = mp((x1 + x2) / 2 + (y1 - y2) / 2, (y1 + y2) / 2 + (x2 - x1) / 2);
    std::pair<ld, ld> ans12 = mp((x1 + x2) / 2 + (y2 - y1) / 2, (y1 + y2) / 2 + (x1 - x2) / 2);
    std::pair<ld, ld> ans1;
    ans1.first = Out.ReadDouble();
    ans1.second = Out.ReadDouble();
    std::pair<ld, ld> ans2;
    ans2.first = Out.ReadDouble();
    ans2.second = Out.ReadDouble();
    if ((ans1 == ans11 && ans2 == ans12) || (ans1 == ans12 && ans2 == ans11))
        QuitWith(AC, "OK");
    QuitWith(WA, "WA");
}
