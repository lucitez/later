import React from 'react'
import { StyleSheet, View, Text, TouchableOpacity, Image } from 'react-native'
import { Tag, Icon } from '../common'
import { colors } from '../../assets/colors'

export default function ContentPreviewHeader({ content, onDotPress, onTagPress }) {
    return (
        <View style={styles.bannerContainer}>
            <View style={styles.senderContainer}>
                <Image style={styles.userImage} source={{ uri: 'https://www.washingtonpost.com/resizer/uwlkeOwC_3JqSUXeH8ZP81cHx3I=/arc-anglerfish-washpost-prod-washpost/public/HB4AT3D3IMI6TMPTWIZ74WAR54.jpg' }} />
                {content.sentByUsername && <Text style={styles.senderName}>@{content.sentByUsername}</Text>}
            </View>
            {
                content.tag && content.savedAt &&
                <View>
                    <TouchableOpacity style={styles.tagContainer} onPress={() => onTagPress(content.tag)}>
                        <Tag name={content.tag} />
                    </TouchableOpacity>
                </View>
            }
            <TouchableOpacity onPress={() => onDotPress(content)} style={styles.dotsContainer}>
                <Icon type='dots' size={20} color={colors.black} />
            </TouchableOpacity>
        </View>
    )
}

const styles = StyleSheet.create({
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
})