import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { StockKeepingUnitService } from '../../../services/StockKeepingUnit.service';
import { SubBaseComponent } from '../../StockKeepingUnit/sub.base.component';


@Component({
    selector: 'app-edit-stockKeepingUnit',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditStockKeepingUnitComponent extends SubBaseComponent implements OnInit {

    title = 'Edit StockKeepingUnit';

    stockKeepingUnitForm: FormGroup;
    stockKeepingUnit: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: StockKeepingUnitService,
        private fb: FormBuilder
) {
        super(http);
        this.stockKeepingUnitForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  skuCode: ['', Validators.required],
      name: ['', Validators.required],
      weight: ['', Validators.required],
      weightUnit: ['', Validators.required],
      volume: ['', Validators.required],
      volumeUnit: ['', Validators.required],
      shelfLifeDays: ['', Validators.required],
      hazardousMaterial: ['', Validators.required],
      InventoryItems: ['', ],
      UomConversions: ['', ],
      ReplenishmentPolicies: ['', ],
      Lots: ['', ],
      SerialNumbers: ['', ],
      ItemType: ['', ],
      UnitOfMeasure: ['', ]
        });
    }

    
    updateStockKeepingUnit(skuCode, name, weight, weightUnit, volume, volumeUnit, shelfLifeDays, hazardousMaterial, InventoryItems, UomConversions, ReplenishmentPolicies, Lots, SerialNumbers, ItemType, UnitOfMeasure): void {
        this.route.params.subscribe((params) => {

                        this.service.updateStockKeepingUnit(skuCode, name, weight, weightUnit, volume, volumeUnit, shelfLifeDays, hazardousMaterial, InventoryItems, UomConversions, ReplenishmentPolicies, Lots, SerialNumbers, ItemType, UnitOfMeasure, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexStockKeepingUnit']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getStockKeepingUnit(params['id']).subscribe(res => {
                this.stockKeepingUnit = res;
            });
        });
    }
}