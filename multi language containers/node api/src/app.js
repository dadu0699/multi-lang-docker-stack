import path from 'node:path';

import cors from 'cors';
import express from 'express';
import morgan from 'morgan';

import { __dirname } from './utils/consts.js';

// Create Express server
export const app = express();

// Disable x-powered-by header for security reasons
app.disable('x-powered-by');

// Middlewares
app.use(morgan('dev'));
app.use(cors());
app.use(express.json());
app.use(express.urlencoded({ extended: true }));
app.use(express.static(path.join(__dirname, 'public')));

// Favicon
app.get('/favicon.ico', (_req, res) => res.status(204));

// Routes
app.get('/', (_req, res) =>
  res.status(200).json({ message: 'Welcome to the Node.js API!' })
);

// Port assignment
const PORT = process.env.PORT || 3000;
export const server = app.listen(PORT, () => {
  console.log(`Server is running on port ${PORT}`);
});
