import express from 'express';
import mongoose from 'mongoose';
import path from 'path';
import cookieParser from 'cookie-parser';
import cors from 'cors';
import dotenv from 'dotenv';
import bcrypt from 'bcrypt';

// Importar rutas
import adminRoutes from './routes/admin.js';
import orderRoutes from './routes/order.js';
import User from './models/User.js';

dotenv.config();

const app = express();

// Configuración del motor de vistas
app.set('view engine', 'pug');
app.set('views', path.join(path.resolve(), 'backend/views'));

// Middlewares
app.use(express.json());
app.use(express.urlencoded({ extended: true }));
app.use(cookieParser());
app.use(cors({ origin: 'http://localhost:8080', credentials: true }));

// Redirección a login cuando accede a la raíz
app.get('/', (req, res) => {
  res.redirect('/admin/login');
});

// Rutas
app.use('/admin', adminRoutes);
app.use('/admin/orders', orderRoutes);

// Conexión a MongoDB
mongoose.connect(process.env.MONGO_URI)
  .then(() => {
    console.log('Conexión exitosa a MongoDB');
    initializeSuperuser();
  })
  .catch(err => console.error('Error conectando a MongoDB:', err));

// Crear superusuario Paula
const initializeSuperuser = async () => {
  try {
    const hashedPassword = await bcrypt.hash('123456789', 10); // Hashear contraseña
    const existingUser = await User.findOne({ username: 'Paula' });

    if (!existingUser) {
      // Crear superusuario si no existe
      const superuser = new User({
        username: 'Paula',
        password: hashedPassword,
        role: 'admin',
      });
      await superuser.save();
      console.log('Superusuario Paula creado con éxito.');
    } else {
      // Actualizar contraseña si el usuario ya existe
      existingUser.password = hashedPassword;
      await existingUser.save();
      console.log('Contraseña del superusuario Paula actualizada.');
    }
  } catch (err) {
    console.error('Error al inicializar el superusuario Paula:', err.message);
  }
};


// Iniciar servidor
const PORT = process.env.PORT || 3000;
app.listen(PORT, () => {
  console.log(`Servidor corriendo en http://localhost:${PORT}`);
});
