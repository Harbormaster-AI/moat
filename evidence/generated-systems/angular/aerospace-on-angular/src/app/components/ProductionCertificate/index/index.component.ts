
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { ProductionCertificateService } from '../../../services/ProductionCertificate.service';
import { ProductionCertificate } from '../../../models/ProductionCertificate';

@Component({
    selector: 'app-index-productionCertificate',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexProductionCertificateComponent implements OnInit {

    productionCertificates: ProductionCertificate[] = [];

    constructor(
        private router: Router,
        private service: ProductionCertificateService
) {}

    ngOnInit(): void {
        this.getProductionCertificates();
}

    getProductionCertificates(): void {
        this.service.getProductionCertificates().subscribe((res) => {
        this.productionCertificates = res;
    });
}

    deleteProductionCertificate(id: any): void {
        this.service.deleteProductionCertificate(id)
            .subscribe(() => {
                this.getProductionCertificates();
            });
    }
}