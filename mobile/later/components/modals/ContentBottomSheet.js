import React, { useState, useEffect } from 'react'
import BottomSheet from './BottomSheet'
import SaveContentBottomSheet from './SaveContentBottomSheet'
import EditTagBottomSheet from './EditTagBottomSheet'
import BottomSheetContainer from './BottomSheetContainer'
import { Button, ButtonGroup } from '../common';

function ContentBottomSheet(props) {
    const [optionsActive, setOptionsActive] = useState(false)
    const [saveActive, setSaveActive] = useState(false)
    const [tagActive, setTagActive] = useState(false)

    useEffect(() => {
        props.setOptionsActive(optionsActive)
    }, [optionsActive])

    useEffect(() => {
        setOptionsActive(props.optionsActive)
    }, [props.optionsActive])

    const Options = () => {
        switch (props.type) {
            case 'home':
                return (
                    <>
                        <Button theme='primary' name='Forward' size='medium' onPress={() => {
                            props.onForward()
                            setOptionsActive(false)
                        }} />
                        <Button theme='primary' name='Save' size='medium' onPress={() => {
                            setOptionsActive(false)
                            setTimeout(() => { setSaveActive(true) }, 400)
                        }} />
                        <Button theme='light' name='Cancel' size='medium' onPress={() => setOptionsActive(false)} />
                    </>
                )
            case 'save':
                return (
                    <>
                        <Button theme='primary' name='Forward' size='medium' onPress={() => {
                            props.onForward()
                            setOptionsActive(false)
                        }} />
                        <Button theme='primary' name='Edit Tag' size='medium' onPress={() => {
                            setOptionsActive(false)
                            setTimeout(() => { setTagActive(true) }, 400)
                        }} />
                        <Button theme='light' name='Cancel' size='medium' onPress={() => setOptionsActive(false)} />
                    </>
                )
        }
    }

    return (
        <>
            <BottomSheet
                visible={optionsActive}
                onHide={() => setOptionsActive(false)}
            >
                <BottomSheetContainer>
                    <ButtonGroup theme='primary'>
                        <Options />
                    </ButtonGroup>
                </BottomSheetContainer>

            </BottomSheet>
            <SaveContentBottomSheet
                active={saveActive}
                onHide={() => setSaveActive(false)}
                {...props}
            />
            <EditTagBottomSheet
                active={tagActive}
                onHide={() => setTagActive(false)}
                {...props}
            />
        </>
    )
}

export default ContentBottomSheet