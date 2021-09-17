import React from "react";
import { Grid, makeStyles } from "@material-ui/core";
import Discussion from "./Discussion";

const useStyles = makeStyles((theme) => ({
  discussionsContainer: {
    paddingBottom: theme.spacing(5),
    paddingRight: theme.spacing(20),
    [theme.breakpoints.down("md")]: {
      display: "none",
    },
  },
}));

const Discussions = ({ discussions }) => {
  const classes = useStyles();
  return (
    <Grid container className={classes.discussionsContainer}>
      {discussions.map((disc) => (
        <Grid item xs={12} sm={12} md={12} lg={12} xl={12}>
          <Discussion
            username={disc.user}
            community={disc.community}
            creationTime={new Date().getTime()}
            title={disc.title}
            commentsCount={disc.numComments}
          />
        </Grid>
      ))}
    </Grid>
  );
};

export default Discussions;
