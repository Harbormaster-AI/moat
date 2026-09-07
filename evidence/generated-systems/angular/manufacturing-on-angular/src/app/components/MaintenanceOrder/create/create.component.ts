import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { MaintenanceOrderService } from '../../../services/MaintenanceOrder.service';
import { MaintenanceOrder } from '../../../models/MaintenanceOrder';
import { SubBaseComponent } from '../../MaintenanceOrder/sub.base.component';

@Component({
    selector: 'app-create-maintenanceOrder',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateMaintenanceOrderComponent extends SubBaseComponent implements OnInit {

    title = 'Add MaintenanceOrder';

    maintenanceOrderForm: FormGroup;
    maintenanceOrder: MaintenanceOrder;

    constructor( http: HttpClient,
        private maintenanceOrderService: MaintenanceOrderService,
        private fb: FormBuilder,
        private router: Router
) {
        super(http);
        this.maintenanceOrderForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  orderNumber: ['', Validators.required],
      priority: ['', Validators.required],
      requestedDate: ['', Validators.required],
      completionDate: ['', Validators.required],
      Asset: ['', ],
      Plan: ['', ],
      WorkCenter: ['', ],
      Status: ['', ]
        });
    }

    
    addMaintenanceOrder(orderNumber, priority, requestedDate, completionDate, Asset, Plan, WorkCenter, Status): void {
        this.maintenanceOrderService
        .addMaintenanceOrder(orderNumber, priority, requestedDate, completionDate, Asset, Plan, WorkCenter, Status)
            .subscribe(() => {
                this.router.navigate(['/indexMaintenanceOrder']);
            });
    }

    ngOnInit(): void {
    }
}