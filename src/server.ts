import express from 'express';
import cors from 'cors';
import routes from './routes';

export const app = express();

app.use(express.json());
app.use(express.urlencoded({ extended: false }));
app.use(cors());
app.use('/api', routes);
