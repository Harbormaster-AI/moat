import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { OrderItemService } from '../../../services/OrderItem.service';
import { OrderItem } from '../../../models/OrderItem';
import { SubBaseComponent } from '../../OrderItem/sub.base.component';

@Component({
    selector: 'app-create-orderItem',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateOrderItemComponent extends SubBaseComponent implements OnInit {

    title = 'Add OrderItem';

    orderItemForm: FormGroup;
    orderItem: OrderItem;

    constructor( http: HttpClient,
        private orderItemService: OrderItemService,
        private fb: FormBuilder,
        private router: Router
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

    
    addOrderItem(quantity, unitPrice, discountAmount, taxAmount, totalAmount, Order, Product, PriceBookEntry): void {
        this.orderItemService
        .addOrderItem(quantity, unitPrice, discountAmount, taxAmount, totalAmount, Order, Product, PriceBookEntry)
            .subscribe(() => {
                this.router.navigate(['/indexOrderItem']);
            });
    }

    ngOnInit(): void {
    }
}