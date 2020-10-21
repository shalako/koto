import React from 'react'
import Button from '@material-ui/core/Button'
import Dialog from '@material-ui/core/Dialog'
import DialogActions from '@material-ui/core/DialogActions'
import DeleteIcon from '@material-ui/icons/Delete'
import IconButton from '@material-ui/core/IconButton'
import { connect } from 'react-redux'
import Tooltip from '@material-ui/core/Tooltip'
import { ApiTypes } from 'src/types'
import Actions from '@store/actions'
import VisibilityOffIcon from '@material-ui/icons/VisibilityOff'
import { DialogTitleStyled } from '@view/shared/styles'

interface Props {
  id: string
  sourceHost: string

  onHideMessage: (data: ApiTypes.Messages.Hide) => void
}

const HideMessageDialog: React.SFC<Props> = (props) => {
  const [open, setOpen] = React.useState(false)
  const { onHideMessage, id, sourceHost } = props

  const onHide = () => {
    onHideMessage({
      host: sourceHost,
      id: id
    })
    setOpen(false)
  }

  return (
    <div>
      <Tooltip title={`Hide`} onClick={() => setOpen(true)}>
        <IconButton>
          <VisibilityOffIcon />
        </IconButton>
      </Tooltip>
      <Dialog
        open={open}
        onClose={() => setOpen(false)}
        aria-labelledby="alert-dialog-title"
        aria-describedby="alert-dialog-description"
      >
        <DialogTitleStyled id="alert-dialog-title">Are you sure, there is no way to undo this?</DialogTitleStyled>
        <DialogActions>
          <Button onClick={() => setOpen(false)}>
            Cancel
          </Button>
          <Button color="secondary" onClick={onHide} autoFocus>
            Hide
          </Button>
        </DialogActions>
      </Dialog>
    </div>
  )
}

type DispatchProps = Pick<Props, 'onHideMessage'>
const mapDispatchToProps = (dispatch): DispatchProps => ({
  onHideMessage: (data: ApiTypes.Messages.Hide) => dispatch(Actions.messages.hideMessageRequest(data)),
})

export default connect(null, mapDispatchToProps)(HideMessageDialog)