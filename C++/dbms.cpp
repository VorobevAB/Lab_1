#include "dbms.h"
#include <iostream>
#include <fstream>

using namespace std;

vector<string> split(const string& s) {
    vector<string> out;
    string cur;
    for (char c : s) {
        if (c == ' ') {
            if (!cur.empty()) {
                out.push_back(cur);
                cur.clear();
            }
        }
        else
            cur.push_back(c);
    }
    if (!cur.empty())
        out.push_back(cur);
    return out;
}

bool isInt(const string& s) {
    if (s.empty())
        return false;
    int i = 0;
    if (s[0] == '-') {
        if (s.size() == 1)
            return false;
        i = 1;
    }
    for (; i < (int)s.size(); ++i)
        if (s[i] < '0' || s[i] > '9')
            return false;
    return true;
}

string exec(DB& db, const string& line) {
    auto t = split(line);
    if (t.empty())
        return "ОШИБКА: пустая команда";
    
    string op = t[0];
	
    if (op == "PRINT") {
        if (t.size() != 2) return "ОШИБКА: PRINT <name>";
        string name = t[1];
        if (db.arrays.count(name)) { db.arrays[name].MPrint(); return "OK"; }
        if (db.flists.count(name)) { db.flists[name].FPrint(); return "OK"; }
        if (db.lists.count(name))  { db.lists[name].LPrint();  return "OK"; }
        if (db.stacks.count(name)) { db.stacks[name].SPrint(); return "OK"; }
        if (db.queues.count(name)) { db.queues[name].QPrint(); return "OK"; }
        if (db.trees.count(name))  { db.trees[name].TPrint();  return "OK"; }
        return "ОШИБКА: контейнер '" + name + "' не найден";
    }
    
    switch (op[0]) {
    case 'M': {
        if (op == "MPUSH") {
            if (t.size() != 3) return "ОШИБКА: MPUSH <name> <value>";
            if (t[2].empty())  return "ОШИБКА: пустое значение";
            db.arrays[t[1]].MPush(t[2]);
            return t[2];
        }
        if (op == "MPUSHIDX") {
            if (t.size() != 4) return "ОШИБКА: MPUSHIDX <name> <idx> <value>";
            if (!isInt(t[2]))  return "ОШИБКА: idx — не число";
            if (t[3].empty())  return "ОШИБКА: пустое значение";
            db.arrays[t[1]].MPushIdx(stoi(t[2]), t[3]);
            return t[3];
        }
        if (op == "MGET") {
            if (t.size() != 3) return "ОШИБКА: MGET <name> <idx>";
            if (!isInt(t[2]))  return "ОШИБКА: idx — не число";
            return db.arrays[t[1]].MGet(stoi(t[2]));
        }
        if (op == "MDEL") {
            if (t.size() != 3) return "ОШИБКА: MDEL <name> <idx>";
            if (!isInt(t[2]))  return "ОШИБКА: idx — не число";
            db.arrays[t[1]].MDelInd(stoi(t[2]));
            return "OK";
        }
        if (op == "MSET") {
            if (t.size() != 4) return "ОШИБКА: MSET <name> <idx> <value>";
            if (!isInt(t[2]))  return "ОШИБКА: idx — не число";
            if (t[3].empty())  return "ОШИБКА: пустое значение";
            db.arrays[t[1]].MSetInd(stoi(t[2]), t[3]);
            return t[3];
        }
        if (op == "MLEN") {
            if (t.size() != 2) return "ОШИБКА: MLEN <name>";
            return to_string(db.arrays[t[1]].MLen());
        }
        return "ОШИБКА: неизвестная команда '" + op + "'";
    }

    case 'F': {
        if (op == "FPUSH") {
            if (t.size() != 3) return "ОШИБКА: FPUSH <name> <value>";
            if (t[2].empty())  return "ОШИБКА: пустое значение";
            db.flists[t[1]].FPushFront(t[2]);
            return t[2];
        }
        if (op == "FPUSHB") {
            if (t.size() != 3) return "ОШИБКА: FPUSHB <name> <value>";
            if (t[2].empty())  return "ОШИБКА: пустое значение";
            db.flists[t[1]].FPushBack(t[2]);
            return t[2];
        }
        if (op == "FDEL") {
            if (t.size() != 2) return "ОШИБКА: FDEL <name>";
            db.flists[t[1]].FDelFront();
            return "OK";
        }
        if (op == "FDELBACK") {
            if (t.size() != 2) return "ОШИБКА: FDELBACK <name>";
            db.flists[t[1]].FDelBack();
            return "OK";
        }
        if (op == "FPUSHAFTER") {
	    if (t.size() != 4) return "ОШИБКА: FPUSHAFTER <name> <ref> <value>";
	    Node<string>* node = nullptr;
	    for (Node<string>* c = db.flists[t[1]].head; c; c = c->next)
		if (c->value == t[2]) { node = c; break; }
	    if (!node) return "ОШИБКА: узла '" + t[2] + "' нет в списке";
	    db.flists[t[1]].FPushAfter(node, t[3]);
	    return t[3];
	}
		if (op == "FPUSHBEFORE") {
	    if (t.size() != 4) return "ОШИБКА: FPUSHBEFORE <name> <ref> <value>";
	    Node<string>* node = nullptr;
	    for (Node<string>* c = db.flists[t[1]].head; c; c = c->next)
		if (c->value == t[2]) { node = c; break; }
	    if (!node) return "ОШИБКА: узла '" + t[2] + "' нет в списке";
	    db.flists[t[1]].FPushBefore(node, t[3]);
	    return t[3];
	}
	if (op == "FDELAFTER") {
	    if (t.size() != 3) return "ОШИБКА: FDELAFTER <name> <ref>";
	    Node<string>* node = nullptr;
	    for (Node<string>* c = db.flists[t[1]].head; c; c = c->next)
		if (c->value == t[2]) { node = c; break; }
	    if (!node) return "ОШИБКА: узла '" + t[2] + "' нет в списке";
	    db.flists[t[1]].FDelAfter(node);
	    return "OK";
	}
	if (op == "FDELBEFORE") {
	    if (t.size() != 3) return "ОШИБКА: FDELBEFORE <name> <ref>";
	    Node<string>* node = nullptr;
	    for (Node<string>* c = db.flists[t[1]].head; c; c = c->next)
		if (c->value == t[2]) { node = c; break; }
	    if (!node) return "ОШИБКА: узла '" + t[2] + "' нет в списке";
	    db.flists[t[1]].FDelBefore(node);
	    return "OK";
	}
        if (op == "FDELVAL") {
            if (t.size() != 3) return "ОШИБКА: FDELVAL <name> <value>";
            if (t[2].empty())  return "ОШИБКА: пустое значение";
            db.flists[t[1]].FDelValue(t[2]);
            return "OK";
        }
        if (op == "FFIND") {
            if (t.size() != 3) return "ОШИБКА: FFIND <name> <value>";
            if (t[2].empty())  return "ОШИБКА: пустое значение";
            db.flists[t[1]].FFind(t[2]);
            return "OK";
        }
        if (op == "FPRINT") {
            if (t.size() != 2) return "ОШИБКА: FPRINT <name>";
            db.flists[t[1]].FPrint();
            return "OK";
        }
        return "ОШИБКА: неизвестная команда '" + op + "'";
    }

    case 'L': {
        if (op == "LPUSH") {
            if (t.size() != 3) return "ОШИБКА: LPUSH <name> <value>";
            if (t[2].empty())  return "ОШИБКА: пустое значение";
            db.lists[t[1]].LPushFront(t[2]);
            return t[2];
        }
        if (op == "LPUSHB") {
            if (t.size() != 3) return "ОШИБКА: LPUSHB <name> <value>";
            if (t[2].empty())  return "ОШИБКА: пустое значение";
            db.lists[t[1]].LPushBack(t[2]);
            return t[2];
        }
                if (op == "LPUSHAFTER") {
            if (t.size() != 4) return "ОШИБКА: LPUSHAFTER <name> <ref> <value>";
            if (t[3].empty())  return "ОШИБКА: пустое значение";
            LNode<string>* node = nullptr;
            for (LNode<string>* c = db.lists[t[1]].head; c; c = c->next)
                if (c->value == t[2]) { node = c; break; }
            if (!node) return "ОШИБКА: узла '" + t[2] + "' нет в списке";
            db.lists[t[1]].LPushAfter(node, t[3]);
            return t[3];
        }
        if (op == "LPUSHBEFORE") {
            if (t.size() != 4) return "ОШИБКА: LPUSHBEFORE <name> <ref> <value>";
            if (t[3].empty())  return "ОШИБКА: пустое значение";
            LNode<string>* node = nullptr;
            for (LNode<string>* c = db.lists[t[1]].head; c; c = c->next)
                if (c->value == t[2]) { node = c; break; }
            if (!node) return "ОШИБКА: узла '" + t[2] + "' нет в списке";
            db.lists[t[1]].LPushBefore(node, t[3]);
            return t[3];
        }
        if (op == "LDELAFTER") {
            if (t.size() != 3) return "ОШИБКА: LDELAFTER <name> <ref>";
            LNode<string>* node = nullptr;
            for (LNode<string>* c = db.lists[t[1]].head; c; c = c->next)
                if (c->value == t[2]) { node = c; break; }
            if (!node) return "ОШИБКА: узла '" + t[2] + "' нет в списке";
            db.lists[t[1]].LDelAfter(node);
            return "OK";
        }
        if (op == "LDELBEFORE") {
            if (t.size() != 3) return "ОШИБКА: LDELBEFORE <name> <ref>";
            LNode<string>* node = nullptr;
            for (LNode<string>* c = db.lists[t[1]].head; c; c = c->next)
                if (c->value == t[2]) { node = c; break; }
            if (!node) return "ОШИБКА: узла '" + t[2] + "' нет в списке";
            db.lists[t[1]].LDelBefore(node);
            return "OK";
        }
        if (op == "LDEL") {
            if (t.size() != 2) return "ОШИБКА: LDEL <name>";
            db.lists[t[1]].LDelFront();
            return "OK";
        }
        if (op == "LDELBACK") {
            if (t.size() != 2) return "ОШИБКА: LDELBACK <name>";
            db.lists[t[1]].LDelBack();
            return "OK";
        }
        if (op == "LDELVAL") {
            if (t.size() != 3) return "ОШИБКА: LDELVAL <name> <value>";
            if (t[2].empty())  return "ОШИБКА: пустое значение";
            db.lists[t[1]].LDelValue(t[2]);
            return "OK";
        }
        if (op == "LFIND") {
            if (t.size() != 3) return "ОШИБКА: LFIND <name> <value>";
            if (t[2].empty())  return "ОШИБКА: пустое значение";
            db.lists[t[1]].LFind(t[2]);
            return "OK";
        }
        if (op == "LPRINT") {
            if (t.size() != 2) return "ОШИБКА: LPRINT <name>";
            db.lists[t[1]].LPrint();
            return "OK";
        }
        if (op == "LPRINTR") {
            if (t.size() != 2) return "ОШИБКА: LPRINTR <name>";
            db.lists[t[1]].LPrintReverse();
            return "OK";
        }
        return "ОШИБКА: неизвестная команда '" + op + "'";
    }

    case 'S': {
        if (op == "SPUSH") {
            if (t.size() != 3) return "ОШИБКА: SPUSH <name> <value>";
            if (t[2].empty())  return "ОШИБКА: пустое значение";
            db.stacks[t[1]].SPush(t[2]);
            return t[2];
        }
        if (op == "SPOP") {
            if (t.size() != 2) return "ОШИБКА: SPOP <name>";
            if (db.stacks[t[1]].SEmpty()) return "ОШИБКА: стек пуст";
            string v = db.stacks[t[1]].STop();
            db.stacks[t[1]].SPop();
            return v;
        }
        if (op == "STOP") {
            if (t.size() != 2) return "ОШИБКА: STOP <name>";
            if (db.stacks[t[1]].SEmpty()) return "ОШИБКА: стек пуст";
            return db.stacks[t[1]].STop();
        }
        if (op == "SPRINT") {
            if (t.size() != 2) return "ОШИБКА: SPRINT <name>";
            db.stacks[t[1]].SPrint();
            return "OK";
        }
        return "ОШИБКА: неизвестная команда '" + op + "'";
    }

    case 'Q': {
        if (op == "QPUSH") {
            if (t.size() != 3) return "ОШИБКА: QPUSH <name> <value>";
            if (t[2].empty())  return "ОШИБКА: пустое значение";
            db.queues[t[1]].QPush(t[2]);
            return t[2];
        }
        if (op == "QPOP") {
            if (t.size() != 2) return "ОШИБКА: QPOP <name>";
            if (db.queues[t[1]].QEmpty()) return "ОШИБКА: очередь пуста";
            string v = db.queues[t[1]].QFront();
            db.queues[t[1]].QPop();
            return v;
        }
        if (op == "QGET") {
            if (t.size() != 2) return "ОШИБКА: QGET <name>";
            if (db.queues[t[1]].QEmpty()) return "ОШИБКА: очередь пуста";
            return db.queues[t[1]].QFront();
        }
        if (op == "QPRINT") {
            if (t.size() != 2) return "ОШИБКА: QPRINT <name>";
            db.queues[t[1]].QPrint();
            return "OK";
        }
        return "ОШИБКА: неизвестная команда '" + op + "'";
    }

    case 'T': {
        if (op == "TINSERT") {
            if (t.size() != 3) return "ОШИБКА: TINSERT <name> <int>";
            if (!isInt(t[2]))  return "ОШИБКА: значение — не число";
            db.trees[t[1]].TInsert(stoi(t[2]));
            return t[2];
        }
        if (op == "TFIND") {
            if (t.size() != 3) return "ОШИБКА: TFIND <name> <int>";
            if (!isInt(t[2]))  return "ОШИБКА: значение — не число";
            db.trees[t[1]].TFind(stoi(t[2]));
            return "OK";
        }
        if (op == "TPRINT") {
            if (t.size() != 2) return "ОШИБКА: TPRINT <name>";
            db.trees[t[1]].TPrint();
            return "OK";
        }
        if (op == "TPRINTTREE") {
            if (t.size() != 2) return "ОШИБКА: TPRINTTREE <name>";
            db.trees[t[1]].TPrintTree();
            return "OK";
        }
        if (op == "TISCOMPLETE") {
            if (t.size() != 2) return "ОШИБКА: TISCOMPLETE <name>";
            return db.trees[t[1]].TIsComplete() ? "TRUE" : "FALSE";
        }
        if (op == "TISFULL") {
            if (t.size() != 2) return "ОШИБКА: TISFULL <name>";
            return db.trees[t[1]].TIsFull() ? "TRUE" : "FALSE";
        }
        return "ОШИБКА: неизвестная команда '" + op + "'";
    }

    default:
        return "ОШИБКА: неизвестная команда '" + op + "'";
    }
}

void load(DB& db, const string& path) {
    ifstream in(path);
    string line;
    while (getline(in, line))
        if (!line.empty())
            exec(db, line);
}

void save(const string& path, const string& cmd) {
    ofstream out(path, ios::app);
    out << cmd << "\n";
}
