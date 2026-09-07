
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { PaymentService } from '../../../services/Payment.service';
import { Payment } from '../../../models/Payment';

@Component({
    selector: 'app-index-payment',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexPaymentComponent implements OnInit {

    payments: Payment[] = [];

    constructor(
        private router: Router,
        private service: PaymentService
) {}

    ngOnInit(): void {
        this.getPayments();
}

    getPayments(): void {
        this.service.getPayments().subscribe((res) => {
        this.payments = res;
    });
}

    deletePayment(id: any): void {
        this.service.deletePayment(id)
            .subscribe(() => {
                this.getPayments();
            });
    }
}