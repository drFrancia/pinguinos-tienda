import express from 'express';
import { login } from '../controllers/authController.js';
import { authenticate } from '../middlewares/auth.js';
import { getDashboard } from '../controllers/adminController.js';
import productRoutes from './product.js';
import { getOrdersPage } from "../controllers/adminController.js";



const router = express.Router();

// Login
router.get('/login', (req, res) => res.render('login'));
router.post('/login', login);

// Dashboard protegido
router.get('/dashboard', authenticate, getDashboard);

// Rutas de productos
router.use('/products', authenticate, productRoutes);

//Rutas de ordenes
router.get("/orders", authenticate, getOrdersPage);


export default router;
