import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { SettlementBatchService } from '../../../services/SettlementBatch.service';
import { SubBaseComponent } from '../../SettlementBatch/sub.base.component';


@Component({
    selector: 'app-edit-settlementBatch',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditSettlementBatchComponent extends SubBaseComponent implements OnInit {

    title = 'Edit SettlementBatch';

    settlementBatchForm: FormGroup;
    settlementBatch: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: SettlementBatchService,
        private fb: FormBuilder
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

    
    updateSettlementBatch(batchId, periodStart, periodEnd, totalVolume, totalCount, Processor, Merchant, Payouts, Transactions, Status): void {
        this.route.params.subscribe((params) => {

                        this.service.updateSettlementBatch(batchId, periodStart, periodEnd, totalVolume, totalCount, Processor, Merchant, Payouts, Transactions, Status, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexSettlementBatch']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getSettlementBatch(params['id']).subscribe(res => {
                this.settlementBatch = res;
            });
        });
    }
}