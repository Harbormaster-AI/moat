
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { PaymentMethodService } from '../../../services/PaymentMethod.service';
import { PaymentMethod } from '../../../models/PaymentMethod';

@Component({
    selector: 'app-index-paymentMethod',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexPaymentMethodComponent implements OnInit {

    paymentMethods: PaymentMethod[] = [];

    constructor(
        private router: Router,
        private service: PaymentMethodService
) {}

    ngOnInit(): void {
        this.getPaymentMethods();
}

    getPaymentMethods(): void {
        this.service.getPaymentMethods().subscribe((res) => {
        this.paymentMethods = res;
    });
}

    deletePaymentMethod(id: any): void {
        this.service.deletePaymentMethod(id)
            .subscribe(() => {
                this.getPaymentMethods();
            });
    }
}