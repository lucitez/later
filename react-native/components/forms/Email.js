import React from 'react'
import PlainText from './PlainText'

export default function Email(props) {
    return <PlainText {...props} hasError={value => hasError(props.required, value)} />
}

const hasError = (required, value) => {
    if (required && value == '') {
        return 'email is required'
    }

    const pattern = /^(([^<>()\[\]\\.,;:\s@"]+(\.[^<>()\[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/

    let valid = pattern.test(value)

    if (valid) {
        return null
    }

    return 'Invalid email format'
}