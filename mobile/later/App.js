import React, { useEffect, useState, useMemo } from 'react'
import { AsyncStorage } from 'react-native'
import { Provider } from 'react-redux'
import store from './store'
import { registerRootComponent } from 'expo'
import { NavigationContainer } from '@react-navigation/native'
import { createBottomTabNavigator } from '@react-navigation/bottom-tabs'
import { createStackNavigator } from '@react-navigation/stack'
import { Icon } from './components/common'
import jwtDecode from 'jwt-decode'
import {
  ByTagScreen,
  ContentScreen,
  SavedScreen,
  SharePreviewScreen,
  SendShareScreen,
  DiscoverScreen,
  ProfileScreen,
  EditProfileScreen,
  UserScreen,
  SplashScreen,
  ChatDisplayScreen
} from './screens/index'
import {
  SignupScreen,
  LoginScreen,
  SMSCodeScreen,
  CreatePasswordScreen,
} from './screens/onboarding/index'
import { colors } from './assets/colors'
import Network from './util/Network'
import * as actions from './actions'
import { AuthContext } from './context'
import { authHeader } from './util/headers'

const ApplicationStack = createStackNavigator()
const ApplicationTabs = createBottomTabNavigator()
const ContentStack = createStackNavigator()
const DiscoverStack = createStackNavigator()
const ShareStack = createStackNavigator()
const ProfileStack = createStackNavigator()

function CreateContentStack() {
  return (
    <ContentStack.Navigator initialRouteName='Home' headerMode='none'>
      <ContentStack.Screen name='Home' component={ContentScreen} />
      <ContentStack.Screen name='Saved' component={SavedScreen} />
      <ContentStack.Screen name='Tag' component={ByTagScreen} />
      <ContentStack.Screen name='Forward' component={SendShareScreen} />
    </ContentStack.Navigator>
  )
}

function CreateDiscoverStack() {
  return (
    <DiscoverStack.Navigator initialRouteName='Discover' headerMode='none'>
      <DiscoverStack.Screen name='Discover' component={DiscoverScreen} />
      <DiscoverStack.Screen name='User' component={UserScreen} />
    </DiscoverStack.Navigator>
  )
}

function CreateShareStack() {
  return (
    <ShareStack.Navigator initialRouteName='Share' headerMode='none'>
      <ShareStack.Screen name='Share' component={SharePreviewScreen} />
      <ShareStack.Screen name='Send Share' component={SendShareScreen} />
    </ShareStack.Navigator>
  )
}

function CreateProfileStack() {
  return (
    <ProfileStack.Navigator initialRouteName='Profile' headerMode='none'>
      <ProfileStack.Screen name='Profile' component={ProfileScreen} />
      <ProfileStack.Screen name='Edit' component={EditProfileScreen} />
      <ProfileStack.Screen name='Chat' component={ChatDisplayScreen} />
    </ProfileStack.Navigator>
  )
}

function CreateApplicationTabs() {
  return (
    <ApplicationTabs.Navigator
      screenOptions={({ route }) => ({
        tabBarIcon: ({ _, color, size }) => (
          <Icon type={route.name.toLowerCase()} size={size} color={color} />
        )
      })}
      tabBarOptions={{
        activeTintColor: colors.primary,
        inactiveTintColor: 'gray',
      }}
    >
      <ApplicationTabs.Screen name='Home' component={CreateContentStack} />
      <ApplicationTabs.Screen name='Search' component={CreateDiscoverStack} />
      <ApplicationTabs.Screen name='Share' component={CreateShareStack} />
      <ApplicationTabs.Screen name='Profile' component={CreateProfileStack} />
    </ApplicationTabs.Navigator>
  )
}

const updateTokens = async (tokens) => {
  try {
    let userId = jwtDecode(tokens.accessToken).sub
    store.dispatch(actions.setTokens(tokens))
    store.dispatch(actions.setUserId(userId))
    await AsyncStorage.setItem('refresh_token', tokens.refreshToken)
  } catch (e) {
    console.error(e)
  }
}

const autoLogin = async () => {
  let refreshToken = await AsyncStorage.getItem('refresh_token')

  if (refreshToken) {
    let refreshAuthHeader = authHeader(refreshToken)
    let newTokens = await Network.POST('/auth/refresh', {}, refreshAuthHeader)

    updateTokens(newTokens)
    return
  }

  throw new Error("No refresh token")
}

function App() {
  const [isLoading, setIsLoading] = useState(true)
  const [isSignedIn, setIsSignedIn] = useState(false)

  useEffect(() => {
    const init = async () => {
      await autoLogin()
        .then(setIsSignedIn(true))
        .catch(err => {
          setIsSignedIn(false)
        })
      setIsLoading(false)
    }

    init()
  }, [])

  const authContext = useMemo(
    () => ({
      signIn: () => { setIsSignedIn(true) },
      signOut: () => { setIsSignedIn(false) },
    }),
    []
  );

  return (
    <AuthContext.Provider value={authContext}>
      <Provider store={store}>
        <NavigationContainer>
          <ApplicationStack.Navigator headerMode='none'>
            {
              isLoading ? (
                <ApplicationStack.Screen name='Splash' component={SplashScreen} />
              ) : isSignedIn ? (
                <ApplicationStack.Screen name='Home' component={CreateApplicationTabs} />
              ) : (
                    <>
                      <ApplicationStack.Screen name='Sign Up' component={SignupScreen} />
                      <ApplicationStack.Screen name='Login' component={LoginScreen} />
                      <ApplicationStack.Screen name='SMS' component={SMSCodeScreen} />
                      <ApplicationStack.Screen name='Password' component={CreatePasswordScreen} />
                    </>
                  )
            }
          </ApplicationStack.Navigator>
        </NavigationContainer>
      </Provider>
    </AuthContext.Provider>
  )
}

registerRootComponent(App)