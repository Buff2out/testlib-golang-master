#include <checkers/testlib.h>
#include <iostream>
#include <fstream>
#include <stdlib.h>
#include <vector>
#include <iomanip>
#include <sstream>
#include <cmath>

using namespace NTestlib;

double func(double x)
{
    return 2 * std::cos(0.5 * x);
}
int main(int argc, char *argv[])
{
    InitInteractor(argc, argv);
    int n = File.ReadInt();
    double a = File.ReadDouble();
    double b = File.ReadDouble();
    std::vector<double> ans;
    for (int i = 0; i < n; i++)
        ans.push_back(File.ReadDouble());
    std::cout << std::setprecision(20) << n << " " << a << " " << b << std::endl;
    int cnt = 0;
    while (true)
    {
        cnt++;
        char requestType = Out.ReadChar();
        switch (requestType)
        {
        case '!':
        {
            for (int i = 0; i < n; i++)
            {
                double r = Out.ReadDouble();
                if (std::abs(r - ans[i]) > 0.25)
                {
                    std::stringstream msg;
                    msg << "Wrong answer. Abs: " << r << " " << ans[i] << " number: " << i;
                    QuitWith(WA, msg.str());
                }
                QuitWith(AC, "OK");
            }
        }
        case '?':
        {
            double x = Out.ReadDouble();
            std::cout << std::setprecision(20) << func(x) << std::endl;
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
    std::cerr << "INT\n";
    return 0;
}