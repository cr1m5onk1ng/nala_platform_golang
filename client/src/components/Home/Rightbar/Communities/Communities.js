import { Grid, makeStyles } from "@material-ui/core";
import Community from "./Community";

const useStyles = makeStyles((theme) => ({
  communitiesContainer: {
    paddingBottom: theme.spacing(5),
    paddingRight: theme.spacing(30),
  },
}));

const Communities = ({ communities }) => {
  const classes = useStyles();
  return (
    <Grid container className={classes.communitiesContainer}>
      {communities.map((com) => (
        <Grid item xs={12} sm={6} md={4}>
          <Community
            title={com.title}
            thumbnail={com.thumbnail}
            members={com.members}
          />
        </Grid>
      ))}
    </Grid>
  );
};

export default Communities;
