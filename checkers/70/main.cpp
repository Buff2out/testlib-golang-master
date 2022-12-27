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
    for (int i = 0; i < n; i++)
    {
        a.push_back(File.ReadInt());
    }
    int ind = Out.ReadInt(1, n);
    long long int cnt1 = 0, cnt2 = 0;
    for (int i = 0; i < ind - 1; i++)
    {
        cnt1 += a[i];
    }
    for (int i = ind; i < n; i++)
    {
        cnt2 += a[i];
    }
    if (cnt1 == cnt2)
        QuitWith(AC, "OK");
    std::stringstream msg;
    msg << "First part: " << cnt1 << "; Second part: " << cnt2;
    QuitWith(WA, msg.str());
}
