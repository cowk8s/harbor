import click

@click.group()
def cli():
    pass

cli.add_command(prepare)
cli.add_command(gencert)
cli.add_command(migrate)

if __main__ = '__main__':
    cli()