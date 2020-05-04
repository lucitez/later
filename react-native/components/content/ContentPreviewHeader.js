import React from 'react'
import { StyleSheet, View, Text, TouchableOpacity, Image } from 'react-native'
import { Tag, Icon } from '../common'
import { colors, contentTypes } from '../../assets/colors'
import { timeSince } from '../../util/time';

export default function ContentPreviewHeader({ content, onDotPress, onTagPress }) {
    return (
        <View style={styles.bannerContainer}>
            <View style={styles.senderContainer}>
                <Image style={styles.userImage} source={{ uri: 'https://www.washingtonpost.com/resizer/uwlkeOwC_3JqSUXeH8ZP81cHx3I=/arc-anglerfish-washpost-prod-washpost/public/HB4AT3D3IMI6TMPTWIZ74WAR54.jpg' }} />
                <View>
                    <View style={{ flexDirection: 'row' }}>
                        <Text style={styles.senderName}>{content.sentByName}</Text>
                        <Text style={styles.senderUsername}>- @{content.sentByUsername}</Text>
                    </View>
                    {content.createdAt && <Text style={{ fontSize: 13, fontWeight: '300' }}>{timeSince(Date.parse(content.createdAt))}</Text>}
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
            {content.contentType && <Icon style={{ marginLeft: 10 }} type={content.contentType} color={contentTypes[content.contentType].color} size={25} />}
            {/* <TouchableOpacity onPress={() => onDotPress(content)} style={styles.dotsContainer}>
                <Icon type='dots' size={20} color={colors.black} />
            </TouchableOpacity> */}
        </View>
    )
}

const styles = StyleSheet.create({
    bannerContainer: {
        padding: 10,
        flexDirection: 'row',
        justifyContent: 'flex-end',
    },
    senderContainer: {
        flexDirection: 'row',
        alignItems: 'center',
        flexGrow: 1,
    },
    senderName: {
        fontWeight: '600',
        fontSize: 14,
    },
    senderUsername: {
        fontWeight: '300',
        fontSize: 14,
        marginLeft: 5,
    },
    userImage: {
        height: 35,
        width: 35,
        marginRight: 5,
        borderRadius: 20,
    },
    dotsContainer: {
        justifyContent: 'center',
        paddingLeft: 10,
        paddingBottom: 4,
    },
})