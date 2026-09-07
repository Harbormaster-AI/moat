
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { PurchaseAgreementService } from '../../../services/PurchaseAgreement.service';
import { PurchaseAgreement } from '../../../models/PurchaseAgreement';

@Component({
    selector: 'app-index-purchaseAgreement',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexPurchaseAgreementComponent implements OnInit {

    purchaseAgreements: PurchaseAgreement[] = [];

    constructor(
        private router: Router,
        private service: PurchaseAgreementService
) {}

    ngOnInit(): void {
        this.getPurchaseAgreements();
}

    getPurchaseAgreements(): void {
        this.service.getPurchaseAgreements().subscribe((res) => {
        this.purchaseAgreements = res;
    });
}

    deletePurchaseAgreement(id: any): void {
        this.service.deletePurchaseAgreement(id)
            .subscribe(() => {
                this.getPurchaseAgreements();
            });
    }
}