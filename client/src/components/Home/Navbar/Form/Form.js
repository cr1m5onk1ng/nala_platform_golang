import { React, useState } from "react";

import {
  Button,
  Container,
  FormControl,
  FormControlLabel,
  FormLabel,
  InputLabel,
  makeStyles,
  MenuItem,
  Modal,
  Radio,
  RadioGroup,
  Select,
  Snackbar,
  TextField,
} from "@material-ui/core";
import MuiAlert from "@material-ui/lab/Alert";

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

function Alert(props) {
  return <MuiAlert elevation={6} variant="filled" {...props} />;
}

const Form = ({ isOpen, setOpen }) => {
  const classes = useStyles();

  const [selectedCheckBok, setSelectedCheckBox] = useState("Communities");
  const [selectedDifficulty, setSelectedDifficulty] = useState("Beginner");
  const [selectedMedia, setSelectedMedia] = useState("Web/Article");
  const [openAlert, setOpenAlert] = useState(false); // snackbar state

  // taken straight from material-ui
  const handleSnackbarClose = (event, reason) => {
    if (reason === "clickaway") {
      return;
    }
    setOpenAlert(false);
  };

  return (
    <Modal open={isOpen} onClose={() => setOpen(false)}>
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
              defaultValue="Help your fellow learners by describing the content..."
              variant="outlined"
              label="Description"
              size="small"
              style={{ width: "100%" }}
            />
          </div>
          <div className={classes.item}>
            <TextField
              select
              label="Difficulty"
              value={selectedDifficulty}
              onChange={(e) => setSelectedDifficulty(e.target.value)}
              required
            >
              <MenuItem value="Beginner">Beginner</MenuItem>
              <MenuItem value="Intermediate">Intermediate</MenuItem>
              <MenuItem value="Advanced">Advanced</MenuItem>
              <MenuItem value="Native">Native</MenuItem>
            </TextField>
          </div>
          <div className={classes.item}>
            <TextField
              select
              label="Media"
              value={selectedMedia}
              onChange={(e) => setSelectedMedia(e.target.value)}
              required
            >
              <MenuItem value="Web/Article">Web/Article</MenuItem>
              <MenuItem value="Video">Video</MenuItem>
              <MenuItem value="Movie/Series">Movie/Series</MenuItem>
              <MenuItem value="Podcast/Music">Podcast/Music</MenuItem>
              <MenuItem value="Video">Video</MenuItem>
            </TextField>
          </div>
          <div className={classes.item}>
            <FormLabel component="legend">Share with</FormLabel>
            <RadioGroup
              value={selectedCheckBok}
              onChange={(e) => setSelectedCheckBox(e.target.value)}
            >
              <FormControlLabel
                value="Communities"
                control={<Radio size="small" />}
                label="Communities"
              />
              <FormControlLabel
                value="Followers"
                control={<Radio size="small" />}
                label="Followers"
              />
              <FormControlLabel
                value="Nobody"
                control={<Radio size="small" />}
                label="Nobody"
              />
            </RadioGroup>
          </div>
          <div className={classes.item}>
            <Button
              variant="outlined"
              color="primary"
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
        <Snackbar
          open={openAlert}
          autoHideDuration={4000}
          onClose={handleSnackbarClose}
          anchorOrigin={{ vertical: "bottom", horizontal: "right" }}
        >
          <Alert onClose={handleSnackbarClose} severity="success">
            Post addedd succefully
          </Alert>
        </Snackbar>
      </Container>
    </Modal>
  );
};

export default Form;
