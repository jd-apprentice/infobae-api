import { Router } from "express";
const infobaeRoutes = Router();
import InfobaeController from "../controllers/infobaeController";

infobaeRoutes.get("/economia", InfobaeController.getPost);

export default infobaeRoutes;
