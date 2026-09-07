import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { ItemService } from '../../../services/Item.service';
import { Item } from '../../../models/Item';
import { SubBaseComponent } from '../../Item/sub.base.component';

@Component({
    selector: 'app-create-item',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateItemComponent extends SubBaseComponent implements OnInit {

    title = 'Add Item';

    itemForm: FormGroup;
    item: Item;

    constructor( http: HttpClient,
        private itemService: ItemService,
        private fb: FormBuilder,
        private router: Router
) {
        super(http);
        this.itemForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  itemNumber: ['', Validators.required],
      name: ['', Validators.required],
      standardCost: ['', Validators.required],
      weight: ['', Validators.required],
      asSerialControlled: ['', Validators.required],
      BusinessUnit: ['', ],
      Boms: ['', ],
      Routings: ['', ],
      Suppliers: ['', ],
      QualitySpecifications: ['', ],
      InventoryItems: ['', ],
      ItemType: ['', ],
      ProcurementType: ['', ],
      UnitOfMeasure: ['', ],
      LifecycleStatus: ['', ]
        });
    }

    
    addItem(itemNumber, name, standardCost, weight, asSerialControlled, BusinessUnit, Boms, Routings, Suppliers, QualitySpecifications, InventoryItems, ItemType, ProcurementType, UnitOfMeasure, LifecycleStatus): void {
        this.itemService
        .addItem(itemNumber, name, standardCost, weight, asSerialControlled, BusinessUnit, Boms, Routings, Suppliers, QualitySpecifications, InventoryItems, ItemType, ProcurementType, UnitOfMeasure, LifecycleStatus)
            .subscribe(() => {
                this.router.navigate(['/indexItem']);
            });
    }

    ngOnInit(): void {
    }
}