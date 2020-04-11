import React from 'react';
import { registerRootComponent } from 'expo';
import { NavigationContainer } from '@react-navigation/native';
import { createBottomTabNavigator } from '@react-navigation/bottom-tabs';
import { createStackNavigator } from '@react-navigation/stack';
import AntIcons from 'react-native-vector-icons/AntDesign';
import ContentScreen from './screens/ContentScreen';
import FriendScreen from './screens/FriendScreen';
import AddFriendScreen from './screens/AddFriendScreen';
import Colors from './assets/colors';

const Tab = createBottomTabNavigator();
const FriendStack = createStackNavigator();

function CreateFriendStack() {
  return (
    <FriendStack.Navigator initialRouteName="Friends" headerMode="none">
      <FriendStack.Screen name="Friends" component={FriendScreen} />
      <FriendStack.Screen name="Test" component={AddFriendScreen} />
    </FriendStack.Navigator>
  )
}

function CreateContentScreen() {
  return <ContentScreen archived={false} />
}

function CreateArchiveScreen() {
  return <ContentScreen archived={true} />
}

class App extends React.Component {
  render() {
    return (
      <NavigationContainer>
        <Tab.Navigator
          screenOptions={({ route }) => ({
            tabBarIcon: ({ _, color, size }) => {
              let iconName;

              if (route.name === 'Home') {
                iconName = 'home'
              } else if (route.name === 'Archive') {
                iconName = 'lock';
              } else if (route.name === 'Friends') {
                iconName = 'user';
              }

              return <AntIcons name={iconName} size={size} color={color} />;
            },
          })}
          tabBarOptions={{
            activeTintColor: Colors.primary,
            inactiveTintColor: 'gray',
          }}
        >
          <Tab.Screen name='Home' component={CreateContentScreen} />
          <Tab.Screen name='Archive' component={CreateArchiveScreen} />
          <Tab.Screen name='Friends' component={CreateFriendStack} />
        </Tab.Navigator>
      </NavigationContainer>
    )
  }
}

registerRootComponent(App)