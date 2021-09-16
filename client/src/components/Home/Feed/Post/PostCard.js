import React from "react";
import Card from "@material-ui/core/Card";
import CardContent from "@material-ui/core/CardContent";
import CardMedia from "@material-ui/core/CardMedia";
import IconButton from "@material-ui/core/IconButton";
import Typography from "@material-ui/core/Typography";
import ShareIcon from "@material-ui/icons/Share";
import ReactTimeAgo from "react-time-ago";
import LinkIcon from "@material-ui/icons/Link";
import CardActionArea from "@material-ui/core/CardActionArea";
import PropTypes from "prop-types";
import { Rating } from "@material-ui/lab";
import SentimentVeryDissatisfiedIcon from "@material-ui/icons/SentimentVeryDissatisfied";
import SentimentDissatisfiedIcon from "@material-ui/icons/SentimentDissatisfied";
import SentimentSatisfiedIcon from "@material-ui/icons/SentimentSatisfied";
import SentimentSatisfiedAltIcon from "@material-ui/icons/SentimentSatisfiedAltOutlined";
import SentimentVerySatisfiedIcon from "@material-ui/icons/SentimentVerySatisfied";
import { CardHeader, makeStyles } from "@material-ui/core";
import { red, blue, green, blueGrey } from "@material-ui/core/colors";

const customIcons = {
  1: {
    icon: <SentimentVeryDissatisfiedIcon />,
    label: "Very Dissatisfied",
  },
  2: {
    icon: <SentimentDissatisfiedIcon />,
    label: "Dissatisfied",
  },
  3: {
    icon: <SentimentSatisfiedIcon />,
    label: "Neutral",
  },
  4: {
    icon: <SentimentSatisfiedAltIcon />,
    label: "Satisfied",
  },
  5: {
    icon: <SentimentVerySatisfiedIcon />,
    label: "Very Satisfied",
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
  cardContainer: {
    display: "flex",
    flexDirection: "column",
    borderRadius: 12,
    maxWidth: "500px",
    padding: theme.spacing(3),
    marginBottom: theme.spacing(5),
    [theme.breakpoints.down("xs")]: {
      padding: theme.spacing(0),
      borderRadius: 0,
    },
  },
  description: {
    minHeight: "30px",
    maxHeight: "60px",
    maxLines: 3,
    textOverflow: "ellipsis",
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
    margin: theme.spacing(0.3),
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
    marginLeft: theme.spacing(0.5),
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
    fontSize: "14px",
    fontWeight: "400",
    color: "grey",
  },
  pink: {
    backgroundColor: blueGrey[400],
    color: "#fff",
  },
  green: {
    backgroundColor: green[400],
    color: "#fff",
  },
  blue: {
    backgroundColor: blue[400],
    color: "#fff",
  },
  red: {
    backgroundColor: red[400],
    color: "#fff",
  },
}));

export default function PostCard({
  title,
  mediaType,
  description,
  url,
  creationTime,
  imageUrl,
  username,
  likesCount,
  easyCount,
  mediumCount,
  advancedCount,
  nativeCount,
}) {
  const classes = useStyles();

  const tag = (text, count, customClass) => {
    return (
      <div className={`${classes.customTag} ${customClass}`}>
        <Typography className={classes.customTag__text}>{text}</Typography>
        <div className={classes.countBadge}>
          <Typography className={classes.countValue}>{count}</Typography>
        </div>
      </div>
    );
  };

  const difficultiesRowCustom = () => {
    return (
      <div className={classes.customDifficultiesRow}>
        {tag("Easy", 16, classes.green)}
        {tag("Medium", 4, classes.blue)}
        {tag("Advanced", 1, classes.pink)}
        {tag("Native", 0, classes.red)}
      </div>
    );
  };

  const buttonsRow = () => {
    return (
      <div className={classes.buttonsContainer}>
        <IconButton className={classes.button}>
          <LinkIcon />
        </IconButton>
      </div>
    );
  };

  const difficultiesRow = () => {
    return (
      <div className={classes.customDifficultiesRow}>
        <Rating
          name="difficulty"
          defaultValue={2}
          IconContainerComponent={IconContainer}
          highlightSelectedOnly
        />
      </div>
    );
  };

  const userRow = (username) => {
    return (
      <CardActionArea>
        <div className={classes.userRow}>
          <Typography
            className={`${classes.userRow__content} ${classes.username}`}
            variant="subtitle2"
          >
            {username}
          </Typography>
          <Typography className={classes.userRow__content} variant="subtitle2">
            {" - "}
          </Typography>
          <ReactTimeAgo date={creationTime} locale="it-IT" />
        </div>
      </CardActionArea>
    );
  };

  return (
    <Card className={classes.cardContainer} elevation={6}>
      {userRow(username)}
      <CardHeader
        action={
          <IconButton>
            <ShareIcon />
          </IconButton>
        }
        title={title}
        subheader={mediaType}
      />
      <CardMedia
        className={classes.image}
        component="img"
        alt="Contemplative Reptile"
        height="140"
        image={imageUrl || "https://source.unsplash.com/random"}
        title={title}
      />
      <CardContent>
        <Typography
          variant="body1"
          component="p"
          color="textSecondary"
          className={classes.description}
        >
          {description}
        </Typography>
      </CardContent>
      {difficultiesRow()}
      {buttonsRow()}
    </Card>
  );
}
