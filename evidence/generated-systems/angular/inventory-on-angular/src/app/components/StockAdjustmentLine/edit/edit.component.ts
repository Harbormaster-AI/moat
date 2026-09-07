import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { StockAdjustmentLineService } from '../../../services/StockAdjustmentLine.service';
import { SubBaseComponent } from '../../StockAdjustmentLine/sub.base.component';


@Component({
    selector: 'app-edit-stockAdjustmentLine',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditStockAdjustmentLineComponent extends SubBaseComponent implements OnInit {

    title = 'Edit StockAdjustmentLine';

    stockAdjustmentLineForm: FormGroup;
    stockAdjustmentLine: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: StockAdjustmentLineService,
        private fb: FormBuilder
) {
        super(http);
        this.stockAdjustmentLineForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  lineNumber: ['', Validators.required],
      quantity: ['', Validators.required],
      Adjustment: ['', ],
      Sku: ['', ],
      Lot: ['', ],
      Location: ['', ],
      SerialNumbers: ['', ],
      UnitOfMeasure: ['', ],
      StockStatus: ['', ]
        });
    }

    
    updateStockAdjustmentLine(lineNumber, quantity, Adjustment, Sku, Lot, Location, SerialNumbers, UnitOfMeasure, StockStatus): void {
        this.route.params.subscribe((params) => {

                        this.service.updateStockAdjustmentLine(lineNumber, quantity, Adjustment, Sku, Lot, Location, SerialNumbers, UnitOfMeasure, StockStatus, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexStockAdjustmentLine']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getStockAdjustmentLine(params['id']).subscribe(res => {
                this.stockAdjustmentLine = res;
            });
        });
    }
}