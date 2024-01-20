#include <iostream>
#include <vector>
#include <unordered_map>
#include <unordered_set>
#include <tuple>

using namespace std;

class PDATransition {
public:
    string state;
    string input;
    string pop;
    string push;
    string nextState;

    PDATransition(string s, string i, string p, string ps, string ns) : state(s), input(i), pop(p), push(ps), nextState(ns) {}
};

class PDATransitionTable {
public:
    vector<PDATransition> transitions;
    string startState;
    unordered_set<string> acceptStates;
};

class PDAGenerator {
private:
    unordered_map<string, vector<string>> cfg;
    int currentStateCounter;
    int nextStateCounter;

public:
    PDAGenerator(unordered_map<string, vector<string>> cfg) : cfg(cfg), currentStateCounter(0), nextStateCounter(0) {}

    PDATransitionTable convert() {
        PDATransitionTable pda;
        pda.startState = getNextState();
        pda.acceptStates.insert(getNextState());

        for (const auto& entry : cfg) {
            const string& nonTerminal = entry.first;
            const vector<string>& productions = entry.second;

            for (const string& production : productions) {
                addTransition(pda, pda.startState, "", "", nonTerminal, production);
            }
        }

        return pda;
    }

private:
    void addTransition(PDATransitionTable& pda, const string& currentState, const string& currentInput,
                       const string& currentPop, const string& nextState, const string& pushSymbols) {
        pda.transitions.emplace_back(currentState, currentInput, currentPop, pushSymbols, nextState);
    }

    string getNextState() {
        return "q" + to_string(nextStateCounter++);
    }
};

int main() {
    // Example CFG
    unordered_map<string, vector<string>> cfgExample = {
        {"S", {"aSb", ""}}
    };

    // Convert CFG to PDA
    PDAGenerator pdaGenerator(cfgExample);
    PDATransitionTable pda = pdaGenerator.convert();

    // Display PDA details
    cout << "PDA Transition Table:" << endl;
    for (const auto& transition : pda.transitions) {
        cout << transition.state << " | " << transition.input << " | " << transition.pop << " | "
             << transition.push << " | " << transition.nextState << endl;
    }

    cout << "Start State: " << pda.startState << endl;
    cout << "Accept States: ";
    for (const auto& acceptState : pda.acceptStates) {
        cout << acceptState << " ";
    }
    cout << endl;

    return 0;
}