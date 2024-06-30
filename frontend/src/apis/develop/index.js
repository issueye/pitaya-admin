import request from '../axios';

/**
 * 获取所有表
 * @returns 
 */
export const apiGetAllTable = () => {
    return request({
        url: 'api/v1/develop/getAllTable',
        method: 'get',
    })
}


/**
 * 获取所有表
 * @returns 
 */
export const apiGetTableInfo = (name) => {
    return request({
        url: `api/v1/develop/getTableColumns/${name}`,
        method: 'get',
    })
}
