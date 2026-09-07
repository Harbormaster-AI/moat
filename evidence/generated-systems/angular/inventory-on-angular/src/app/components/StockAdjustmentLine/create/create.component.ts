import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { StockAdjustmentLineService } from '../../../services/StockAdjustmentLine.service';
import { StockAdjustmentLine } from '../../../models/StockAdjustmentLine';
import { SubBaseComponent } from '../../StockAdjustmentLine/sub.base.component';

@Component({
    selector: 'app-create-stockAdjustmentLine',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateStockAdjustmentLineComponent extends SubBaseComponent implements OnInit {

    title = 'Add StockAdjustmentLine';

    stockAdjustmentLineForm: FormGroup;
    stockAdjustmentLine: StockAdjustmentLine;

    constructor( http: HttpClient,
        private stockAdjustmentLineService: StockAdjustmentLineService,
        private fb: FormBuilder,
        private router: Router
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

    
    addStockAdjustmentLine(lineNumber, quantity, Adjustment, Sku, Lot, Location, SerialNumbers, UnitOfMeasure, StockStatus): void {
        this.stockAdjustmentLineService
        .addStockAdjustmentLine(lineNumber, quantity, Adjustment, Sku, Lot, Location, SerialNumbers, UnitOfMeasure, StockStatus)
            .subscribe(() => {
                this.router.navigate(['/indexStockAdjustmentLine']);
            });
    }

    ngOnInit(): void {
    }
}