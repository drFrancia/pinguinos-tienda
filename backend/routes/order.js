import express from "express";
import { createOrder, getOrders } from "../controllers/orderController.js";

const router = express.Router();

// Ruta para registrar un pedido
router.post("/create", createOrder);

// Ruta para obtener todos los pedidos
router.get("/", getOrders);

export default router;
