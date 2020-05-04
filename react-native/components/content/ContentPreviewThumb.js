import React from 'react';
import { StyleSheet, Text, View, Image } from 'react-native';
import { Icon, Link, Tag } from '../common';
import { colors } from '../../assets/colors';
import ContentPreviewHeader from './ContentPreviewHeader'
import ContentPreviewFooter from './ContentPreviewFooter'

export default function ContentPreviewThumb({
    kind,
    onForwardPress,
    onDeletePress,
    onTagPress,
    onEditTagPress,
    onSavePress,
    content,
    linkActive,
    imageAR,
    includeFooter
}) {

    return (
        <View style={styles.container}>
            <ContentPreviewHeader content={content} onTagPress={onTagPress} />
            <View style={styles.contentContainer}>
                <View style={{ flexBasis: 0, flexGrow: 1 }}>
                    <View style={styles.titleContainer}>
                        <Text numberOfLines={8} style={styles.title}>{content.title}</Text>
                    </View>
                    <Text numberOfLines={2} style={styles.description}>{content.description}</Text>
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
            {includeFooter && <ContentPreviewFooter
                kind={kind}
                content={content}
                onForwardPress={onForwardPress}
                onDeletePress={onDeletePress}
                onSavePress={onSavePress}
                onEditTagPress={onEditTagPress}
            />}
        </View>
    );
}

const styles = StyleSheet.create({
    container: {
        backgroundColor: colors.white,
        paddingTop: 5,
        paddingBottom: 5,
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
        fontSize: 16,
    },
    description: {
        fontSize: 12,
    },
    imageContainer: {
        marginLeft: 10,
        height: 100,
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