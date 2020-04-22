import React from 'react'
import PlainText from './PlainText'


export default function Email(props) {

    return <PlainText {...props} isValid={isValid} />

}

const isValid = (value) => {
    const pattern = /^(([^<>()\[\]\\.,;:\s@"]+(\.[^<>()\[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/

    return pattern.test(value)
}