export class SessionUserBase {
    user_id: number;
    username: string;
    email: string;
    realname: string;
    role_name?: string;
    role_id?: number;
    comment: string;
    oidc_user_meta?: OidcUserMeta;
}

export class OidcUserMeta {
    id: number;
    user_id: number;
    secret: string;
    subiss: string;
    creation_time: Date;
    update_time: Date;
}