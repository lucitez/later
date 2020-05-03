import axios from 'axios';
import _ from 'lodash';
import store from '../store'
import env from '../environment'

const client = axios.create({
    baseURL: env.apiUrl,
    timeout: 10000
})

client.interceptors.response.use(response => {
    return { ...response, ['data']: dataToCamelCase(response.data) }
})

const addAuthHeaders = (url, headers) => {
    if (url.startsWith("/auth")) {
        return {
            ...headers,
            'Client-ID': env.clientId
        }
    } else {
        let accessToken = store.getState().auth.tokens.accessToken
        return {
            ...headers,
            'Authorization': `Basic ${accessToken}`
        }
    }
}

function request(options) {
    options = {
        ...options,
        headers: addAuthHeaders(options.url, options.headers),
    }
    let currTime = Date.now()
    return new Promise((resolve, reject) => {
        client.request(options)
            .then(response => {
                console.log(`Millis elapsed: ${Date.now() - currTime}`)
                return resolve(response.data)
            })
            .catch((error) => {
                console.log(`Millis elapsed: ${Date.now() - currTime}`)
                if (error.response) {
                    // The request was made and the server responded with a status code
                    // that falls out of the range of 2xx
                    console.log(error.response.data);
                    console.log(error.response.status);
                    console.log(error.response.headers);
                    return reject(error.response.data)
                } else if (error.request) {
                    // The request was made but no response was received
                    // `error.request` is an instance of XMLHttpRequest in the browser and an instance of
                    // http.ClientRequest in node.js
                    console.log(error.request);
                    return reject(error.request)
                } else {
                    // Something happened in setting up the request that triggered an Error
                    console.log('Error', error.message);
                    return reject(error.message)
                }
            })
    });
}

const objToSnakeCase = (obj) => {
    let snakeCaseParams = {}

    for (let [key, value] of Object.entries(obj)) {
        snakeCaseParams[_.snakeCase(key)] = value
    }

    return snakeCaseParams
}

const objToCamelCase = (data) => {
    if (typeof data == 'object') {
        let camelCaseParams = {}

        for (let [key, value] of Object.entries(data)) {
            if (typeof value == 'object' && value) {
                value = objToCamelCase(value)
            }
            camelCaseParams[_.camelCase(key)] = value
        }

        return camelCaseParams
    }

    return data
}

// THIS IS BRITTLE AF
const dataToCamelCase = (data) => {
    if (!data) {
        return data
    }
    switch (typeof data) {
        case 'object':
            if (Array.isArray(data)) {
                return data.map(obj => objToCamelCase(obj))
            } else {
                return objToCamelCase(data)
            }
        default:
            return data
    }
}

const Network = {
    POST(url, body = {}, headers = {}) {
        let options = {
            method: 'post',
            url: url,
            headers: headers,
            data: objToSnakeCase(body),
        }
        console.log(options)
        return request(options)
    },
    PUT(url, body = {}) {
        let options = {
            method: 'put',
            url: url,
            data: objToSnakeCase(body),
        }
        return request(options)
    },
    GET(url, params = {}) {
        let options = {
            method: 'get',
            url: url,
            params: objToSnakeCase(params),
        }
        return request(options)
    }
}

export default Network