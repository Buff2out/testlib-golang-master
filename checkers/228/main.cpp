#include <checkers/testlib.h>
#include <iostream>
#include <fstream>
#include <stdlib.h>
#include <vector>
#include <iomanip>
#include <sstream>

using namespace NTestlib;

#define ld long double
#define eps 0.001
ld f1(ld x, ld y)
{
    return x * x + y * y - 0.1 * std::abs(1 - x) - 0.1 * std::abs(1 - y);
}
ld f2(ld x, ld y)
{
    return 20 * std::abs(x - 50) * std::abs(y - 25) + 10 * (std::abs(x - 10) * std::abs(y - 10) + std::abs(x - 50));
}
int main(int argc, char *argv[])
{
    InitInteractor(argc, argv);
    int f = File.ReadInt();
    int k = 0;
    while (Out.HasInput())
    {
        k++;
        if (k > 100000)
            QuitWith(WA, "Too many requests");
        char requestType = Out.ReadChar();
        ld x = Out.ReadDouble();
        ld y = Out.ReadDouble();
        switch (requestType)
        {
        case '!':
        {
            if (f)
                if (std::abs(x + 0.05) < eps && abs(y + 0.05) < eps)
                    QuitWith(AC, "Full solution");
                else
                {
                    std::stringstream msg;
                    msg << "Incorrect output: " << x << " " << y;
                    QuitWith(WA, msg.str());
                }
            else if (std::abs(50 - x) < eps && std::abs(10 - y) < eps)
                QuitWith(AC, "Full solution");
            else
            {
                std::stringstream msg;
                msg << "Incorrect output: " << x << " " << y;
                QuitWith(WA, msg.str());
            }
        }
        case '?':
        {
            if (f)
                std::cout << f1(x, y) << std::endl;
            else
                std::cout << f2(x, y) << std::endl;
            break;
        }
        default:
        {
            std::stringstream msg;
            msg << "Incorrect input command: " << requestType;
            QuitWith(PE, msg.str());
        }
        break;
        }
        Out.Match('\n');
    }
    QuitWith(EF, "Incorrect: Input after EOF");
    std::cerr << "INT\n";
    return 0;
}