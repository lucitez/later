import React from 'react';
import { StyleSheet, Text, View, Image, TouchableOpacity } from 'react-native';
import { Icon, Link, Tag } from '../common';
import { colors } from '../../assets/colors';
import ContentPreviewHeader from './ContentPreviewHeader'

function ContentPreviewThumb({ onDotPress, onTagPress, content, linkActive, imageAR }) {

    return (
        <View style={{ backgroundColor: colors.white, paddingTop: 5, paddingBottom: 5 }}>
            <View style={styles.topContainer}>
                <ContentPreviewHeader content={content} onDotPress={onDotPress} onTagPress={onTagPress} />
            </View>
            <View style={styles.contentContainer}>
                <View style={{ flexBasis: 0, flexGrow: 1 }}>
                    <View style={styles.titleContainer}>
                        <Text numberOfLines={8} style={styles.title}>{content.title}</Text>
                    </View>
                </View>
                <Link url={content.url} active={linkActive}>
                    <View style={[styles.imageContainer, { aspectRatio: imageAR }]}>
                        <Image style={styles.image} source={content.imageUrl ? { uri: content.imageUrl } : {}} />
                        <View style={styles.hostnameContainer}>
                            <Text style={{ color: colors.white, opacity: 0 }} numberOfLines={1}>{content.hostname}</Text>
                        </View>
                        <View style={[styles.hostnameContainer, { backgroundColor: null, opacity: 1 }]}>
                            <Text style={{ color: colors.white, opacity: 1 }} numberOfLines={1}>{content.hostname.replace('www.', '')}</Text>
                        </View>
                    </View>
                </Link>
            </View>
        </View>
    );
}

const styles = StyleSheet.create({
    topContainer: {
        padding: 10,
        paddingTop: 0,
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
    imageContainer: {
        marginLeft: 10,
        height: 115,
    },
    image: {
        borderRadius: 10,
        height: '100%',
        width: '100%',
    },
    hostnameContainer: {
        position: 'absolute',
        bottom: 0,
        opacity: 0.75,
        width: '100%',
        backgroundColor: colors.black,
        padding: 5,
        borderBottomLeftRadius: 10,
        borderBottomRightRadius: 10,
    },
    tagContainer: {
        paddingTop: 10,
        alignItems: 'flex-start'
    },
});

export default ContentPreviewThumb