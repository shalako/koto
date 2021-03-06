import React, { ChangeEvent } from 'react'
import ListItem from '@material-ui/core/ListItem'
import Paper from '@material-ui/core/Paper'
import Divider from '@material-ui/core/Divider'
import ListItemText from '@material-ui/core/ListItemText'
import ListItemAvatar from '@material-ui/core/ListItemAvatar'
import { connect } from 'react-redux'
import Actions from '@store/actions'
import { StoreTypes, ApiTypes } from 'src/types'
import selectors from '@selectors/index'
import { getAvatarUrl } from '@services/avatarUrl'
import {
  UsersWrapper,
  ListStyled,
  SearchWrapper,
  SearchInput,
  EmptyMessage,
  UserNameLink,
  PageWrapper,
  ListItemWrapper,
  AvatarStyled,
  SearchIconStyled,
} from './styles'
import { AvatarWrapperLink } from '@view/pages/MessagesPage/styles'

export interface Props {
  friends: ApiTypes.Friends.Friend[]
  onGetFriends: () => void
  onAddFriend: (data: ApiTypes.Friends.Request) => void
}

interface State {
  searchValue: string
  searchResult: ApiTypes.Friends.Friend[]
}

class Friends extends React.Component<Props, State> {

  state = {
    searchValue: '',
    searchResult: [],
  }

  searchInputRef = React.createRef<HTMLInputElement>()

  showEmptyListMessage = () => {
    const { searchValue } = this.state

    if (searchValue) {
      return <EmptyMessage>No one's been found.</EmptyMessage>
    } else {
      return <EmptyMessage>No friends yet.</EmptyMessage>
    }
  }

  mapFriends = (friends: ApiTypes.Friends.Friend[]) => {

    if (!friends || !friends?.length) {
      return this.showEmptyListMessage()
    }

    return friends.map(item => (
      <ListItemWrapper key={item.user.id}>
        <ListItem>
          <ListItemAvatar>
            <AvatarWrapperLink to={`/profile/user?id=${item.user.id}`}>
              <AvatarStyled
                alt={item.user.name}
                src={getAvatarUrl(item.user.id)} />
            </AvatarWrapperLink>
          </ListItemAvatar>
          <ListItemText
            primary={<UserNameLink to={`/profile/user?id=${item.user.id}`}>{item.user.name}</UserNameLink>} />
        </ListItem>
        <Divider variant="inset" component="li" />
      </ListItemWrapper>
    ))
  }

  onSearch = (event: ChangeEvent<HTMLInputElement>) => {
    const { friends } = this.props
    const { value } = event.currentTarget

    this.setState({
      searchValue: value,
      searchResult: friends.filter(item => item.user.name.toLowerCase().includes(value.toLowerCase()))
    })
  }

  componentDidMount() {
    this.props.onGetFriends()
  }

  render() {
    const { friends } = this.props
    const { searchResult, searchValue } = this.state

    return (
      <PageWrapper>
        <UsersWrapper>
          <Paper>
            <SearchWrapper>
              <SearchIconStyled onClick={() => this.searchInputRef?.current?.focus()} />
              <SearchInput
                ref={this.searchInputRef}
                id="filter"
                placeholder="Filter"
                onChange={this.onSearch}
                value={searchValue}
              />
            </SearchWrapper>
            <ListStyled>
              {this.mapFriends((searchValue) ? searchResult : friends)}
            </ListStyled>
          </Paper>
        </UsersWrapper>
      </PageWrapper>
    )
  }
}

type StateProps = Pick<Props, 'friends'>
const mapStateToProps = (state: StoreTypes): StateProps => ({
  friends: selectors.friends.friends(state),
})

type DispatchProps = Pick<Props, 'onGetFriends' | 'onAddFriend'>
const mapDispatchToProps = (dispatch): DispatchProps => ({
  onGetFriends: () => dispatch(Actions.friends.getFriendsRequest()),
  onAddFriend: (data: ApiTypes.Friends.Request) => dispatch(Actions.friends.addFriendRequest(data)),
})

export default connect(mapStateToProps, mapDispatchToProps)(Friends)