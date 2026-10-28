#pragma once
#include <string>
#include <vector>
#include <map>
#include "array.h"
#include "forward_list.h"
#include "list.h"
#include "stack.h"
#include "queue.h"
#include "tree.h"

using namespace std;

struct DB {
    map<string, Array<string>> arrays;
    map<string, ForwardList<string>> flists;
    map<string, List<string>> lists;
    map<string, Stack<string>> stacks;
    map<string, Queue<string>> queues;
    map<string, CompleteBinaryTree<int>> trees;
};

vector<string> split(const string& s);
bool isInt(const string& s);
string exec(DB& db, const string& line);
void load(DB& db, const string& path);
void save(const string& path, const string& cmd);
