import { Grid, makeStyles } from "@material-ui/core";
import Feed from "./Feed/Feed";
import Leftbar from "./Leftbar/Leftbar";
import Navbar from "./Navbar/Navbar";
import Rightbar from "./Rightbar/Rightbar";

const useStyles = makeStyles((theme) => ({
  right: {
    [theme.breakpoints.down("xs")]: {
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
        <Grid item sm={7} xs={10}>
          <Feed />
        </Grid>
        <Grid item sm={3} className={classes.right}>
          <Rightbar />
        </Grid>
      </Grid>
    </div>
  );
};

export default Home;
