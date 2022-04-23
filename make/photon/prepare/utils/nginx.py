import os, shutil

host_ngx_real_cert_dir = Path(os.path.join(data_dir, 'secret', 'cert'))

def prepare_nginx(config_dict):
    prepare_dir(nginx_conf_dir, uid=DEFAULT_UID, gid=DEFAULT_UID)
    render_nginx_template(config_dict)

def prepare_nginx_certs(cert_key_path, cert_path):
    """
    Prepare the certs file with proper owenership
    """
    