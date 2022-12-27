#include <checkers/testlib.h>

using namespace NTestlib;

int main(int argc, char *argv[])
{
    InitChecker(argc, argv);
    int i;
    for (i = 0; Ans.HasInput() && Out.HasInput(); ++i)
    {
        double x = Ans.ReadDouble();
        double y = Out.ReadDouble();
        if (!DoubleEq(y, x, 1e-7))
        {
            std::stringstream msg;
            msg << i + 1 << " doubles don't match. Expected: " << x << ", got: " << y;
            msg << ". Error: " << std::min(AbsError(x, y), RelError(y, x));
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
    msg << "Ok. " << i << " tokens equal.";
    QuitWith(AC, msg.str());
}
