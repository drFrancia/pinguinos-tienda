import express from 'express';
import { getAllProductsAPI } from '../controllers/productController.js';

const router = express.Router();

router.get('/admin/products', getAllProductsAPI);

export default router;
