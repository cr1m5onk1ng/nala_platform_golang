import { makeStyles } from "@material-ui/core";
import { green } from "@material-ui/core/colors";
import { red } from "@material-ui/core/colors";
import { blue } from "@material-ui/core/colors";
import { lightGreen } from "@material-ui/core/colors";
import { lightBlue } from "@material-ui/core/colors";
import { blueGrey } from "@material-ui/core/colors";

export default makeStyles((theme) => ({
  container: {
    display: "flex",
    alignItems: "center",
    paddingLeft: theme.spacing(1),
    paddingBottom: theme.spacing(1),
  },
  tag: {
    display: "flex",
    maxWidth: "50px",
    alignItems: "space-between",
    borderRadius: 15,
    padding: theme.spacing(3),
  },
  text: {
    fontSize: 12,
  },
  number: {
    fontSize: 12,
  },
  easyTag: {
    backgroundColor: lightGreen[200],
  },
  mediumTag: {
    backgroundColor: lightBlue[200],
  },
  advancedTag: {
    backgroundColor: red[200],
  },
  nativeTag: {
    backgroundColor: blueGrey[200],
  },
}));
