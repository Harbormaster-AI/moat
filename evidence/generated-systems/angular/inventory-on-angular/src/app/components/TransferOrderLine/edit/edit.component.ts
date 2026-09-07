import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { TransferOrderLineService } from '../../../services/TransferOrderLine.service';
import { SubBaseComponent } from '../../TransferOrderLine/sub.base.component';


@Component({
    selector: 'app-edit-transferOrderLine',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditTransferOrderLineComponent extends SubBaseComponent implements OnInit {

    title = 'Edit TransferOrderLine';

    transferOrderLineForm: FormGroup;
    transferOrderLine: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: TransferOrderLineService,
        private fb: FormBuilder
) {
        super(http);
        this.transferOrderLineForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  lineNumber: ['', Validators.required],
      quantity: ['', Validators.required],
      TransferOrder: ['', ],
      Sku: ['', ],
      Lot: ['', ],
      SerialNumbers: ['', ],
      FromLocation: ['', ],
      ToLocation: ['', ],
      UnitOfMeasure: ['', ],
      StockStatus: ['', ]
        });
    }

    
    updateTransferOrderLine(lineNumber, quantity, TransferOrder, Sku, Lot, SerialNumbers, FromLocation, ToLocation, UnitOfMeasure, StockStatus): void {
        this.route.params.subscribe((params) => {

                        this.service.updateTransferOrderLine(lineNumber, quantity, TransferOrder, Sku, Lot, SerialNumbers, FromLocation, ToLocation, UnitOfMeasure, StockStatus, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexTransferOrderLine']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getTransferOrderLine(params['id']).subscribe(res => {
                this.transferOrderLine = res;
            });
        });
    }
}