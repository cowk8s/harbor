
/**
 * Deeper clone all
 *
 **
 *  ** deprecated param {*} srcObj
 * returns {*}
 */
export function clone(srcObj: any): any {
    if (!srcObj) { return null; }
    return JSON.parse(JSON.stringify(srcObj));
}