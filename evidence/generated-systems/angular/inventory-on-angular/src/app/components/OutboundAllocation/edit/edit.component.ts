import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { OutboundAllocationService } from '../../../services/OutboundAllocation.service';
import { SubBaseComponent } from '../../OutboundAllocation/sub.base.component';


@Component({
    selector: 'app-edit-outboundAllocation',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditOutboundAllocationComponent extends SubBaseComponent implements OnInit {

    title = 'Edit OutboundAllocation';

    outboundAllocationForm: FormGroup;
    outboundAllocation: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: OutboundAllocationService,
        private fb: FormBuilder
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

    
    updateOutboundAllocation(allocationNumber, allocatedQuantity, allocationDate, Warehouse, Sku, InventoryItem, Reservation, Lot, SerialNumbers, SourceLocation, Status): void {
        this.route.params.subscribe((params) => {

                        this.service.updateOutboundAllocation(allocationNumber, allocatedQuantity, allocationDate, Warehouse, Sku, InventoryItem, Reservation, Lot, SerialNumbers, SourceLocation, Status, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexOutboundAllocation']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getOutboundAllocation(params['id']).subscribe(res => {
                this.outboundAllocation = res;
            });
        });
    }
}