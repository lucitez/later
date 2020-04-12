import React from 'react';
import SimpleLineIcons from 'react-native-vector-icons/SimpleLineIcons';
import AntIcons from 'react-native-vector-icons/AntDesign';
import FeatherIcons from 'react-native-vector-icons/Feather';
import FontAwesomeIcons from 'react-native-vector-icons/FontAwesome';
import MaterialIcons from 'react-native-vector-icons/MaterialIcons';

function Icon(props) {
    switch (props.type) {

        case 'watch':
            return <SimpleLineIcons name='screen-desktop' {...props} />
        case 'read':
            return <SimpleLineIcons name='eyeglass' {...props} />
        case 'listen':
            return <SimpleLineIcons name='earphones' {...props} />
        case 'add_friend':
            return <AntIcons name='adduser' {...props} />
        case 'search':
            return <AntIcons name='search1' {...props} />
        case 'plus':
            return <AntIcons name='pluscircleo' {...props} />
        case 'check_filled':
            return <AntIcons name='checkcircle' {...props} />
        case 'back':
            return <AntIcons name='left' {...props} />
        case 'chat':
            return <FeatherIcons name='message-circle' {...props} />
        case 'home':
            return <AntIcons name='home' {...props} />
        case 'archive':
            return <AntIcons name='inbox' {...props} />
        case 'friends':
            return <AntIcons name='user' {...props} />
        case 'share':
            return <FontAwesomeIcons name='send' {...props} />
        case 'paste':
            return <MaterialIcons name='content-paste' {...props} />
        case 'close':
            return <AntIcons name='closecircle' {...props} />
        default:
            return null
    }
}

export default Icon