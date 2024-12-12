import express from "express";
import { getAllProducts, createProduct } from "../controllers/productController.js"; 

const router = express.Router();

// Rutas CRUD
router.get("/", getAllProducts); // Ruta para obtener productos
router.post("/", createProduct); // Ruta para crear productos

export default router;
