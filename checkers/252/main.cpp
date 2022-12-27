#define _CRT_SECURE_NO_DEPRECATE
#define _USE_MATH_DEFINES
#include <iostream>
#include <cmath>
#include <checkers/testlib.h>
#include <vector>
#include <algorithm>
typedef long long int64;
typedef long double ld;
const double eps = 1e-6;
#define mp std::make_pair
// const int maxn = 1e+6;
#define pb push_back

using namespace NTestlib;

typedef std::pair<ld, ld> vec;
typedef std::vector<int64> vint;
typedef std::vector<vint> vvint;
vint primes;
vint HH, pp;
vec operator-(const vec &lhs, const vec &rhs)
{
    return mp(lhs.first - rhs.first, lhs.second - rhs.second);
}
vec operator+(const vec &lhs, const vec &rhs)
{
    return mp(lhs.first + rhs.first, lhs.second + rhs.second);
}
vec operator+(const vec &lhs, const int &rhs)
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
int main(int argc, char **argv)
{
    InitChecker(argc, argv);
    ld lena = File.ReadDouble(), lenb = File.ReadDouble(), lenc = File.ReadDouble();
    vec a = mp(Out.ReadDouble(), Out.ReadDouble()), b = mp(Out.ReadDouble(), Out.ReadDouble()), c = mp(Out.ReadDouble(), Out.ReadDouble());
    std::vector<ld> t;
    t.pb(len(a - b));
    t.pb(len(b - c));
    t.pb(len(c - a));
    std::vector<ld> ans;
    ans.pb(lena);
    ans.pb(lenb);
    ans.pb(lenc);
    std::sort(ans.begin(), ans.end());
    std::sort(t.begin(), t.end());
    for (int i = 0; i < 3; ++i)
    {
        if (std::abs(ans[i] - t[i]) > eps)
            QuitWith(WA, "Wrong answer");
    }
    QuitWith(AC, "OK");
}