
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { BillingAccountService } from '../../../services/BillingAccount.service';
import { BillingAccount } from '../../../models/BillingAccount';

@Component({
    selector: 'app-index-billingAccount',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexBillingAccountComponent implements OnInit {

    billingAccounts: BillingAccount[] = [];

    constructor(
        private router: Router,
        private service: BillingAccountService
) {}

    ngOnInit(): void {
        this.getBillingAccounts();
}

    getBillingAccounts(): void {
        this.service.getBillingAccounts().subscribe((res) => {
        this.billingAccounts = res;
    });
}

    deleteBillingAccount(id: any): void {
        this.service.deleteBillingAccount(id)
            .subscribe(() => {
                this.getBillingAccounts();
            });
    }
}