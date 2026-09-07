import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { BillingProfileService } from '../../../services/BillingProfile.service';
import { SubBaseComponent } from '../../BillingProfile/sub.base.component';


@Component({
    selector: 'app-edit-billingProfile',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditBillingProfileComponent extends SubBaseComponent implements OnInit {

    title = 'Edit BillingProfile';

    billingProfileForm: FormGroup;
    billingProfile: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: BillingProfileService,
        private fb: FormBuilder
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

    
    updateBillingProfile(billingName, taxId, billingAddress, Advertiser, PaymentMethods, AdAccounts, PaymentTerms): void {
        this.route.params.subscribe((params) => {

                        this.service.updateBillingProfile(billingName, taxId, billingAddress, Advertiser, PaymentMethods, AdAccounts, PaymentTerms, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexBillingProfile']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getBillingProfile(params['id']).subscribe(res => {
                this.billingProfile = res;
            });
        });
    }
}