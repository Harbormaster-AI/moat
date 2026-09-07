import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { InventoryTransactionService } from '../../../services/InventoryTransaction.service';
import { SubBaseComponent } from '../../InventoryTransaction/sub.base.component';


@Component({
    selector: 'app-edit-inventoryTransaction',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditInventoryTransactionComponent extends SubBaseComponent implements OnInit {

    title = 'Edit InventoryTransaction';

    inventoryTransactionForm: FormGroup;
    inventoryTransaction: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: InventoryTransactionService,
        private fb: FormBuilder
) {
        super(http);
        this.inventoryTransactionForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  transactionNumber: ['', Validators.required],
      quantity: ['', Validators.required],
      unitCost: ['', Validators.required],
      transactionDate: ['', Validators.required],
      reasonCode: ['', Validators.required],
      Sku: ['', ],
      Warehouse: ['', ],
      Location: ['', ],
      Lot: ['', ],
      SerialNumbers: ['', ],
      RelatedReservation: ['', ],
      TransferOrder: ['', ],
      Adjustment: ['', ],
      CycleCount: ['', ],
      TransactionType: ['', ],
      UnitOfMeasure: ['', ],
      Status: ['', ]
        });
    }

    
    updateInventoryTransaction(transactionNumber, quantity, unitCost, transactionDate, reasonCode, Sku, Warehouse, Location, Lot, SerialNumbers, RelatedReservation, TransferOrder, Adjustment, CycleCount, TransactionType, UnitOfMeasure, Status): void {
        this.route.params.subscribe((params) => {

                        this.service.updateInventoryTransaction(transactionNumber, quantity, unitCost, transactionDate, reasonCode, Sku, Warehouse, Location, Lot, SerialNumbers, RelatedReservation, TransferOrder, Adjustment, CycleCount, TransactionType, UnitOfMeasure, Status, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexInventoryTransaction']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getInventoryTransaction(params['id']).subscribe(res => {
                this.inventoryTransaction = res;
            });
        });
    }
}