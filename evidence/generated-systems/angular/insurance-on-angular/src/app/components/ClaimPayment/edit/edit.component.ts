import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { ClaimPaymentService } from '../../../services/ClaimPayment.service';
import { SubBaseComponent } from '../../ClaimPayment/sub.base.component';


@Component({
    selector: 'app-edit-claimPayment',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditClaimPaymentComponent extends SubBaseComponent implements OnInit {

    title = 'Edit ClaimPayment';

    claimPaymentForm: FormGroup;
    claimPayment: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: ClaimPaymentService,
        private fb: FormBuilder
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

    
    updateClaimPayment(paymentNumber, amount, paymentDate, Claim, Exposure, Beneficiary, ServiceProvider, Customer, PayeeType, Method, Status): void {
        this.route.params.subscribe((params) => {

                        this.service.updateClaimPayment(paymentNumber, amount, paymentDate, Claim, Exposure, Beneficiary, ServiceProvider, Customer, PayeeType, Method, Status, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexClaimPayment']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getClaimPayment(params['id']).subscribe(res => {
                this.claimPayment = res;
            });
        });
    }
}