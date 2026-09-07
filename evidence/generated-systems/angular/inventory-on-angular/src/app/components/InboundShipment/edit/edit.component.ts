import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { InboundShipmentService } from '../../../services/InboundShipment.service';
import { SubBaseComponent } from '../../InboundShipment/sub.base.component';


@Component({
    selector: 'app-edit-inboundShipment',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditInboundShipmentComponent extends SubBaseComponent implements OnInit {

    title = 'Edit InboundShipment';

    inboundShipmentForm: FormGroup;
    inboundShipment: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: InboundShipmentService,
        private fb: FormBuilder
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

    
    updateInboundShipment(shipmentNumber, expectedArrivalDate, arrivalDate, carrierName, Warehouse, Lines, Transactions, Status): void {
        this.route.params.subscribe((params) => {

                        this.service.updateInboundShipment(shipmentNumber, expectedArrivalDate, arrivalDate, carrierName, Warehouse, Lines, Transactions, Status, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexInboundShipment']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getInboundShipment(params['id']).subscribe(res => {
                this.inboundShipment = res;
            });
        });
    }
}