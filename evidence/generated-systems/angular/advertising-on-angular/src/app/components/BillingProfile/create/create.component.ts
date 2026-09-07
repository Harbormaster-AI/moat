import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { BillingProfileService } from '../../../services/BillingProfile.service';
import { BillingProfile } from '../../../models/BillingProfile';
import { SubBaseComponent } from '../../BillingProfile/sub.base.component';

@Component({
    selector: 'app-create-billingProfile',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateBillingProfileComponent extends SubBaseComponent implements OnInit {

    title = 'Add BillingProfile';

    billingProfileForm: FormGroup;
    billingProfile: BillingProfile;

    constructor( http: HttpClient,
        private billingProfileService: BillingProfileService,
        private fb: FormBuilder,
        private router: Router
) {
        super(http);
        this.billingProfileForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  billingName: ['', Validators.required],
      taxId: ['', Validators.required],
      billingAddress: ['', Validators.required],
      Advertiser: ['', ],
      PaymentMethods: ['', ],
      AdAccounts: ['', ],
      PaymentTerms: ['', ]
        });
    }

    
    addBillingProfile(billingName, taxId, billingAddress, Advertiser, PaymentMethods, AdAccounts, PaymentTerms): void {
        this.billingProfileService
        .addBillingProfile(billingName, taxId, billingAddress, Advertiser, PaymentMethods, AdAccounts, PaymentTerms)
            .subscribe(() => {
                this.router.navigate(['/indexBillingProfile']);
            });
    }

    ngOnInit(): void {
    }
}