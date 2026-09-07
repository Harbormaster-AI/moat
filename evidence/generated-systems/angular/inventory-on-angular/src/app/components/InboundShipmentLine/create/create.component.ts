import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { InboundShipmentLineService } from '../../../services/InboundShipmentLine.service';
import { InboundShipmentLine } from '../../../models/InboundShipmentLine';
import { SubBaseComponent } from '../../InboundShipmentLine/sub.base.component';

@Component({
    selector: 'app-create-inboundShipmentLine',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateInboundShipmentLineComponent extends SubBaseComponent implements OnInit {

    title = 'Add InboundShipmentLine';

    inboundShipmentLineForm: FormGroup;
    inboundShipmentLine: InboundShipmentLine;

    constructor( http: HttpClient,
        private inboundShipmentLineService: InboundShipmentLineService,
        private fb: FormBuilder,
        private router: Router
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

    
    addInboundShipmentLine(lineNumber, quantity, InboundShipment, Sku, Lot, SerialNumbers, DestinationLocation, UnitOfMeasure, StockStatus): void {
        this.inboundShipmentLineService
        .addInboundShipmentLine(lineNumber, quantity, InboundShipment, Sku, Lot, SerialNumbers, DestinationLocation, UnitOfMeasure, StockStatus)
            .subscribe(() => {
                this.router.navigate(['/indexInboundShipmentLine']);
            });
    }

    ngOnInit(): void {
    }
}