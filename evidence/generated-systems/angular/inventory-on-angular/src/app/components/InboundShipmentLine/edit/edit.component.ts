import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { InboundShipmentLineService } from '../../../services/InboundShipmentLine.service';
import { SubBaseComponent } from '../../InboundShipmentLine/sub.base.component';


@Component({
    selector: 'app-edit-inboundShipmentLine',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditInboundShipmentLineComponent extends SubBaseComponent implements OnInit {

    title = 'Edit InboundShipmentLine';

    inboundShipmentLineForm: FormGroup;
    inboundShipmentLine: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: InboundShipmentLineService,
        private fb: FormBuilder
) {
        super(http);
        this.inboundShipmentLineForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  lineNumber: ['', Validators.required],
      quantity: ['', Validators.required],
      InboundShipment: ['', ],
      Sku: ['', ],
      Lot: ['', ],
      SerialNumbers: ['', ],
      DestinationLocation: ['', ],
      UnitOfMeasure: ['', ],
      StockStatus: ['', ]
        });
    }

    
    updateInboundShipmentLine(lineNumber, quantity, InboundShipment, Sku, Lot, SerialNumbers, DestinationLocation, UnitOfMeasure, StockStatus): void {
        this.route.params.subscribe((params) => {

                        this.service.updateInboundShipmentLine(lineNumber, quantity, InboundShipment, Sku, Lot, SerialNumbers, DestinationLocation, UnitOfMeasure, StockStatus, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexInboundShipmentLine']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getInboundShipmentLine(params['id']).subscribe(res => {
                this.inboundShipmentLine = res;
            });
        });
    }
}