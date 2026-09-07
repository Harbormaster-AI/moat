import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { OrderService } from '../../../services/Order.service';
import { SubBaseComponent } from '../../Order/sub.base.component';


@Component({
    selector: 'app-edit-order',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditOrderComponent extends SubBaseComponent implements OnInit {

    title = 'Edit Order';

    orderForm: FormGroup;
    order: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: OrderService,
        private fb: FormBuilder
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

    
    updateOrder(orderNumber, orderDate, totalAmount, taxAmount, shippingAmount, Organization, Account, Opportunity, Quote, Owner, Items, Contract, PriceBook, Status): void {
        this.route.params.subscribe((params) => {

                        this.service.updateOrder(orderNumber, orderDate, totalAmount, taxAmount, shippingAmount, Organization, Account, Opportunity, Quote, Owner, Items, Contract, PriceBook, Status, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexOrder']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getOrder(params['id']).subscribe(res => {
                this.order = res;
            });
        });
    }
}