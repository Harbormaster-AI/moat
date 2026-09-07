import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { StockKeepingUnitService } from '../../../services/StockKeepingUnit.service';
import { StockKeepingUnit } from '../../../models/StockKeepingUnit';
import { SubBaseComponent } from '../../StockKeepingUnit/sub.base.component';

@Component({
    selector: 'app-create-stockKeepingUnit',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateStockKeepingUnitComponent extends SubBaseComponent implements OnInit {

    title = 'Add StockKeepingUnit';

    stockKeepingUnitForm: FormGroup;
    stockKeepingUnit: StockKeepingUnit;

    constructor( http: HttpClient,
        private stockKeepingUnitService: StockKeepingUnitService,
        private fb: FormBuilder,
        private router: Router
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

    
    addStockKeepingUnit(skuCode, name, weight, weightUnit, volume, volumeUnit, shelfLifeDays, hazardousMaterial, InventoryItems, UomConversions, ReplenishmentPolicies, Lots, SerialNumbers, ItemType, UnitOfMeasure): void {
        this.stockKeepingUnitService
        .addStockKeepingUnit(skuCode, name, weight, weightUnit, volume, volumeUnit, shelfLifeDays, hazardousMaterial, InventoryItems, UomConversions, ReplenishmentPolicies, Lots, SerialNumbers, ItemType, UnitOfMeasure)
            .subscribe(() => {
                this.router.navigate(['/indexStockKeepingUnit']);
            });
    }

    ngOnInit(): void {
    }
}