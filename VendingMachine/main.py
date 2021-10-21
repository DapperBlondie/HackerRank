# This is a sample Python script.

# Press Shift+F10 to execute it or replace it with your code.
# Press Double Shift to search everywhere for classes, files, tool windows, actions, and settings.

class VendingMachine:
    def __init__(self, num_items, item_price):
        self.NumItems = num_items
        self.ItemPrice = item_price

    def buy(self, num_items, num_coins) -> int:
        if self.NumItems < num_items:
            raise ValueError("Not enough items in the machine")
        if self.ItemPrice*num_items > num_coins:
            raise ValueError("Not enough coins")

        self.NumItems -= num_items

        return num_coins - num_items*self.ItemPrice

def transformSentence(sentence: str) -> str:
    inputs = sentence.split(sep=' ')
    outputs = []
    for input in inputs:
        output = input[0]
        for i in range(1, len(input)):
            if input[i] > input[i - 1]:
                output += input[i].upper()
            elif input[i] < input[i - 1]:
                output += input[i].lower()
            else:
                output += input[i]
        outputs.append(output)

    return ' '.join(outputs)
# Press the green button in the gutter to run the script.
if __name__ == '__main__':
    i = 'c'
    h = 'b'
    print(i > h)
    input = "a blue MOON"
    print(transformSentence(input))
# See PyCharm help at https://www.jetbrains.com/help/pycharm/
