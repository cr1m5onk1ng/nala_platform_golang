import { Grid, makeStyles } from "@material-ui/core";
import Feed from "./Feed/Feed";
import Leftbar from "./Leftbar/Leftbar";
import Navbar from "./Navbar/Navbar";
import Rightbar from "./Rightbar/Rightbar";

const useStyles = makeStyles((theme) => ({
  right: {
    [theme.breakpoints.down("md")]: {
      display: "none",
    },
  },
}));

const Home = ({ darkMode, setDarkMode }) => {
  const classes = useStyles();
  return (
    <div>
      <Navbar />
      <Grid container>
        <Grid item sm={2} xs={2}>
          <Leftbar />
        </Grid>
        <Grid item sm={7} xs={10} sm={10} md={10} lg={7} xl={7}>
          <Feed />
        </Grid>
        <Grid item sm={3} lg={3} xl={3} className={classes.right}>
          <Rightbar />
        </Grid>
      </Grid>
    </div>
  );
};

export default Home;
