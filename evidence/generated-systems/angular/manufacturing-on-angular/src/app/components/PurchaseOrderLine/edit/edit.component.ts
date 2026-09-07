import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { PurchaseOrderLineService } from '../../../services/PurchaseOrderLine.service';
import { SubBaseComponent } from '../../PurchaseOrderLine/sub.base.component';


@Component({
    selector: 'app-edit-purchaseOrderLine',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditPurchaseOrderLineComponent extends SubBaseComponent implements OnInit {

    title = 'Edit PurchaseOrderLine';

    purchaseOrderLineForm: FormGroup;
    purchaseOrderLine: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: PurchaseOrderLineService,
        private fb: FormBuilder
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

    
    updatePurchaseOrderLine(lineNumber, quantity, unitPrice, dueDate, PurchaseOrder, Item): void {
        this.route.params.subscribe((params) => {

                        this.service.updatePurchaseOrderLine(lineNumber, quantity, unitPrice, dueDate, PurchaseOrder, Item, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexPurchaseOrderLine']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getPurchaseOrderLine(params['id']).subscribe(res => {
                this.purchaseOrderLine = res;
            });
        });
    }
}