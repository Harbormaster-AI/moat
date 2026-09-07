
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { PaymentProcessorService } from '../../../services/PaymentProcessor.service';
import { PaymentProcessor } from '../../../models/PaymentProcessor';

@Component({
    selector: 'app-index-paymentProcessor',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexPaymentProcessorComponent implements OnInit {

    paymentProcessors: PaymentProcessor[] = [];

    constructor(
        private router: Router,
        private service: PaymentProcessorService
) {}

    ngOnInit(): void {
        this.getPaymentProcessors();
}

    getPaymentProcessors(): void {
        this.service.getPaymentProcessors().subscribe((res) => {
        this.paymentProcessors = res;
    });
}

    deletePaymentProcessor(id: any): void {
        this.service.deletePaymentProcessor(id)
            .subscribe(() => {
                this.getPaymentProcessors();
            });
    }
}