#define _CRT_SECURE_NO_DEPRECATE
#define _USE_MATH_DEFINES
#include <iostream>
#include <cmath>
#include <vector>
#include <algorithm>
#include <checkers/testlib.h>
typedef long long int64;
typedef long double ld;
// const int maxn = 1e+6;
#define mp std::make_pair
#define eps 1e-6
#define pb push_back

using namespace NTestlib;

typedef std::pair<ld, ld> vec;
typedef std::vector<int64> vint;
typedef std::vector<vint> vvint;
vec operator-(const vec &lhs, const vec &rhs)
{
    return mp(lhs.first - rhs.first, lhs.second - rhs.second);
}
ld dist(vec a, vec b)
{
    return sqrtl((a.first - b.first) * (a.first - b.first) + (a.second - b.second) * (a.second - b.second));
}
ld len(vec x)
{
    return sqrtl(x.first * x.first + x.second * x.second);
}
ld angleaa(vec a, vec b)
{
    return acosl((a.first * b.first + a.second * b.second) / (len(a) * len(b)));
}
ld ar_tr(vec a, vec b)
{
    return a.first * b.second - b.first * a.second;
}
int main(int argc, char **argv)
{
    InitChecker(argc, argv);
    ld lena = File.ReadDouble(), anglea = File.ReadDouble(), angleb = File.ReadDouble();
    vec a = mp(Out.ReadDouble(), Out.ReadDouble()), b = mp(Out.ReadDouble(), Out.ReadDouble());
    vec c = mp(Out.ReadDouble(), Out.ReadDouble());
    std::vector<vec> t;
    t.pb(a);
    t.pb(b);
    t.pb(c);
    for (int i = 0; i < 3; ++i)
    {
        vec first = (t[(i + 1) % 3] - t[i]);
        vec second = (t[(i + 2) % 3] - t[i]);
        vec third = (t[(i + 1) % 3] - t[(i + 2) % 3]);
        ld angle1 = angleaa(first, third) * 180 / M_PI;
        ld angle2 = angleaa(second, first) * 180 / M_PI;
        if (std::abs(len(first) - lena) < eps && ((std::abs(angle1 - anglea) && std::abs(angle2 - angleb) < eps) || (std::abs(angle1 - angleb) < eps && std::abs(angle2 - anglea) < eps)))
            QuitWith(AC, "OK");
    }
    QuitWith(WA, "Wrong answer");
}
