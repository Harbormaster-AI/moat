import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { SalesOrderService } from '../../../services/SalesOrder.service';
import { SubBaseComponent } from '../../SalesOrder/sub.base.component';


@Component({
    selector: 'app-edit-salesOrder',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditSalesOrderComponent extends SubBaseComponent implements OnInit {

    title = 'Edit SalesOrder';

    salesOrderForm: FormGroup;
    salesOrder: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: SalesOrderService,
        private fb: FormBuilder
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

    
    updateSalesOrder(orderNumber, orderDate, totalAmount, Customer, Plant, Lines, WorkOrders, Status): void {
        this.route.params.subscribe((params) => {

                        this.service.updateSalesOrder(orderNumber, orderDate, totalAmount, Customer, Plant, Lines, WorkOrders, Status, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexSalesOrder']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getSalesOrder(params['id']).subscribe(res => {
                this.salesOrder = res;
            });
        });
    }
}