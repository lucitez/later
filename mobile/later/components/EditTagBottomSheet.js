
import React from 'react'
import { Text } from 'react-native'
import BottomSheet from './BottomSheet';

function EditTagBottomSheet(props) {

    return (
        <BottomSheet
            visible={props.active}
            onHide={() => props.onHide()}
            avoidKeyboard={true}
        >
            <Text>TAGGG</Text>
        </BottomSheet>
    )
}


export default EditTagBottomSheet