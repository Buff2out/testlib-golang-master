#include <checkers/testlib.h>
#include <cmath>
#include <iostream>
#include <vector>
#include <set>
#include <algorithm>
#include <map>
#include <set>
#include <string>
#include <sstream>

using namespace NTestlib;

double dist(std::pair<int, int> a, std::pair<int, int> b)
{
    return std::sqrt(pow(a.first - b.first, 2) + pow(a.second - b.second, 2));
}
int main(int argc, char *argv[])
{
    InitChecker(argc, argv);
    int n = File.ReadInt();
    double answer = Ans.ReadDouble();
    double userAnswer = Out.ReadDouble();
    std::vector<std::pair<int, int>> vec(n);
    for (int i = 0; i < n; i++)
    {
        vec[i].first = File.ReadInt();
        vec[i].second = File.ReadInt();
    }
    std::vector<int> ind(n);
    std::vector<int> u(n, 0);
    for (int i = 0; i < n; i++)
    {
        ind[i] = Out.ReadInt();
        if (u[ind[i] - 1])
        {
            QuitWith(PE, "Value already exist");
        }
        u[ind[i] - 1]++;
    }
    double res = 0;
    for (int i = 0; i < n; i++)
    {
        res += dist(vec[ind[i] - 1], vec[ind[(i + 1) % n] - 1]);
    }
    if (std::abs(res - userAnswer) > 0.5)
    {
        std::stringstream msg;
        msg << "Incorrect output. User answer: " << userAnswer << ". Correct output: " << res;
        QuitWith(WA, msg.str());
    }
    if (userAnswer - answer > answer * 0.1)
    {
        std::stringstream msg;
        msg << "Wrong answer. Absolute diff: " << userAnswer;
        QuitWith(WA, msg.str());
    }
    QuitWith(AC, "Full solution");
}