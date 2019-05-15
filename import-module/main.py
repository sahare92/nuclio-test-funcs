import module_a


def handler(context, event):
    get_sum = module_a.TestClass(event).get_sum()
    return context.Response(body=str(get_sum), headers=None, content_type='text/plain', status_code=200)
