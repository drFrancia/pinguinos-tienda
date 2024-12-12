import Product from "../models/Product.js";

export const getAllProducts = async (req, res) => {
    try {
        const products = await Product.find();
        res.status(200).json(products); // Devuelve la lista de productos
    } catch (e) {
        res.status(500).json({ message: "Error al obtener productos", error: e });
    }
};

export const createProduct = async (req, res) => {
    try {
        const { name, price, description, stock } = req.body;
        const newProduct = new Product({ name, price, description, stock });
        await newProduct.save();
        res.status(201).json(newProduct); // Devuelve el producto creado
    } catch (e) {
        res.status(500).json({ message: "Error al crear producto", error: e });
    }
};
