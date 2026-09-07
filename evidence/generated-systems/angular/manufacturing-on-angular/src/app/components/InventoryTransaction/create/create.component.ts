import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { InventoryTransactionService } from '../../../services/InventoryTransaction.service';
import { InventoryTransaction } from '../../../models/InventoryTransaction';
import { SubBaseComponent } from '../../InventoryTransaction/sub.base.component';

@Component({
    selector: 'app-create-inventoryTransaction',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateInventoryTransactionComponent extends SubBaseComponent implements OnInit {

    title = 'Add InventoryTransaction';

    inventoryTransactionForm: FormGroup;
    inventoryTransaction: InventoryTransaction;

    constructor( http: HttpClient,
        private inventoryTransactionService: InventoryTransactionService,
        private fb: FormBuilder,
        private router: Router
) {
        super(http);
        this.inventoryTransactionForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  transactionNumber: ['', Validators.required],
      quantity: ['', Validators.required],
      transactionDateTime: ['', Validators.required],
      referenceDocument: ['', Validators.required],
      Item: ['', ],
      Location: ['', ],
      WorkOrder: ['', ],
      PurchaseOrder: ['', ],
      SalesOrder: ['', ],
      TransactionType: ['', ]
        });
    }

    
    addInventoryTransaction(transactionNumber, quantity, transactionDateTime, referenceDocument, Item, Location, WorkOrder, PurchaseOrder, SalesOrder, TransactionType): void {
        this.inventoryTransactionService
        .addInventoryTransaction(transactionNumber, quantity, transactionDateTime, referenceDocument, Item, Location, WorkOrder, PurchaseOrder, SalesOrder, TransactionType)
            .subscribe(() => {
                this.router.navigate(['/indexInventoryTransaction']);
            });
    }

    ngOnInit(): void {
    }
}