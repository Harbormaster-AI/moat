import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { PolicyService } from '../../../services/Policy.service';
import { Policy } from '../../../models/Policy';
import { SubBaseComponent } from '../../Policy/sub.base.component';

@Component({
    selector: 'app-create-policy',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreatePolicyComponent extends SubBaseComponent implements OnInit {

    title = 'Add Policy';

    policyForm: FormGroup;
    policy: Policy;

    constructor( http: HttpClient,
        private policyService: PolicyService,
        private fb: FormBuilder,
        private router: Router
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

    
    addPolicy(policyNumber, effectivePeriod, totalPremium, Insurer, Customer, Product, Agent, Coverages, InsuredObjects, Endorsements, BillingAccount, Beneficiaries, Claims, ReinsuranceAgreements, Status, PaymentPlan): void {
        this.policyService
        .addPolicy(policyNumber, effectivePeriod, totalPremium, Insurer, Customer, Product, Agent, Coverages, InsuredObjects, Endorsements, BillingAccount, Beneficiaries, Claims, ReinsuranceAgreements, Status, PaymentPlan)
            .subscribe(() => {
                this.router.navigate(['/indexPolicy']);
            });
    }

    ngOnInit(): void {
    }
}