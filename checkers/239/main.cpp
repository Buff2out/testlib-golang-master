#define _CRT_SECURE_NO_DEPRECATE
#define _USE_MATH_DEFINES
#include <iostream>
#include <checkers/testlib.h>
#include <time.h>
#include <vector>
#include <algorithm>
#include <cmath>
#include <string>
#include <set>
#include <vector>
#include <map>
#include <sstream>
#include <iomanip>
#include <stack>
#include <cstdio>
#include <fstream>
#include <cstdlib>
#include <numeric>
#include <cstring>
#include <complex>
#include <cassert>
#include <iterator>
#include <functional>

using namespace NTestlib;

struct node
{
    int key;
    unsigned char height;
    node *left;
    node *right;
    node(int k)
    {
        key = k;
        left = right = 0;
        height = 1;
    }
};

unsigned char height(node *p)
{
    return p ? p->height : 0;
}

int bfactor(node *p)
{
    return height(p->right) - height(p->left);
}

void fixheight(node *p)
{
    unsigned char hl = height(p->left);
    unsigned char hr = height(p->right);
    p->height = (hl > hr ? hl : hr) + 1;
}

node *rotateright(node *p)
{
    node *q = p->left;
    p->left = q->right;
    q->right = p;
    fixheight(p);
    fixheight(q);
    return q;
}

node *rotateleft(node *q)
{
    node *p = q->right;
    q->right = p->left;
    p->left = q;
    fixheight(q);
    fixheight(p);
    return p;
}

node *balance(node *p)
{
    fixheight(p);
    if (bfactor(p) == 2)
    {
        if (bfactor(p->right) < 0)
            p->right = rotateright(p->right);
        return rotateleft(p);
    }
    if (bfactor(p) == -2)
    {
        if (bfactor(p->left) > 0)
            p->left = rotateleft(p->left);
        return rotateright(p);
    }
    return p;
}

node *insert(node *p, int k)
{
    if (!p)
        return new node(k);
    if (k < p->key)
        p->left = insert(p->left, k);
    else
        p->right = insert(p->right, k);
    return balance(p);
}

node *findmin(node *p)
{
    return p->left ? findmin(p->left) : p;
}

node *removemin(node *p)
{
    if (p->left == 0)
        return p->right;
    p->left = removemin(p->left);
    return balance(p);
}

node *remove(node *p, int k)
{
    if (!p)
        return 0;
    if (k < p->key)
        p->left = remove(p->left, k);
    else if (k > p->key)
        p->right = remove(p->right, k);
    else // k == p->key
    {
        node *q = p->left;
        node *r = p->right;
        delete p;
        if (!r)
            return q;
        node *min = findmin(r);
        min->right = removemin(r);
        min->left = q;
        return balance(min);
    }
    return balance(p);
}
int nk = 0;
std::map<int, int> ch, temp_ch;
void input(node *k, node *before, bool isleft)
{
    int x = Out.ReadInt();
    if (x == -1)
        return;
    nk++;
    if (!k)
    {
        k = new node(x);
        if (isleft)
            before->left = k;
        else
            before->right = k;
    }
    else
        k->key = x;
    input(k->left, k, true);
    input(k->right, k, false);
}
void check(node *k)
{
    if (!k)
        return;
    if (bfactor(k) > 1)
        QuitWith(WA, "" Not balanced "");
    if (k->left != NULL)
    {
        if (k->key < k->left->key)
            QuitWith(WA, "" Tree not AVL "");
        else
            check(k->left);
    }
    if (k->right != NULL)
    {
        if (k->key > k->right->key)
            QuitWith(WA, "" Tree not AVL "");
        else
            check(k->right);
    }
    return;
}
int main(int argc, char **argv)
{

    InitChecker(argc, argv);

    int n = File.ReadInt();
    int x;
    std::stringstream firstRow;
    firstRow << File.ReadLine();
    while (firstRow >> x)
    {
        if (x != -1)
            ch[x]++;
    }
    int m = File.ReadInt();
    for (int i = 0; i < m; ++i)
    {
        nk = 0;
        int q = File.ReadInt();
        x = File.ReadInt();
        if (q == 1)
        {
            ch[x]++;
            n++;
        }
        else
        {
            ch[x]--;
            n--;
        }
        temp_ch = ch;
        node *k = new node(0);
        input(k, 0, 0);
        if (nk != n)
        {
            std::stringstream msg;
            msg << "" Less vertices "" << x << ""
                                               ""
                << nk << ""
                         ""
                << n;
            QuitWith(WA, msg.str());
        }
        check(k);
        ch = temp_ch;
    }
    QuitWith(AC, "" OK "");
}