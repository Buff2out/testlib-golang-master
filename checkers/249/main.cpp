#define _CRT_SECURE_NO_DEPRECATE
#define _USE_MATH_DEFINES
#include <iostream>
#include <cmath>
#include <vector>
#include <checkers/testlib.h>
#include <algorithm>
typedef long long int64;
typedef long double ld;
// const int maxn = 1e+6;
#define eps 1e-6
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
    ld lena = File.ReadDouble(), lenb = File.ReadDouble(), angle = File.ReadDouble();
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
        if (((std::abs(len(first) - lena) < eps && std::abs(len(second) - lenb) < eps) || (std::abs(len(first) - lenb) < eps && std::abs(len(second) - lena) < eps)) && std::abs(angleaa(first, second) * 180 / M_PI - angle) < eps)
            QuitWith(AC, "OK");
    }
    QuitWith(WA, "Wrong answer");
}
