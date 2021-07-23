import { Typography } from "@material-ui/core";
import { mergeClasses } from "@material-ui/styles";
import React from "react";
import useStyles from "./styles";

function DifficultiesSection({
  easyCount,
  mediumCount,
  advancedCount,
  nativeCount,
}) {
  const classes = useStyles();

  return (
    <div className={classes.container}>
      <div className={`${classes.tag} ${classes.easyTag}`}>
        <Typography variant="button">Easy</Typography>
        <Typography variant="body1">{easyCount}</Typography>
      </div>
      <div className={`${classes.tag} ${classes.mediumTag}`}>
        <Typography variant="button">Medium</Typography>
        <Typography variant="body1">{mediumCount}</Typography>
      </div>
      <div className={`${classes.tag} ${classes.advancedTag}`}>
        <Typography variant="button">Advanced</Typography>
        <Typography variant="body1">{advancedCount}</Typography>
      </div>
      <div className={`${classes.tag} ${classes.nativeTag}`}>
        <Typography variant="button">Native</Typography>
        <Typography variant="body1">{nativeCount}</Typography>
      </div>
    </div>
  );
}

export default DifficultiesSection;
