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
    std::vector<int> a;
    for (int i = 0; i < n; i++)
        a.push_back(File.ReadInt());
    int cnt = 0;
    int answer;
    answer = Ans.ReadInt();
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
            int userAnswer = Out.ReadInt();
            if (userAnswer != answer)
            {
                std::stringstream msg;
                msg << "Incorrect user output. User answer: " << userAnswer << ". Correct answer: " << answer;
                QuitWith(WA, msg.str());
            }
            QuitWith(AC, "OK");
            break;
        }
        case '?':
        {
            int userAnswer = Out.ReadInt(1, n);
            if (userAnswer > 0 && userAnswer <= n)
                std::cout << a[userAnswer - 1] << std::endl;
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