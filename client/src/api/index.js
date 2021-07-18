import axios from 'axios';

const API = axios.create({ baseURL: 'http://localhost:5000' });

API.interceptors.request.use((req) => {
  if (localStorage.getItem('profile')) {
    req.headers.Authorization = `Bearer ${JSON.parse(localStorage.getItem('profile')).token}`;
  }
  return req;
});

export const fetchPost = (id) => API.get(`/posts/${id}`);
export const fetchPostsByLanguage = (lang, page) => API.get(`/posts?lang=${lang}?page=${page}`);
export const fetchPostsByLanguageAndMedia = (lang, media, page) => API.get(`/posts?lang=${lang}?media=${media}?page=${page}`);
export const fetchPostsByLanguageAndCategory = (lang, cat, page) => API.get(`/posts?lang=${lang}?cat=${cat}?page=${page}`);
export const fetchPostsByLanguageAndMediaAndCategory = (lang, media, cat, page) => 
                                                        API.get(`/posts?lang=${lang}?media=${media}?cat=${cat}?page=${page}`);
export const fetchPostsByCreator = (creator) => API.get(`/posts/creator?name=${creator}`);
export const fetchPostsBySearch = (query) => API.get(`/posts/search?searchQuery=${query.search || 'none'}&tags=${query.tags}`);
export const createPost = (post) => API.post('/posts', post);
export const likePost = (id) => API.patch(`/posts/${id}/likePost`);
export const commentPost = (value, id) => API.post(`/posts/${id}/commentPost`, { value });
export const updatePost = (id, updatedPost) => API.patch(`/posts/${id}`, updatedPost);
export const deletePost = (id) => API.delete(`/posts/${id}`);
export const likeComment = (id) => API.patch(`/comments/${id}/likeComment`);

export const signIn = (data) => API.post('/user/signin', data);
export const signUp = (data) => API.post('/user/signup', data);