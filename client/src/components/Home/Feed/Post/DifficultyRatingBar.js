import { React, useState } from "react";
import { Rating } from "@material-ui/lab";
import { makeStyles, Typography, Box } from "@material-ui/core";
import { Star } from "@material-ui/icons";

const labels = {
  1: "Beginner",
  2: "Intermediate",
  3: "Advanced",
  4: "Native",
};

const useStyles = makeStyles((theme) => ({
  difficultyRatingTitle: {
    margin: theme.spacing(2),
  },
  customDifficultiesRow: {
    display: "flex",
    alignItems: "center",
    justifyContent: "center",
  },
}));

const DifficultyRatingBar = () => {
  const classes = useStyles();
  const [value, setValue] = useState(2);
  const [hover, setHover] = useState(-1);
  return (
    <div className={classes.customDifficultiesRow}>
      <Typography className={classes.difficultyRatingTitle} variant="body1">
        {" "}
        Difficulty{" "}
      </Typography>
      <Rating
        name="difficulty"
        value={value}
        precision={1}
        onChange={(event, newValue) => {
          setValue(newValue);
        }}
        onChangeActive={(event, newHover) => {
          setHover(newHover);
        }}
        emptyIcon={<Star style={{ opacity: 0.55 }} fontSize="inherit" />}
      />
      {value !== null && (
        <Box sx={{ ml: 2 }}>{labels[hover !== -1 ? hover : value]}</Box>
      )}
    </div>
  );
};

export default DifficultyRatingBar;
