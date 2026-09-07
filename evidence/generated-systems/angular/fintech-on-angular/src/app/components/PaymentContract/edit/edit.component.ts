import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { PaymentContractService } from '../../../services/PaymentContract.service';
import { SubBaseComponent } from '../../PaymentContract/sub.base.component';


@Component({
    selector: 'app-edit-paymentContract',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditPaymentContractComponent extends SubBaseComponent implements OnInit {

    title = 'Edit PaymentContract';

    paymentContractForm: FormGroup;
    paymentContract: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: PaymentContractService,
        private fb: FormBuilder
) {
        super(http);
        this.paymentContractForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  contractNumber: ['', Validators.required],
      pricingPlanCode: ['', Validators.required],
      Merchant: ['', ],
      Acquirer: ['', ],
      Status: ['', ]
        });
    }

    
    updatePaymentContract(contractNumber, pricingPlanCode, Merchant, Acquirer, Status): void {
        this.route.params.subscribe((params) => {

                        this.service.updatePaymentContract(contractNumber, pricingPlanCode, Merchant, Acquirer, Status, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexPaymentContract']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getPaymentContract(params['id']).subscribe(res => {
                this.paymentContract = res;
            });
        });
    }
}