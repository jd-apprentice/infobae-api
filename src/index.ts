import { app } from './server';
import { config } from '#config';

const startupMessage = 'Infobae API';

app.listen(config.app.port, (): void => console.info(startupMessage));
