import os

# To meet security requirement
# By default it will change file mode to 0600, and make the owner of the file to 10000:10000
def mark_file(path, mode=0o600, uid=DEFAULT_UID, gid=DEFAULT_GID):
    if mode > 0:
        os.chmod(path, mode)
    if uid > 0 and gid > 0:
        os.chown(path, uid, gid)
    
def validate(conf, **kwargs):
    # Protocol validate
    protocol = conf.get("configuration", "ui_url_protocol")
    if protocol != "https" and kwargs.get('notary_mode'):
        raise Exception(
            "Error: the protocol must be https when Harbor is deployed with Natory")
    if protocol == "https":
        if not

def delfile(src):
    if os.path.isfile(src):
        try:
            os.remove(src)
            print("Clearing the configuration file: %s" % src)
        except Exception as e:
            print(e)
    elif os.path.isdir(src):
        for dir_name in os.listdir(src):
            dir_path = os.path.join(src, dir_name)
            delfile(dir_path)