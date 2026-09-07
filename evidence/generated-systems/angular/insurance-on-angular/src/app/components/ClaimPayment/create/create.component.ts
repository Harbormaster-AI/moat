import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { ClaimPaymentService } from '../../../services/ClaimPayment.service';
import { ClaimPayment } from '../../../models/ClaimPayment';
import { SubBaseComponent } from '../../ClaimPayment/sub.base.component';

@Component({
    selector: 'app-create-claimPayment',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateClaimPaymentComponent extends SubBaseComponent implements OnInit {

    title = 'Add ClaimPayment';

    claimPaymentForm: FormGroup;
    claimPayment: ClaimPayment;

    constructor( http: HttpClient,
        private claimPaymentService: ClaimPaymentService,
        private fb: FormBuilder,
        private router: Router
) {
        super(http);
        this.claimPaymentForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  paymentNumber: ['', Validators.required],
      amount: ['', Validators.required],
      paymentDate: ['', Validators.required],
      Claim: ['', ],
      Exposure: ['', ],
      Beneficiary: ['', ],
      ServiceProvider: ['', ],
      Customer: ['', ],
      PayeeType: ['', ],
      Method: ['', ],
      Status: ['', ]
        });
    }

    
    addClaimPayment(paymentNumber, amount, paymentDate, Claim, Exposure, Beneficiary, ServiceProvider, Customer, PayeeType, Method, Status): void {
        this.claimPaymentService
        .addClaimPayment(paymentNumber, amount, paymentDate, Claim, Exposure, Beneficiary, ServiceProvider, Customer, PayeeType, Method, Status)
            .subscribe(() => {
                this.router.navigate(['/indexClaimPayment']);
            });
    }

    ngOnInit(): void {
    }
}