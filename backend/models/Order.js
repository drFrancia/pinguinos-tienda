import mongoose from "mongoose";

const orderSchema = new mongoose.Schema({
  name: { type: String, required: true },
  address: { type: String, required: true },
  products: { type: Array, required: true },
  date: { type: Date, default: Date.now },
});

export default mongoose.model("Order", orderSchema);
