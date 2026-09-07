import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { OrderService } from '../../../services/Order.service';
import { Order } from '../../../models/Order';
import { SubBaseComponent } from '../../Order/sub.base.component';

@Component({
    selector: 'app-create-order',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateOrderComponent extends SubBaseComponent implements OnInit {

    title = 'Add Order';

    orderForm: FormGroup;
    order: Order;

    constructor( http: HttpClient,
        private orderService: OrderService,
        private fb: FormBuilder,
        private router: Router
) {
        super(http);
        this.orderForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  orderNumber: ['', Validators.required],
      orderDate: ['', Validators.required],
      totalAmount: ['', Validators.required],
      taxAmount: ['', Validators.required],
      shippingAmount: ['', Validators.required],
      Organization: ['', ],
      Account: ['', ],
      Opportunity: ['', ],
      Quote: ['', ],
      Owner: ['', ],
      Items: ['', ],
      Contract: ['', ],
      PriceBook: ['', ],
      Status: ['', ]
        });
    }

    
    addOrder(orderNumber, orderDate, totalAmount, taxAmount, shippingAmount, Organization, Account, Opportunity, Quote, Owner, Items, Contract, PriceBook, Status): void {
        this.orderService
        .addOrder(orderNumber, orderDate, totalAmount, taxAmount, shippingAmount, Organization, Account, Opportunity, Quote, Owner, Items, Contract, PriceBook, Status)
            .subscribe(() => {
                this.router.navigate(['/indexOrder']);
            });
    }

    ngOnInit(): void {
    }
}