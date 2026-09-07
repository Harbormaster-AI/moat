import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { WarehouseService } from '../../../services/Warehouse.service';
import { Warehouse } from '../../../models/Warehouse';
import { SubBaseComponent } from '../../Warehouse/sub.base.component';

@Component({
    selector: 'app-create-warehouse',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateWarehouseComponent extends SubBaseComponent implements OnInit {

    title = 'Add Warehouse';

    warehouseForm: FormGroup;
    warehouse: Warehouse;

    constructor( http: HttpClient,
        private warehouseService: WarehouseService,
        private fb: FormBuilder,
        private router: Router
) {
        super(http);
        this.warehouseForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  name: ['', Validators.required],
      code: ['', Validators.required],
      address: ['', Validators.required],
      timeZone: ['', Validators.required],
      allowsOverAllocation: ['', Validators.required],
      StorageLocations: ['', ],
      InventoryItems: ['', ],
      InboundShipments: ['', ],
      OutboundAllocations: ['', ],
      OriginTransfers: ['', ],
      DestinationTransfers: ['', ],
      CycleCounts: ['', ]
        });
    }

    
    addWarehouse(name, code, address, timeZone, allowsOverAllocation, StorageLocations, InventoryItems, InboundShipments, OutboundAllocations, OriginTransfers, DestinationTransfers, CycleCounts): void {
        this.warehouseService
        .addWarehouse(name, code, address, timeZone, allowsOverAllocation, StorageLocations, InventoryItems, InboundShipments, OutboundAllocations, OriginTransfers, DestinationTransfers, CycleCounts)
            .subscribe(() => {
                this.router.navigate(['/indexWarehouse']);
            });
    }

    ngOnInit(): void {
    }
}