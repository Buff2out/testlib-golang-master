#include <checkers/testlib.h>
#include <iostream>
#include <vector>
#include <algorithm>
#include <string>

using namespace NTestlib;

int main(int argc, char *argv[])
{
    InitChecker(argc, argv);
    int n = File.ReadInt();
    std::vector<int> a;
    std::vector<int> b;
    for (int i = 0; i < n; i++)
    {
        int x = File.ReadInt();
        a.push_back(x);
        b.push_back(x);
    }
    int ind1, ind2, cnt1 = 0, cnt2 = 0, max1 = 0, max2 = 0;
    ind1 = Out.ReadInt(1, n);
    ind2 = Ans.ReadInt();
    a[ind1 - 1] = 1;
    b[ind2 - 1] = 1;
    for (int i = 0; i < n; i++)
    {
        if (a[i] == 0)
        {
            max1 = std::max(max1, cnt1);
            cnt1 = 0;
        }
        else
            cnt1++;
    }
    max1 = std::max(max1, cnt1);
    for (int i = 0; i < n; i++)
    {
        if (b[i] == 0)
        {
            max2 = std::max(max2, cnt2);
            cnt2 = 0;
        }
        else
            cnt2++;
    }
    max2 = std::max(max2, cnt2);
    if (max1 < max2)
    {
        std::stringstream msg;
        msg << "User max value " << max1 << " less than correct max value " << max2;
        QuitWith(WA, msg.str());
    }
    if (max1 == max2)
        QuitWith(AC, "OK");
    if (max1 > max2)
    {
        std::stringstream msg;
        msg << "STUDENT SOLUTION RESULT " << max1 << " BETTER THAN SYSTEM " << max2;
        QuitWith(EF, msg.str());
    }
}