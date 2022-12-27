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

int main(int argc, char *argv[])
{
    InitChecker(argc, argv);
    int counter = 0;
    for (int i = 0; i < 5; i++)
    {
        int a1 = Out.ReadInt();
        int b1 = Out.ReadInt();
        int a2 = Ans.ReadInt();
        int b2 = Ans.ReadInt();
        if (a1 == a2 && b1 == b2)
            counter++;
    }
    if (counter < 4)
    {
        std::stringstream msg;
        msg << "Wrong answer. Only " << counter << " of 5 was right";
        QuitWith(WA, msg.str());
    }
    QuitWith(AC, "Full solution");
}