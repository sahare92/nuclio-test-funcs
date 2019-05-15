import os


def handler(context, event):
    if os.getenv('NAIPI_VAR') != 'naipi_var':
        raise Exception('Missing env variable NAIPI_VAR')
    num_sum = sum(event.body.values())
    return context.Response(body=str(num_sum), headers=None, content_type='text/plain', status_code=200)