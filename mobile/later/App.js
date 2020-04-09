import React from 'react';
import ContentScreen from './screens/Content';
import { registerRootComponent } from 'expo';
import { NavigationContainer } from '@react-navigation/native';
import { createBottomTabNavigator } from '@react-navigation/bottom-tabs';
import AntIcons from 'react-native-vector-icons/AntDesign';
import ArchiveScreen from './screens/Archive';
import Colors from './assets/colors';

const Tab = createBottomTabNavigator();

class App extends React.Component {
  render() {
    return (
      <NavigationContainer>
        <Tab.Navigator
          screenOptions={({ route }) => ({
            tabBarIcon: ({ focused, color, size }) => {
              let iconName;

              if (route.name === 'Home') {
                iconName = 'home'
              } else if (route.name === 'Archive') {
                iconName = 'lock';
              }

              return <AntIcons name={iconName} size={size} color={color} />;
            },
          })}
          tabBarOptions={{
            activeTintColor: Colors.primary,
            inactiveTintColor: 'gray',
          }}
        >
          <Tab.Screen name="Home" component={ContentScreen} />
          <Tab.Screen name='Archive' component={ArchiveScreen} />
        </Tab.Navigator>
      </NavigationContainer>
    )
  }
}

registerRootComponent(App)