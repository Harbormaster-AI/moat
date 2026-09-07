
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { BillingProfileService } from '../../../services/BillingProfile.service';
import { BillingProfile } from '../../../models/BillingProfile';

@Component({
    selector: 'app-index-billingProfile',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexBillingProfileComponent implements OnInit {

    billingProfiles: BillingProfile[] = [];

    constructor(
        private router: Router,
        private service: BillingProfileService
) {}

    ngOnInit(): void {
        this.getBillingProfiles();
}

    getBillingProfiles(): void {
        this.service.getBillingProfiles().subscribe((res) => {
        this.billingProfiles = res;
    });
}

    deleteBillingProfile(id: any): void {
        this.service.deleteBillingProfile(id)
            .subscribe(() => {
                this.getBillingProfiles();
            });
    }
}