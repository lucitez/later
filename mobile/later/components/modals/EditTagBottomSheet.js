
import React, { useState, useEffect } from 'react'
import { StyleSheet, View, Keyboard } from 'react-native'
import BottomSheet from './BottomSheet'
import { Button, SearchBar } from '../common';
import { colors } from '../../assets/colors';
import Network from '../../util/Network';

function EditTagBottomSheet({ value, isVisible, onSubmit, onHide }) {
    const [tagValue, setTagValue] = useState(value)
    const [visible, setVisible] = useState(isVisible)

    useEffect(() => { setVisible(isVisible) }, [isVisible])

    return (
        <BottomSheet
            visible={visible}
            onHide={() => onHide()}
            avoidKeyboard={true}
        >
            <View style={styles.contentContainer}>
                <View style={styles.addTagContainer}>
                    <SearchBar
                        placeholder='Add Tag...'
                        iconName='tag'
                        value={value}
                        onChange={value => setTagValue(value)}
                        onCancel={() => onHide()}
                        autocompleteFunc={filterTags}
                    />
                </View>
                <View style={styles.saveButtonContainer}>
                    <Button theme='primary' name='Submit' size='medium' onPress={() => {
                        Keyboard.dismiss()
                        onHide()
                        onSubmit(tagValue)
                    }} />
                </View>
            </View>
        </BottomSheet>
    )
}

const filterTags = search => {
    params = { search }
    return Network.GET('/user-content/tags/filter', params)
}

const styles = StyleSheet.create({
    contentContainer: {
        height: '40%',
        backgroundColor: colors.primary,
        paddingBottom: 10
    },
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

export default EditTagBottomSheet