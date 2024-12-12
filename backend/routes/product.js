import express from 'express';
import {
  getAllProducts,
  createProduct,
  editProduct,
  deleteProduct,
} from '../controllers/productController.js';

const router = express.Router();

router.get('/', getAllProducts);
router.post('/create', createProduct);
router.post('/edit/:id', editProduct);
router.post('/delete/:id', deleteProduct);

export default router;
