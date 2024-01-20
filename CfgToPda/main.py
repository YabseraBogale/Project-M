class CFGtoPDAConverter:
    def init(self, cfg):
        self.cfg = cfg
        self.pda_states = set()
        self.pda_alphabet = set()
        self.pda_stack_alphabet = set()
        self.pda_transitions = []
        self.pda_start_state = None
        self.pda_stack_start_symbol = None
        self.pda_accept_states = set()

    def convert(self):
        self.extract_pda_details()
        self.generate_pda_transitions()

    def extract_pda_details(self):
        # Extract non-terminals, terminals, and productions from the CFG
        non_terminals = set(self.cfg.keys())
        terminals = set()
        productions = set()

        for non_terminal, expressions in self.cfg.items():
            for expr in expressions:
                for symbol in expr:
                    if symbol.isalpha() and symbol.islower():
                        terminals.add(symbol)
                    productions.add((non_terminal, tuple(expr)))

        # PDA details
        self.pda_states = non_terminals.union({"q_accept", "q_reject"})
        self.pda_alphabet = terminals
        self.pda_stack_alphabet = terminals.union({"Z0"})
        self.pda_start_state = "S"
        self.pda_stack_start_symbol = "Z0"
        self.pda_accept_states = {"q_accept"}

    def generate_pda_transitions(self):
        # Generate PDA transitions based on CFG productions
        for non_terminal, expressions in self.cfg.items():
            for expr in expressions:
                self.pda_transitions.append((non_terminal, expr))

        # Add transitions for acceptance and rejection
        for terminal in self.pda_alphabet:
            self.pda_transitions.append(("S", (terminal, "S")))
            self.pda_transitions.append(("S", (terminal,)))
            self.pda_transitions.append(("S", ("",)))

        # Add transitions to acceptance and rejection states
        for terminal in self.pda_alphabet:
            self.pda_transitions.append(("S", (terminal, "q_accept")))
            self.pda_transitions.append(("S", (terminal, "q_reject")))

    def display_pda(self):
        print("PDA States:", self.pda_states)
        print("PDA Alphabet:", self.pda_alphabet)
        print("PDA Stack Alphabet:", self.pda_stack_alphabet)
        print("PDA Transitions:", self.pda_transitions)
        print("PDA Start State:", self.pda_start_state)
        print("PDA Stack Start Symbol:", self.pda_stack_start_symbol)
        print("PDA Accept States:", self.pda_accept_states)

# Example CFG
cfg_example = {
    "S": [["a", "S", "b"], [""]],
}

# Convert CFG to PDA
converter = CFGtoPDAConverter(cfg_example)
converter.convert()

# Display PDA details
converter.display_pda()