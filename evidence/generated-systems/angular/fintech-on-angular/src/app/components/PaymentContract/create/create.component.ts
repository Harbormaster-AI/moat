import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { PaymentContractService } from '../../../services/PaymentContract.service';
import { PaymentContract } from '../../../models/PaymentContract';
import { SubBaseComponent } from '../../PaymentContract/sub.base.component';

@Component({
    selector: 'app-create-paymentContract',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreatePaymentContractComponent extends SubBaseComponent implements OnInit {

    title = 'Add PaymentContract';

    paymentContractForm: FormGroup;
    paymentContract: PaymentContract;

    constructor( http: HttpClient,
        private paymentContractService: PaymentContractService,
        private fb: FormBuilder,
        private router: Router
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

    
    addPaymentContract(contractNumber, pricingPlanCode, Merchant, Acquirer, Status): void {
        this.paymentContractService
        .addPaymentContract(contractNumber, pricingPlanCode, Merchant, Acquirer, Status)
            .subscribe(() => {
                this.router.navigate(['/indexPaymentContract']);
            });
    }

    ngOnInit(): void {
    }
}