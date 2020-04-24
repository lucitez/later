
import React, { useState, useEffect } from 'react'
import { StyleSheet, View, ScrollView, TouchableOpacity, Text, Keyboard } from 'react-native'
import BottomSheet from './BottomSheet'
import BottomSheetContainer from './BottomSheetContainer';
import { Button, SearchBar } from '../common';
import { colors } from '../../assets/colors';
import Network from '../../util/Network';

function SavedContentBottomSheet(props) {
    const [tagValue, setTagValue] = useState('')

    return (
        <BottomSheet
            visible={props.active}
            onHide={() => props.onHide()}
            avoidKeyboard={true}
        >
            <BottomSheetContainer height='40%'>
                <View style={styles.addTagContainer}>
                    <SearchBar
                        placeholder='Add Tag...'
                        iconName='tag'
                        val
                        onChange={value => setTagValue(value)}
                        onCancel={() => props.onHide()}
                        autocomplete={true}
                        autocompleteFunc={filterTags}
                    />
                </View>
                <View style={styles.saveButtonContainer}>
                    <Button theme='primary' name='Save' size='medium' onPress={() => {
                        Keyboard.dismiss()
                        props.onHide()
                        props.onSave(tagValue)
                    }} />
                </View>
            </BottomSheetContainer>
        </BottomSheet>
    )
}

const filterTags = search => {
    params = { search }
    return Network.GET('/user-content/tags/filter', params)
}

const styles = StyleSheet.create({
    addTagContainer: {
        margin: 10,
        flexGrow: 1,
    },
    saveButtonContainer: {
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

export default SavedContentBottomSheet