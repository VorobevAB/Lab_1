#include <iostream>
#include <string>
#include "dbms.h"

using namespace std;

int main(int argc, char** argv) {
    string file, query;
    for (int i = 1; i < argc; ++i) {
        string a = argv[i];
        if (a == "--file" && i + 1 < argc) file = argv[++i];
        if (a == "--query" && i + 1 < argc) query = argv[++i];
    }

    DB db;
    if (!file.empty())
        load(db, file);

    if (!query.empty()) {
        string res = exec(db, query);
        cout << res << endl;
        if (!file.empty() && res.rfind("ОШИБКА", 0) != 0)
            save(file, query);
    }
    return 0;
}
