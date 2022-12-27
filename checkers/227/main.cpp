#include <checkers/testlib.h>
#include <cmath>
#include <iostream>
#include <fstream>
#include <stdlib.h>
#include <algorithm>
#include <vector>
#include <iomanip>
#include <sstream>

using namespace NTestlib;

double func(int a, int b, double x, double y)
{
    return pow(x + a, 2) + pow(y + b, 2);
}
double dx(int a, double x)
{
    return 2 * (x + a);
}
double dy(int b, double y)
{
    return 2 * (b + y);
}
double dist(double anX, double anY, double x, double y)
{
    return std::sqrt(pow(anX - x, 2) + pow(anY - y, 2));
}
int main(int argc, char *argv[])
{
    InitInteractor(argc, argv);
    int a = File.ReadInt();
    int b = File.ReadInt();
    double anX = -a;
    double anY = -b;
    int cnt = 0;
    double x, y;
    while (1)
    {
        char requestType = Out.ReadChar();
        if (cnt > 1000000000)
            QuitWith(WA, "Too many requests");
        switch (requestType)
        {
        case '!':
        {
            x = Out.ReadDouble();
            y = Out.ReadDouble();
            if (dist(anX, anY, x, y) > 0.1)
            {
                std::stringstream msg;
                msg << "Incorrect user output. User answer: " << x << ", " << y << ". abs2: " << dist(anX, anY, x, y);
                QuitWith(WA, msg.str());
            }
            else
                QuitWith(AC, "OK");
            break;
        }
        case '?':
        {
            x = Out.ReadDouble();
            y = Out.ReadDouble();
            std::cout << std::setprecision(20) << func(a, b, x, y) << " " << dx(a, x) << " " << dy(b, y) << std::endl;
            break;
        }
        default:
        {
            std::stringstream msg;
            msg << "Wrong query: " << requestType;
            QuitWith(PE, msg.str());
            break;
        }
        }
        Out.Match('\n');
        cnt++;
    }
    std::cerr << "INT\n";
    return 0;
}