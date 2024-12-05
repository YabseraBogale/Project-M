from chapa import Chapa
from chapa import AsyncChapa
from chapa import get_testing_cards, get_testing_mobile 

# Get a list of testing cards
test_cards = get_testing_cards()
print(test_cards)

# Get a list of testing mobile numbers
test_mobiles = get_testing_mobile()
print(test_mobiles)
