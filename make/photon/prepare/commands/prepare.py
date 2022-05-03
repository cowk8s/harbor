# pylint: disable=no-value-for-parameter

import sys
import logging

import click

from utils.misc import delfile
from utils.configs import validate, parse_yaml_config
from g import (config_dir, input_config_path)

@click.command()
@click.option('--conf', default=input_config_path, help="the path of Harbor configuration file")
@click.option('--with-notary', is_flag=True, help="the Harbor instance is to be deployed with notary")
@click.option('--with-trivy', is_flag=True, help="the Harbor instance is to be deployed with Trivy")
@click.option('--with-chartmuseum', is_flag=True, help="the Harbor instance is to be deployed with chart repository supporting")
def prepare(conf, with_notary, with_trivy, with_chartmuseum):

    delfile(config_dir)
    config_dict = parse_yaml_config(conf, with_notary=with_notary, with_trivy=with_trivy, with_chartmuseum=with_chartmuseum)
    try:
        validate(config_dict, notary_mode=with_notary)
    except Exception as e:
        click.echo('Error happened in config validation...')
        logging.error(e)
        sys.exit(-1)

    prepare_portal(config_dir)