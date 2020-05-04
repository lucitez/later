
import React, { useState, useEffect } from 'react'
import { StyleSheet, View, Keyboard, SafeAreaView } from 'react-native'
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
            <SafeAreaView style={styles.container}>
                <View style={styles.addTagContainer}>
                    <SearchBar
                        placeholder='Add or Edit Tag...'
                        iconName='tag'
                        value={value}
                        onChange={value => setTagValue(value)}
                        onCancel={() => onHide()}
                        autocompleteFunc={filterTags}
                    />
                </View>
                <View style={styles.saveButtonContainer}>
                    <Button theme='light' name='Submit' size='medium' onPress={() => {
                        Keyboard.dismiss()
                        onHide()
                        onSubmit(tagValue)
                    }} />
                </View>
            </SafeAreaView>
        </BottomSheet>
    )
}

const filterTags = search => {
    params = { search }
    return Network.GET('/user-content/tags/filter', params)
}

const styles = StyleSheet.create({
    container: {
        height: '50%',
        backgroundColor: colors.primary,
    },
    addTagContainer: {
        flexBasis: 0,
        flexGrow: 1,
        borderColor: 'white'
    },
    saveButtonContainer: {
        paddingLeft: 15,
        paddingRight: 15,
    },
});

export default EditTagBottomSheet