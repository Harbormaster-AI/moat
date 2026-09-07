
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { ProductOfferingService } from '../../../services/ProductOffering.service';
import { ProductOffering } from '../../../models/ProductOffering';

@Component({
    selector: 'app-index-productOffering',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexProductOfferingComponent implements OnInit {

    productOfferings: ProductOffering[] = [];

    constructor(
        private router: Router,
        private service: ProductOfferingService
) {}

    ngOnInit(): void {
        this.getProductOfferings();
}

    getProductOfferings(): void {
        this.service.getProductOfferings().subscribe((res) => {
        this.productOfferings = res;
    });
}

    deleteProductOffering(id: any): void {
        this.service.deleteProductOffering(id)
            .subscribe(() => {
                this.getProductOfferings();
            });
    }
}