import React, { ChangeEvent } from 'react'
import {
  PageWrapper,
  Header,
  SidebarWrapper,
  ContentWrapper,
  ListStyled,
  SearchWrapper,
  SearchIconStyled,
  ContainerTitle,
  EmptyFriendsText,
  UserNoteUnderlined,
  UserName,
} from './styles'
import { Tabs } from './Tabs'
import TopBar from '@view/shared/TopBar'
import ListItem from '@material-ui/core/ListItem'
import InputAdornment from '@material-ui/core/InputAdornment'
import FormControl from '@material-ui/core/FormControl'
import Input from '@material-ui/core/Input'
import Paper from '@material-ui/core/Paper'
import List from '@material-ui/core/List'
import Divider from '@material-ui/core/Divider'
import ListItemText from '@material-ui/core/ListItemText'
import ListItemAvatar from '@material-ui/core/ListItemAvatar'
import Avatar from '@material-ui/core/Avatar'
import { ApiTypes, FriendsTypes } from './../../../types/index'

export interface Props {
  friends: ApiTypes.User[]
  friendsOfFriends: ApiTypes.FriendsOfFriend[]
  onGetFriends: () => void
  onGetFriendsOfFriends: () => void
}

interface State {
  filteredFriends: ApiTypes.User[]
  filteredFriendsOfFriends: ApiTypes.FriendsOfFriend[]
  filterValue: string
  currentTab: FriendsTypes.CurrentTab
}

export class FriendsList extends React.Component<Props, State> {

  state = {
    filteredFriends: [],
    filteredFriendsOfFriends: [],
    filterValue: '',
    currentTab: 'friends' as FriendsTypes.CurrentTab
  }

  mainContent = () => {
    return (
      <>
        <ContainerTitle>Content</ContainerTitle>
        <Divider />
        <List>
          {/* <ListItem>
            <ListItemAvatar>
              <Avatar alt="User Name" />
            </ListItemAvatar>
            <ListItemText
              primary="User Name"
              secondary={null}
            />
          </ListItem> */}
        </List>
      </>
    )
  }

  renderEmptyListMessage = () => {
    const { filterValue, currentTab } = this.state

    if (filterValue) {
      return <EmptyFriendsText>No one's been found.</EmptyFriendsText>
    }

    switch (currentTab) {
      case 'friends': return <EmptyFriendsText>You don't have any friends yet.</EmptyFriendsText>
      case 'friends-of-friends': return <EmptyFriendsText>You don't have any friends of friends yet.</EmptyFriendsText>
      default: return null
    }

  }

  mapFriends = (friends: ApiTypes.User[]) => {

    if (!friends.length) {
      return this.renderEmptyListMessage()
    }

    return friends.map(item => (
      <div key={item.id}>
        <ListItem>
          <ListItemAvatar>
            <Avatar alt={item.name} />
          </ListItemAvatar>
          <ListItemText primary={<UserName>{item.name}</UserName>} />
        </ListItem>
        <Divider variant="inset" component="li" />
      </div>
    ))
  }

  mapFriendsOfFriends = (friendsOfFriends: ApiTypes.FriendsOfFriend[]) => {

    if (!friendsOfFriends.length) {
      return this.renderEmptyListMessage()
    }

    return friendsOfFriends.map(item => {
      const { user, friends } = item
      return (
        <div key={user.id}>
          <ListItem alignItems={friends.length ? 'flex-start' : 'center'}>
            <ListItemAvatar>
              <Avatar alt={user.name} />
            </ListItemAvatar>
            <ListItemText
              primary={<UserName>{user.name}</UserName>}
              secondary={(friends.length) ? <UserNoteUnderlined>You have {friends.length} in common</UserNoteUnderlined> : null}
            />
          </ListItem>
          <Divider variant="inset" component="li" />
        </div>
      )
    })
  }

  onTabSwitch = (value: FriendsTypes.CurrentTab) => {
    this.setState({
      currentTab: value,
      filterValue: '',
      filteredFriends: [],
    })
  }

  onFilterValueChange = (event: ChangeEvent<HTMLInputElement>) => {
    const { value } = event.currentTarget
    const { currentTab } = this.state

    this.setState({
      filterValue: value,
    })

    switch (currentTab) {
      case 'friends': this.setFilteredFriends(value); break
      case 'friends-of-friends': this.setFilteredFriendsOfFriends(value); break
      default: return null
    }
  }

  setFilteredFriends = (value: string) => {
    const { friends } = this.props

    this.setState({
      filteredFriends: friends.filter(item => item.name.includes(value))
    })
  }

  setFilteredFriendsOfFriends = (value: string) => {
    const { friendsOfFriends } = this.props

    this.setState({
      filteredFriendsOfFriends: friendsOfFriends.filter(item => item.user.name.includes(value))
    })
  }

  componentDidMount() {
    this.props.onGetFriends()
    this.props.onGetFriendsOfFriends()
  }

  render() {
    const { friends, friendsOfFriends } = this.props
    const { filteredFriends, filteredFriendsOfFriends, filterValue, currentTab } = this.state

    return (
      <PageWrapper>
        <TopBar />
        <Header>
          <Tabs onTabClick={this.onTabSwitch} />
        </Header>
        <SidebarWrapper>
          <Paper>
            <SearchWrapper>
              <FormControl fullWidth>
                <Input
                  id="filter"
                  placeholder="Filter"
                  onChange={this.onFilterValueChange}
                  value={filterValue}
                  startAdornment={<InputAdornment position="start"><SearchIconStyled /></InputAdornment>}
                />
              </FormControl>
            </SearchWrapper>
            <ListStyled>
              {(currentTab === 'friends') && this.mapFriends((filterValue) ? filteredFriends : friends)}
              {(currentTab === 'friends-of-friends') && this.mapFriendsOfFriends((filterValue) ? filteredFriendsOfFriends : friendsOfFriends)}
            </ListStyled>
          </Paper>
        </SidebarWrapper>
        <ContentWrapper>
          {this.mainContent()}
        </ContentWrapper>
      </PageWrapper>
    )
  }
}