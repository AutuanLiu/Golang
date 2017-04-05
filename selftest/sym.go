/*
*  求 11-999 的回文数 使 m = m^2 =m^3
*
*/

#include <iostream>
using namespace std;

bool sym(unsigned m)
{
	unsigned n = 0;
	unsigned i = m;
	while (i != 0)  //÷ 10 取余 构造新数
	{
		n = n * 10 + i % 10;
		i /= 10;
	}
	return m == n;
}

int main()
{
	unsigned m;
	for (m = 10;m < 1000; m++)
	{
		if (sym(m) && sym(m * m) && sym(m * m * m))
		{
			cout<<"m = "<<m<< "  m^2 = "<<m*m <<"  m^3 = "<<m*m*m<<endl;
 		}
	}
	return 0;
}