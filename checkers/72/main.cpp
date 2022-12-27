#include <checkers/testlib.h>
#include <iostream>
#include <vector>
#include <set>
#include <algorithm>
#include <map>
#include <string>
#include <sstream>

using namespace NTestlib;

std::string inttostring(int Number)
{
    std::stringstream convert;
    convert << Number;
    return convert.str();
}
int main(int argc, char *argv[])
{
    InitChecker(argc, argv);
    std::multiset<int> a;
    std::set<std::string> b;
    int n = File.ReadInt();
    for (int i = 0; i < n; i++)
    {
        a.insert(File.ReadInt());
    }
    int m = File.ReadInt();
    int cnt = 0, cnt1 = 0;
    for (int i = 0; Ans.HasInput(); i++)
    {
        for (int j = 0; j < m; j++)
        {
            int x = Ans.ReadInt();
        }
        cnt++;
    }
    for (int i = 0; Out.HasInput(); i++)
    {
        std::map<int, int> cntmap;
        std::string check = "";
        for (int j = 0; j < m; j++)
        {
            int x = Out.ReadInt();
            cntmap[x]++;
            if (cntmap[x] > a.count(x))
            {
                std::stringstream msg;
                msg << "Map: " << cntmap[x] << "; Count: " << a.count(x) << "; Number: " << x;
                QuitWith(WA, msg.str());
            }
            check += inttostring(x) + '|';
        }
        cnt1++;
    }
    if (cnt1 < cnt)
        QuitWith(WA, "Users count less than correct");
    if (cnt1 == cnt)
        QuitWith(AC, "OK");
    if (cnt1 > cnt)
        QuitWith(WA, "Users count greater than correct");
}