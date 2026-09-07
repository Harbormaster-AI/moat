import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { InventoryThresholdAlertService } from '../../../services/InventoryThresholdAlert.service';
import { SubBaseComponent } from '../../InventoryThresholdAlert/sub.base.component';


@Component({
    selector: 'app-edit-inventoryThresholdAlert',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditInventoryThresholdAlertComponent extends SubBaseComponent implements OnInit {

    title = 'Edit InventoryThresholdAlert';

    inventoryThresholdAlertForm: FormGroup;
    inventoryThresholdAlert: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: InventoryThresholdAlertService,
        private fb: FormBuilder
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

    
    updateInventoryThresholdAlert(alertNumber, detectedAt, message, Sku, Warehouse, Location, RelatedPolicy, AlertType, Status): void {
        this.route.params.subscribe((params) => {

                        this.service.updateInventoryThresholdAlert(alertNumber, detectedAt, message, Sku, Warehouse, Location, RelatedPolicy, AlertType, Status, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexInventoryThresholdAlert']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getInventoryThresholdAlert(params['id']).subscribe(res => {
                this.inventoryThresholdAlert = res;
            });
        });
    }
}