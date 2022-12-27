#define _CRT_SECURE_NO_DEPRECATE
#define _USE_MATH_DEFINES
#include <checkers/testlib.h>
#include <iostream>

using namespace NTestlib;

#define mp std::make_pair
int main(int argc, char *argv[])
{
    InitChecker(argc, argv);
    int x1 = File.ReadInt();
    int y1 = File.ReadInt();
    int x2 = File.ReadInt();
    int y2 = File.ReadInt();
    std::pair<int, int> ans11 = mp(x1 + (y2 - y1), y1 - (x2 - x1));
    std::pair<int, int> ans21 = mp(x1 - (y2 - y1), y1 + (x2 - x1));
    std::pair<int, int> ans12 = mp(x2 + (y2 - y1), y2 - (x2 - x1));
    std::pair<int, int> ans22 = mp(x2 - (y2 - y1), y2 + (x2 - x1));
    std::pair<int, int> ans1;
    ans1.first = Out.ReadInt();
    ans1.second = Out.ReadInt();
    std::pair<int, int> ans2;
    ans2.first = Out.ReadInt();
    ans2.second = Out.ReadInt();
    if (((ans1 == ans11 && ans2 == ans12) || (ans1 == ans12 && ans2 == ans11)) || ((ans1 == ans22 && ans2 == ans21) || (ans2 == ans21 && ans1 == ans22)))
        QuitWith(AC, "OK");
    std::stringstream msg;
    msg << "WA ans1: " << ans1.first << ", " << ans1.second << "; ans2: " << ans2.first << ", " << ans2.second << "; ans11: " << ans11.first << ", " << ans11.second << "; ans12: " << ans12.first << ", " << ans12.second << "; ans21: " << ans21.first << ", " << ans21.second << "; ans22: " << ans22.first << ", " << ans22.second;
    QuitWith(WA, msg.str());
}