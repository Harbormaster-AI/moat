import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { PaymentOrderService } from '../../../services/PaymentOrder.service';
import { PaymentOrder } from '../../../models/PaymentOrder';
import { SubBaseComponent } from '../../PaymentOrder/sub.base.component';

@Component({
    selector: 'app-create-paymentOrder',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreatePaymentOrderComponent extends SubBaseComponent implements OnInit {

    title = 'Add PaymentOrder';

    paymentOrderForm: FormGroup;
    paymentOrder: PaymentOrder;

    constructor( http: HttpClient,
        private paymentOrderService: PaymentOrderService,
        private fb: FormBuilder,
        private router: Router
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

    
    addPaymentOrder(orderReference, requestedExecutionDate, purpose, SourceAccount, DestinationAccount, Beneficiary, Transactions, FxDeal, Fees, PaymentMethod, Status, Priority): void {
        this.paymentOrderService
        .addPaymentOrder(orderReference, requestedExecutionDate, purpose, SourceAccount, DestinationAccount, Beneficiary, Transactions, FxDeal, Fees, PaymentMethod, Status, Priority)
            .subscribe(() => {
                this.router.navigate(['/indexPaymentOrder']);
            });
    }

    ngOnInit(): void {
    }
}