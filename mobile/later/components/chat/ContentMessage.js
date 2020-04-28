import React, { useState, useEffect } from 'react'
import { View, Text, StyleSheet, Image } from 'react-native'
import { colors } from '../../assets/colors'

export default ContentMessage = ({ title, imageUrl }) => {

    const [imageAR, setImageAR] = useState(1)

    useEffect(() => {
        Image.getSize(imageUrl, (width, height) => setImageAR(width / height))
    })

    return (
        <View style={styles.container}>
            <View style={styles.detailsContainer}>
                {imageUrl &&
                    <Image
                        style={[styles.image, { aspectRatio: imageAR }]}
                        source={imageUrl ? { uri: imageUrl } : {}}
                    />
                }
                <View style={styles.titleContainer}>
                    <Text numberOfLines={2} style={styles.title}>{title}</Text>
                </View>
            </View>
        </View>

    )
}

const styles = StyleSheet.create({
    container: {
        width: '60%',
        marginTop: 10,
        marginLeft: 5,
        marginRight: 5,
    },
    detailsContainer: {
        backgroundColor: colors.black,
        borderRadius: 10,
    },
    image: {
        width: '100%',
        borderTopLeftRadius: 10,
        borderTopRightRadius: 10,
    },
    titleContainer: {
        padding: 10,
    },
    title: {
        color: colors.white,
        fontWeight: 'bold',
    }
})