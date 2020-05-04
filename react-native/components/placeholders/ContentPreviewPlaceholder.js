import React from 'react';
import { StyleSheet, View } from 'react-native';
import { colors } from '../../assets/colors';
import ContentPreviewHeaderPlaceholder from './ContentPreviewHeaderPlaceholder'

export default function ContentPreviewPlaceholder() {
    return (
        <View style={styles.container}>
            <ContentPreviewHeaderPlaceholder />
            <View style={styles.contentContainer}>
                <View style={{ flexBasis: 0, flexGrow: 1 }}>
                    <View style={styles.titlePlaceholder} />
                    <View style={[styles.titlePlaceholder, { width: '80%' }]} />
                    <View style={[styles.titlePlaceholder, { width: '30%' }]} />
                    <View style={styles.descriptionPlaceholder} />
                    <View style={[styles.descriptionPlaceholder, { width: '60%' }]} />
                </View>
                <View style={[styles.imageContainer, { aspectRatio: 1.5 }]} />
            </View>
        </View>
    );
}

const styles = StyleSheet.create({
    container: {
        paddingTop: 10,
        paddingBottom: 5,
    },
    contentContainer: {
        flexDirection: 'row',
        padding: 10,
        paddingTop: 0,
    },
    titlePlaceholder: {
        height: 13,
        backgroundColor: colors.darkGray,
        borderRadius: 15,
        margin: 5
    },
    descriptionPlaceholder: {
        height: 9,
        backgroundColor: colors.darkGray,
        opacity: 0.25,
        borderRadius: 15,
        margin: 5
    },
    imageContainer: {
        marginLeft: 10,
        height: 100,
        aspectRatio: 1.5,
        borderRadius: 20,
        backgroundColor: colors.darkGray,
        opacity: 0.5
    },
});