import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { InventoryThresholdAlertService } from '../../../services/InventoryThresholdAlert.service';
import { InventoryThresholdAlert } from '../../../models/InventoryThresholdAlert';
import { SubBaseComponent } from '../../InventoryThresholdAlert/sub.base.component';

@Component({
    selector: 'app-create-inventoryThresholdAlert',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateInventoryThresholdAlertComponent extends SubBaseComponent implements OnInit {

    title = 'Add InventoryThresholdAlert';

    inventoryThresholdAlertForm: FormGroup;
    inventoryThresholdAlert: InventoryThresholdAlert;

    constructor( http: HttpClient,
        private inventoryThresholdAlertService: InventoryThresholdAlertService,
        private fb: FormBuilder,
        private router: Router
) {
        super(http);
        this.inventoryThresholdAlertForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  alertNumber: ['', Validators.required],
      detectedAt: ['', Validators.required],
      message: ['', Validators.required],
      Sku: ['', ],
      Warehouse: ['', ],
      Location: ['', ],
      RelatedPolicy: ['', ],
      AlertType: ['', ],
      Status: ['', ]
        });
    }

    
    addInventoryThresholdAlert(alertNumber, detectedAt, message, Sku, Warehouse, Location, RelatedPolicy, AlertType, Status): void {
        this.inventoryThresholdAlertService
        .addInventoryThresholdAlert(alertNumber, detectedAt, message, Sku, Warehouse, Location, RelatedPolicy, AlertType, Status)
            .subscribe(() => {
                this.router.navigate(['/indexInventoryThresholdAlert']);
            });
    }

    ngOnInit(): void {
    }
}