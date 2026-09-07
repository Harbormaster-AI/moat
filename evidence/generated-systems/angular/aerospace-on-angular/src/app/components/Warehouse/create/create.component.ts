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
      InventoryItems: ['', ]
        });
    }

    
    addWarehouse(name, InventoryItems): void {
        this.warehouseService
        .addWarehouse(name, InventoryItems)
            .subscribe(() => {
                this.router.navigate(['/indexWarehouse']);
            });
    }

    ngOnInit(): void {
    }
}