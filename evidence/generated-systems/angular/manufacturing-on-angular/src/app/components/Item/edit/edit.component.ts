import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { ItemService } from '../../../services/Item.service';
import { SubBaseComponent } from '../../Item/sub.base.component';


@Component({
    selector: 'app-edit-item',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditItemComponent extends SubBaseComponent implements OnInit {

    title = 'Edit Item';

    itemForm: FormGroup;
    item: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: ItemService,
        private fb: FormBuilder
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

    
    updateItem(itemNumber, name, standardCost, weight, asSerialControlled, BusinessUnit, Boms, Routings, Suppliers, QualitySpecifications, InventoryItems, ItemType, ProcurementType, UnitOfMeasure, LifecycleStatus): void {
        this.route.params.subscribe((params) => {

                        this.service.updateItem(itemNumber, name, standardCost, weight, asSerialControlled, BusinessUnit, Boms, Routings, Suppliers, QualitySpecifications, InventoryItems, ItemType, ProcurementType, UnitOfMeasure, LifecycleStatus, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexItem']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getItem(params['id']).subscribe(res => {
                this.item = res;
            });
        });
    }
}