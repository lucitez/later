import React, { useState, useEffect } from 'react';
import { StyleSheet, Text, View, Image, TouchableOpacity } from 'react-native';
import { Icon, Link, Tag } from '../common';
import { contentTypes, colors } from '../../assets/colors';

function ContentPreviewFull({ onDotPress, onTagPress, content, linkActive, imageAR }) {

    return (
        <View>
            <View style={styles.topContainer}>
                <View style={styles.bannerContainer}>
                    <View style={styles.senderContainer}>
                        <Image style={styles.userImage} source={{ uri: 'https://www.washingtonpost.com/resizer/uwlkeOwC_3JqSUXeH8ZP81cHx3I=/arc-anglerfish-washpost-prod-washpost/public/HB4AT3D3IMI6TMPTWIZ74WAR54.jpg' }} />
                        <Text style={styles.senderName}>@{content.sentByUsername}</Text>
                    </View>
                    <TouchableOpacity onPress={() => onDotPress(content)} style={styles.dotsContainer}>
                        <Icon type='dots' size={20} color={colors.black} />
                    </TouchableOpacity>
                </View>
                <View>
                    <View style={styles.titleContainer}>
                        <Text numberOfLines={2} style={styles.title}>{content.title}</Text>
                    </View>
                    <View style={styles.descriptionContainer}>
                        <Text numberOfLines={1} style={styles.description}>{content.description}</Text>
                    </View>
                </View>
                {
                    content.tag && content.savedAt &&
                    <View>
                        <TouchableOpacity style={styles.tagContainer} onPress={() => onTagPress(content.tag)}>
                            <Tag name={content.tag} />
                        </TouchableOpacity>
                    </View>
                }
            </View>
            <View style={styles.bottomContainer}>
                <Link url={content.url} active={linkActive}>
                    <>
                        <View style={styles.imageContainer}>
                            <Image style={[styles.image, { aspectRatio: imageAR }]} source={{ uri: content.imageUrl }} />
                        </View>
                        <View style={styles.hostnameContainer}>
                            <Text style={{ color: colors.white, opacity: 0 }}>{content.hostname}</Text>
                        </View>
                        <View style={[styles.hostnameContainer, { backgroundColor: null, opacity: 1 }]}>
                            <Text style={{ color: colors.white, opacity: 1 }}>{content.hostname}</Text>
                        </View>
                    </>
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
    bannerContainer: {
        flexDirection: 'row',
        justifyContent: 'space-between',
    },
    senderContainer: {
        flexDirection: 'row',
        alignItems: 'center',
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
    },
    titleContainer: {
        marginBottom: 5,
    },
    title: {
        fontWeight: 'bold',
        fontSize: 18,
    },
    description: {
        fontSize: 12,
    },
    tagContainer: {
        paddingTop: 10,
        alignItems: 'flex-start'
    },
    hostnameContainer: {
        position: 'absolute',
        bottom: 0,
        opacity: 0.75,
        width: '100%',
        backgroundColor: colors.black,
        padding: 10,
    }
});

export default ContentPreviewFull