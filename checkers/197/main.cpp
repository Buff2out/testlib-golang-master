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

int f[8] = {-2, -2, -1, -1, 1, 1, 2, 2};
int s[8] = {1, -1, 2, -2, -2, 2, -1, 1};

int main(int argc, char *argv[])
{
  InitChecker(argc, argv);
  int n = File.ReadInt(), m = File.ReadInt();
  std::vector<std::vector<int>> c(n + 1, std::vector<int>(m + 1, 0));
  int temp = Out.ReadInt();
  if (temp == -1)
    if (n == 4 && m == 4)
      QuitWith(AC, "" Full solution "");
    else
      QuitWith(WA, "" Wrong answer "");
  if (temp < 1 || temp > n)
  {
    QuitWith(PE, "" Index out of range "");
  }
  int x = temp, y = Out.ReadInt(1, m);
  c[x][y] = 1;
  int x2 = x, y2 = y;
  for (int i = 0; i < n * m - 1; i++)
  {
    bool isFind = false;
    int x1 = Out.ReadInt(1, n), y1 = Out.ReadInt(1, m);
    for (int j = 0; j < 8; j++)
    {
      if (x1 - x == f[j] && y1 - y == s[j])
      {
        isFind = true;
        break;
      }
    }
    if (!isFind)
    {
      std::stringstream msg;
      msg << "" WA Incorrect step : x = "" << x1 << "";
      y = "" << y1;
      QuitWith(WA, msg.str());
    }
    if (c[x1][y1] == 1 && !(x1 == x2 && y1 == y2))
      QuitWith(WA, "" Twice "");
    c[x1][y1] = 1;
    x = x1;
    y = y1;
  }
  QuitWith(AC, "" Full solution "");
}