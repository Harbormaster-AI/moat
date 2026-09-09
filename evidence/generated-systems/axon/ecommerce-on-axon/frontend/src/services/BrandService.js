import axios from 'axios';

const BRAND_API_BASE_URL = "/Brand";

class BrandService {

    getBrands(){
        return axios.get(BRAND_API_BASE_URL + '/' );
    }

    createBrand(brand){
        return axios.post(BRAND_API_BASE_URL  + '/create', brand);
    }

    getBrandById(brandId){
        return axios.get(BRAND_API_BASE_URL + '/load?brandId=' + brandId);
    }

    updateBrand(brand){
        return axios.put(BRAND_API_BASE_URL + '/update', brand);
    }

    deleteBrand(brandId){
        return axios.delete(BRAND_API_BASE_URL + '/delete?brandId=' + brandId);
    }
}

export default new BrandService()