#include <checkers/testlib.h>
#include <iostream>
#include <string>
#include <vector>

using namespace NTestlib;

int main(int argc, char *argv[])
{
    InitChecker(argc, argv);
    int n = File.ReadInt();
    int sum = File.ReadInt();
    std::vector<int> a;
    for (int i = 0; i < n; i++)
        a.push_back(File.ReadInt());
    int answer = Ans.ReadInt();
    if (answer == -1)
    {
        Out.ReadInt(-1, -1);
        QuitWith(AC, "Correct output");
    }
    int i1 = Out.ReadInt(1, n), i2 = Out.ReadInt(1, n);
    if (i1 == i2)
    {
        QuitWith(WA, "First variable equal second");
    }
    if (a[i1 - 1] + a[i2 - 1] != sum)
    {
        std::stringstream msg;
        msg << "Incorrect sum. Get " << a[i1 - 1] + a[i2 - 1] << ". Correct answer: " << sum;
        QuitWith(WA, msg.str());
    }
    QuitWith(AC, "Correct output");
}