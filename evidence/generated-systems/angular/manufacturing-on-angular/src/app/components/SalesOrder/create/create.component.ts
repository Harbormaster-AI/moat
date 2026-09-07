import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { SalesOrderService } from '../../../services/SalesOrder.service';
import { SalesOrder } from '../../../models/SalesOrder';
import { SubBaseComponent } from '../../SalesOrder/sub.base.component';

@Component({
    selector: 'app-create-salesOrder',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateSalesOrderComponent extends SubBaseComponent implements OnInit {

    title = 'Add SalesOrder';

    salesOrderForm: FormGroup;
    salesOrder: SalesOrder;

    constructor( http: HttpClient,
        private salesOrderService: SalesOrderService,
        private fb: FormBuilder,
        private router: Router
) {
        super(http);
        this.salesOrderForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  orderNumber: ['', Validators.required],
      orderDate: ['', Validators.required],
      totalAmount: ['', Validators.required],
      Customer: ['', ],
      Plant: ['', ],
      Lines: ['', ],
      WorkOrders: ['', ],
      Status: ['', ]
        });
    }

    
    addSalesOrder(orderNumber, orderDate, totalAmount, Customer, Plant, Lines, WorkOrders, Status): void {
        this.salesOrderService
        .addSalesOrder(orderNumber, orderDate, totalAmount, Customer, Plant, Lines, WorkOrders, Status)
            .subscribe(() => {
                this.router.navigate(['/indexSalesOrder']);
            });
    }

    ngOnInit(): void {
    }
}