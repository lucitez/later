import React from 'react';
import { StyleSheet, Text, View } from 'react-native'
import { Link } from '../common'
import { colors } from '../../assets/colors'
import ContentPreviewHeader from './ContentPreviewHeader'
import ContentPreviewFooter from './ContentPreviewFooter'

function ContentPreviewNoImage({
    kind,
    onForwardPress,
    onDeletePress,
    onTagPress,
    onEditTagPress,
    onSavePress,
    content,
    linkActive,
    includeFooter
}) {

    return (
        <View style={styles.container}>
            <ContentPreviewHeader content={content} onTagPress={onTagPress} />
            <Link url={content.url} active={linkActive}>
                <View style={styles.contentContainer}>
                    <View style={{ flexBasis: 0, flexGrow: 1 }}>
                        <View style={styles.titleContainer}>
                            <Text numberOfLines={6} style={styles.title}>{content.title}</Text>
                        </View>
                    </View>
                </View>
            </Link>
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
        fontSize: 17,
    },
});
export default ContentPreviewNoImage