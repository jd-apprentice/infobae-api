import { Router } from "express";
const infobaeRoutes = Router();
import InfobaeController from "../controllers/infobaeController";

infobaeRoutes.get("/", InfobaeController.getPosts);

export default infobaeRoutes;
