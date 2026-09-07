import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { StorageLocationService } from '../../../services/StorageLocation.service';
import { SubBaseComponent } from '../../StorageLocation/sub.base.component';


@Component({
    selector: 'app-edit-storageLocation',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditStorageLocationComponent extends SubBaseComponent implements OnInit {

    title = 'Edit StorageLocation';

    storageLocationForm: FormGroup;
    storageLocation: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: StorageLocationService,
        private fb: FormBuilder
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

    
    updateStorageLocation(code, temperatureControlled, capacity, capacityUnit, Warehouse, ParentLocation, ChildLocations, InventoryItems, LocationType): void {
        this.route.params.subscribe((params) => {

                        this.service.updateStorageLocation(code, temperatureControlled, capacity, capacityUnit, Warehouse, ParentLocation, ChildLocations, InventoryItems, LocationType, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexStorageLocation']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getStorageLocation(params['id']).subscribe(res => {
                this.storageLocation = res;
            });
        });
    }
}