import React from 'react'
import { StyleSheet, View, Text, SafeAreaView } from 'react-native'
import { Icon } from '../../components/common'
import { colors } from '../../assets/colors'
import { ScrollView } from 'react-native-gesture-handler'

function OnboardingFormWrapper({ leftIcon, rightIcon, title, description, children }) {
    return (
        <SafeAreaView style={{ flex: 1 }}>
            <ScrollView contentContainerStyle={styles.container} keyboardShouldPersistTaps='handled'>
                <View style={styles.header}>
                    <View style={styles.leftIconContainer}>
                        {leftIcon}
                    </View>
                    <View style={styles.rightIconContainer}>
                        {rightIcon}
                    </View>
                </View>
                <View style={styles.details}>
                    <View style={styles.logoContainer}>
                        <Icon type='share' size={50} color={colors.primary} />
                    </View>
                    <View style={styles.titleContainer}>
                        <Text style={styles.title}>{title}</Text>
                    </View>
                    {description && <View style={styles.descriptionContainer}>
                        <Text style={styles.description}>{description}</Text>
                    </View>}
                </View>
                <View style={styles.formContainer}>
                    <View style={styles.formContentsContainer}>
                        {children}
                    </View>
                </View>
            </ScrollView>
        </SafeAreaView>
    )
}

const styles = StyleSheet.create({
    container: {
        flex: 1,
        backgroundColor: colors.lightGray,
    },
    header: {
        height: '7%',
        flexDirection: 'row',
        justifyContent: 'space-between',
        alignItems: 'center',
    },
    leftIconContainer: {
        flex: 1,
        alignItems: 'flex-start',
        paddingLeft: 20
    },
    rightIconContainer: {
        flex: 1,
        alignItems: 'flex-end',
        paddingRight: 20
    },
    details: {
        alignItems: 'center',
        marginBottom: 30,
    },
    logoContainer: {
        padding: 10,
    },
    titleContainer: {
        padding: 10,
    },
    descriptionContainer: {
        width: '80%',
        padding: 5,
    },
    title: {
        fontSize: 24,
        fontWeight: '300'
    },
    description: {
        fontSize: 16,
        fontWeight: '300',
        textAlign: 'center',
    },
    formContainer: {
        alignItems: 'center',
    },
    formContentsContainer: {
        width: '80%',
        backgroundColor: colors.white,
        borderRadius: 10,
        padding: 20,
    },
})

export default OnboardingFormWrapper