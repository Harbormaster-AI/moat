import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { OutboundAllocationService } from '../../../services/OutboundAllocation.service';
import { OutboundAllocation } from '../../../models/OutboundAllocation';
import { SubBaseComponent } from '../../OutboundAllocation/sub.base.component';

@Component({
    selector: 'app-create-outboundAllocation',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateOutboundAllocationComponent extends SubBaseComponent implements OnInit {

    title = 'Add OutboundAllocation';

    outboundAllocationForm: FormGroup;
    outboundAllocation: OutboundAllocation;

    constructor( http: HttpClient,
        private outboundAllocationService: OutboundAllocationService,
        private fb: FormBuilder,
        private router: Router
) {
        super(http);
        this.outboundAllocationForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  allocationNumber: ['', Validators.required],
      allocatedQuantity: ['', Validators.required],
      allocationDate: ['', Validators.required],
      Warehouse: ['', ],
      Sku: ['', ],
      InventoryItem: ['', ],
      Reservation: ['', ],
      Lot: ['', ],
      SerialNumbers: ['', ],
      SourceLocation: ['', ],
      Status: ['', ]
        });
    }

    
    addOutboundAllocation(allocationNumber, allocatedQuantity, allocationDate, Warehouse, Sku, InventoryItem, Reservation, Lot, SerialNumbers, SourceLocation, Status): void {
        this.outboundAllocationService
        .addOutboundAllocation(allocationNumber, allocatedQuantity, allocationDate, Warehouse, Sku, InventoryItem, Reservation, Lot, SerialNumbers, SourceLocation, Status)
            .subscribe(() => {
                this.router.navigate(['/indexOutboundAllocation']);
            });
    }

    ngOnInit(): void {
    }
}