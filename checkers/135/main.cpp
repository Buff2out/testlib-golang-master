#include <checkers/testlib.h>
#include <iostream>
#include <fstream>
#include <stdlib.h>
#include <vector>
#include <sstream>

using namespace NTestlib;

int main(int argc, char *argv[])
{
    InitInteractor(argc, argv);
    int n = File.ReadInt();
    std::cout << n << std::endl;
    int cnt = 0;
    int answer;
    answer = Ans.ReadInt();
    std::vector<int> a;
    for (int i = 0; i < n; i++)
    {
        a.push_back(File.ReadInt());
    }
    while (true)
    {
        if (cnt >= 50)
        {
            QuitWith(WA, "Too many requests");
        }
        char requestType = Out.ReadChar();
        switch (requestType)
        {
        case '!':
        {
            int userAnswer = Out.ReadInt(1, n);
            int res = userAnswer - 1;
            if (res == 0)
            {
                if (a[0] >= a[1])
                    QuitWith(AC, "OK");
                else
                    QuitWith(WA, "Element isn't a peak one");
            }
            else if (res == n - 1)
            {
                if (a[n - 1] >= a[n - 2])
                    QuitWith(AC, "OK");
                else
                    QuitWith(WA, "Element isn't a peak one");
            }
            else if (a[res] >= a[res - 1] && a[res] >= a[res + 1])
            {
                QuitWith(AC, "OK");
            }
            else
            {
                QuitWith(WA, "Element isn't a peak one");
            }
        }
        case '?':
        {
            int userAnswer = Out.ReadInt(1, n);
            if (userAnswer > 0 && userAnswer <= n)
            {
                std::cout << a[userAnswer - 1] << std::endl;
            }
            else
            {
                std::stringstream msg;
                msg << "Wrong index " << userAnswer;
                QuitWith(PE, msg.str());
            }
            break;
        }
        default:
        {
            std::stringstream msg;
            msg << "Incorrect input command: " << requestType;
            QuitWith(PE, msg.str());
        }
        break;
        }
        Out.Match('\n');
        cnt++;
    }
    std::cerr << "INT\n";
    return 0;
}