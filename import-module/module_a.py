class TestClass:
    def __init__(self, event):
        self.nums = event.body.values()

    def get_sum(self):
        return sum(self.nums)
