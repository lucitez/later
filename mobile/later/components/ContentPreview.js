import React from 'react';
import { StyleSheet, Text, View, Image, TouchableOpacity } from 'react-native';
import Icon from './Icon';
import Link from './Link';
import Tag from './Tag';
import { contentTypes, colors } from '../assets/colors';

function ContentPreview({ onDotPress, content, linkActive }) {
    return (
        <View style={styles.contentContainer}>
            <View style={styles.imageContainer}>
                <Image style={styles.thumb} source={content.imageUrl ? { uri: content.imageUrl } : {}} />
            </View>
            <View style={styles.detailsContainer}>
                <View style={styles.topDetailsContainer}>
                    <View style={styles.titleAndDescriptionContainer}>
                        <Link url={content.url} active={linkActive}>
                            <View>
                                <Text style={styles.title} numberOfLines={2}>{content.title}</Text>
                                <Text style={styles.description} numberOfLines={1}>{content.description}</Text>
                            </View>
                        </Link>
                    </View>
                    <TouchableOpacity onPress={() => onDotPress()}>
                        <Icon type='dots' size={20} color={colors.black} />
                    </TouchableOpacity>
                </View>
                <View style={styles.bottomDetailsContainer}>
                    <View style={styles.usernameContianer}>
                        {content.sentByUsername ? <Text>From @{content.sentByUsername}</Text> : null}
                    </View>
                    <View style={styles.contentTypeIconContainer}>
                        {content.contentType ?
                            <Icon type={content.contentType} size={25} color={contentTypes[content.contentType].color} /> :
                            null
                        }
                    </View>
                    <TouchableOpacity style={styles.tagContainer} onPress={() => onTagPress(content.tag)}>
                        {content.tag ? <Tag name={content.tag} /> : null}
                    </TouchableOpacity>
                </View>
            </View>
        </View>
    );
}

const renderTag = (archived, tag) => {
    if (!archived) return null
    if (!tag) return null
    return (
        <TouchableOpacity style={styles.tagContainer} onPress={() => onTagPress(tag)}>
            <Tag name={tag} />
        </TouchableOpacity>
    )
}

const styles = StyleSheet.create({
    contentContainer: {
        flexDirection: 'row',
        height: 120,
        width: '100%',
        marginTop: 5,
        marginBottom: 5,
    },
    imageContainer: {
        width: 120,
        justifyContent: 'center',
        alignItems: 'center',
        padding: 5,
    },
    detailsContainer: {
        flexDirection: 'column',
        flexGrow: 1,
        flexBasis: 0,
        padding: 5,
    },
    topDetailsContainer: {
        flexDirection: 'row',
        justifyContent: 'space-between',
    },
    titleAndDescriptionContainer: {
        flexGrow: 1,
        flexBasis: 0,
    },
    thumb: {
        height: '100%',
        width: '100%',
        borderRadius: 5,
    },
    title: {
        fontWeight: 'bold',
        fontSize: 16,
    },
    description: {
        fontSize: 12,
    },
    bottomDetailsContainer: {
        flexGrow: 1,
        flexDirection: 'row',
        alignItems: 'flex-end',
        justifyContent: 'space-between',
    },
    usernameContianer: {
        flex: 2,
        alignItems: 'flex-start'
    },
    tagContainer: {
        flex: 2,
        alignItems: 'flex-end'
    },
    contentTypeIconContainer: {
        flex: 1,
        alignItems: 'center',
        marginBottom: -3,
    },
});

export default ContentPreview