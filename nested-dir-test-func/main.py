import lib1
import lib2


def handler(context, event):
    nums = lib1.A.get_nums(event)
    sum_of_nums = lib2.B.get_sum(nums)
    return context.Response(body=str(sum_of_nums), headers=None, content_type='text/plain', status_code=200)
