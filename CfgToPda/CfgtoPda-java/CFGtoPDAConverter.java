import java.util.HashSet;
import java.util.Set;

public class CFGtoPDAConverter {

    private final Set<String> cfg;
    private Set<String> pdaStates;
    private Set<Character> pdaAlphabet;
    private Set<Character> pdaStackAlphabet;
    private Set<PDATransition> pdaTransitions;
    private String pdaStartState;
    private Character pdaStackStartSymbol;
    private Set<String> pdaAcceptStates;

    public CFGtoPDAConverter(Set<String> cfg) {
        this.cfg = cfg;
        pdaStates = new HashSet<>();
        pdaAlphabet = new HashSet<>();
        pdaStackAlphabet = new HashSet<>();
        pdaTransitions = new HashSet<>();
        pdaAcceptStates = new HashSet<>();
    }

    public void convert() {
        extractPDADetails();
        generatePDATransitions();
    }

    private void extractPDADetails() {
        // Extract non-terminals, terminals, and productions from the CFG
        Set<String> nonTerminals = cfg.keySet();
        Set<Character> terminals = new HashSet<>();
        Set<Production> productions = new HashSet<>();

        for (String nonTerminal : cfg.keySet()) {
            for (String expr : cfg.get(nonTerminal)) {
                for (char symbol : expr.toCharArray()) {
                    if (Character.isLowerCase(symbol)) {
                        terminals.add(symbol);
                    }
                    productions.add(new Production(nonTerminal, expr.toCharArray()));
                }
            }
        }

        // PDA details
        pdaStates.addAll(nonTerminals);
        pdaStates.add("q_accept");
        pdaStates.add("q_reject");
        pdaAlphabet = terminals;
        pdaStackAlphabet.addAll(terminals);
        pdaStackAlphabet.add('Z'); // Using 'Z' instead of 'Z0' for a single character stack symbol
        pdaStartState = "S";
        pdaStackStartSymbol = 'Z';
        pdaAcceptStates.add("q_accept");
    }

    private void generatePDATransitions() {
        // Generate PDA transitions based on CFG productions
        for (Production production : productions) {
            pdaTransitions.add(new PDATransition(production.nonTerminal, production.expression));
        }

        // Add transitions for acceptance and rejection
        for (char terminal : pdaAlphabet) {
            pdaTransitions.add(new PDATransition("S", new char[]{terminal, 'S'}));
            pdaTransitions.add(new PDATransition("S", new char[]{terminal}));
            pdaTransitions.add(new PDATransition("S", new char[]{}));
            pdaTransitions.add(new PDATransition("S", new char[]{terminal}, "q_accept"));
            pdaTransitions.add(new PDATransition("S", new char[]{terminal}, "q_reject"));
        }
    }

    public void displayPDA() {
        System.out.println("PDA States: " + pdaStates);
        System.out.println("PDA Alphabet: " + pdaAlphabet);
        System.out.println("PDA Stack Alphabet: " + pdaStackAlphabet);
        System.out.println("PDA Transitions: " + pdaTransitions);
        System.out.println("PDA Start State: " + pdaStartState);
        System.out.println("PDA Stack Start Symbol: " + pdaStackStartSymbol);
        System.out.println("PDA Accept States: " + pdaAcceptStates);
    }

    public static void main(String[] args) {
        // Example CFG
        Set<String> cfgExample = new HashSet<>();
        cfgExample.add("S aSb");
        cfgExample.add("S");

        // Convert CFG to PDA
        CFGtoPDAConverter converter = new CFGtoPDAConverter(cfgExample);
        converter.convert();

        // Display PDA details
        converter.displayPDA();
    }

