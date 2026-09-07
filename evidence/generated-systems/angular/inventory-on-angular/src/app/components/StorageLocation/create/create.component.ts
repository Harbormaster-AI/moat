import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { StorageLocationService } from '../../../services/StorageLocation.service';
import { StorageLocation } from '../../../models/StorageLocation';
import { SubBaseComponent } from '../../StorageLocation/sub.base.component';

@Component({
    selector: 'app-create-storageLocation',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateStorageLocationComponent extends SubBaseComponent implements OnInit {

    title = 'Add StorageLocation';

    storageLocationForm: FormGroup;
    storageLocation: StorageLocation;

    constructor( http: HttpClient,
        private storageLocationService: StorageLocationService,
        private fb: FormBuilder,
        private router: Router
) {
        super(http);
        this.storageLocationForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  code: ['', Validators.required],
      temperatureControlled: ['', Validators.required],
      capacity: ['', Validators.required],
      capacityUnit: ['', Validators.required],
      Warehouse: ['', ],
      ParentLocation: ['', ],
      ChildLocations: ['', ],
      InventoryItems: ['', ],
      LocationType: ['', ]
        });
    }

    
    addStorageLocation(code, temperatureControlled, capacity, capacityUnit, Warehouse, ParentLocation, ChildLocations, InventoryItems, LocationType): void {
        this.storageLocationService
        .addStorageLocation(code, temperatureControlled, capacity, capacityUnit, Warehouse, ParentLocation, ChildLocations, InventoryItems, LocationType)
            .subscribe(() => {
                this.router.navigate(['/indexStorageLocation']);
            });
    }

    ngOnInit(): void {
    }
}