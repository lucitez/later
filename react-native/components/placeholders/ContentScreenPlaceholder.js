import React from 'react'
import { View, Text, StyleSheet } from 'react-native'
import ContentPreviewPlaceholder from './ContentPreviewPlaceholder'
import { TouchableOpacity } from 'react-native-gesture-handler'
import { colors } from '../../assets/colors'
import { Icon } from '../common'

export default function ContentScreenPlaceholder({ navigation, kind }) {

    const HomeMessage = () => (
        <>
            <Text style={styles.title}>Looks like you don't have any new content.</Text>
            <Text style={styles.subtitle}>Content that friends send you will show up here.</Text>
            <TouchableOpacity style={styles.buttonContainer} onPress={() => navigation.navigate('Search')}>
                <Text style={styles.button}>Add and invite friends</Text>
            </TouchableOpacity>
        </>
    )

    const SavedMessage = () => (
        <>
            <Text style={styles.title}>
                Looks like you don't have any saved content.
            </Text>
            <Text style={styles.subtitle}>Save content by hitting the inbox icon.</Text>
        </>
    )

    const _renderMessage = () => {
        switch (kind) {
            case 'home':
                return <HomeMessage />
            case 'saved':
                return <SavedMessage />
        }
    }

    return (
        <View>
            <ContentPreviewPlaceholder />
            <View style={styles.messageContainer}>
                {_renderMessage()}
            </View>
        </View>

    )
}

const styles = StyleSheet.create({
    messageContainer: {
        alignItems: 'center',
        padding: 20
    },
    title: {
        textAlign: 'center',
        fontSize: 18,
        fontWeight: '600',
        marginBottom: 10,
    },
    subtitle: {
        textAlign: 'center',
        fontSize: 14,
        fontWeight: '300',
        marginBottom: 30,
    },
    buttonContainer: {
        padding: 15,
        paddingTop: 10,
        paddingBottom: 10,
        backgroundColor: colors.primary,
        borderRadius: 50,
    },
    button: {
        color: colors.white,
        fontSize: 16,
        fontWeight: '400'
    }
})