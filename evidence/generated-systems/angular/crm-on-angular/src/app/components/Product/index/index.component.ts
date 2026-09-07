
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { ProductService } from '../../../services/Product.service';
import { Product } from '../../../models/Product';

@Component({
    selector: 'app-index-product',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexProductComponent implements OnInit {

    products: Product[] = [];

    constructor(
        private router: Router,
        private service: ProductService
) {}

    ngOnInit(): void {
        this.getProducts();
}

    getProducts(): void {
        this.service.getProducts().subscribe((res) => {
        this.products = res;
    });
}

    deleteProduct(id: any): void {
        this.service.deleteProduct(id)
            .subscribe(() => {
                this.getProducts();
            });
    }
}