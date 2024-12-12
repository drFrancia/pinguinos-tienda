import Order from "../models/Order.js";

export const getOrdersPage = async (req, res) => {
  try {
    const orders = await Order.find().sort({ date: -1 });
    res.render("orders", { orders });
  } catch (err) {
    res.status(500).send("Error cargando los pedidos: " + err.message);
  }
};

export const getDashboard = (req, res) => {
  res.render('dashboard', { user: req.user });
};
