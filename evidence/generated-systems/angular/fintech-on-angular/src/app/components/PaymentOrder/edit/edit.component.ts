import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { PaymentOrderService } from '../../../services/PaymentOrder.service';
import { SubBaseComponent } from '../../PaymentOrder/sub.base.component';


@Component({
    selector: 'app-edit-paymentOrder',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditPaymentOrderComponent extends SubBaseComponent implements OnInit {

    title = 'Edit PaymentOrder';

    paymentOrderForm: FormGroup;
    paymentOrder: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: PaymentOrderService,
        private fb: FormBuilder
) {
        super(http);
        this.paymentOrderForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  orderReference: ['', Validators.required],
      requestedExecutionDate: ['', Validators.required],
      purpose: ['', Validators.required],
      SourceAccount: ['', ],
      DestinationAccount: ['', ],
      Beneficiary: ['', ],
      Transactions: ['', ],
      FxDeal: ['', ],
      Fees: ['', ],
      PaymentMethod: ['', ],
      Status: ['', ],
      Priority: ['', ]
        });
    }

    
    updatePaymentOrder(orderReference, requestedExecutionDate, purpose, SourceAccount, DestinationAccount, Beneficiary, Transactions, FxDeal, Fees, PaymentMethod, Status, Priority): void {
        this.route.params.subscribe((params) => {

                        this.service.updatePaymentOrder(orderReference, requestedExecutionDate, purpose, SourceAccount, DestinationAccount, Beneficiary, Transactions, FxDeal, Fees, PaymentMethod, Status, Priority, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexPaymentOrder']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getPaymentOrder(params['id']).subscribe(res => {
                this.paymentOrder = res;
            });
        });
    }
}