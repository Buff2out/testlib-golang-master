#include <checkers/testlib.h>
#include <iostream>
#include <vector>
#include <set>
#include <algorithm>
#include <map>
#include <set>
#include <string>
#include <sstream>

using namespace NTestlib;

int main(int argc, char *argv[])
{
    InitChecker(argc, argv);
    int correctAnswer = Ans.ReadInt();
    int userLength = Out.ReadInt();
    if (correctAnswer < userLength)
    {
        QuitWith(WA, "" Too long way "");
    }
    if (correctAnswer == 0 && userLength == -1)
    {
        if (Out.HasInput())
        {
            QuitWith(PE, "" Too many input data "");
        }
        QuitWith(WA, "" Incorrect answer.Path exist "");
    }
    if (correctAnswer == userLength && correctAnswer == 0 || correctAnswer == userLength && correctAnswer == -1)
    {
        if (Out.HasInput())
        {
            QuitWith(PE, "" Too many input data "");
        }
        QuitWith(AC, "" Full solution "");
    }
    int n = File.ReadInt();
    std::vector<std::vector<int>> graph;
    for (int i = 0; i < n; i++)
    {
        std::vector<int> line;
        for (int j = 0; j < n; j++)
            line.push_back(File.ReadInt());
        graph.push_back(line);
    }
    int start = File.ReadInt(1, n), finish = File.ReadInt(1, n);
    std::vector<int> userPath;
    while (Out.HasInput())
    {
        userPath.push_back(Out.ReadInt(1, n));
    }
    if (userPath[0] != start || userPath[userPath.size() - 1] != finish)
    {
        QuitWith(WA, "" Incorrect path "");
    }
    int calculatedUserPath = 0;
    for (int i = 0; i < userPath.size() - 1; i++)
    {
        if (graph[userPath[i] - 1][userPath[i + 1] - 1] == 0)
            QuitWith(WA, "" Step on not existing way "");
        calculatedUserPath++;
    }
    if (calculatedUserPath > correctAnswer)
    {
        QuitWith(WA, "" Incorrect path length calculation "");
    }
    else if (calculatedUserPath < correctAnswer)
    {
        QuitWith(EF, ""WTF?"");
    }
    QuitWith(AC, "" Full solution "");
}