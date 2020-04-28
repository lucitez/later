import React, { useState, useEffect } from 'react';
import { StyleSheet, Text, View, Image, TouchableOpacity } from 'react-native';
import { Icon, Link, Tag } from '../common';
import { contentTypes, colors } from '../../assets/colors';

function ContentPreview({ onDotPress, onTagPress, content, linkActive }) {

    const [imageAR, setImageAR] = useState(1)

    useEffect(() => {
        content.imageUrl && Image.getSize(content.imageUrl, (width, height) => setImageAR(width / height))
    })

    return (
        <View style={styles.container}>
            <View style={styles.topContainer}>
                <View style={styles.imageContainer}>
                    <Image style={[styles.image, { aspectRatio: imageAR }]} source={content.imageUrl ? { uri: content.imageUrl } : {}} />
                </View>
                <View style={styles.rightContainer}>
                    <View style={styles.actionContainer}>
                        <View style={styles.contentTypeIconContainer}>
                            {content.contentType ?
                                <Icon type={content.contentType} size={25} color={contentTypes[content.contentType].color} /> :
                                null
                            }
                        </View>
                        <TouchableOpacity style={styles.tagContainer} onPress={() => onTagPress(content.tag)}>
                            {content.tag && content.savedAt ? <Tag name={content.tag} /> : null}
                        </TouchableOpacity>
                        <TouchableOpacity onPress={() => onDotPress(content)} style={styles.dotsContainer}>
                            <Icon type='dots' size={20} color={colors.black} />
                        </TouchableOpacity>
                    </View>
                    <View style={styles.sentByContainer}>
                        <Image style={styles.sentByImage} source={{ uri: 'https://www.washingtonpost.com/resizer/uwlkeOwC_3JqSUXeH8ZP81cHx3I=/arc-anglerfish-washpost-prod-washpost/public/HB4AT3D3IMI6TMPTWIZ74WAR54.jpg' }} />
                        <View style={styles.sentByUsernameContainer} >
                            <Text style={styles.sentByUsername}>@{content.sentByUsername}</Text>
                        </View>
                    </View>
                </View>
            </View>
            <View style={styles.bottomContainer}>
                <Link url={content.url} active={linkActive}>
                    <View>
                        <View style={styles.titleContainer}>
                            <Text numberOfLines={1} style={styles.title}>{content.title}</Text>
                        </View>
                        <View style={styles.descriptionContainer}>
                            <Text numberOfLines={2} style={styles.description}>{content.description}</Text>
                        </View>
                    </View>
                </Link>
            </View>
        </View>
    );
}

const styles = StyleSheet.create({
    container: {
        height: 165,
        width: '100%',
        paddingTop: 5,
        paddingBottom: 5,
        backgroundColor: colors.white,
    },
    topContainer: {
        height: '69%',
        flexDirection: 'row',
        justifyContent: 'space-between',
    },
    imageContainer: {
        padding: 5,
        maxWidth: '50%',
    },
    rightContainer: {
        width: '50%',
    },
    actionContainer: {
        flex: 1,
        flexDirection: 'row',
        justifyContent: 'space-between',
        alignItems: 'flex-start',
        paddingTop: 10,
    },
    sentByContainer: {
        flex: 1,
        justifyContent: 'flex-start',
        alignItems: 'flex-end',
        flexDirection: 'row',
        paddingBottom: 10,
    },
    sentByImage: {
        borderRadius: 15,
        height: 30,
        width: 30,
    },
    sentByUsernameContainer: {
        marginLeft: 5,
        marginBottom: -2,
    },
    sentByUsername: {
        fontWeight: '600'
    },
    image: {
        height: '100%',
        borderRadius: 10,
    },
    bottomContainer: {
        height: '31%',
        paddingLeft: 10,
    },
    descriptionContainer: {
        overflow: 'hidden'
    },
    title: {
        fontWeight: 'bold',
        fontSize: 18,
    },
    description: {
        fontSize: 12,
    },
    contentTypeIconContainer: {
        flex: 1,
        marginTop: -2,
    },
    tagContainer: {
        flex: 1,
        alignItems: 'flex-start'
    },
    dotsContainer: {
        flex: 1,
        marginTop: 2,
        paddingLeft: 10,
        paddingRight: 10,
        alignItems: 'flex-end'
    },
});

export default ContentPreview