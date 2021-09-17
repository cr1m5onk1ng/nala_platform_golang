import { Avatar, Badge, IconButton, makeStyles } from "@material-ui/core";

const useStyles = makeStyles((theme) => ({
  badge: {
    margin: theme.spacing(3),
  },
  img: {
    width: theme.spacing(10),
    height: theme.spacing(10),
    border: "1px solid lightgray",
  },
}));

const Community = ({ title, thumbnail, members, onClick }) => {
  const classes = useStyles();
  return (
    <IconButton onClick={onClick}>
      <Badge
        className={classes.badge}
        badgeContent={members}
        max={10000}
        color="secondary"
      >
        <Avatar className={classes.img} alt={title} src={thumbnail} />
      </Badge>
    </IconButton>
  );
};

export default Community;
