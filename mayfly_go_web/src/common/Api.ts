import request from './request'

/**
 * 可用于各模块定义各自api请求
 */
class Api {
    /**
     * 请求url
     */
    url: string;

    /**
     * 请求方法
     */
    method: string;

    constructor(url: string, method: string) {
        this.url = url;
        this.method = method;
    }

    /**
     * 获取权限的完整url
     */
    getUrl() {
        return request.getApiUrl(this.url);
    }

    /**
     * 操作该权限，即请求对应的url
     * @param {Object} param 请求该权限的参数
     */
    request(param: any = null, options: any = null): Promise<any> {
        return request.send(this, param, options);
    }

    /**
    * 操作该权限，即请求对应的url
    * @param {Object} param 请求该权限的参数
    */
    requestWithHeaders(param: any, headers: any): Promise<any> {
        return request.sendWithHeaders(this, param, headers);
    }


    /**    静态方法     **/

    /**
     * 静态工厂，返回Api对象，并设置url与method属性
     * @param url url
     * @param method 请求方法(get,post,put,delete...)
     */
    static create(url: string, method: string) {
        return new Api(url, method);
    }
}


export default Api
