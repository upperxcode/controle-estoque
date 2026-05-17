import api from "./api";

export default {
  register(movement) {
    return api.post("/movements", movement);
  },
  getByProduct(productId) {
    return api.get(`/movements?product_id=${productId}`);
  },
  getAll() {
    return api.get("/movements");
  },
};
