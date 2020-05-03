import React from 'react';
import SimpleLineIcons from 'react-native-vector-icons/SimpleLineIcons';
import AntIcons from 'react-native-vector-icons/AntDesign';
import FeatherIcons from 'react-native-vector-icons/Feather';
import FontAwesomeIcons from 'react-native-vector-icons/FontAwesome';
import MaterialIcons from 'react-native-vector-icons/MaterialIcons';
import EntypoIcons from 'react-native-vector-icons/Entypo';
import Ionicons from 'react-native-vector-icons/Ionicons';
import { TouchableOpacity } from 'react-native-gesture-handler';

function RenderIcon(props) {
    switch (props.type) {

        /**
        Content Type
        */

        case 'watch':
            return <SimpleLineIcons name='screen-desktop' {...props} />
        case 'read':
            return <SimpleLineIcons name='eyeglass' {...props} />
        case 'listen':
            return <SimpleLineIcons name='earphones' {...props} />
        case 'login':
            return <SimpleLineIcons name='earphones' {...props} />

        /**
        Tab Nav
        */

        case 'home':
            return <AntIcons name='home' {...props} />
        case 'discover':
            return <Ionicons name='md-globe' {...props} />
        case 'share':
            return <FontAwesomeIcons name='send' {...props} />
        case 'profile':
            return <AntIcons name='user' {...props} />
        case 'chat':
            return <SimpleLineIcons name='bubble' {...props} />

        /**
        Utility
        */

        case 'back':
            return <AntIcons name='left' {...props} />
        case 'search':
            return <AntIcons name='search1' {...props} />
        case 'circle':
            return <FeatherIcons name='circle' {...props} />
        case 'check_filled':
            return <FeatherIcons name='check-circle' {...props} />
        case 'close':
            return <AntIcons name='closecircle' {...props} />
        case 'dots':
            return <EntypoIcons name='dots-three-vertical' {...props} />
        case 'gear':
            return <SimpleLineIcons name='settings' {...props} />

        /**
        App
        */

        case 'tag':
            return <AntIcons name='tag' {...props} />
        case 'save':
            return <AntIcons name='inbox' {...props} />

        /**
        Misc
        */

        case 'paste':
            return <MaterialIcons name='content-paste' {...props} />
        case 'next':
            return <AntIcons name='rightcircle' {...props} />
        case 'new-message':
            return <AntIcons name='form' {...props} />


        default:
            return null
    }
}

function Icon(props) {
    return (
        <TouchableOpacity onPress={props.onPress}>
            {RenderIcon(props)}
        </TouchableOpacity>
    )
}

export default Icon