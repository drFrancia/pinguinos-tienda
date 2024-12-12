import jwt from 'jsonwebtoken';

export const authenticate = (req, res, next) => {
  const token = req.cookies.token;

  if (!token) {
    return res.redirect('/admin/login');
  }

  try {
    const decoded = jwt.verify(token, process.env.JWT_SECRET);
    req.user = decoded;
    next();
  } catch (err) {
    res.redirect('/admin/login');
  }
};
