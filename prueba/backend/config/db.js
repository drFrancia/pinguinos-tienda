import mongoose from "mongoose";

const connectDB = async () => {
    try {
        await mongoose.connect(process.env.MONGO_URI, {
            useNewUrlParser: true,
            useUnifiedTopology: true,
        });
        console.log("MongoDB conectado");
    } catch (e) {
        console.error("Error al intentar conectarse a MongoDB", e.message);
        process.exit(1);
    }
};

export default connectDB;
