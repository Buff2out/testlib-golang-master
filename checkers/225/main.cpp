#include <checkers/testlib.h>
#include <iostream>
#include <cmath>
#include <fstream>
#include <stdlib.h> 
#include <vector>
#include <iomanip>
#include <sstream>

using namespace NTestlib;
double func1(double a, double b, double x)
{
	return pow(x+a, 2) + b;
}

double func2(double a, double b, double c, double x)
{
	return std::sin(x) + std::abs(x + a) - std::abs(x + b) + std::abs(x + c);
}

int main(int argc, char* argv[]) {
	InitInteractor(argc, argv);
	int test = File.ReadInt();
	int cnt = 0;
	if (test > 10)
	{
		double a = File.ReadDouble();
		double b = File.ReadDouble();
		double c = File.ReadDouble();
		double an;
		if (std::abs(a - b) > std::abs(b - c))
			an = -a;
		else
			an = -c;
		while (1)
		{
			char requestType = Out.ReadChar();
			if (cnt > 1000000000)
				QuitWith(WA, "WA2");	
			switch (requestType)
			{
				case '!':
				{
					double x = Out.ReadDouble();
					if (std::abs(x - an) > 0.25 && (std::abs(func2(a,b,c,x) - func2(a,b,c,an)) > 1.9))
					{
						std::stringstream msg;
						msg << "WA x:" << x << ", ans: " << an << ", abs1: " << std::abs(x - an) << ", abs2: " << func2(a, b, c, an) - func2(a, b, c, x);
						QuitWith(WA, msg.str());
					}
					else
						QuitWith(AC, "OK");	
					break;
				}
				case '?':
				{
					double x = Out.ReadDouble();
					std::cout << std::setprecision(20) << func2(a, b, c, x) << std::endl;
					break;
				}
				default:
				{
					std::stringstream msg;
					msg << "Wrong query: " << requestType;
					QuitWith(PE, msg.str());
					break;
				}
			}
			Out.Match('\n');
			cnt++;
		}
	}
	else
	{
		double a = File.ReadDouble();
		double b = File.ReadDouble();
		double an = -a;
		while (1)
		{
			char requestType = Out.ReadChar();
			if (cnt > 1000000000)
				QuitWith(WA, "Too many requests");	
			switch (requestType)
			{
				case '!':
				{
					double x = Out.ReadDouble();
					if (std::abs(x - an) > 0.25)
					{
						std::stringstream msg;
						msg << "WA x: " <<  x << "ans: " << an << "abs: " << std::abs(x - an);
						QuitWith(WA, msg.str());
					}
					else
						QuitWith(AC, "OK");	
					break;
				}
				case '?':
				{
					double x = Out.ReadDouble();
					std::cout << std::setprecision(20) << func1(a, b, x) << std::endl;
					break;
				}
				default:
				{
					std::stringstream msg;
					msg << "Wrong query: " << requestType;
					QuitWith(PE, msg.str());
					break;
				}
			}
			Out.Match('\n');
			cnt++;
		}
	}
	std::cerr << "INT\n";
	return 0;
}