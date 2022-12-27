#include <checkers/testlib.h>

using namespace NTestlib;

int main(int argc, char *argv[])
{
    InitChecker(argc, argv);
    std::stringstream msg;
    msg << "Ok";
    QuitWith(AC, msg.str());
}