import React, { useState } from 'react';
import { StyleSheet, View, Text, TouchableWithoutFeedback } from 'react-native';
// import { colors } from '../assets/colors';
import Icon from '../components/Icon';
import { colors, contentTypes } from '../assets/colors';

function FilterOption(props) {

    return (
        <View style={styles.filterValueContainer}>
            <TouchableWithoutFeedback onPress={props.onPress}>
                <View style={styles.filterIconContainer}>
                    <Icon type={props.name} size={30} color={props.active ? contentTypes[props.name].color : colors.gray} />
                </View>
            </TouchableWithoutFeedback >
        </View>

    )
}

function ArchiveFilter(props) {

    let contentTypes = ['watch', 'read', 'listen']

    const [filter, setFilter] = useState({
        'content_type': null
    })

    return (
        <View style={styles.filterContainer}>
            {
                contentTypes.map((contentType, index) => (
                    <FilterOption
                        key={index}
                        name={contentType}
                        active={filter['content_type'] == contentType}
                        onPress={() => {
                            let updatedFilter = {
                                ...filter,
                                ['content_type']: filter['content_type'] == contentType ? null : contentType
                            }
                            setFilter(updatedFilter)
                            props.onChange(updatedFilter)
                        }}
                    />
                ))
            }
        </View>
    )
}

const styles = StyleSheet.create({
    filterContainer: {
        width: '100%',
        height: 50,
        flexDirection: 'row',
        justifyContent: 'space-between',
        backgroundColor: colors.lightGray,
    },
    filterValueContainer: {
        flexGrow: 1,
    },
    filterIconContainer: {
        height: '100%',
        justifyContent: 'center',
        alignItems: 'center',
    },

});

export default ArchiveFilter