const host = {
    local: 'http://192.168.254.64:8000'
}

/**
 * Parses the JSON returned by a network request
 *
 * @param  {object} response A response from a network request
 *
 * @return {object}          The parsed JSON, status from the response
 */
function parseJSON(response) {
    return new Promise((resolve) => response.json()
        .then((json) => resolve({
            status: response.status,
            ok: response.ok,
            json,
        }))
        .catch(() => {
            resolve({
                status: 500,
                ok: false,
                json: { error: "Something went wrong" }
            })
        })
    );
}

/**
 * Requests a URL, returning a promise
 *
 * @param  {string} endpoint     The URL we want to request
 * @param  {object} [options] The options we want to pass to "fetch"
 *
 * @return {Promise}           The request promise
 */
function request(endpoint, options) {
    return new Promise((resolve, reject) => {
        fetch(host.local + endpoint, options)
            .then(parseJSON)
            .then((response) => {
                if (response.ok) {
                    return resolve(response.json);
                }
                // extract the error from the server's json
                return reject({
                    status: response.status,
                    error: response.json.error
                });
            })
            .catch((error) => reject({
                networkError: error.message,
            }));
    });
}

const Network = {
    POST(endpoint, body, headers = {}) {
        options = {
            method: 'POST',
            headers: headers,
            body: JSON.stringify(body)
        }
        return request(endpoint, options)
    },
    PUT(endpoint, body, headers = {}) {
        options = {
            method: 'GET',
            headers: headers,
            body: JSON.stringify(body)
        }
        return request(endpoint, options)
    },
    GET(endpoint, headers = {}) {
        options = {
            method: 'GET',
            headers: headers,
        }
        return request(endpoint, options)
    }
}

export default Network