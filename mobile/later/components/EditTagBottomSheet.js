import React, { useState, useEffect } from 'react'
import { StyleSheet, View, ScrollView, TouchableOpacity, Text, Keyboard } from 'react-native'
import BottomSheet from './BottomSheet';
import Button from './Button';
import SearchBar from './SearchBar';
import BottomSheetContainer from './BottomSheetContainer';
import { colors } from '../assets/colors';

function EditTagBottomSheet(props) {
    const [tagValue, setTagValue] = useState(props.content.tag)
    const [tags, setTags] = useState([])

    useEffect(() => {
        setTags(tagValue == '' ? [] : [tagValue])
    }, [tagValue])

    return (
        <BottomSheet
            visible={props.active}
            onHide={() => props.onHide()}
            avoidKeyboard={true}
        >
            <BottomSheetContainer height='40%'>
                <View style={styles.addTagContainer}>
                    <SearchBar
                        startingValue={props.content.tag}
                        iconName='tag'
                        onChange={value => setTagValue(value)}
                        onCancel={() => props.onHide()}
                    />
                </View>
                <ScrollView style={styles.tagsContainer}>
                    {tags.map((tag, index) => (
                        <TouchableOpacity key={index} style={styles.tagContainer}>
                            <Text style={styles.tag}>{tag}</Text>
                        </TouchableOpacity>
                    ))}
                </ScrollView>
                <View style={styles.archiveButtonContainer}>
                    <Button theme='primary' name='Update Tag' size='medium' onPress={() => {
                        Keyboard.dismiss()
                        props.onHide()
                        props.onUpdateTag(tagValue)
                    }} />
                </View>
            </BottomSheetContainer>
        </BottomSheet>
    )
}

const styles = StyleSheet.create({
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

export default EditTagBottomSheet