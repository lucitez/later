import React from 'react'
import PlainText from './PlainText'

export default function Password(props) {
    return <PlainText {...props} secureTextEntry={true} />
}