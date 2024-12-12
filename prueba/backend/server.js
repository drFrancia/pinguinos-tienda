import express from "express";
import dotenv from "dotenv";
import connectDB from "./config/db.js";
import productRoutes from "./routes/productRoutes.js"; // Importa las rutas correctamente

dotenv.config();

const app = express();
const PORT = process.env.PORT || 3000;

// Middleware
app.use(express.json());
app.use(express.urlencoded({ extended: true }));

// Configuración de rutas
app.use("/api/products", productRoutes); // Configura las rutas de productos

// Conexión con la base de datos
connectDB();

// Inicia el servidor
app.listen(PORT, () => {
    console.log(`Server running on http://localhost:${PORT}`);
});
