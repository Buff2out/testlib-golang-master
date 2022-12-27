#include <checkers/testlib.h>
#include <iostream>
#include <vector>
#include <algorithm>
#include <string>

using namespace NTestlib;

int main(int argc, char *argv[])
{
    InitChecker(argc, argv);
    int a = File.ReadInt(), b = File.ReadInt(), c = File.ReadInt(), d = File.ReadInt();
    double x = Out.ReadDouble();
    if (std::abs(a * x * x * x + b * x * x + c * x + d) < 0.000001)
    {
        QuitWith(AC, "OK");
    }
    QuitWith(WA, "Incorrect user output");
}
