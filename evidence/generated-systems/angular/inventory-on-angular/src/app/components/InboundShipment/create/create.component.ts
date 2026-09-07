import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { InboundShipmentService } from '../../../services/InboundShipment.service';
import { InboundShipment } from '../../../models/InboundShipment';
import { SubBaseComponent } from '../../InboundShipment/sub.base.component';

@Component({
    selector: 'app-create-inboundShipment',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateInboundShipmentComponent extends SubBaseComponent implements OnInit {

    title = 'Add InboundShipment';

    inboundShipmentForm: FormGroup;
    inboundShipment: InboundShipment;

    constructor( http: HttpClient,
        private inboundShipmentService: InboundShipmentService,
        private fb: FormBuilder,
        private router: Router
) {
        super(http);
        this.inboundShipmentForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  shipmentNumber: ['', Validators.required],
      expectedArrivalDate: ['', Validators.required],
      arrivalDate: ['', Validators.required],
      carrierName: ['', Validators.required],
      Warehouse: ['', ],
      Lines: ['', ],
      Transactions: ['', ],
      Status: ['', ]
        });
    }

    
    addInboundShipment(shipmentNumber, expectedArrivalDate, arrivalDate, carrierName, Warehouse, Lines, Transactions, Status): void {
        this.inboundShipmentService
        .addInboundShipment(shipmentNumber, expectedArrivalDate, arrivalDate, carrierName, Warehouse, Lines, Transactions, Status)
            .subscribe(() => {
                this.router.navigate(['/indexInboundShipment']);
            });
    }

    ngOnInit(): void {
    }
}