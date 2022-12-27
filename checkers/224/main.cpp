#include <checkers/testlib.h>
#include <iostream>
#include <fstream>
#include <stdlib.h>
#include <vector>
#include <iomanip>
#include <sstream>

using namespace NTestlib;

double func(double a, double b, double c, double d, double x)
{
    return x * x * x * a + x * x * b + x * c + d;
}
double diff(double a, double b, double c, double d, double x)
{
    return 3 * x * x * a + 2 * x * b + c;
}
int main(int argc, char *argv[])
{
    InitInteractor(argc, argv);
    double a = File.ReadDouble();
    double b = File.ReadDouble();
    double c = File.ReadDouble();
    double d = File.ReadDouble();
    double an = File.ReadDouble();
    an *= -1;
    int cnt = 0;
    while (true)
    {
        if (cnt > 50)
            QuitWith(WA, "Too many requests");
        char requestType = Out.ReadChar();
        switch (requestType)
        {
        case '!':
        {
            double x = Out.ReadDouble();
            if (std::min(std::min(std::abs(x - an), std::abs(func(a, b, c, d, x) - func(a, b, c, d, an))), std::abs(diff(a, b, c, d, x) - diff(a, b, c, d, an))) > 1e-6)
            {
                std::stringstream msg;
                msg << "Wrong answer. x: " << x << " ans: " << an << " abs: " << std::abs(x - an);
                QuitWith(WA, msg.str());
            }
            else
            {
                QuitWith(AC, "OK");
            }
        }
        case '?':
        {
            double x = Out.ReadDouble();
            std::cout << std::setprecision(20) << func(a, b, c, d, x) << " " << diff(a, b, c, d, x) << std::endl;
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
        cnt++;
        Out.Match('\n');
    }
    std::cerr << "INT\n";
    return 0;
}