import React, { useEffect, useState } from 'react'
import { View } from 'react-native'
import BottomSheet from './BottomSheet'
import { colors } from '../../assets/colors'
import ButtonGroup from '../common/ButtonGroup'

export default function ButtonBottomSheet({ isVisible, onHide, children }) {

    const [visible, setVisible] = useState(isVisible)

    useEffect(() => { setVisible(isVisible) }, [isVisible])

    return (
        <BottomSheet
            visible={visible}
            onHide={() => onHide()}
        >
            <View style={{ backgroundColor: colors.primary, paddingBottom: 10 }}>
                <ButtonGroup>
                    {children}
                </ButtonGroup>
            </View>
        </BottomSheet>
    )
}