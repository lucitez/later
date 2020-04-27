import React from 'react'
import { StyleSheet, View, Text, SafeAreaView } from 'react-native'
import { colors } from '../assets/colors'

export default ChatDisplayScreen = ({ navigation, route }) => {
    let chatDetails = route.params.chatDetails

    return (
        <SafeAreaView style={styles.container}>
            <View><Text>hi</Text></View>
        </SafeAreaView>
    )
}

const styles = StyleSheet.create({
    container: {
        flex: 1,
        backgroundColor: colors.lightGray,
    }
})