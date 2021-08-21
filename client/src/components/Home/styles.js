import { makeStyles } from "@material-ui/core";
import { DRAWER_WIDTH } from "../../constants/layout";

export default makeStyles((theme) => ({
  home: {
    backgroundColor: theme.palette.background.dark,
  },
  appContainer: {
    marginLeft: "7vw",
    marginRight: "7vw",
  },
  bodyContainer: {
    display: "flex",
    height: "100vh",
    alignItems: "start",
    justifyContent: "center",
    marginLeft: "7vw",
    marginRight: "7vw",
  },
}));
