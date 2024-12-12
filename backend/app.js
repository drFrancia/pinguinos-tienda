import express from 'express';
import mongoose from 'mongoose';
import path from 'path';
import cookieParser from 'cookie-parser';
import cors from 'cors'; // Importar CORS
import dotenv from 'dotenv';
import adminRoutes from './routes/admin.js';
import orderRoutes from './routes/order.js';
import User from './models/User.js';

dotenv.config();

const app = express();

// Configuración básica
app.set('view engine', 'pug');
app.set('views', path.join(path.resolve(), 'backend/views'));

app.use(express.json());
app.use(express.urlencoded({ extended: true }));
app.use(cookieParser());

// Configuración de CORS
app.use(cors({
  origin: 'http://localhost:8080', // Frontend URL
  methods: ['GET', 'POST'],
  credentials: true,
}));

// Rutas
app.use('/admin', adminRoutes);
app.use('/api/orders', orderRoutes);

// Conexión a MongoDB
mongoose.connect(process.env.MONGO_URI, { useNewUrlParser: true, useUnifiedTopology: true })
  .then(() => {
    console.log('Conexión exitosa a MongoDB');
    initializeSuperuser();
  })
  .catch((err) => {
    console.error('Error conectando a MongoDB:', err.message);
  });

// Función para crear el superusuario Paula
const initializeSuperuser = async () => {
  try {
    const existingUser = await User.findOne({ username: 'Paula' });
    if (!existingUser) {
      const superuser = new User({
        username: 'Paula',
        password: '123456789', // Cambiar contraseña en producción
        role: 'admin',
      });
      await superuser.save();
      console.log('Superusuario Paula creado con éxito.');
    } else {
      console.log('El superusuario Paula ya existe.');
    }
  } catch (err) {
    console.error('Error al crear el superusuario Paula:', err);
  }
};

// Iniciar el servidor
const PORT = process.env.PORT || 3000;
app.listen(PORT, () => {
  console.log(`Servidor corriendo en http://localhost:${PORT}`);
});
