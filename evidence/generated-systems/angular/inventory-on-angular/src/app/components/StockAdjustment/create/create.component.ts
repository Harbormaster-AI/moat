import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { StockAdjustmentService } from '../../../services/StockAdjustment.service';
import { StockAdjustment } from '../../../models/StockAdjustment';
import { SubBaseComponent } from '../../StockAdjustment/sub.base.component';

@Component({
    selector: 'app-create-stockAdjustment',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateStockAdjustmentComponent extends SubBaseComponent implements OnInit {

    title = 'Add StockAdjustment';

    stockAdjustmentForm: FormGroup;
    stockAdjustment: StockAdjustment;

    constructor( http: HttpClient,
        private stockAdjustmentService: StockAdjustmentService,
        private fb: FormBuilder,
        private router: Router
) {
        super(http);
        this.stockAdjustmentForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  adjustmentNumber: ['', Validators.required],
      reason: ['', Validators.required],
      adjustmentDate: ['', Validators.required],
      Warehouse: ['', ],
      Lines: ['', ],
      Transactions: ['', ],
      AdjustmentType: ['', ],
      Status: ['', ]
        });
    }

    
    addStockAdjustment(adjustmentNumber, reason, adjustmentDate, Warehouse, Lines, Transactions, AdjustmentType, Status): void {
        this.stockAdjustmentService
        .addStockAdjustment(adjustmentNumber, reason, adjustmentDate, Warehouse, Lines, Transactions, AdjustmentType, Status)
            .subscribe(() => {
                this.router.navigate(['/indexStockAdjustment']);
            });
    }

    ngOnInit(): void {
    }
}