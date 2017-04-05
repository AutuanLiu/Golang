/*
*    二分幂法 求x^n
*/

#include <iostream>
using std::cin;
using std::cout;
using std::endl;

double powerf(double x, unsigned n)
{
	double ans = 1.0;

	while (n != 0)
	{
		if (n % 2 == 1)
		{
			ans *= x ; 
		}
		x *= x;
		n /= 2;
	}
	return ans;
}

/*
*    递归法 求x^n
*/

double powerf2(double x, unsigned n)
{
	double ans = 1.0;

	if (n == 0)
	{
		return 1;
	}
	else
	{
		return x * powerf2(x,n-1);
	}
}

/*
*    循环法 求x^n
*/

double powerf3(double x, unsigned n)
{
	double ans = 1.0;

	while (n != 0)
	{
		ans *= x;
		n--;
	}
	return ans;
}

int main()
{
	
	double x;
	unsigned n;
	cin >> x >> n;
	cout << x <<"^" << n << " = " << powerf(x,n) << endl;
	cout << x <<"^" << n << " = " << powerf2(x,n) << endl;
	cout << x <<"^" << n << " = " << powerf3(x,n) << endl;
	return 0;
}