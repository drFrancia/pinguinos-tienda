import Product from '../models/Product.js';

// Obtener todos los productos para la API y vistas
export const getAllProducts = async (req, res) => {
  try {
    const products = await Product.find();
    res.render('products', { products });
  } catch (err) {
    res.status(500).send("Error al obtener productos: " + err.message);
  }
};

// API para obtener productos en JSON
export const getProductsAPI = async (req, res) => {
  try {
    const products = await Product.find();
    res.status(200).json(products);
  } catch (err) {
    res.status(500).json({ message: err.message });
  }
};

// Crear producto
export const createProduct = async (req, res) => {
  try {
    const { name, price, description } = req.body;
    await Product.create({ name, price, description });
    res.redirect('/admin/products');
  } catch (err) {
    res.status(500).send("Error al crear producto: " + err.message);
  }
};

// Editar producto
export const editProduct = async (req, res) => {
  try {
    const { id } = req.params;
    const { name, price, description } = req.body;
    await Product.findByIdAndUpdate(id, { name, price, description });
    res.redirect('/admin/products');
  } catch (err) {
    res.status(500).send("Error al editar producto: " + err.message);
  }
};

// Eliminar producto
export const deleteProduct = async (req, res) => {
  try {
    const { id } = req.params;
    await Product.findByIdAndDelete(id);
    res.redirect('/admin/products');
  } catch (err) {
    res.status(500).send("Error al eliminar producto: " + err.message);
  }
};
