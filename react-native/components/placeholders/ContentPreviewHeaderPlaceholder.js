import React from 'react'
import { StyleSheet, View, Text } from 'react-native'
import { colors } from '../../assets/colors'
export default function ContentPreviewHeaderPlaceholder() {
    return (
        <View style={styles.bannerContainer}>
            <View style={styles.userImage} />
            <View style={{ flexGrow: 1, alignItems: 'flex-start' }}>
                <View style={{ flexDirection: 'row', alignItems: 'flex-end' }}>
                    <View style={{ width: '30%', height: 12, borderRadius: 10, backgroundColor: colors.darkGray, marginRight: 10, }} />
                    <View style={{ width: '25%', height: 10, borderRadius: 10, backgroundColor: colors.darkGray, opacity: 0.5 }} />
                </View>
            </View>
        </View>
    )
}

const styles = StyleSheet.create({
    bannerContainer: {
        padding: 10,
        flexDirection: 'row',
        justifyContent: 'flex-start',
        alignItems: 'center',
    },
    userImage: {
        height: 35,
        width: 35,
        marginRight: 5,
        borderRadius: 20,
        backgroundColor: colors.darkGray,
    },
    dotsContainer: {
        justifyContent: 'center',
        paddingLeft: 10,
        paddingBottom: 4,
    },
})