import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { InventoryItemService } from '../../../services/InventoryItem.service';
import { SubBaseComponent } from '../../InventoryItem/sub.base.component';


@Component({
    selector: 'app-edit-inventoryItem',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditInventoryItemComponent extends SubBaseComponent implements OnInit {

    title = 'Edit InventoryItem';

    inventoryItemForm: FormGroup;
    inventoryItem: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: InventoryItemService,
        private fb: FormBuilder
) {
        super(http);
        this.inventoryItemForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  quantityOnHand: ['', Validators.required],
      quantityAvailable: ['', Validators.required],
      quantityReserved: ['', Validators.required],
      unitCost: ['', Validators.required],
      lastUpdated: ['', Validators.required],
      Sku: ['', ],
      Warehouse: ['', ],
      Location: ['', ],
      Lot: ['', ],
      SerialNumbers: ['', ],
      Transactions: ['', ],
      Reservations: ['', ],
      StockStatus: ['', ]
        });
    }

    
    updateInventoryItem(quantityOnHand, quantityAvailable, quantityReserved, unitCost, lastUpdated, Sku, Warehouse, Location, Lot, SerialNumbers, Transactions, Reservations, StockStatus): void {
        this.route.params.subscribe((params) => {

                        this.service.updateInventoryItem(quantityOnHand, quantityAvailable, quantityReserved, unitCost, lastUpdated, Sku, Warehouse, Location, Lot, SerialNumbers, Transactions, Reservations, StockStatus, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexInventoryItem']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getInventoryItem(params['id']).subscribe(res => {
                this.inventoryItem = res;
            });
        });
    }
}