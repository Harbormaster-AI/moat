import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { TransferOrderService } from '../../../services/TransferOrder.service';
import { TransferOrder } from '../../../models/TransferOrder';
import { SubBaseComponent } from '../../TransferOrder/sub.base.component';

@Component({
    selector: 'app-create-transferOrder',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateTransferOrderComponent extends SubBaseComponent implements OnInit {

    title = 'Add TransferOrder';

    transferOrderForm: FormGroup;
    transferOrder: TransferOrder;

    constructor( http: HttpClient,
        private transferOrderService: TransferOrderService,
        private fb: FormBuilder,
        private router: Router
) {
        super(http);
        this.transferOrderForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  orderNumber: ['', Validators.required],
      requestedShipDate: ['', Validators.required],
      requestedReceiveDate: ['', Validators.required],
      shippedDate: ['', Validators.required],
      receivedDate: ['', Validators.required],
      OriginWarehouse: ['', ],
      DestinationWarehouse: ['', ],
      Lines: ['', ],
      Transactions: ['', ],
      Status: ['', ]
        });
    }

    
    addTransferOrder(orderNumber, requestedShipDate, requestedReceiveDate, shippedDate, receivedDate, OriginWarehouse, DestinationWarehouse, Lines, Transactions, Status): void {
        this.transferOrderService
        .addTransferOrder(orderNumber, requestedShipDate, requestedReceiveDate, shippedDate, receivedDate, OriginWarehouse, DestinationWarehouse, Lines, Transactions, Status)
            .subscribe(() => {
                this.router.navigate(['/indexTransferOrder']);
            });
    }

    ngOnInit(): void {
    }
}