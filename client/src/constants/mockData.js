import { FETCH_POSTS_BY_LANGUAGE_AND_CATEGORY_AND_MEDIA } from "./actionTypes";
import {
  PlayCircleFilled,
  Language,
  Radio,
  Home,
  Settings,
  List,
  Movie,
  ExitToApp,
} from "@material-ui/icons";

import { red, blue, green, blueGrey } from "@material-ui/core/colors";

export const leftbarNavIcons = [
  {
    icon: Home,
    title: "Home",
  },
  {
    icon: List,
    title: "Study Lists",
  },
];

export const leftbarMediaIcons = [
  {
    icon: PlayCircleFilled,
    title: "Videos",
    color: red[400],
  },
  {
    icon: Language,
    title: "Web Articles",
    color: blue[400],
  },
  {
    icon: Radio,
    title: "Music /\nPodcasts",
    color: green[400],
  },
  {
    icon: Movie,
    title: "Movies /\nSeries",
    color: blueGrey[400],
  },
];

export const leftbarFooterIcons = [
  {
    icon: Settings,
    title: "Settings",
  },
  {
    icon: ExitToApp,
    title: "Logout",
  },
];

export const posts = [
  {
    title:
      "Best youtube video ever let's just make it a little bit longer to see what happens",
    mediaType: "Video",
    description:
      "whatever a description may look like. Something like this video is so cool bruh you sould watch it let's just make it longer for the sake of testing",
    url: "https://www.google.com",
    creationTime: Date.now(),
    imageUrl: undefined,
    username: "cr1m5onk1ng",
    likesCount: "500",
    easyCount: "120",
    mediumCount: "250",
    advancedCount: "11",
    nativeCount: "0",
  },
  {
    title: "Best youtube video ever",
    mediaType: "Video",
    description:
      "whatever a description may look like. Something like this video is so cool bruh you sould watch it",
    url: "https://www.google.com",
    creationTime: Date.now(),
    imageUrl: undefined,
    username: "cr1m5onk1ng",
    likesCount: "500",
    easyCount: "120",
    mediumCount: "250",
    advancedCount: "11",
    nativeCount: "0",
  },
  {
    title: "Best youtube video ever",
    mediaType: "Video",
    description:
      "whatever a description may look like. Something like this video is so cool bruh you sould watch it",
    url: "https://www.google.com",
    creationTime: Date.now(),
    imageUrl: undefined,
    username: "cr1m5onk1ng",
    likesCount: "500",
    easyCount: "120",
    mediumCount: "250",
    advancedCount: "11",
    nativeCount: "0",
  },
  {
    title: "Best youtube video ever",
    mediaType: "Video",
    description:
      "whatever a description may look like. Something like this video is so cool bruh you sould watch it",
    url: "https://www.google.com",
    creationTime: Date.now(),
    imageUrl: undefined,
    username: "cr1m5onk1ng",
    likesCount: "500",
    easyCount: "120",
    mediumCount: "250",
    advancedCount: "11",
    nativeCount: "0",
  },
  {
    title: "Best youtube video ever",
    mediaType: "Video",
    description:
      "whatever a description may look like. Something like this video is so cool bruh you sould watch it",
    url: "https://www.google.com",
    creationTime: Date.now(),
    imageUrl: undefined,
    username: "cr1m5onk1ng",
    likesCount: "500",
    easyCount: "120",
    mediumCount: "250",
    advancedCount: "11",
    nativeCount: "0",
  },
  {
    title: "Best youtube video ever",
    mediaType: "Video",
    description:
      "whatever a description may look like. Something like this video is so cool bruh you sould watch it",
    url: "https://www.google.com",
    creationTime: Date.now(),
    imageUrl: undefined,
    username: "cr1m5onk1ng",
    likesCount: "500",
    easyCount: "120",
    mediumCount: "250",
    advancedCount: "11",
    nativeCount: "0",
  },
  {
    title: "Best youtube video ever",
    mediaType: "Video",
    description:
      "whatever a description may look like. Something like this video is so cool bruh you sould watch it",
    url: "https://www.google.com",
    creationTime: Date.now(),
    imageUrl: undefined,
    username: "cr1m5onk1ng",
    likesCount: "500",
    easyCount: "120",
    mediumCount: "250",
    advancedCount: "11",
    nativeCount: "0",
  },
  {
    title: "Best youtube video ever",
    mediaType: "Video",
    description:
      "whatever a description may look like. Something like this video is so cool bruh you sould watch it",
    url: "https://www.google.com",
    creationTime: Date.now(),
    imageUrl: undefined,
    username: "cr1m5onk1ng",
    likesCount: "500",
    easyCount: "120",
    mediumCount: "250",
    advancedCount: "11",
    nativeCount: "0",
  },
  {
    title: "Best youtube video ever",
    mediaType: "Video",
    description:
      "whatever a description may look like. Something like this video is so cool bruh you sould watch it",
    url: "https://www.google.com",
    creationTime: Date.now(),
    imageUrl: undefined,
    username: "cr1m5onk1ng",
    likesCount: "500",
    easyCount: "120",
    mediumCount: "250",
    advancedCount: "11",
    nativeCount: "0",
  },
];

export const communities = [
  {
    title: "en",
    thumbnail: "https://freesvg.org/img/tobias-Flag-of-the-United-Kingdom.png",
    members: 1345,
  },
  {
    title: "es",
    thumbnail:
      "https://cdn.pixabay.com/photo/2012/04/11/15/33/spain-28530_1280.png",
    members: 8765,
  },
  {
    title: "ja",
    thumbnail:
      "https://cdn.pixabay.com/photo/2012/04/13/12/23/flag-32177_1280.png",
    members: 467,
  },
  {
    title: "fr",
    thumbnail:
      "https://upload.wikimedia.org/wikipedia/commons/1/1d/Flag_of_France_%28bordered%29.svg",
    members: 2451,
  },
];

export const members = [
  {
    username: "John Smith",
    avatar: "https://material-ui.com/static/images/avatar/1.jpg",
  },
  {
    username: "Gary Paige",
    avatar: "https://material-ui.com/static/images/avatar/2.jpg",
  },
  {
    username: "William Ford",
    avatar: "https://material-ui.com/static/images/avatar/3.jpg",
  },
  {
    username: "Vitalik Buterin",
    avatar: "https://material-ui.com/static/images/avatar/6.jpg",
  },
  {
    username: "John Ford",
    avatar: "https://material-ui.com/static/images/avatar/7.jpg",
  },
  {
    username: "Your Mom",
    avatar: "https://material-ui.com/static/images/avatar/8.jpg",
  },
  {
    username: "John Smith",
    avatar: "https://material-ui.com/static/images/avatar/1.jpg",
  },
  {
    username: "Gary Paige",
    avatar: "https://material-ui.com/static/images/avatar/2.jpg",
  },
  {
    username: "William Ford",
    avatar: "https://material-ui.com/static/images/avatar/3.jpg",
  },
  {
    username: "Vitalik Buterin",
    avatar: "https://material-ui.com/static/images/avatar/6.jpg",
  },
  {
    username: "John Ford",
    avatar: "https://material-ui.com/static/images/avatar/7.jpg",
  },
  {
    username: "Your Mom",
    avatar: "https://material-ui.com/static/images/avatar/8.jpg",
  },
  {
    username: "John Smith",
    avatar: "https://material-ui.com/static/images/avatar/1.jpg",
  },
  {
    username: "Gary Paige",
    avatar: "https://material-ui.com/static/images/avatar/2.jpg",
  },
  {
    username: "William Ford",
    avatar: "https://material-ui.com/static/images/avatar/3.jpg",
  },
  {
    username: "Vitalik Buterin",
    avatar: "https://material-ui.com/static/images/avatar/6.jpg",
  },
  {
    username: "John Ford",
    avatar: "https://material-ui.com/static/images/avatar/7.jpg",
  },
  {
    username: "Your Mom",
    avatar: "https://material-ui.com/static/images/avatar/8.jpg",
  },
];

const sidebarIcons = [];
