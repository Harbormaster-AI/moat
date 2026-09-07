import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { TransferOrderLineService } from '../../../services/TransferOrderLine.service';
import { TransferOrderLine } from '../../../models/TransferOrderLine';
import { SubBaseComponent } from '../../TransferOrderLine/sub.base.component';

@Component({
    selector: 'app-create-transferOrderLine',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateTransferOrderLineComponent extends SubBaseComponent implements OnInit {

    title = 'Add TransferOrderLine';

    transferOrderLineForm: FormGroup;
    transferOrderLine: TransferOrderLine;

    constructor( http: HttpClient,
        private transferOrderLineService: TransferOrderLineService,
        private fb: FormBuilder,
        private router: Router
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

    
    addTransferOrderLine(lineNumber, quantity, TransferOrder, Sku, Lot, SerialNumbers, FromLocation, ToLocation, UnitOfMeasure, StockStatus): void {
        this.transferOrderLineService
        .addTransferOrderLine(lineNumber, quantity, TransferOrder, Sku, Lot, SerialNumbers, FromLocation, ToLocation, UnitOfMeasure, StockStatus)
            .subscribe(() => {
                this.router.navigate(['/indexTransferOrderLine']);
            });
    }

    ngOnInit(): void {
    }
}