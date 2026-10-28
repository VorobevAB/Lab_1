#pragma once
#include <iostream>

using namespace std;

template <typename T>
struct LNode {
    T value;
    LNode<T>* next = nullptr;
    LNode<T>* prev = nullptr;
};

template <typename T>
struct List {
    LNode<T>* head = nullptr;
    LNode<T>* tail = nullptr;

    void LPushFront(const T& v) {
        LNode<T>* n = new LNode<T>{ v, head, nullptr };
        if (head)
            head->prev = n;
        else
            tail = n;
        head = n;
    }

    void LPushBack(const T& v) {
        LNode<T>* n = new LNode<T>{ v, nullptr, tail };
        if (tail)
            tail->next = n;
        else
            head = n;
        tail = n;
    }

    void LPushBefore(LNode<T>* node, const T& v) {
        if (node == nullptr || head == node) {
            LPushFront(v);
            return;
        }
        LNode<T>* cur = head;
        while (cur && cur->next != node)
            cur = cur->next;
        if (cur) {
            LNode<T>* n = new LNode<T>{ v, node, node->prev };
            node->prev->next = n;
            node->prev = n;
        }
        else
            cout << "ОШИБКА: Узла " << node->value << " нет в списке" << endl;
    }

    void LPushAfter(LNode<T>* node, const T& v) {
        if (node == nullptr) {
            LPushBack(v);
            return;
        }
        LNode<T>* cur = head;
        while (cur && cur != node)
            cur = cur->next;
        if (cur) {
            LNode<T>* n = new LNode<T>{ v, node->next, node };
            if (node->next)
                node->next->prev = n;
            else
                tail = n;
            node->next = n;
        }
        else
            cout << "ОШИБКА: Узла " << node->value << " нет в списке" << endl;
    }

    void LDelFront() {
        if (!head) {
            cout << "ОШИБКА: В списке нет элементов, удалять нечего" << endl;
            return;
        }
        LNode<T>* d = head;
        head = d->next;
        if (head)
            head->prev = nullptr;
        else
            tail = nullptr;
        delete d;
    }

    void LDelBack() {
        if (!tail) {
            cout << "ОШИБКА: В списке нет элементов, удалять нечего" << endl;
            return;
        }
        LNode<T>* d = tail;
        tail = d->prev;
        if (tail)
            tail->next = nullptr;
        else
            head = nullptr;
        delete d;
    }

    void LDelNode(LNode<T>* node) {
        if (node == nullptr) {
            cout << "ОШИБКА: Узел не может быть nullptr" << endl;
            return;
        }
        LNode<T>* cur = head;
        while (cur && cur != node)
            cur = cur->next;
        if (!cur) {
            cout << "ОШИБКА: Узла " << node->value << " нет в списке" << endl;
            return;
        }
        if (node->prev)
            node->prev->next = node->next;
        else
            head = node->next;
        if (node->next)
            node->next->prev = node->prev;
        else
            tail = node->prev;
        delete node;
    }

    void LDelAfter(LNode<T>* node) {
        if (node == nullptr) {
            cout << "ОШИБКА: Узел не может быть nullptr" << endl;
            return;
        }
        else if (node == tail) {
            cout << "ОШИБКА: За последним элементом, нет элементов для удаления" << endl;
            return;
        }
        LNode<T>* cur = head;
        while (cur && cur != node)
            cur = cur->next;
        if (!cur) {
            cout << "ОШИБКА: Узла " << node->value << " нет в списке" << endl;
            return;
        }
        else if (cur->next == tail)
            LDelBack();
        else {
            LNode<T>* d = cur->next;
            cur->next = d->next;
            cur->next->prev = d->prev;
            delete d;
        }
    }

    void LDelBefore(LNode<T>* node) {
        if (node == nullptr) {
            cout << "ОШИБКА: Узел не может быть nullptr" << endl;
            return;
        }
        else if (node == head) {
            cout << "ОШИБКА: Перед первым элементом, нет элементов для удаления" << endl;
            return;
        }
        LNode<T>* cur = head;
        while (cur && cur != node)
            cur = cur->next;
        if (!cur) {
            cout << "ОШИБКА: Узла " << node->value << " нет в списке" << endl;
            return;
        }
        else if (cur->prev == head)
            LDelFront();
        else {
            LNode<T>* d = cur->prev;
            cur->prev = d->prev;
            cur->prev->next = d->next;
            delete d;
        }
    }

    void LDelValue(const T& v) {
        LNode<T>* cur = head;
        while (cur && cur->value != v)
            cur = cur->next;
        if (cur)
            LDelNode(cur);
        else
            cout << "ОШИБКА: Значения " << v << " нет в списке" << endl;
    }

    void LFind(const T& v) {
        int i = 0;
        for (LNode<T>* c = head; c; c = c->next, ++i) {
            if (c->value == v) {
                cout << "Элемент был найден на позиции " << i << endl;
                return;
            }
        }
        cout << "Элемент не был найден в списке" << endl;
    }

    void LPrint() {
        cout << "L[";
        for (LNode<T>* c = head; c; c = c->next)
            cout << c->value << " ";
        cout << "]" << endl;
    }

    void LPrintReverse() {
        cout << "Lr[";
        for (LNode<T>* c = tail; c; c = c->prev)
            cout << c->value << " ";
        cout << "]" << endl;
    }

    void LPrintRec(LNode<T>* n) {
        if (!n)
            return;
        cout << n->value << " ";
        LPrintRec(n->next);
    }

    void clear() {
        LNode<T>* c = head;
        while (c) {
            LNode<T>* n = c->next;
            delete c;
            c = n;
        }
        head = nullptr;
        tail = nullptr;
    }
};
