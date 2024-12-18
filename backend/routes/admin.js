import express from 'express';
import { authenticate } from '../middlewares/auth.js';
import { login } from '../controllers/authController.js';
import {
  getAllProducts,
  createProduct,
  editProduct,
  deleteProduct,
} from '../controllers/productController.js';

const router = express.Router();

// Login
router.get('/login', (req, res) => res.render('login'));
router.post('/login', login);

// Dashboard protegido
router.get('/dashboard', authenticate, (req, res) => {
  res.render('dashboard', { user: req.user });
});

// Rutas de productos
router.get('/products', authenticate, getAllProducts);
router.post('/products/create', authenticate, createProduct);
router.post('/products/edit/:id', authenticate, editProduct);
router.post('/products/delete/:id', authenticate, deleteProduct);

export default router;
