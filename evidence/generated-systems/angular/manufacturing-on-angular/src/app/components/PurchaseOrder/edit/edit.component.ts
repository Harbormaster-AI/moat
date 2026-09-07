import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { PurchaseOrderService } from '../../../services/PurchaseOrder.service';
import { SubBaseComponent } from '../../PurchaseOrder/sub.base.component';


@Component({
    selector: 'app-edit-purchaseOrder',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditPurchaseOrderComponent extends SubBaseComponent implements OnInit {

    title = 'Edit PurchaseOrder';

    purchaseOrderForm: FormGroup;
    purchaseOrder: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: PurchaseOrderService,
        private fb: FormBuilder
) {
        super(http);
        this.purchaseOrderForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  poNumber: ['', Validators.required],
      orderDate: ['', Validators.required],
      totalAmount: ['', Validators.required],
      Supplier: ['', ],
      Plant: ['', ],
      Lines: ['', ],
      GoodsReceipts: ['', ],
      Status: ['', ]
        });
    }

    
    updatePurchaseOrder(poNumber, orderDate, totalAmount, Supplier, Plant, Lines, GoodsReceipts, Status): void {
        this.route.params.subscribe((params) => {

                        this.service.updatePurchaseOrder(poNumber, orderDate, totalAmount, Supplier, Plant, Lines, GoodsReceipts, Status, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexPurchaseOrder']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getPurchaseOrder(params['id']).subscribe(res => {
                this.purchaseOrder = res;
            });
        });
    }
}