import React from 'react'
import Button from '@material-ui/core/Button'
import Dialog from '@material-ui/core/Dialog'
import DialogActions from '@material-ui/core/DialogActions'
import { connect } from 'react-redux'
import { ApiTypes } from 'src/types'
import Actions from '@store/actions'
import ErrorOutlineIcon from '@material-ui/icons/ErrorOutline'
import MenuItem from '@material-ui/core/MenuItem'
import ListItemText from '@material-ui/core/ListItemText'
import {
  ListItemIconStyled,
  DialogTextWrapper,
  DialogTitleStyled,
  DialogContentStyled,
  ErrorMessage,
} from '@view/shared/styles'
import { TextFieldStyled } from './styles'

interface Props {
  message: string
  id: string
  sourceHost: string
  onComplainContent: (data: ApiTypes.Messages.ReportMessageHub) => void
}

const ComplainContentDialog: React.SFC<Props> = (props) => {
  const [open, setOpen] = React.useState<boolean>(false)
  const [description, setDescription] = React.useState<string>('')
  const [isError, setError] = React.useState<boolean>(false)
  const { onComplainContent, id, sourceHost } = props

  const onComplane = () => {

    if (!Boolean(description)) {
      setError(true)
    } else {
      onComplainContent({
        host: sourceHost,
        body: {
          message_id: id,
          report: description,
        }
      })
      setOpen(false)
      setError(false)
    }
    
  }

  return (
    <div>
      <MenuItem onClick={() => setOpen(true)}>
        <ListItemIconStyled>
          <ErrorOutlineIcon fontSize="small" />
        </ListItemIconStyled>
        <ListItemText primary="Report" />
      </MenuItem>
      <Dialog
        open={open}
        onClose={() => setOpen(false)}
        aria-labelledby="alert-dialog-title"
        aria-describedby="alert-dialog-description"
      >
        <DialogTitleStyled id="alert-dialog-title">Tell us why you find this content objectionable:</DialogTitleStyled>
        <DialogContentStyled>
          <DialogTextWrapper>
            <TextFieldStyled
              id="outlined-multiline-static"
              label="Description"
              multiline
              rows={4}
              onChange={(event) => setDescription(event?.target?.value)}
              defaultValue={description}
              variant="outlined"
            />
          </DialogTextWrapper>
          {isError && <ErrorMessage>The message cannot be empty</ErrorMessage>}
        </DialogContentStyled>
        <DialogActions>
          <Button onClick={() => setOpen(false)}>
            Cancel
          </Button>
          <Button color="secondary" onClick={onComplane} autoFocus>
            Report
          </Button>
        </DialogActions>
      </Dialog>
    </div>
  )
}

type DispatchProps = Pick<Props, 'onComplainContent'>
const mapDispatchToProps = (dispatch): DispatchProps => ({
  onComplainContent: (data: ApiTypes.Messages.ReportMessageHub) => dispatch(Actions.messages.reportMessageHubRequest(data)),
})

export default connect(null, mapDispatchToProps)(ComplainContentDialog)