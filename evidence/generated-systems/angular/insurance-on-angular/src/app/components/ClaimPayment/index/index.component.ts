
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { ClaimPaymentService } from '../../../services/ClaimPayment.service';
import { ClaimPayment } from '../../../models/ClaimPayment';

@Component({
    selector: 'app-index-claimPayment',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexClaimPaymentComponent implements OnInit {

    claimPayments: ClaimPayment[] = [];

    constructor(
        private router: Router,
        private service: ClaimPaymentService
) {}

    ngOnInit(): void {
        this.getClaimPayments();
}

    getClaimPayments(): void {
        this.service.getClaimPayments().subscribe((res) => {
        this.claimPayments = res;
    });
}

    deleteClaimPayment(id: any): void {
        this.service.deleteClaimPayment(id)
            .subscribe(() => {
                this.getClaimPayments();
            });
    }
}