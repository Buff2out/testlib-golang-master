#include <checkers/testlib.h>
#include <cmath>
#include <iostream>
#include <string>
#include <vector>
#include <algorithm>

using namespace NTestlib;

std::vector<std::pair<int, int>> dots;
int main(int argc, char *argv[])
{
    InitChecker(argc, argv);
    double cnt = 0;
    int n = File.ReadInt();
    int k = File.ReadInt();
    for (int i = 0; i < n; i++)
    {
        int x = File.ReadInt();
        int y = File.ReadInt();
        dots.push_back(std::make_pair(x, y));
    }
    double reference = Ans.ReadDouble();
    double result = Out.ReadDouble();
    if (result > reference * 1.3)
        QuitWith(WA, "Not correct clasterisation");
    double cluster_sum = 0;
    for (int i = 0; i < k; i++)
    {
        double x = Out.ReadDouble();
        double y = Out.ReadDouble();
        int num = Out.ReadInt();
        double tmp = 0;
        for (int j = 0; j < num; j++)
        {
            int cluster = Out.ReadInt() - 1;
            tmp += std::abs(std::sqrt(pow(x - dots[cluster].first, 2) + pow(y - dots[cluster].second, 2)));
        }
        cluster_sum += tmp;
    }
    if (std::abs(result - cluster_sum) < 10)
        QuitWith(AC, "OK");
    std::stringstream msg;
    msg << "WTF. Result:" << result << ", Clister sum:" << cluster_sum;
    QuitWith(WA, msg.str());
}