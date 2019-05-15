class A(object):

    @staticmethod
    def get_nums(event):
        return event.body.values()
