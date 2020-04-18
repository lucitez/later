import React from 'react'
import Icon from './Icon'
import { colors } from '../../assets/colors'

const BackIcon = ({ navigation }) => {
    return <Icon type='back' color={colors.white} size={25} onPress={() => navigation.pop()} />
}

export default BackIcon