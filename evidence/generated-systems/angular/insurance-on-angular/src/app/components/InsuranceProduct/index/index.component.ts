
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { InsuranceProductService } from '../../../services/InsuranceProduct.service';
import { InsuranceProduct } from '../../../models/InsuranceProduct';

@Component({
    selector: 'app-index-insuranceProduct',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexInsuranceProductComponent implements OnInit {

    insuranceProducts: InsuranceProduct[] = [];

    constructor(
        private router: Router,
        private service: InsuranceProductService
) {}

    ngOnInit(): void {
        this.getInsuranceProducts();
}

    getInsuranceProducts(): void {
        this.service.getInsuranceProducts().subscribe((res) => {
        this.insuranceProducts = res;
    });
}

    deleteInsuranceProduct(id: any): void {
        this.service.deleteInsuranceProduct(id)
            .subscribe(() => {
                this.getInsuranceProducts();
            });
    }
}