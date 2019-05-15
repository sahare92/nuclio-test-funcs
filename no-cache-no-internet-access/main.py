def handler(context, event):
    num_sum = sum(event.body.values())
    return context.Response(body=str(num_sum), headers=None, content_type='text/plain', status_code=200)