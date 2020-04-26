import React from 'react'
import { StyleSheet, View, Text, Image } from 'react-native'
import { colors } from '../../assets/colors'

function UserDetails({ user }) {
    return (
        <View style={styles.container}>
            <View style={styles.leftContainer}>
                <View style={styles.imageContainer} >
                    <Image style={styles.thumb} source={{ uri: 'https://www.washingtonpost.com/resizer/uwlkeOwC_3JqSUXeH8ZP81cHx3I=/arc-anglerfish-washpost-prod-washpost/public/HB4AT3D3IMI6TMPTWIZ74WAR54.jpg' }} />
                </View>
                <View style={styles.tasteContainer}>
                    <Text style={styles.taste}>{user.taste} taste</Text>
                </View>
            </View>
            <View style={styles.infoContainer}>
                <Text style={styles.name}>{user.name}</Text>
                <Text style={styles.username}>@{user.username}</Text>
            </View>
        </View>
    )
}

const styles = StyleSheet.create({
    container: {
        flexDirection: 'row',
        padding: 10,
    },
    leftContainer: {
        width: 90,
        alignItems: 'center',
    },
    imageContainer: {
        height: 75,
        width: 75,
        borderRadius: 50,
        alignItems: 'center',
        justifyContent: 'center',
        marginBottom: 10,
        borderWidth: 2,
        borderColor: colors.white,
        padding: 2,
    },
    thumb: {
        aspectRatio: 1,
        height: '100%',
        width: '100%',
        borderRadius: 100,
        backgroundColor: 'coral',
    },
    infoContainer: {
        paddingLeft: 10,
        justifyContent: 'center',
    },
    taste: {
        color: colors.white,
    },
    name: {
        fontSize: 18,
        fontWeight: 'bold',
        color: colors.white
    },
    username: {
        color: colors.white
    },
})

export default UserDetails