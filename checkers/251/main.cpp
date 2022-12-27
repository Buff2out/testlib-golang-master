#define _CRT_SECURE_NO_DEPRECATE
#define _USE_MATH_DEFINES
#include <iostream>
#include <cmath>
#include <vector>
#include <algorithm>
#include <checkers/testlib.h>

typedef long long int64;
typedef long double ld;

const double eps = 5 * 1e-3;

// const int maxn = 1e+6;

#define mp std::make_pair
#define pb push_back

using namespace NTestlib;

typedef std::pair<ld, ld> vec;
typedef std::vector<int64> vint;
typedef std::vector<vint> vvint;

vec operator-(const vec &lhs, const vec &rhs)
{
    return mp(lhs.first - rhs.first, lhs.second - rhs.second);
}
vec operator+(const vec &lhs, const vec &rhs)
{
    return mp(lhs.first + rhs.first, lhs.second + rhs.second);
}
vec operator/(const vec &lhs, const int &rhs)
{
    return mp(lhs.first / rhs, lhs.second / rhs);
}

ld dist(vec a, vec b)
{
    return sqrtl((a.first - b.first) * (a.first - b.first) + (a.second - b.second) * (a.second - b.second));
}
ld len(vec x)
{
    return sqrtl(x.first * x.first + x.second * x.second);
}
ld angle(vec a, vec b)
{
    return acosl((a.first * b.first + a.second * b.second) / (len(a) * len(b)));
}
ld ar_tr(vec a, vec b)
{
    return a.first * b.second - b.first * a.second;
}

bool CheckBySinusTheoreme(std::vector<ld> sortedSideLengths, std::vector<ld> sortedAngles)
{
    bool flag = false;
    ld tempValue = sortedSideLengths[0] / sin(sortedAngles[0] * M_PI / 180);
    for (int i = 1; i < 3; i++)
    {
        if (std::abs(sortedSideLengths[i] / sin(sortedAngles[i] * M_PI / 180) - tempValue) > eps)
        {
            return false;
        }
    }
    return true;
}

int main(int argc, char **argv)
{
    // TODO: Napishi menya normalno
    InitChecker(argc, argv);
    ld lena = File.ReadDouble(), anglea = File.ReadDouble(), angleb = File.ReadDouble(), anglec = 180.0 - anglea - angleb;
    vec a = mp(Out.ReadDouble(), Out.ReadDouble()), b = mp(Out.ReadDouble(), Out.ReadDouble()), c = mp(Out.ReadDouble(), Out.ReadDouble());
    // sort angle by ASC = vector<ld> sortedAngles
    // Find length of all lines and sort = vector <ld> sortedLines
    // Check by sinus theoreme
    // If previous == true =>
    // Check cosinus theroeme
    // If previous == true =>
    // OK
    // else =>
    std::vector<vec> t;
    std::vector<ld> sortedAngles;
    sortedAngles.pb(anglea);
    sortedAngles.pb(angleb);
    sortedAngles.pb(anglec);
    std::sort(sortedAngles.begin(), sortedAngles.end());
    std::vector<ld> sortedSideLengths;
    sortedSideLengths.pb(len(b - a));
    sortedSideLengths.pb(len(c - a));
    sortedSideLengths.pb(len(c - b));
    std::sort(sortedSideLengths.begin(), sortedSideLengths.end());
    int gammaAngleIndex = 0;
    for (int i = 0; i < 3; i++)
    {
        if (sortedAngles[i] == anglec)
        {
            gammaAngleIndex = i;
            break;
        }
    }
    if (!CheckBySinusTheoreme(sortedSideLengths, sortedAngles))
    {
        QuitWith(WA, "" Wrong answer.Does not match the sinus theorem "");
    }
    int notMedianAngleIndex = (gammaAngleIndex + 1) % 3;
    ld medianSideLength = sortedSideLengths[gammaAngleIndex] / 2;
    int lastAngleIndex = (notMedianAngleIndex + 1) % 3;
    // By cosinus theorem
    if (std::abs((medianSideLength * medianSideLength + sortedSideLengths[lastAngleIndex] * sortedSideLengths[lastAngleIndex] - lena * lena) / (2 * medianSideLength * sortedSideLengths[lastAngleIndex]) - cos(sortedAngles[notMedianAngleIndex] * M_PI / 180)) < eps)
    {
        QuitWith(AC, "" OK "");
    }
    QuitWith(WA, "" Wrong answer.Does not match the cosinus theorem "");
}