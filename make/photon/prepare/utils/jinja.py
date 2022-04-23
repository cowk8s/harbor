import json

from jinja2 import Environment, FileSystemLoader
from .misc import mark_file

jinja_env = Environment(loader=FileSystemLoader('/'), trim_blocks=True, lstrip_blocks=True)

def to_json(value):
    return json.dumps(value)

jinja_env.filters['to_json'] = to_json

def render_jinja(src, dest,mode=0o640)