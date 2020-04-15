
import React, { useState, useEffect } from 'react'
import { StyleSheet, View, Keyboard, ScrollView, TouchableOpacity, Text } from 'react-native'
import BottomSheet from './BottomSheet';
import Button from './Button';
import SearchBar from './SearchBar';
import { colors } from '../assets/colors';

function ArchiveContentBottomSheet(props) {
    const [tagValue, setTagValue] = useState('')
    const [tags, setTags] = useState([])

    const [isKeyboardShowing, setKeyboardShowing] = useState(false)

    Keyboard.addListener('keyboardWillShow', function () {
        setKeyboardShowing(true)
    })
    Keyboard.addListener('keyboardWillHide', function () {
        setKeyboardShowing(false)
    })

    useEffect(() => {
        setTags(tagValue == '' ? [] : [tagValue])
    }, [tagValue])

    return (
        <BottomSheet
            visible={props.active}
            onHide={() => props.onHide()}
            avoidKeyboard={true}
        >
            <View style={[styles.archiveBottomSheet, isKeyboardShowing ? { paddingBottom: 5 } : { paddingBottom: 30 }]}>
                <View style={styles.addTagContainer}>
                    <SearchBar iconName='tag' onChange={value => setTagValue(value)} onCancel={() => props.onHide()} />
                </View>
                <ScrollView style={styles.tagsContainer}>
                    {tags.map((tag, index) => (
                        <TouchableOpacity key={index} style={styles.tagContainer}>
                            <Text style={styles.tag}>{tag}</Text>
                        </TouchableOpacity>
                    ))}
                </ScrollView>
                <View style={styles.archiveButtonContainer}>
                    <Button theme='primary' name='Archive' size='medium' onPress={() => {
                        props.onHide(false)
                        props.onArchive(tagValue)
                    }} />
                </View>

            </View>
        </BottomSheet>
    )
}

const styles = StyleSheet.create({
    archiveBottomSheet: {
        height: '40%',
        backgroundColor: colors.primary,
        justifyContent: 'flex-start'
    },
    addTagContainer: {
        margin: 10,
    },
    archiveButtonContainer: {
        paddingLeft: 15,
        paddingRight: 15,
    },
    tagsContainer: {
        flexGrow: 1,
        marginLeft: 15,
        marginRight: 15,
    },
    tagContainer: {
        padding: 5,
        borderTopWidth: 0.5,
        borderBottomWidth: 0.5,
        borderColor: colors.white,
    },
    tag: {
        color: colors.white,
        fontSize: 16,
    }
});

export default ArchiveContentBottomSheet