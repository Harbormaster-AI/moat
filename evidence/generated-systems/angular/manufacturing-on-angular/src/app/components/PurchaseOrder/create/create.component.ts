import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { PurchaseOrderService } from '../../../services/PurchaseOrder.service';
import { PurchaseOrder } from '../../../models/PurchaseOrder';
import { SubBaseComponent } from '../../PurchaseOrder/sub.base.component';

@Component({
    selector: 'app-create-purchaseOrder',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreatePurchaseOrderComponent extends SubBaseComponent implements OnInit {

    title = 'Add PurchaseOrder';

    purchaseOrderForm: FormGroup;
    purchaseOrder: PurchaseOrder;

    constructor( http: HttpClient,
        private purchaseOrderService: PurchaseOrderService,
        private fb: FormBuilder,
        private router: Router
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

    
    addPurchaseOrder(poNumber, orderDate, totalAmount, Supplier, Plant, Lines, GoodsReceipts, Status): void {
        this.purchaseOrderService
        .addPurchaseOrder(poNumber, orderDate, totalAmount, Supplier, Plant, Lines, GoodsReceipts, Status)
            .subscribe(() => {
                this.router.navigate(['/indexPurchaseOrder']);
            });
    }

    ngOnInit(): void {
    }
}