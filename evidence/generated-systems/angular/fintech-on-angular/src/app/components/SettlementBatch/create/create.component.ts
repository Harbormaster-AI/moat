import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { SettlementBatchService } from '../../../services/SettlementBatch.service';
import { SettlementBatch } from '../../../models/SettlementBatch';
import { SubBaseComponent } from '../../SettlementBatch/sub.base.component';

@Component({
    selector: 'app-create-settlementBatch',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateSettlementBatchComponent extends SubBaseComponent implements OnInit {

    title = 'Add SettlementBatch';

    settlementBatchForm: FormGroup;
    settlementBatch: SettlementBatch;

    constructor( http: HttpClient,
        private settlementBatchService: SettlementBatchService,
        private fb: FormBuilder,
        private router: Router
) {
        super(http);
        this.settlementBatchForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  batchId: ['', Validators.required],
      periodStart: ['', Validators.required],
      periodEnd: ['', Validators.required],
      totalVolume: ['', Validators.required],
      totalCount: ['', Validators.required],
      Processor: ['', ],
      Merchant: ['', ],
      Payouts: ['', ],
      Transactions: ['', ],
      Status: ['', ]
        });
    }

    
    addSettlementBatch(batchId, periodStart, periodEnd, totalVolume, totalCount, Processor, Merchant, Payouts, Transactions, Status): void {
        this.settlementBatchService
        .addSettlementBatch(batchId, periodStart, periodEnd, totalVolume, totalCount, Processor, Merchant, Payouts, Transactions, Status)
            .subscribe(() => {
                this.router.navigate(['/indexSettlementBatch']);
            });
    }

    ngOnInit(): void {
    }
}