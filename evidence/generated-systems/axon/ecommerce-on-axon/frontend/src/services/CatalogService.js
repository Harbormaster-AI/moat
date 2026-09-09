import axios from 'axios';

const CATALOG_API_BASE_URL = "/Catalog";

class CatalogService {

    getCatalogs(){
        return axios.get(CATALOG_API_BASE_URL + '/' );
    }

    createCatalog(catalog){
        return axios.post(CATALOG_API_BASE_URL  + '/create', catalog);
    }

    getCatalogById(catalogId){
        return axios.get(CATALOG_API_BASE_URL + '/load?catalogId=' + catalogId);
    }

    updateCatalog(catalog){
        return axios.put(CATALOG_API_BASE_URL + '/update', catalog);
    }

    deleteCatalog(catalogId){
        return axios.delete(CATALOG_API_BASE_URL + '/delete?catalogId=' + catalogId);
    }
}

export default new CatalogService()