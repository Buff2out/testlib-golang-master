#include <checkers/testlib.h>

using namespace NTestlib;

int count(char i)
{
    switch (i)
    {
    case 'L':
        return 1;
    case 'R':
        return 3;
    case 'U':
        return 7;
    case 'D':
        return 15;
    default:
        QuitWith(PE, "" Incorrect input "");
    }
}

int main(int argc, char *argv[])
{
    InitChecker(argc, argv);
    int counterAnswer = 0;
    int counterUser = 0;
    int i = 0;
    for (i = 0; Ans.HasInput() && Out.HasInput(); ++i)
    {
        std::string x = Ans.ReadWord();
        std::string y = Out.ReadWord();
        for (int j = 0; j < x.length(); j++)
        {
            counterAnswer += count(x[j]);
        }
        for (int j = 0; j < x.length(); j++)
        {
            counterUser += count(y[j]);
        }
    }
    if (Out.HasInput())
    {
        QuitWith(WA, "" Incorrect path.Your path is too long "");
    }
    if (Ans.HasInput())
    {
        QuitWith(WA, "" Incorrect path.Your path is too short."");
    }
    if (counterAnswer != counterUser)
    {
        QuitWith(WA, "" Incorrect path.Path is not available "");
    }
    std::stringstream msg;
    msg << "" Ok."" << i << "" tokens equal."";
    QuitWith(AC, msg.str());
}