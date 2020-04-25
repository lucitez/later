import React, { useEffect, useState } from 'react';
import { StyleSheet, View, Alert, SafeAreaView } from 'react-native';
import { colors } from '../assets/colors';
import Network from '../util/Network';
import { Header, Icon, Button } from '../components/common';
import { ContentGroup2 } from '../components/content';
import { ButtonBottomSheet, EditTagBottomSheet } from '../components/modals';

function ContentScreen({ navigation }) {
    const [content, setContents] = useState([])
    const [offset, setOffset] = useState(0)
    const [loading, setLoading] = useState(true)
    const [contentBottomSheetVisible, setContentBottomSheetVisible] = useState(false)
    const [editTagBottomSheetVisible, setEditTagBottomSheetVisible] = useState(false)
    const [selectedContent, setSelectedContent] = useState(null)

    useEffect(() => {
        setLoading(true)
        getContent()
            .then(content => {
                setContents(content)
                setLoading(false)
            })
            .catch(error => console.error(error))
    }, [offset])

    const onSave = (tag) => {
        saveContent(selectedContent, tag)
            .then(setContents(content.filter(c => c.id != selectedContent.id)))
            .catch(error => Alert.alert(error))
    }

    return (
        <SafeAreaView style={styles.container}>
            <Header
                name="Later"
                rightIcon={<Icon
                    type='save'
                    size={25}
                    color={colors.white}
                    onPress={() => navigation.navigate('Saved')}
                />}
            />
            <View style={styles.contentContainer}>
                <ContentGroup2
                    type='home'
                    contents={content}
                    linkActive={!contentBottomSheetVisible && !editTagBottomSheetVisible}
                    onDotPress={content => {
                        setSelectedContent(content)
                        setContentBottomSheetVisible(true)
                    }}
                />
            </View>
            <ButtonBottomSheet isVisible={contentBottomSheetVisible} onHide={() => setContentBottomSheetVisible(false)}>
                <Button
                    theme='primary'
                    name='Forward'
                    size='medium'
                    onPress={() => {
                        setContentBottomSheetVisible(false)
                        navigation.navigate('Forward', { contentPreview: selectedContent, previousScreen: 'Home' })
                    }}
                />
                <Button
                    theme='primary'
                    name='Save'
                    size='medium'
                    onPress={() => {
                        setContentBottomSheetVisible(false)
                        setTimeout(() => { setEditTagBottomSheetVisible(true) }, 400)
                    }}
                />
                <Button
                    theme='light'
                    name='Cancel'
                    size='medium'
                    onPress={() => setContentBottomSheetVisible(false)}
                />
            </ButtonBottomSheet>
            <EditTagBottomSheet
                isVisible={editTagBottomSheetVisible}
                onSubmit={onSave}
                onHide={() => setEditTagBottomSheetVisible(false)}
            />
        </SafeAreaView>
    );
}

const getContent = () => {
    return Network.GET(`/user-content/filter`)
}

const saveContent = (content, tag) => {
    let params = {
        id: content.id,
        tag: tag
    }

    return Network.PUT(`/user-content/save`, params)
}

const styles = StyleSheet.create({
    container: {
        flex: 1,
        backgroundColor: colors.primary,
    },
    contentContainer: {
        backgroundColor: colors.white,
        flexGrow: 1,
    }
});

export default ContentScreen