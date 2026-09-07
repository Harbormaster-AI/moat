import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { InventoryItemService } from '../../../services/InventoryItem.service';
import { InventoryItem } from '../../../models/InventoryItem';
import { SubBaseComponent } from '../../InventoryItem/sub.base.component';

@Component({
    selector: 'app-create-inventoryItem',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateInventoryItemComponent extends SubBaseComponent implements OnInit {

    title = 'Add InventoryItem';

    inventoryItemForm: FormGroup;
    inventoryItem: InventoryItem;

    constructor( http: HttpClient,
        private inventoryItemService: InventoryItemService,
        private fb: FormBuilder,
        private router: Router
) {
        super(http);
        this.inventoryItemForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  quantityOnHand: ['', Validators.required],
      quantityReserved: ['', Validators.required],
      lotNumber: ['', Validators.required],
      Component: ['', ],
      Warehouse: ['', ]
        });
    }

    
    addInventoryItem(quantityOnHand, quantityReserved, lotNumber, Component, Warehouse): void {
        this.inventoryItemService
        .addInventoryItem(quantityOnHand, quantityReserved, lotNumber, Component, Warehouse)
            .subscribe(() => {
                this.router.navigate(['/indexInventoryItem']);
            });
    }

    ngOnInit(): void {
    }
}