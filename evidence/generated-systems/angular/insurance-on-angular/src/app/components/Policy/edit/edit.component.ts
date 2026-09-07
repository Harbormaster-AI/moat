import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { PolicyService } from '../../../services/Policy.service';
import { SubBaseComponent } from '../../Policy/sub.base.component';


@Component({
    selector: 'app-edit-policy',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditPolicyComponent extends SubBaseComponent implements OnInit {

    title = 'Edit Policy';

    policyForm: FormGroup;
    policy: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: PolicyService,
        private fb: FormBuilder
) {
        super(http);
        this.policyForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  policyNumber: ['', Validators.required],
      effectivePeriod: ['', Validators.required],
      totalPremium: ['', Validators.required],
      Insurer: ['', ],
      Customer: ['', ],
      Product: ['', ],
      Agent: ['', ],
      Coverages: ['', ],
      InsuredObjects: ['', ],
      Endorsements: ['', ],
      BillingAccount: ['', ],
      Beneficiaries: ['', ],
      Claims: ['', ],
      ReinsuranceAgreements: ['', ],
      Status: ['', ],
      PaymentPlan: ['', ]
        });
    }

    
    updatePolicy(policyNumber, effectivePeriod, totalPremium, Insurer, Customer, Product, Agent, Coverages, InsuredObjects, Endorsements, BillingAccount, Beneficiaries, Claims, ReinsuranceAgreements, Status, PaymentPlan): void {
        this.route.params.subscribe((params) => {

                        this.service.updatePolicy(policyNumber, effectivePeriod, totalPremium, Insurer, Customer, Product, Agent, Coverages, InsuredObjects, Endorsements, BillingAccount, Beneficiaries, Claims, ReinsuranceAgreements, Status, PaymentPlan, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexPolicy']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getPolicy(params['id']).subscribe(res => {
                this.policy = res;
            });
        });
    }
}