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
                  sku: ['', Validators.required],
      name: ['', Validators.required],
      quantityOnHand: ['', Validators.required],
      quantityReserved: ['', Validators.required],
      Facility: ['', ],
      Supplier: ['', ]
        });
    }

    
    updateInventoryItem(sku, name, quantityOnHand, quantityReserved, Facility, Supplier): void {
        this.route.params.subscribe((params) => {

                        this.service.updateInventoryItem(sku, name, quantityOnHand, quantityReserved, Facility, Supplier, params['id'])
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