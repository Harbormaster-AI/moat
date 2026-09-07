import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { OrderItemService } from '../../../services/OrderItem.service';
import { SubBaseComponent } from '../../OrderItem/sub.base.component';


@Component({
    selector: 'app-edit-orderItem',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditOrderItemComponent extends SubBaseComponent implements OnInit {

    title = 'Edit OrderItem';

    orderItemForm: FormGroup;
    orderItem: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: OrderItemService,
        private fb: FormBuilder
) {
        super(http);
        this.orderItemForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  quantity: ['', Validators.required],
      unitPrice: ['', Validators.required],
      discountAmount: ['', Validators.required],
      taxAmount: ['', Validators.required],
      totalAmount: ['', Validators.required],
      Order: ['', ],
      Product: ['', ],
      PriceBookEntry: ['', ]
        });
    }

    
    updateOrderItem(quantity, unitPrice, discountAmount, taxAmount, totalAmount, Order, Product, PriceBookEntry): void {
        this.route.params.subscribe((params) => {

                        this.service.updateOrderItem(quantity, unitPrice, discountAmount, taxAmount, totalAmount, Order, Product, PriceBookEntry, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexOrderItem']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getOrderItem(params['id']).subscribe(res => {
                this.orderItem = res;
            });
        });
    }
}