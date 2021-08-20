import { makeStyles } from "@material-ui/core";
import { DRAWER_WIDTH } from "../../constants/layout";

export default makeStyles((theme) => ({
  home: {
    backgroundColor: theme.palette.background.dark,
    padding: "60px",
  },
  container: {
    display: "flex",
    alignItems: "center",
  },
  bodyContainer: {
    display: "flex",
    alignItems: "center",
    justifyContent: "center",
    //marginLeft: `${DRAWER_WIDTH}px`,
  },
}));
