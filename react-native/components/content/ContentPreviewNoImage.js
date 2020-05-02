import React from 'react';
import { StyleSheet, Text, View } from 'react-native'
import { Link } from '../common'
import { colors } from '../../assets/colors'
import ContentPreviewHeader from './ContentPreviewHeader'

function ContentPreviewNoImage({ onDotPress, onTagPress, content, linkActive }) {

    return (
        <View style={{ backgroundColor: colors.white, paddingTop: 5, paddingBottom: 5 }}>
            <View style={styles.topContainer}>
                <ContentPreviewHeader content={content} onDotPress={onDotPress} onTagPress={onTagPress} />
            </View>
            <Link url={content.url} active={linkActive}>
                <View style={styles.contentContainer}>
                    <View style={{ flexBasis: 0, flexGrow: 1 }}>
                        <View style={styles.titleContainer}>
                            <Text numberOfLines={8} style={styles.title}>{content.title}</Text>
                        </View>
                    </View>
                </View>
            </Link>
        </View>
    );
}

const styles = StyleSheet.create({
    topContainer: {
        padding: 10,
        paddingTop: 0,
    },
    bannerContainer: {
        flexDirection: 'row',
        justifyContent: 'flex-end',
    },
    senderContainer: {
        flexDirection: 'row',
        alignItems: 'center',
        flexGrow: 1,
    },
    senderName: {
        fontWeight: '400',
        fontSize: 16,
    },
    userImage: {
        height: 40,
        width: 40,
        margin: 5,
        marginLeft: 0,
        borderRadius: 20,
    },
    dotsContainer: {
        justifyContent: 'center',
        paddingLeft: 10,
        paddingBottom: 4,
    },
    contentContainer: {
        flexDirection: 'row',
        padding: 10,
        paddingTop: 0,
    },
    titleContainer: {
        marginBottom: 5,
    },
    title: {
        fontWeight: '600',
        fontSize: 17,
    },
    description: {
        fontSize: 12,
    },
    tagContainer: {
        paddingTop: 10,
        alignItems: 'flex-start'
    },
});
export default ContentPreviewNoImage