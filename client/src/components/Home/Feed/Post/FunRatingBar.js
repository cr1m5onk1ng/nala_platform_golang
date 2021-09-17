import React from "react";
import { Rating } from "@material-ui/lab";
import { makeStyles, Typography } from "@material-ui/core";
import SentimentVeryDissatisfiedIcon from "@material-ui/icons/SentimentVeryDissatisfied";
import SentimentDissatisfiedIcon from "@material-ui/icons/SentimentDissatisfied";
import SentimentSatisfiedIcon from "@material-ui/icons/SentimentSatisfied";
import SentimentSatisfiedAltIcon from "@material-ui/icons/SentimentSatisfiedAltOutlined";
import SentimentVerySatisfiedIcon from "@material-ui/icons/SentimentVerySatisfied";
import PropTypes from "prop-types";

const customIcons = {
  1: {
    icon: <SentimentVeryDissatisfiedIcon />,
    label: "Not fun at all",
  },
  2: {
    icon: <SentimentDissatisfiedIcon />,
    label: "Not particularly fuin",
  },
  3: {
    icon: <SentimentSatisfiedIcon />,
    label: "Not bad",
  },
  4: {
    icon: <SentimentSatisfiedAltIcon />,
    label: "Quite some fun",
  },
  5: {
    icon: <SentimentVerySatisfiedIcon />,
    label: "Lots of fun",
  },
};

const IconContainer = (props) => {
  const { value, ...other } = props;
  return <span {...other}>{customIcons[value].icon}</span>;
};

IconContainer.propTypes = {
  value: PropTypes.number.isRequired,
};

const useStyles = makeStyles((theme) => ({
  funFactorTitle: {
    margin: theme.spacing(2),
  },
  customDifficultiesRow: {
    display: "flex",
    alignItems: "center",
    justifyContent: "center",
  },
}));

const FunRatingBar = () => {
  const classes = useStyles();
  return (
    <div className={classes.customDifficultiesRow}>
      <Typography className={classes.funFactorTitle} variant="body1">
        {" "}
        Fun Factor{" "}
      </Typography>
      <Rating
        name="fun factor"
        defaultValue={3}
        IconContainerComponent={IconContainer}
        highlightSelectedOnly
      />
    </div>
  );
};

export default FunRatingBar;
