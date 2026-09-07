
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { PaymentOrderService } from '../../../services/PaymentOrder.service';
import { PaymentOrder } from '../../../models/PaymentOrder';

@Component({
    selector: 'app-index-paymentOrder',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexPaymentOrderComponent implements OnInit {

    paymentOrders: PaymentOrder[] = [];

    constructor(
        private router: Router,
        private service: PaymentOrderService
) {}

    ngOnInit(): void {
        this.getPaymentOrders();
}

    getPaymentOrders(): void {
        this.service.getPaymentOrders().subscribe((res) => {
        this.paymentOrders = res;
    });
}

    deletePaymentOrder(id: any): void {
        this.service.deletePaymentOrder(id)
            .subscribe(() => {
                this.getPaymentOrders();
            });
    }
}