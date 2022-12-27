#include <checkers/testlib.h>
#include <algorithm>
#include <string>
#include <iostream>
#include <map>
#include <sstream>
#include <vector>

using namespace NTestlib;

std::vector<int> ReadInts(TInputStream &in)
{
    std::stringstream ss(in.ReadLine());
    std::vector<int> res;
    int x;
    while (ss >> x)
    {
        res.push_back(x);
    }
    return res;
}

std::string Msg1(int line, int expected, int got)
{
    std::stringstream out;
    out << "" Wrong number of ints on line : "" << line;
    out << "".Expected : "" << expected << "", got : "" << got;
    return out.str();
}

std::string Msg2(int key)
{
    std::stringstream out;
    out << "" Key : "" << key << "" is not a valid key "";
    return out.str();
}

int main(int argc, char *argv[])
{
    InitChecker(argc, argv);
    int n = File.ReadUInt();
    std::map<int, int> mp;
    for (int i = 0; i < n; ++i)
    {
        int key = File.ReadUInt();
        int val = File.ReadUInt();
        mp[key] += val;

        auto a = ReadInts(Ans);
        auto b = ReadInts(Out);

        // check output for duplicates
        sort(begin(a), end(a));
        sort(begin(b), end(b));
        if (unique(begin(a), end(a)) != end(a))
        {
            QuitWith(EF, "" Some key was printed more than once !"");
        }
        if (unique(begin(b), end(b)) != end(b))
        {
            QuitWith(WA, "" Some key was printed more than once !"");
        }

        // check size
        if (a.size() != b.size())
        {
            QuitWith(WA, Msg1(i, a.size(), b.size()));
        }

        // check that all keys are the keys that was already given, not arbitary ones.
        int n = (int)a.size();
        for (int i = 0; i < n; ++i)
        {
            if (mp.find(a[i]) == end(mp))
            {
                QuitWith(EF, Msg2(a[i]));
            }
            if (mp.find(b[i]) == end(mp))
            {
                QuitWith(WA, Msg2(b[i]));
            }
        }

        // now sort and compare values. we can't compare just keys because there may be keys with same value
        // and in such cases it is allowed to output any value
        // in other words multiset of values corresponding to ans keys and solution keys show be equal
        auto f = [&](int x)
        {
            return mp[x];
        };
        std::transform(begin(a), end(a), begin(a), f);
        std::transform(begin(b), end(b), begin(b), f);
        std::sort(begin(a), end(a));
        std::sort(begin(b), end(b));
        if (a > b)
        {
            QuitWith(WA, "" Solution not optimal "");
        }
        if (b > a)
        {
            QuitWith(EF, "" Solution not optimal "");
        }
    }
    QuitWith(AC);
    return 0;
}