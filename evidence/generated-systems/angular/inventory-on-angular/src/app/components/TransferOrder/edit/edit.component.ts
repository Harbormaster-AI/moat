import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { TransferOrderService } from '../../../services/TransferOrder.service';
import { SubBaseComponent } from '../../TransferOrder/sub.base.component';


@Component({
    selector: 'app-edit-transferOrder',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditTransferOrderComponent extends SubBaseComponent implements OnInit {

    title = 'Edit TransferOrder';

    transferOrderForm: FormGroup;
    transferOrder: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: TransferOrderService,
        private fb: FormBuilder
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

    
    updateTransferOrder(orderNumber, requestedShipDate, requestedReceiveDate, shippedDate, receivedDate, OriginWarehouse, DestinationWarehouse, Lines, Transactions, Status): void {
        this.route.params.subscribe((params) => {

                        this.service.updateTransferOrder(orderNumber, requestedShipDate, requestedReceiveDate, shippedDate, receivedDate, OriginWarehouse, DestinationWarehouse, Lines, Transactions, Status, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexTransferOrder']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getTransferOrder(params['id']).subscribe(res => {
                this.transferOrder = res;
            });
        });
    }
}