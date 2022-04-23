# pylint: disable=no-value-for-parameter

import sys
import logging

import click
from utils.misc import delfile

@click.command()
@click.option('--conf', default=input_config_path, help="the path of Harbor configuration file")
@click.option('--with-notary', is_flag=True, help="the Harbor instance is to be deployed with notary")
@click.option('--with-trivy', is_flag=True, help="the Harbor instance is to be deployed with Trivy")
def prepare(conf, with_notary, with_trivy, with_chartmuseum):

    delfile(config_dir)
    config_dir = parse_yaml_config(conf, with_notary=with_notary, with_trivy=with_trivy, with_chartmuseum=with_chartmuseum)
    try:
        validate(config_dir, notary_mode=with_notary)
    except Exception as e:
        click.echo('Error happened in config validation...')
        logging.error(e)
        sys.exit(-1)

    prepare_portal(config_dir)