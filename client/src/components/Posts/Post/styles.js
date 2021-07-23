import { makeStyles } from "@material-ui/core";
import { pink, lightGreen, lightBlue, red } from "@material-ui/core/colors";

export default makeStyles((theme) => ({
  cardContainer: {
    display: "flex",
    flexDirection: "column",
    borderRadius: 12,
    minWidth: "250px",
    maxWidth: "500px",
    padding: theme.spacing(6),
  },
  image: {
    paddingLeft: theme.spacing(3),
    paddingRight: theme.spacing(3),
  },
  difficultiesContainer: {
    display: "flex",
    alignItems: "center",
    justifyContent: "center",
  },
  tagContainer: {
    margin: theme.spacing(0.5),
  },
  tagRow: {
    display: "flex",
    alignItems: "space-between",
    justifyContent: "center",
    margin: theme.spacing(3),
  },
  buttonsContainer: {
    display: "flex",
    alignItems: "center",
    justifyContent: "end",
  },
  button: {
    margin: theme.spacing(1),
  },
  userRow: {
    display: "flex",
    alignItems: "center",
    justifyContent: "start",
    marginTop: theme.spacing(1),
  },
  userRow__content: {
    margin: theme.spacing(2),
  },
  username: {
    fontSize: "16px",
  },
  customDifficultiesRow: {
    display: "flex",
    alignItems: "center",
    justifyContent: "center",
  },
  customTag: {
    display: "flex",
    padding: theme.spacing(1.5),
    marginLeft: theme.spacing(1),
    marginRight: theme.spacing(1),
    alignItems: "center",
    justifyContent: "center",
    height: "40px",
    borderRadius: 46,
  },
  customTag__text: {
    fontSize: "16px",
    fontWeight: "500",
  },
  countBadge: {
    display: "flex",
    alignItems: "center",
    justifyContent: "center",
    borderRadius: 16,
    width: "30px",
    height: "30px",
    padding: theme.spacing(0.5),
    margin: theme.spacing(0.8),
    backgroundColor: "#fff",
    color: "#000",
  },
  countValue: {
    fontSize: "16px",
  },
  pink: {
    backgroundColor: pink[400],
    color: "#fff",
  },
  green: {
    backgroundColor: lightGreen[400],
    color: "#fff",
  },
  blue: {
    backgroundColor: lightBlue[400],
    color: "#fff",
  },
  red: {
    backgroundColor: red[400],
    color: "#fff",
  },
}));