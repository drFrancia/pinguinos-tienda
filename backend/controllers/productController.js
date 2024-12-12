import Product from '../models/Product.js';

export const getAllProducts = async (req, res) => {
  const products = await Product.find();
  res.render('products', { products });
};

export const createProduct = async (req, res) => {
  const { name, price, description } = req.body;
  await Product.create({ name, price, description });
  res.redirect('/admin/products');
};

export const editProduct = async (req, res) => {
  const { id } = req.params;
  const { name, price, description } = req.body;
  await Product.findByIdAndUpdate(id, { name, price, description });
  res.redirect('/admin/products');
};

export const deleteProduct = async (req, res) => {
  const { id } = req.params;
  await Product.findByIdAndDelete(id);
  res.redirect('/admin/products');
};
