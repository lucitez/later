import React from 'react';
import { Dimensions } from 'react-native';
import Modal from 'react-native-modal';

function BottomSheet(props) {
    const deviceWidth = Dimensions.get("window").width
    const deviceHeight = Dimensions.get("window").height

    return (
        <Modal
            isVisible={props.visible}
            backdropOpacity={props.backdropOpacity ? props.backdropOpacity : 0}
            onBackdropPress={() => props.onHide()}
            animationIn='slideInUp'
            animationOut='slideOutDown'
            animationOutTiming={200}
            deviceHeight={deviceHeight}
            deviceWidth={deviceWidth}
            avoidKeyboard={props.avoidKeyboard}
            style={{ justifyContent: 'flex-end', margin: 0 }}
        >
            {props.children}
        </Modal>
    )
}

export default BottomSheet