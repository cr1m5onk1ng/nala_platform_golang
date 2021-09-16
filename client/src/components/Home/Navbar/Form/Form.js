import { React, useState } from "react";

import {
  Button,
  Container,
  FormControlLabel,
  FormLabel,
  makeStyles,
  MenuItem,
  Modal,
  Radio,
  RadioGroup,
  TextField,
} from "@material-ui/core";

const useStyles = makeStyles((theme) => ({
  container: {
    width: "80vw",
    height: "80vh",
    backgroundColor: "white",
    position: "absolute",
    top: 0,
    bottom: 0,
    left: 0,
    right: 70,
    margin: "auto",
    [theme.breakpoints.down("xs")]: {
      width: "100vw",
      height: "100vh",
    },
  },
  form: {
    padding: theme.spacing(2),
  },
  item: {
    marginBottom: theme.spacing(3),
  },
}));

const Form = ({ isOpen, setOpen, setOpenAlert }) => {
  const classes = useStyles();

  const [selectedCheckBok, setSelectedCheckBox] = useState(0);

  const onChangeCheckBoxState = (checked, position) => {
    if (checked) setSelectedCheckBox(position);
  };

  return (
    <Modal open={isOpen}>
      <Container className={classes.container}>
        <form className={classes.form} autoComplete="off">
          <div className={classes.item}>
            <TextField
              id="standard-basic"
              label="Url"
              size="small"
              style={{ width: "100%" }}
              required
            />
          </div>
          <div className={classes.item}>
            <TextField
              id="standard-basic"
              label="Title"
              size="small"
              style={{ width: "100%" }}
              required
            />
          </div>
          <div className={classes.item}>
            <TextField
              id="outlined-multiline-static"
              multiline
              rows={4}
              defaultValue="Tell fellow learners about this content..."
              variant="outlined"
              label="Description"
              size="small"
              style={{ width: "100%" }}
            />
          </div>
          <div className={classes.item}>
            <TextField required select label="Difficulty" value="Public">
              <MenuItem value="Beginner">Beginner</MenuItem>
              <MenuItem value="Intermediate">Intermediate</MenuItem>
              <MenuItem value="Advanced">Advanced</MenuItem>
              <MenuItem value="Native">Native</MenuItem>
            </TextField>
          </div>
          <div className={classes.item}>
            <TextField required select label="Media" value="Public">
              <MenuItem value="Video">Video</MenuItem>
              <MenuItem value="Movie/Series">Movie/Series</MenuItem>
              <MenuItem value="Music/Podcast">Music/Podcast</MenuItem>
              <MenuItem value="Web/Article">Web/Article</MenuItem>
            </TextField>
          </div>
          <div className={classes.item}>
            <FormLabel component="legend">Share with</FormLabel>
            <RadioGroup>
              <FormControlLabel
                value="My Communities"
                control={<Radio size="small" />}
                label="My Communities"
                onChange={(event) =>
                  onChangeCheckBoxState(event.target.checked, 0)
                }
                checked={selectedCheckBok == 0}
              />
              <FormControlLabel
                value="My Followers"
                control={<Radio size="small" />}
                label="My Followers"
                onChange={(event) =>
                  onChangeCheckBoxState(event.target.checked, 1)
                }
                checked={selectedCheckBok == 1}
              />
              <FormControlLabel
                value="Nobody"
                control={<Radio size="small" />}
                label="Nobody"
                onChange={(event) =>
                  onChangeCheckBoxState(event.target.checked, 2)
                }
                checked={selectedCheckBok == 2}
              />
            </RadioGroup>
          </div>
          <div className={classes.item}>
            <Button
              variant="outlined"
              color="inherit"
              style={{ marginRight: 20 }}
              onClick={() => setOpenAlert(true)}
            >
              Create
            </Button>
            <Button
              variant="outlined"
              color="secondary"
              onClick={() => setOpen(false)}
            >
              Cancel
            </Button>
          </div>
        </form>
      </Container>
    </Modal>
  );
};

export default Form;
