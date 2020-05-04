import React from 'react'
import { StyleSheet, View } from 'react-native'
import { Icon } from '../common'
import { BaseHighlight } from '../common/index'
import { colors } from '../../assets/colors'

export default function ContentPreviewFooter({ kind, content, onDeletePress, onSavePress, onEditTagPress, onForwardPress }) {

    const middleIcon = () => {
        switch (kind) {
            case 'home': return (
                <BaseHighlight onPress={() => onSavePress(content)} style={styles.middleIconContainer}>
                    <Icon type='save' size={25} color={colors.darkGray} />
                </BaseHighlight>
            )
            default: return (
                <BaseHighlight onPress={() => onEditTagPress(content)} style={styles.middleIconContainer}>
                    <Icon type='tag' size={30} color={colors.darkGray} />
                </BaseHighlight>
            )
        }
    }

    return (
        <View style={styles.container}>
            <BaseHighlight onPress={() => onDeletePress(content)} style={styles.leftIconContainer}>
                <Icon type='trash' size={30} color={colors.darkGray} />
            </BaseHighlight>
            {middleIcon()}
            <BaseHighlight onPress={() => onForwardPress(content)} style={styles.rightIconContainer}>
                <Icon type='forward' size={25} color={colors.darkGray} />
            </BaseHighlight>
        </View>
    )
}

const styles = StyleSheet.create({
    container: {
        paddingLeft: 10,
        paddingRight: 10,
        flexDirection: 'row',
        justifyContent: 'space-between',
    },
    leftIconContainer: {
        flex: 1,
        alignItems: 'flex-start',
        justifyContent: 'center',
    },
    middleIconContainer: {
        flex: 1,
        alignItems: 'center',
        justifyContent: 'center',
    },
    rightIconContainer: {
        flex: 1,
        alignItems: 'flex-end',
        justifyContent: 'center',
    }
})