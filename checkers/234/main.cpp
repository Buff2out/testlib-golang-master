#include <checkers/testlib.h>
#include <iostream>
#include <vector>
#include <set>
#include <algorithm>
#include <map>
#include <set>
#include <string>
#include <sstream>
#include <set>

using namespace NTestlib;

int main(int argc, char *argv[])
{
    InitChecker(argc, argv);
    double cnt = 0;
    int x;
    for (int i = 0; i < 3; i++)
        x = File.ReadInt();
    for (int i = 0; i < x; i++)
    {
        std::string x = Ans.ReadWord();
        std::string y = Out.ReadWord();
        if (x == y)
            cnt++;
    }
    if (cnt / x > 0.8)
    {
        std::stringstream msg;
        msg << "OK " << cnt;
        QuitWith(AC, msg.str());
    }
    else
    {
        std::stringstream msg;
        msg << "WA " << cnt / 50;
        QuitWith(WA, msg.str());
    }
}
