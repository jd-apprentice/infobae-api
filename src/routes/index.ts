import { Router } from "express";
import infobaeRoutes from "../routes/infobae.routes";
const routes = Router();

routes.use("/infobae", infobaeRoutes);

export default routes;
