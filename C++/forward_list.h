#pragma once
#include <iostream>

using namespace std;

template <typename T>
struct Node {
    T value;
    Node<T>* next = nullptr;
};

template <typename T>
struct ForwardList {
    Node<T>* head = nullptr;

    void FPushFront(const T& v) {
        head = new Node<T>{ v, head };
    }

    void FPushBack(const T& v) {
        Node<T>* n = new Node<T>{ v, nullptr };
        if (!head) {
            head = n;
            return;
        }
        Node<T>* cur = head;
        while (cur->next)
            cur = cur->next;
        cur->next = n;
    }

    void FPushAfter(Node<T>* node, const T& v) {
        if (node == nullptr) {
            FPushBack(v);
            return;
        }
        Node<T>* cur = head;
        while (cur && cur != node)
            cur = cur->next;
        if (cur)
            cur->next = new Node<T>{ v, cur->next };
        else
            cout << "ОШИБКА: Узла " << node->value << " нет в списке" << endl;
    }

    void FPushBefore(Node<T>* node, const T& v) {
        if (head == node || node == nullptr) {
            FPushFront(v);
            return;
        }
        Node<T>* cur = head;
        while (cur && cur->next != node)
            cur = cur->next;
        if (cur)
            cur->next = new Node<T>{ v, cur->next };
        else
            cout << "ОШИБКА: Узла " << node->value << " нет в списке" << endl;
    }

    void FDelFront() {
        if (!head) {
            cout << "ОШИБКА: В списке нет элементов, удалять нечего" << endl;
            return;
        }
        Node<T>* d = head;
        head = d->next;
        delete d;
    }

    void FDelBack() {
        if (!head) {
            cout << "ОШИБКА: В списке нет элементов, удалять нечего" << endl;
            return;
        }
        else if (!head->next) {
            delete head;
            head = nullptr;
            return;
        }
        Node<T>* cur = head;
        while (cur->next->next)
            cur = cur->next;
        delete cur->next;
        cur->next = nullptr;
    }

    void FDelAfter(Node<T>* node) {
        if (node == nullptr) {
            cout << "ОШИБКА: Узел не может быть nullptr" << endl;
            return;
        }
        Node<T>* cur = head;
        while (cur && cur != node)
            cur = cur->next;
        if (cur) {
            if (cur->next == nullptr) {
                cout << "ОШИБКА: После узла " << node->value << " ничего нет" << endl;
                return;
            }
            Node<T>* d = cur->next;
            cur->next = d->next;
            delete d;
        }
        else {
            cout << "ОШИБКА: Узла " << node->value << " нет в списке" << endl;
        }
    }

    void FDelBefore(Node<T>* node) {
        if (node == nullptr) {
            cout << "ОШИБКА: Узел не может быть nullptr" << endl;
            return;
        }
        else if (head == node) {
            cout << "ОШИБКА: Перед головой ничего нет" << endl;
            return;
        }
        else if (head->next == node) {
            FDelFront();
            return;
        }
        Node<T>* cur = head;
        while (cur && cur->next && cur->next->next != node)
            cur = cur->next;
        if (cur) {
            Node<T>* d = cur->next;
            cur->next = d->next;
            delete d;
        }
        else {
            cout << "ОШИБКА: Узла " << node->value << " нет в списке" << endl;
        }
    }

    void FDelNode(Node<T>* node) {
        if (node == nullptr) {
            cout << "ОШИБКА: Узел не может быть nullptr" << endl;
            return;
        }
        else if (head == node) {
            FDelFront();
            return;
        }
        Node<T>* cur = head;
        while (cur && cur->next != node)
            cur = cur->next;
        if (cur) {
            cur->next = node->next;
            delete node;
        }
        else {
            cout << "ОШИБКА: Узла " << node->value << " нет в списке" << endl;
        }
    }

    void FDelValue(const T& v) {
        Node<T>* cur = head;
        while (cur && cur->value != v)
            cur = cur->next;
        if (cur)
            FDelNode(cur);
        else
            cout << "ОШИБКА: Значения " << v << " нет в списке" << endl;
    }

    void FFind(const T& v) {
        int i = 0;
        for (Node<T>* c = head; c; c = c->next, ++i) {
            if (c->value == v) {
                cout << "Элемент был найден на позиции " << i << endl;
                return;
            }
        }
        cout << "Элемент не был найден в списке" << endl;
    }

    void FPrint() {
        cout << "F[";
        for (Node<T>* c = head; c; c = c->next)
            cout << c->value << " ";
        cout << "]" << endl;
    }

    void FPrintRec(Node<T>* n) {
        if (!n)
            return;
        cout << n->value << " ";
        FPrintRec(n->next);
    }

    void clear() {
        while (head) {
            Node<T>* n = head->next;
            delete head;
            head = n;
        }
    }
};
