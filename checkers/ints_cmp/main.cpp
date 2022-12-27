#include <checkers/testlib.h>

using namespace NTestlib;

int main(int argc, char *argv[])
{
    InitChecker(argc, argv);
    int i;
    for (i = 0; Ans.HasInput() && Out.HasInput(); ++i)
    {
        i64 x = Ans.ReadInt();
        i64 y = Out.ReadInt();
        if (x != y)
        {
            std::stringstream msg;
            msg << i + 1 << " integer don't match. Expected: " << x << ", got: " << y;
            QuitWith(WA, msg.str());
        }
    }
    if (Out.HasInput())
    {
        QuitWith(PE, "solution has more data than expected");
    }
    if (Ans.HasInput())
    {
        QuitWith(PE, "solution has less data than expected");
    }
    std::stringstream msg;
    msg << "Ok. " << i << " integers equal.";
    QuitWith(AC, msg.str());
}