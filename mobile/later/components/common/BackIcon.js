import React from 'react'
import Icon from './Icon'
import { colors } from '../../assets/colors'

const BackIcon = ({ navigation, color }) => {
    return <Icon type='back' color={color ? color : colors.white} size={25} onPress={() => navigation.pop()} />
}

export default BackIcon