import Order from "../models/Order.js";

// Registrar un nuevo pedido
export const createOrder = async (req, res) => {
  const { name, address, products } = req.body;

  try {
    const newOrder = new Order({ name, address, products });
    await newOrder.save();
    res.status(201).send("Pedido registrado con éxito");
  } catch (err) {
    res.status(500).send("Error registrando el pedido: " + err.message);
  }
};

// Obtener todos los pedidos
export const getOrdersPage = async (req, res) => {
  try {
    const orders = await Order.find().sort({ date: -1 });
    res.render("orders", { orders });
  } catch (err) {
    res.status(500).send("Error obteniendo los pedidos: " + err.message);
  }
};
