import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { StockAdjustmentService } from '../../../services/StockAdjustment.service';
import { SubBaseComponent } from '../../StockAdjustment/sub.base.component';


@Component({
    selector: 'app-edit-stockAdjustment',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditStockAdjustmentComponent extends SubBaseComponent implements OnInit {

    title = 'Edit StockAdjustment';

    stockAdjustmentForm: FormGroup;
    stockAdjustment: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: StockAdjustmentService,
        private fb: FormBuilder
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

    
    updateStockAdjustment(adjustmentNumber, reason, adjustmentDate, Warehouse, Lines, Transactions, AdjustmentType, Status): void {
        this.route.params.subscribe((params) => {

                        this.service.updateStockAdjustment(adjustmentNumber, reason, adjustmentDate, Warehouse, Lines, Transactions, AdjustmentType, Status, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexStockAdjustment']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getStockAdjustment(params['id']).subscribe(res => {
                this.stockAdjustment = res;
            });
        });
    }
}