import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { WarehouseService } from '../../../services/Warehouse.service';
import { SubBaseComponent } from '../../Warehouse/sub.base.component';


@Component({
    selector: 'app-edit-warehouse',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditWarehouseComponent extends SubBaseComponent implements OnInit {

    title = 'Edit Warehouse';

    warehouseForm: FormGroup;
    warehouse: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: WarehouseService,
        private fb: FormBuilder
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

    
    updateWarehouse(name, code, address, timeZone, allowsOverAllocation, StorageLocations, InventoryItems, InboundShipments, OutboundAllocations, OriginTransfers, DestinationTransfers, CycleCounts): void {
        this.route.params.subscribe((params) => {

                        this.service.updateWarehouse(name, code, address, timeZone, allowsOverAllocation, StorageLocations, InventoryItems, InboundShipments, OutboundAllocations, OriginTransfers, DestinationTransfers, CycleCounts, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexWarehouse']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getWarehouse(params['id']).subscribe(res => {
                this.warehouse = res;
            });
        });
    }
}