
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { PaymentContractService } from '../../../services/PaymentContract.service';
import { PaymentContract } from '../../../models/PaymentContract';

@Component({
    selector: 'app-index-paymentContract',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexPaymentContractComponent implements OnInit {

    paymentContracts: PaymentContract[] = [];

    constructor(
        private router: Router,
        private service: PaymentContractService
) {}

    ngOnInit(): void {
        this.getPaymentContracts();
}

    getPaymentContracts(): void {
        this.service.getPaymentContracts().subscribe((res) => {
        this.paymentContracts = res;
    });
}

    deletePaymentContract(id: any): void {
        this.service.deletePaymentContract(id)
            .subscribe(() => {
                this.getPaymentContracts();
            });
    }
}