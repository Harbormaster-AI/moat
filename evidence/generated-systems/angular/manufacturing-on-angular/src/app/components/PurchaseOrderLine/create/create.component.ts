import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { PurchaseOrderLineService } from '../../../services/PurchaseOrderLine.service';
import { PurchaseOrderLine } from '../../../models/PurchaseOrderLine';
import { SubBaseComponent } from '../../PurchaseOrderLine/sub.base.component';

@Component({
    selector: 'app-create-purchaseOrderLine',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreatePurchaseOrderLineComponent extends SubBaseComponent implements OnInit {

    title = 'Add PurchaseOrderLine';

    purchaseOrderLineForm: FormGroup;
    purchaseOrderLine: PurchaseOrderLine;

    constructor( http: HttpClient,
        private purchaseOrderLineService: PurchaseOrderLineService,
        private fb: FormBuilder,
        private router: Router
) {
        super(http);
        this.purchaseOrderLineForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  lineNumber: ['', Validators.required],
      quantity: ['', Validators.required],
      unitPrice: ['', Validators.required],
      dueDate: ['', Validators.required],
      PurchaseOrder: ['', ],
      Item: ['', ]
        });
    }

    
    addPurchaseOrderLine(lineNumber, quantity, unitPrice, dueDate, PurchaseOrder, Item): void {
        this.purchaseOrderLineService
        .addPurchaseOrderLine(lineNumber, quantity, unitPrice, dueDate, PurchaseOrder, Item)
            .subscribe(() => {
                this.router.navigate(['/indexPurchaseOrderLine']);
            });
    }

    ngOnInit(): void {
    }
}