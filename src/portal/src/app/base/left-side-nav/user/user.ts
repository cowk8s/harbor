/**
 * For user management
 *
 **
 * class User
 */
export class User {
    user_id?: number;
    username?: string;
    realname?: string;
    email?: string;
    password?: string;
    comment?: string;
    deleted?: boolean;
    role_name?: string;
    role_id?: number;
    sysadmin_flag?: boolean;
    reset_uuid?: string;
    creation_time?: string;
    update_time?: string;
}